# inkssg

A small static site generator in Go. Built for sites with one or two pages. One command, sane defaults, no surprises.

## Status

Early. v0.1 in progress.

## Vision

Sites with a couple of pages should not need 20 dependencies and a webpack config. inkssg takes a folder of content and produces a folder of HTML.

- Markdown for prose, raw HTML when you need full control.
- Frontmatter for per-page metadata.
- Theme system: shared layout, per-theme styles.
- Single `assets/` directory. No per-page asset folders.
- Single binary. No node, no plugins, no theme inheritance.

## Quickstart (planned)

```
inkssg new my-site
cd my-site
inkssg build
inkssg serve
```

## Convention

```
my-site/
├── inkssg.yaml              # site config: name, default theme, links
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
└── assets/                  # everything else (images, favicons, fonts)
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

## Why another SSG

Hugo and Astro are great. inkssg is for people who want a small Go binary they can read in an afternoon. Built in public, learning-first.

## License

[MIT](LICENSE)
