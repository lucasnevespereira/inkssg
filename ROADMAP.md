# Roadmap

## v0.1 — Library-first DX

Goal: a user can `go get` inkssg and have a site running in minutes. No config needed for simple sites.

### Library API (primary DX)
- [ ] `inkssg.Build()` — no args, defaults to current directory
- [ ] `inkssg.Build(dir)` — explicit path (same behavior as CLI)
- [ ] Auto-detect site structure (no `ink.yaml` required)
- [ ] Defaults: `minimal` theme, `public/` output, `pages/` dir
- [ ] Clear error if not a valid site: "No `pages/` folder found. Run `inkssg new .` to scaffold."

### CLI surface (wraps library)
- [ ] `inkssg build` — calls `inkssg.Build(".")`
- [ ] `inkssg new <path>` — scaffold from template
- [ ] `inkssg version`

### Build core
- [ ] Discover pages under `pages/<name>/`
- [ ] Parse YAML frontmatter from `content.md` and `content.html`
- [ ] Render markdown with goldmark (autoid for headings)
- [ ] Pass HTML body through unchanged for `content.html`
- [ ] Apply `themes/<name>/layout.html` via `html/template` with `{{.Site}}`, `{{.Page}}`, `{{.Theme}}`, `{{.Content}}`
- [ ] Output `<name>.html` at root of `public/`
- [ ] Build all pages, report all errors, exit non-zero on any failure
- [ ] Output: `✓ <page> → <file>` per page + summary line

### Themes
- [ ] Theme resolution: local `themes/<name>/` → built-in (embedded) → error
- [ ] Built-in `minimal` theme shipped via `go:embed`
- [ ] Theme requires only `layout.html`; styles/script optional
- [ ] Copy theme `styles.css`, `script.js` to `public/themes/<name>/`

### Assets
- [ ] Copy `assets/*` → `public/assets/` (one rule, no per-page folders)

### Examples + CI
- [ ] `examples/library/` — `go run main.go` builds the site
- [ ] `examples/minimal/` — 1 page, default theme
- [ ] `examples/multi-page/` — multiple pages, shared theme
- [ ] CI: build every example on push

Note: `ink.yaml` is optional in v0.1. Add it only when you need site-wide data (name, links, bio).

## v0.2 — Polish + daily iteration

- [ ] `inkssg serve` — local server with file watcher and auto-rebuild
- [ ] Options: `FromConfig`, `WithTheme`, `WithOutputDir`, `WithPagesDir`
- [ ] Hooks: `BeforeBuild`, `AfterBuild`
- [ ] CSS/JS minification
- [ ] Sitemap generation

## v0.3 — Community niceties

- [ ] `inkssg theme eject <name>` — copy built-in theme into local `themes/`
- [ ] `inkssg theme list` — show available themes
- [ ] "Did you mean" suggestions on typos
- [ ] `inkssg validate` — check config without building

## v0.4 — Production niceties

- [ ] `inkssg build --strict` — fail on warnings (broken links, missing alt text)

## v0.5 — Theme sharing

- [ ] `inkssg theme add <git-url>` — clone a theme repo into `themes/`

## Non-goals (forever, unless very strong demand)

- Hot module replacement
- React/JSX / typed templates
- Plugin system
- Theme inheritance
- i18n
- Image optimization pipeline
- Blog post collections / RSS / drafts

If you need those, use Astro, Next, or Hugo.
