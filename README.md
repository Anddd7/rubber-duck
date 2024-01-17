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

- [ ] find a way to reduce the code of kustomization.yaml (or auto generate)
  - [ ] makefile, utilize the envsubs with kustomize
  - ...
- [x] tui engine for daily operations
  - see detials in[rubber-duck-cli](./rubber-duck-cli/README.md)
- [ ] curl as sidecar (kustomize patch)
- [ ] gitclone as sidecar
- [ ] devcontainers
- [x] composed wiki page (patten of readme)

### relavant tools

- TUI
  - charmbracelet / bubbletea
  - hairyhenderson / gomplate
  - ~~jedib0t / go-pretty~~
  
## Usages

- [kw: kubectl wrapper](./kw/README.md):
A simple wrapper for kubectl to make it easier to use.
- [Rubber Duck TUI](./rubber-duck-tui/README.md):
A reuseable TUI component for rubber duck.
- [zsh tools](./zsh/README.md):
Misc tools for zsh
- [excalidraw collabration](./excalidraw-room/README.md):
Deploy the excalidraw with sefl-hosted collabration(room) in your cluster.
- [Rubber Duck CLI](./rubber-duck-cli/README.md):
A simple CLI to help you debug your cluster.
- [tempfile](./tmpnb/README.md):
Create temp notebook(folder) and file
- [base64 shortcuts](./b64/README.md):
Some shortcuts for base64 encoding and decoding.
