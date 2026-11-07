# DSA Pavilion

Docusaurus site for [Data Structures and Algorithms in Go](https://github.com/ghrushneshr25/dsa), plus the **docgen** tool that builds these pages from annotated Go source.

Live site: https://ghrushneshr25.github.io/dsa-pavilion/

## Layout

```
dsa-pavilion/          this repo
  docgen/              Go generator
  docs/                generated MDX (do not edit by hand)
  src/                 theme, CSS, GitHub route map
  sidebars.js          generated
dsa/                   sibling repo — source of truth
```

Clone them as siblings:

```text
some-folder/
  dsa/
  dsa-pavilion/
```

## Generate docs

Requires Go 1.24+ and the `dsa` repo at `../dsa`.

```bash
yarn generate
# or
go -C docgen run .
./generate.sh
```

This wipes and rewrites `docs/`, `sidebars.js`, `src/data/dsaGithubRoutes.json`, and `../dsa/readme.md`.

Readme only (no site files):

```bash
yarn generate:readme
```

From the `dsa` repo: `make readme` (looks for `../dsa-pavilion/docgen`).

## Local development

```bash
yarn
yarn generate
yarn start
```

## Build / deploy

```bash
yarn build
```

Pushes to `master` deploy GitHub Pages via `.github/workflows/deploy.yml`. Docs themselves are regenerated from the `dsa` repo by `.github/workflows/regenerate-docs.yml`.
