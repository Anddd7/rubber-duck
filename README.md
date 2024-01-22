# rubber-duck

:monocle_face: A rubber duck in your cluster, help to debug everything

## Swiss Army Knife or Doraemon Pocke**

To debug a kubernetes cluster / component(middleware) / pod(business), it not similar to debug your local machine, you need to take care of the network, permission, and so on. This project is a collection of tools for k8s debugging, And classify and organize them according to **scenarios**, so hope they can be used right out-of-box.

## Structure

It contains tools - e.g. oauth2-proxy, nginx, httpbin, with its artifacts, source-code, dockerfile and kustomization.

They, combined with kustomize in different scenarios, can be used to debug different problems.

```md
.
├── component1
│   ├── src
│   └── kustomization
├── component2
│   └── kustomization
├── scenario
│   ├── scenario1
│   └── scenario2
├── README.md
└── LICENSE
```

## Features

- [x] ~~find a way to reduce the code of kustomization.yaml (or auto generate)~~
- [x] ~~makefile, utilize the envsubs with kustomize~~
  - keep it simple stupid, declaritive first
- [x] tui engine for daily operations
  - see detials in[rubber-duck-cli](./rubber-duck-cli/README.md)
- [x] curl as sidecar (kustomize patch)
- [x] ~~gitclone as sidecar~~ 
  - ❗not safe
- [ ] devcontainers
- [x] composed wiki page (patten of readme)
- [ ] prod ready: add probe, image pull policy, resource limits ...

### relavant tools

- TUI
  - charmbracelet / bubbletea
  - hairyhenderson / gomplate
  - ~~jedib0t / go-pretty~~
- K8S
  - kustomize
- shell
  - shdoc
  
## Usages

- [kw: kubectl wrapper](./kw/README.md):
A simple wrapper for kubectl to make it easier to use.
- [poker-planning](./poker-planning/README.md):
Deploy poker-planning (<https://github.com/ModPhoenix/poker-planning>) to your cluster as communication tool.
- [Rubber Duck TUI](./rubber-duck-tui/README.md):
A reuseable TUI component for rubber duck.
- [kustomization patches](./kustomization-patches/README.md):
A bunch of common patches (as component), e.g. sidecars, annotations
- [nginx](./nginx/README.md):
A simple echo service to test the connection to your pod.
- [oauth2-proxy](./oauth2-proxy/README.md):
Deploy oauth2-proxy(<https://github.com/oauth2-proxy/oauth2-proxy>) to setup oauth2 authentication for you application.
- [zsh tools](./zsh/README.md):
Misc tools for zsh
- [httpbin](./httpbin/README.md):
Use httpbin(<https://httpbin.org/>) to test the connection to your pod.
- [excalidraw collabration](./excalidraw/README.md):
Deploy excalidraw(<https://github.com/excalidraw/excalidraw>) to your cluster as communication tool.
- [Rubber Duck CLI](./rubber-duck-cli/README.md):
A simple CLI to help you debug your cluster.
- [tempfile](./tmpnb/README.md):
Create temp notebook(folder) and file
- [base64 shortcuts](./b64/README.md):
Some shortcuts for base64 encoding and decoding.
