package inkssg

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

type Site struct {
	Dir    string
	Pages  []Page
	Output string
}

type Page struct {
	Name        string
	Dir         string
	Content     string
	Title       string
	Description string
	Theme       string
	ContentHTML string
}

type SiteConfig struct {
	Name         string `yaml:"name"`
	DefaultTheme string `yaml:"default_theme"`
	OutputDir    string `yaml:"output_dir"`
}

func Build(paths ...string) error {
	dir := "."
	if len(paths) > 0 {
		dir = paths[0]
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	site, err := NewSite(absDir)
	if err != nil {
		return err
	}

	return site.Build()
}

func NewSite(dir string) (*Site, error) {
	site := &Site{
		Dir:    dir,
		Pages:  []Page{},
		Output: "public",
	}

	if err := site.detect(); err != nil {
		return nil, err
	}

	return site, nil
}

func (s *Site) detect() error {
	pagesDir := filepath.Join(s.Dir, "pages")
	if _, err := os.Stat(pagesDir); os.IsNotExist(err) {
		return fmt.Errorf("no pages/ folder found. Run 'inkssg new .' to scaffold a site.")
	}

	entries, err := os.ReadDir(pagesDir)
	if err != nil {
		return fmt.Errorf("cannot read pages/: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pageDir := filepath.Join(pagesDir, entry.Name())
		contentPath := filepath.Join(pageDir, "content.md")
		contentType := "md"
		if _, err := os.Stat(contentPath); os.IsNotExist(err) {
			contentPath = filepath.Join(pageDir, "content.html")
			contentType = "html"
			if _, err := os.Stat(contentPath); os.IsNotExist(err) {
				continue
			}
		}

		page := Page{
			Name:    entry.Name(),
			Dir:     pageDir,
			Content: contentPath,
		}

		if err := page.parseFrontmatter(contentType); err != nil {
			return fmt.Errorf("%s: %w", entry.Name(), err)
		}

		s.Pages = append(s.Pages, page)
	}

	if len(s.Pages) == 0 {
		return fmt.Errorf("no pages found. Add a pages/<name>/content.md or content.html file.")
	}

	return nil
}

func (p *Page) parseFrontmatter(contentType string) error {
	data, err := os.ReadFile(p.Content)
	if err != nil {
		return fmt.Errorf("cannot read content: %w", err)
	}

	frontmatter := &struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Theme       string `yaml:"theme"`
	}{}

	str := string(data)

	if strings.HasPrefix(str, "---") {
		str = strings.TrimPrefix(str, "---\n")
		endIdx := strings.Index(str, "\n---")
		if endIdx > 0 {
			yamlData := str[:endIdx]
			if err := yaml.Unmarshal([]byte(yamlData), frontmatter); err != nil {
				return fmt.Errorf("invalid frontmatter: %w", err)
			}
			content := str[endIdx+4:]
			p.ContentHTML = p.renderMarkdown([]byte(content), contentType)
		} else {
			p.ContentHTML = p.renderMarkdown(data, contentType)
		}
	} else {
		p.ContentHTML = p.renderMarkdown(data, contentType)
	}

	p.Title = frontmatter.Title
	p.Description = frontmatter.Description
	p.Theme = frontmatter.Theme

	if p.Title == "" {
		p.Title = p.Name
	}

	return nil
}

func (p *Page) renderMarkdown(data []byte, contentType string) string {
	if contentType == "html" {
		return string(data)
	}

	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert(data, &buf); err != nil {
		return ""
	}
	return buf.String()
}

func (s *Site) Build() error {
	start := time.Now()

	if err := os.RemoveAll(filepath.Join(s.Dir, s.Output)); err != nil {
		return fmt.Errorf("cannot clean output dir: %w", err)
	}

	if err := os.MkdirAll(filepath.Join(s.Dir, s.Output), 0755); err != nil {
		return fmt.Errorf("cannot create output dir: %w", err)
	}

	failed := 0
	for _, page := range s.Pages {
		if err := s.buildPage(page); err != nil {
			fmt.Fprintf(os.Stderr, "✗ %s: %v\n", page.Name, err)
			failed++
		} else {
			fmt.Printf("✓ %s → %s.html\n", page.Name, page.Name)
		}
	}

	total := len(s.Pages)
	duration := time.Since(start)

	if failed > 0 {
		fmt.Printf("%d page failed (of %d) in %s\n", failed, total, duration)
		return fmt.Errorf("build failed")
	}

	fmt.Printf("built %d pages in %s\n", total, duration)
	return nil
}

func (s *Site) buildPage(page Page) error {
	outputPath := filepath.Join(s.Dir, s.Output, page.Name+".html")

	html := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <title>%s</title>
  <meta name="description" content="%s">
</head>
<body>
%s
</body>
</html>`, page.Title, page.Description, page.ContentHTML)

	return os.WriteFile(outputPath, []byte(html), 0644)
}
