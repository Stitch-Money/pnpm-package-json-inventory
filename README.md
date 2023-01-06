# pnpm-package-json-inventory
Generate a basic package-lock.json package inventory from a pnpm project

## Explanation
This tool uses the pnpm cli to output a list of all packages, then parses and converts that list into a package-lock.json equivalent. The conversion isn't perfect, as some data doesn't exist in the pnpm output that does in the package-lock.json, but it serves its purpose.

The output is useful for dependency scanning tools which don't have pnpm support, such as Snyk or Semgrep Supply Chain.

## Install
`go install github.com/Stitch-Money/pnpm-package-json-inventory@latest`

A linux amd64 build can be retrieved from Releases, which will always reflect the Main branch.

## Why is it slow?
`pnpm list` outputs an unabridged list of every dependency, including duplicates. So if you have 3 packages that each depend on a package with 500 dependencies, your output will have 1500 extra dependencies. This snowballs out of proportion fast, and the only solution is to parse them all. Don't worry, this tool de-dupes them before outputting the package-lock.json, but still needs to parse all of the JSON from pnpm.