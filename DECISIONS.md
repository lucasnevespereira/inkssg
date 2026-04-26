# Decisions

Choices made early to avoid bikeshedding later. The bias is always: simpler over more powerful.

## Assets — single directory

**Decision:** one `assets/` directory at the site root. No per-page asset folders.

```
my-site/
├── pages/
├── themes/
└── assets/                 ← all assets here
    ├── icons/
    ├── img/
    └── snowz/
        └── vikusha.png
```

`assets/*` → `public/assets/*`. One copy rule.

**Why:** one mental model. "Asset URL = `/assets/<path>`". No choice between site-wide vs per-page. Less code to write, less doc to read. For sites under 20 pages (the target audience), per-page co-location is unnecessary; grouping by sub-folder (`assets/snowz/`) is enough.

## Markdown images

**Decision:** absolute paths only. No relative path rewriting.

User writes:
```markdown
![avatar](/assets/snowz/vikusha.png)
```

**Why:** zero magic. User learns the convention once: "image URL = `/assets/<path>`". No surprise rewriting.

## Syntax highlighting

**Decision:** none in v0.1.

Code blocks render as plain `<pre><code class="language-go">`. User can wire prism.js / highlight.js in their theme's `script.js` if they want highlighting.

**Why:** chroma dep adds ~5MB to the binary for a feature most landing-page sites don't need. Easy to add later if demand exists.

## Heading IDs

**Decision:** automatic via goldmark's autoid extension.

`## My Section` → `<h2 id="my-section">My Section</h2>`. Anchor links `#my-section` work.

**Why:** standard SSG behavior, zero config, one line to enable.

## External links

**Decision:** no auto-rewriting.

If a user wants `target="_blank"`, they write it explicitly:
```markdown
[Google](https://google.com)        <!-- opens in same tab -->
<a href="https://google.com" target="_blank" rel="noopener">Google</a>
```

**Why:** the "auto open external in new tab" UX is debatable (some users hate it). Don't impose. User has full control via inline HTML.

## Drafts

**Decision:** no `draft: true` frontmatter field.

If a page is a draft, user removes it from `pages/` (git stash, branch, comment out). Filesystem is the source of truth.

**Why:** drafts are a blogging feature; inkssg targets simple sites, not blogs. If the user needs drafts, they probably need a blog → use Hugo. Avoids feature creep.

## Build error strategy

**Decision:** build all pages, report all errors at the end, exit non-zero.

Don't stop on first error. The user sees every problem in one run.

**Why:** better DX than fail-fast (no "fix, rebuild, find next error" loop). Still CI-friendly because exit code != 0 on any failure.

## Build output format

**Decision:** one line per page + summary.

```
✓ index → index.html
✓ snowz → snowz.html
built 2 pages in 156ms
```

On error:
```
✓ index → index.html
✗ snowz: missing content.md or content.html
1 page failed (of 2) in 134ms
```

Stderr for errors, stdout for the summary line.

**Why:** scannable, no spam. Matches modern tools (esbuild, vite). Easy to grep.

## Default output directory

**Decision:** `public/`.

**Why:** Hugo, Next.js, GH Pages all default to or accept `public/`. Most universally compatible.

## Default theme name

**Decision:** `minimal`.

The built-in theme shipped via `go:embed` is named `minimal`. If `inkssg.yaml` doesn't specify `default_theme`, falls back to `minimal`.

**Why:** clearly communicates "this is a basic theme, customize or replace it".

## Config file name

**Decision:** `inkssg.yaml`.

**Why:** explicit, tool-named-file convention (`astro.config.mjs`, `next.config.js`). Searchable. Not confused with other yaml files in the project.

## Frontmatter format

**Decision:** YAML between `---` delimiters at the top of `content.md` and `content.html`.

```markdown
---
title: My Page
theme: linkinbio
---

content here
```

**Why:** standard across Hugo / Jekyll / Astro / 11ty. Users coming from any SSG already know it.

## Theme: required vs optional files

**Decision:** only `layout.html` is required. Everything else (`styles.css`, `script.js`, `theme.yaml`, `partials/`) is optional.

**Why:** lowest barrier to creating a custom theme. A user can write a 10-line `layout.html` and call it a theme.

## Theme inheritance

**Decision:** none. No theme can extend another.

If you want to base a theme on another, eject it (`inkssg theme eject minimal`) and modify the copy.

**Why:** Hugo theme inheritance is powerful but a debugging nightmare. Eject + modify is explicit and traceable.
