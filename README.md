<div align="center">

# inkssg

***A small static site generator in Go. No node, no plugins, no config required.***

[![Version](https://img.shields.io/github/v/release/snowztech/inkssg?logo=github)](https://github.com/snowztech/inkssg/releases)
[![Stars](https://img.shields.io/github/stars/snowztech/inkssg?logo=github)](https://github.com/snowztech/inkssg/stargazers)
[![Issues](https://img.shields.io/github/issues/snowztech/inkssg?logo=github)](https://github.com/snowztech/inkssg/issues)
[![Contributors](https://img.shields.io/github/contributors/snowztech/inkssg?logo=github)](https://github.com/snowztech/inkssg/graphs/contributors)
[![License](https://img.shields.io/github/license/snowztech/inkssg)](LICENSE)

</div>

- Markdown for prose, raw HTML when you want full control
- Frontmatter for page metadata, `ink.yaml` for site-wide data
- Built-in themes you can override with your own
- Single `assets/` directory
- Works as a CLI or a Go library

## Install

```
go install github.com/snowztech/inkssg/cmd/inkssg@latest
```

Or grab a binary from [Releases](https://github.com/snowztech/inkssg/releases).

## Quick start

```
inkssg new my-site
inkssg build my-site
```

Output goes to `my-site/public/`. Open `index.html` to see it.

To change the page, edit `my-site/pages/index/content.md` and run `inkssg build my-site` again.

## Use as a library

```go
import "github.com/snowztech/inkssg"

inkssg.Build(".") // builds the site in the given directory
```

See [`examples/library`](examples/library) for a runnable version.

## Project layout

```
my-site/
├── ink.yaml             # optional: site-wide data
├── pages/
│   ├── index/
│   │   └── content.md   # or content.html
│   └── about/
│       └── content.md
├── themes/              # optional: your own themes
│   └── custom/
│       ├── layout.html  # Go template with {{.Content}}
│       ├── styles.css
│       └── script.js
└── assets/              # images, favicons, fonts
```

Build output goes to `public/`.

## Themes

inkssg ships with two themes: `minimal` and `devtool`. Pick one in `ink.yaml`:

```yaml
default_theme: devtool
```

To use your own theme, drop it in `themes/<name>/`. A local theme with the same name takes precedence over the built-in.

## Page metadata

Frontmatter goes at the top of `content.md`:

```markdown
---
title: My Page
description: A short blurb
---

# Hello

Content here.
```

Add `theme: devtool` to use a different theme on one page.

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

[Full docs](docs/inkssg.md) · [Examples](examples) · [Principles](PRINCIPLES.md) · [Roadmap](ROADMAP.md) · [MIT](LICENSE)
