# docgen

Go generator that turns annotated source in the sibling [`dsa`](https://github.com/ghrushneshr25/dsa) repo into this site’s MDX, sidebar, and GitHub source links.

Run from the **pavilion repo root** (not this directory):

```bash
go -C docgen run .
# or
yarn generate
./generate.sh
```

Paths are resolved from the Docusaurus root. Override with `DOCGEN_CODE`, `DOCGEN_DOCS`, `DOCGEN_README`, and related env vars if `dsa` is not at `../dsa`.
