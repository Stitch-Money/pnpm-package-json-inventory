# pnpm-package-json-inventory
Generate a basic package-lock.json package inventory from a pnpm project

## Install
1. Configure git to use ssh instead of https: `git config --global --add url."ssh://git@github.com/".insteadOf "https://github.com/"`
2. Ensure your ssh key is unlocked in the ssh agent: `ssh-add`
3. Have Go installed: `brew install go`
4. Install the utility: `GOPRIVATE=github.com/Stitch-Money/pnpm-package-json-inventory go install github.com/Stitch-Money/pnpm-package-json-inventory@latest`
5. Use as needed (you might need to add `~/go/bin` to your PATH): `pnpm-package-json-inventory`