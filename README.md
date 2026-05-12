# rubber-duck

`rubber-duck` is a DevOps operations marketplace plus a CLI.

- **Marketplace assets**: reusable `kustomize/` templates (components, patches, scenarios), kubectl plugins, shell helper scripts, and local app stacks.
- **CLI**: `rubber-duck` command to list/get/install/uninstall those assets in Kubernetes.

## Repository structure

- `cmd/`: Go CLI source code
- `kustomize/`: reusable ops templates
  - `<component>/kustomization/{base,overlays/...}`
  - `_patches/<patch>`
  - `_scenarios/<scenario>`
- `kubectl-plugins/`: kubectl plugin scripts
- `scripts/`: shell helper snippets for sourcing/reuse
- `apps/`: local app stacks (for example Docker Compose)

## Build and install CLI

```bash
make build      # bin/rubber-duck
make install    # ~/bin/rubber-duck
make uninstall
```

Direct build check:

```bash
go build ./cmd/rubber-duck
```

Install via `go install` (recommended for development):

```bash
# from the repository root
go install ./cmd/rubber-duck
# this will install the binary into $GOBIN or $GOPATH/bin
```

If you prefer to install a specific version from the repository remote once a tag is published:

```bash
go install github.com/Anddd7/rubber-duck/cmd/rubber-duck@latest
```

## CLI quick usage

```bash
go run ./cmd/rubber-duck --help
go run ./cmd/rubber-duck kust --help
go run ./cmd/rubber-duck kust-patch --help
go run ./cmd/rubber-duck kust-scenario --help
```

### List assets

```bash
go run ./cmd/rubber-duck kust list
go run ./cmd/rubber-duck kust-patch list
go run ./cmd/rubber-duck kust-scenario list
```

### Install/uninstall component

```bash
go run ./cmd/rubber-duck kust install httpbin --overlay base -n default
go run ./cmd/rubber-duck kust uninstall httpbin --overlay base -n default
```

### Install/uninstall patch

```bash
go run ./cmd/rubber-duck kust-patch install curl --pod <pod-name> -n default
go run ./cmd/rubber-duck kust-patch uninstall curl --deploy <deployment-name> -n default
```

### Install/uninstall scenario

```bash
go run ./cmd/rubber-duck kust-scenario install oauth-github -n default
go run ./cmd/rubber-duck kust-scenario uninstall oauth-github -n default
```

## End-to-end example (verified)

This is the exact flow validated in a real cluster.

```bash
# 1) choose kubeconfig
export KUBECONFIG=~/.kube/config_xcrcli_aws-usea1-sc-qa07

# 2) deploy component
go run ./cmd/rubber-duck kust install httpbin --overlay base -n default

# 3) verify component
kubectl get deploy/httpbin svc/httpbin -n default

# 4) install curl patch to one httpbin pod
POD=$(kubectl get pod -n default | awk '/^httpbin-/{print $1; exit}')
go run ./cmd/rubber-duck kust-patch install curl --pod "$POD" -n default

# 5) verify sidecar added
kubectl get deploy httpbin -n default -o jsonpath='{.spec.template.spec.containers[*].name}{"\n"}'

# 6) run curl inside injected sidecar
POD=$(kubectl get pod -n default | awk '/^httpbin-/{print $1; exit}')
kubectl exec -n default "$POD" -c sidecar-curl -- curl -I -sS https://google.com

# 7) rollback patch
go run ./cmd/rubber-duck kust-patch uninstall curl --deploy httpbin -n default

# 8) verify sidecar removed
kubectl get deploy httpbin -n default -o jsonpath='{.spec.template.spec.containers[*].name}{"\n"}'

# 9) cleanup component
go run ./cmd/rubber-duck kust uninstall httpbin --overlay base -n default
```

## Notes

- `kust install/uninstall` requires `--overlay` (`base` or an overlay name).
- `kust-patch install` requires `--pod` and patches the owning deployment.
- `kust-patch uninstall` performs `rollout undo` (rollback), not restart.
- Asset descriptions in CLI are read from each asset `README.md` first non-header line.

CI

This repository includes GitHub Actions workflows to validate changes automatically:

- .github/workflows/go.yml — runs `go vet`, `go test`, and `golangci-lint` on push/PR
- .github/workflows/kustomize.yml — builds each `kustomization` to validate templates
