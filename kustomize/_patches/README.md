# kustomization patches

A bunch of common patches (as component), e.g. sidecars, annotations

## Usage

```yaml
components:
- ../../_patches/<component>
```

## Components

- `curl`: add `sidecar-curl` to Deployments
- `netshoot`: add `sidecar-netshoot` to Deployments
