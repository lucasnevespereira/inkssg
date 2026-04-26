# inkssg

A small static site generator in Go for sites with a few pages. No node, no plugins, no config required.

- Markdown for prose, raw HTML when you need full control
- Frontmatter for page metadata, `ink.yaml` for site-wide data
- Built-in themes, override with your own
- Single `assets/` directory
- Use as a CLI or a Go library

## Install

```
go install github.com/snowztech/inkssg/cmd/inkssg@latest
```

Or download a binary from [Releases](https://github.com/snowztech/inkssg/releases).

## Quick start

```
inkssg new my-site
inkssg build my-site
```

Output lands in `my-site/public/`. Open `index.html`.

## Use as a library

```go
import "github.com/snowztech/inkssg"

inkssg.Build(".")
```

Same behavior as the CLI. See [`examples/library`](examples/library) for a runnable version.

## Project layout

```
my-site/
├── ink.yaml             # optional: site-wide data (name, links, meta)
├── pages/
│   ├── index/
│   │   └── content.md   # or content.html
│   └── about/
│       └── content.md
├── themes/              # optional: override built-in themes
│   └── custom/
│       ├── layout.html  # Go template with {{.Content}}
│       ├── styles.css
│       └── script.js
└── assets/              # images, favicons, fonts
```

Build output goes to `public/`.

## Themes

Ships with `minimal` and `devtool` embedded — no setup needed. Pick one in `ink.yaml`:

```yaml
default_theme: devtool
```

Drop a folder in `themes/<name>/` to add your own. Local themes win over built-ins.

## Page metadata

Frontmatter at the top of `content.md`:

```markdown
---
title: My Page
description: A short blurb
---

# Hello

Content here.
```

Add `theme: devtool` to override the site theme for one page.

## Site config

`ink.yaml` holds data shared across pages. Optional.

```yaml
name: My Site
meta:
  lang: en
  description: Site description
links:
  - name: GitHub
    url: https://github.com/user
```

## More

[Full docs](docs/inkssg.md) · [Principles](PRINCIPLES.md) · [Roadmap](ROADMAP.md) · [MIT](LICENSE)
