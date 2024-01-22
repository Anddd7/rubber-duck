# excalidraw collabration

Deploy excalidraw(<https://github.com/excalidraw/excalidraw>) to your cluster as communication tool.

## Description

The realtime collaboration feature is comming from community:

- Issue: <https://github.com/excalidraw/excalidraw/discussions/3879>

Solution:

- Blog: <https://blog.alswl.com/2022/10/self-hosted-excalidraw/>
- Code: <https://github.com/alswl/excalidraw-collaboration>
  - forked excalidraw: <https://github.com/alswl/excalidraw>
  - forked storage backend: <https://github.com/alswl/excalidraw-storage-backend>

### TODO: move to forked excalidraw, replace with self build image

to keep the source same, should build the image from code

## Usage

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
