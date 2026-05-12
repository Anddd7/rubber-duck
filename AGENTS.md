# AGENTS.md

## Repo reality (do not guess)

- This repo is both:
  - a Go CLI (`cmd/`) for operating kustomize assets
  - a marketplace-style collection of operational assets (`kustomize/`, `kubectl-plugins/`, `scripts/`, `apps/`)
- Do not scope changes to CLI only unless explicitly requested.
- CLI entrypoint: `cmd/rubber-duck/main.go`; command wiring: `cmd/internal/root.go`.
- CLI command groups:
  - `kust` → normal components under `kustomize/<name>/kustomization/{base,overlays/...}`
  - `kust-patch` → patches under `kustomize/_patches/<name>`
  - `kust-scenario` → scenarios under `kustomize/_scenarios/<name>`
- Component metadata shown by CLI is read from each target `README.md` first non-header line.

## Asset domains beyond CLI

- `kustomize/`: reusable templates for debug/test/verify workflows and self-deploy compositions.
- `kubectl-plugins/`: kubectl plugin scripts; install with `make -C kubectl-plugins install` (copies `kubectl*` to `~/bin`).
- `scripts/`: shell utility snippets (`_base64`, `_gitcz`, `_tmpnb`, `_nocolor`) for sourcing/reuse.
- `apps/chatbot9527/`: local Docker-compose app stack (intentionally outside kustomize).

## Build / install commands

- Build CLI binary: `make build` (outputs `bin/rubber-duck`)
- Install CLI binary: `make install` (copies to `~/bin/rubber-duck`)
- Uninstall CLI binary: `make uninstall`
- Direct CLI build check: `go build ./cmd/rubber-duck`

## Fast verification commands (preferred)

- Show CLI command tree: `go run ./cmd/rubber-duck --help`
- Normal components: `go run ./cmd/rubber-duck kust list`
- Patch components: `go run ./cmd/rubber-duck kust-patch list`
- Scenario components: `go run ./cmd/rubber-duck kust-scenario list`

## Operational gotchas

- CLI `install/uninstall` commands shell out to `kubectl`; use `--dry-run` for safe checks.
- `kust install/uninstall` requires `--overlay` (`base` or overlay name).
- `kust-patch install` requires `--pod`; uninstall requires one of `--pod` or `--deploy`.
- `kust-patch uninstall` uses `kubectl rollout undo deployment ...` (rollback semantics, not restart).
- Some scenarios are intentionally incomplete for `kubectl apply -k` (can return `no objects passed to apply`).
- `apps/chatbot9527` is Docker-compose based; do not force it into `kustomize/` layout.

## Boundaries to keep

- Keep CLI code under `cmd/`.
- Keep marketplace assets in current domain dirs (`kustomize/`, `kubectl-plugins/`, `scripts/`, `apps/`) unless migration is explicitly requested.
- Do not add non-standard top-level fields to `kustomization.yaml` for metadata; use README-based descriptions.

## README description rule (important)

- For every kustomize component/patch/scenario, CLI description comes from `README.md` first non-header line.
- That line must be goal-oriented and concise (what problem it solves), not implementation detail.
- When adding/updating any kustomize asset, update this README description line in the same change.
