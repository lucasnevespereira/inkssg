# inkssg

A small static site generator in Go. Built for sites with a couple of pages. Use as a library or CLI. No config needed for simple sites.

## Status

Early. v0.1 in progress.

## Vision

Sites with a couple of pages should not need 20 dependencies and a webpack config. inkssg takes a folder of content and produces a folder of HTML.

- Use as a Go library or CLI.
- Markdown for prose, raw HTML when you need full control.
- Frontmatter for per-page metadata. `ink.yaml` for site-wide data.
- Theme system: shared layout, per-theme styles.
- Single `assets/` directory. No per-page asset folders.
- No node, no plugins, no theme inheritance.

## Quick start (library)

```go
package main

import inkssg "github.com/snowztech/inkssg"

func main() {
    inkssg.Build()
}
```

```
go run main.go
```

## Quick start (CLI)

```
go install github.com/snowztech/inkssg/cmd/inkssg@latest
inkssg build
```

## Convention

```
my-site/
├── ink.yaml               # optional: site-wide data (name, links, bio)
├── pages/
│   ├── index/
│   │   └── content.md       # or content.html
│   └── about/
│       └── content.md
├── themes/
│   └── minimal/
│       ├── layout.html      # Go template, slot {{.Content}}
│       ├── styles.css       # optional
│       └── script.js        # optional
└── assets/                  # images, favicons, fonts
    ├── icons/
    └── img/
```

Build outputs to `public/`.

## Page metadata via frontmatter

```markdown
---
title: My Page
description: A short blurb
lang: en
theme: minimal
---

# Hello

Content here.
```

The same frontmatter works in `content.html`.

## Site-wide config via ink.yaml

```yaml
name: My Site
meta:
  lang: en
  description: Site description
links:
  - name: GitHub
    url: https://github.com/user
```

Use `ink.yaml` when you need site-wide data shared across pages. Not required for simple sites.

## Why another SSG

Hugo and Astro are great. inkssg is for people who want a small Go binary they can read in an afternoon. Built in public, learning-first.

## License

[MIT](LICENSE)
