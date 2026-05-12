# kustomization patches

## Introduction

Reusable Kustomize `Component` patches shared by multiple apps/overlays.

## Use cases

- Add debug sidecars without changing app base manifests
- Apply the same deployment mutation across multiple overlays

## How to use

```yaml
components:
- ../../_patches/<component>
```

Available components:

- `curl`: add `sidecar-curl` to Deployments
- `netshoot`: add `sidecar-netshoot` to Deployments
