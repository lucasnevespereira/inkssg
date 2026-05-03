package inkssg

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const reloadPath = "/__inkssg/reload"

const reloadScript = `<script>
(function(){
  if (window.__inkssgReload) return;
  window.__inkssgReload = true;
  var es = new EventSource("` + reloadPath + `");
  es.onmessage = function(e){ if (e.data === "reload") location.reload(); };
})();
</script>`

// reloader fans out reload pings to connected SSE clients.
type reloader struct {
	mu      sync.Mutex
	clients map[chan struct{}]struct{}
}

func newReloader() *reloader {
	return &reloader{clients: map[chan struct{}]struct{}{}}
}

func (r *reloader) subscribe() chan struct{} {
	ch := make(chan struct{}, 1)
	r.mu.Lock()
	r.clients[ch] = struct{}{}
	r.mu.Unlock()
	return ch
}

func (r *reloader) unsubscribe(ch chan struct{}) {
	r.mu.Lock()
	delete(r.clients, ch)
	r.mu.Unlock()
}

func (r *reloader) broadcast() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for ch := range r.clients {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}

func serveSSE(r *reloader) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ch := r.subscribe()
		defer r.unsubscribe(ch)

		fmt.Fprintf(w, ": connected\n\n")
		flusher.Flush()

		ctx := req.Context()
		ping := time.NewTicker(30 * time.Second)
		defer ping.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ch:
				fmt.Fprintf(w, "data: reload\n\n")
				flusher.Flush()
			case <-ping.C:
				fmt.Fprintf(w, ": ping\n\n")
				flusher.Flush()
			}
		}
	}
}

// htmlInjector wraps a file server so HTML responses get the reload script
// inserted before </body>. Non-HTML files pass through untouched.
type htmlInjector struct {
	root string
	next http.Handler
}

func (h *htmlInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if urlPath == "/" || strings.HasSuffix(urlPath, "/") {
		urlPath += "index.html"
	}
	if !strings.HasSuffix(urlPath, ".html") {
		h.next.ServeHTTP(w, r)
		return
	}

	clean := filepath.Clean(strings.TrimPrefix(urlPath, "/"))
	abs := filepath.Join(h.root, clean)
	rel, err := filepath.Rel(h.root, abs)
	if err != nil || strings.HasPrefix(rel, "..") {
		http.NotFound(w, r)
		return
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		h.next.ServeHTTP(w, r)
		return
	}

	out := injectReloadScript(data)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Write(out)
}

func injectReloadScript(html []byte) []byte {
	closing := []byte("</body>")
	idx := bytes.LastIndex(html, closing)
	if idx < 0 {
		return append(html, []byte(reloadScript)...)
	}
	out := make([]byte, 0, len(html)+len(reloadScript))
	out = append(out, html[:idx]...)
	out = append(out, []byte(reloadScript)...)
	out = append(out, html[idx:]...)
	return out
}

func runServer(ctx context.Context, addr, root string, r *reloader) error {
	mux := http.NewServeMux()
	mux.Handle(reloadPath, serveSSE(r))
	mux.Handle("/", &htmlInjector{root: root, next: http.FileServer(http.Dir(root))})

	srv := &http.Server{Addr: addr, Handler: mux}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		return srv.Shutdown(shutdownCtx)
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}
