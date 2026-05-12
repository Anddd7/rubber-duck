# excalidraw collabration

## Introduction

Deploy [Excalidraw](https://github.com/excalidraw/excalidraw) with collaboration services in Kubernetes.

## Use cases

- Team whiteboarding in private cluster environments
- Internal collaboration demos with custom ingress/domain

## How to use

- Base:

  ```sh
  kustomize build kustomize/excalidraw/kustomization/base
  ```

- Ingress overlay:

  ```sh
  kustomize build kustomize/excalidraw/kustomization/overlays/ingress
  ```

For custom domain/tls/url values, patch overlay configmap/ingresses from scenario or env overlay.

## Notes

The realtime collaboration feature is comming from community:

- Issue: <https://github.com/excalidraw/excalidraw/discussions/3879>

Solution:

- Blog: <https://blog.alswl.com/2022/10/self-hosted-excalidraw/>
- Code: <https://github.com/alswl/excalidraw-collaboration>
  - forked excalidraw: <https://github.com/alswl/excalidraw>
  - forked storage backend: <https://github.com/alswl/excalidraw-storage-backend>

### TODO: move to forked excalidraw, replace with self build image

to keep the source same, should build the image from code

Overwrite the configurations and compose it with kustomize

```yaml
namespace: rubber-duck
commonAnnotations:
  rubber-duck/scenarios: your-excalidraw

resources:
- ../../../excalidraw/kustomization/base
# overwrite ingress, e.g. hots, tls, annotation ...
- ingress-frontend.yaml   
- ingress-room.yaml
- ingress-storage.yaml

patches:
# overwrite host/url in configmap
- path: configmap.yaml
```
