# Roadmap

## v0.1 — Usable for a real site

Goal: a user can `inkssg new` a site, `inkssg build` it, and deploy. lucasnp.dev runs on this.

### Build core
- [ ] `inkssg build` — read `inkssg.yaml`, build all pages, output to `public/`
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

### CLI surface
- [ ] `inkssg new <path>` — scaffold a site from `examples/minimal/` template
- [ ] `inkssg build`
- [ ] `inkssg version`

### Examples + CI
- [ ] `examples/minimal/` — 1 page, default theme
- [ ] `examples/multi-page/` — multiple pages, shared theme
- [ ] `examples/custom-theme/` — site with its own theme
- [ ] CI: build every example on push

## v0.2 — Make daily iteration painless

- [ ] `inkssg serve` — local server with file watcher and auto-rebuild
- [ ] Stable public library API: `site.New(opts...)`, `site.Build()`
- [ ] Options: `FromConfig`, `WithTheme`, `WithOutputDir`, `WithPagesDir`
- [ ] Hooks: `BeforeBuild`, `AfterBuild`
- [ ] `examples/library/` — programmatic usage from Go

## v0.3 — Polish + community

- [ ] `inkssg theme eject <name>` — copy a built-in theme into local `themes/`
- [ ] `inkssg theme list` — show available themes (built-in + local)
- [ ] "Did you mean" suggestions on errors (typos in theme names, frontmatter fields)
- [ ] `inkssg validate` — check config and frontmatter without building

## v0.4 — Production niceties

- [ ] CSS/JS minification (port from lucasnp.dev's existing build)
- [ ] Sitemap generation
- [ ] `inkssg build --strict` — fail on warnings (broken internal links, missing alt text)

## v0.5 — Theme sharing

- [ ] `inkssg theme add <git-url>` — clone a theme repo into `themes/`
- [ ] Optional `theme.yaml` manifest (name, version, author)

## Non-goals (forever, unless very strong demand)

- Hot module replacement
- React/JSX / typed templates
- Plugin system
- Theme inheritance
- i18n
- Image optimization pipeline
- Blog post collections / RSS / drafts

If you need those, use Astro, Next, or Hugo.
