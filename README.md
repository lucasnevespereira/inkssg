# inkssg

A small static site generator in Go. Built for sites with a couple of pages.
Sites with a couple of pages should not need 20 dependencies and a webpack config.

Use as a library or CLI. No config needed for simple sites.

Features:
- Markdown for prose, raw HTML when you need full control
- Frontmatter for page metadata, ink.yaml for site-wide data
- Theme system with shared layout
- Single assets/ directory
- No node, no plugins

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
theme: minimal
---

# Hello

Content here.
```

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

## Documentation

- [docs/inkssg.md](docs/inkssg.md) - full user docs
- [PRINCIPLES.md](PRINCIPLES.md) - design principles
- [ROADMAP.md](ROADMAP.md) - what's next

## License

[MIT](LICENSE)
