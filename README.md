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
- [x] devcontainers
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

- [oauth2-proxy](./oauth2-proxy/README.md):
Deploy oauth2-proxy(<https://github.com/oauth2-proxy/oauth2-proxy>) to setup oauth2 authentication for you application.
- [Grpcbin](./grpcbin/README.md):
Similar with httpbin(<https://httpbin.org/>) to test the grpc connection.
- [Kubectl Plugins](./kubectl-plugins/README.md):
Some simple plugins to extend kubectl functionality.
- [Chatbot 9527](./chatbot9527/README.md):
Deploy a chatbot for personal usage, it should be able to host locally
- [httpbin](./httpbin/README.md):
Use httpbin(<https://httpbin.org/>) to test the connection to your pod.
- [zsh tools](./zsh/README.md):
Misc tools for zsh
- [tempfile](./tmpnb/README.md):
Create temp notebook(folder) and file
- [kustomization patches](./kustomization-patches/README.md):
A bunch of common patches (as component), e.g. sidecars, annotations
- [nginx](./nginx/README.md):
A simple echo service to test the connection to your pod.
- [base64 shortcuts](./b64/README.md):
Some shortcuts for base64 encoding and decoding.
- [Rubber Duck CLI](./rubber-duck-cli/README.md):
A simple CLI to help you debug your cluster.
- [ls](./ls/README.md):
Use ls to exlore the file system in pod, e.g. verify the mounted volume
- [Dev Container in Kubernetes](./devcontainers/README.md):
Install a dev container inside your cluster, develop and test your application in a real environment.
- [Rubber Duck TUI](./rubber-duck-tui/README.md):
A reuseable TUI component for rubber duck.
- [poker-planning](./poker-planning/README.md):
Deploy poker-planning (<https://github.com/ModPhoenix/poker-planning>) to your cluster as communication tool.
- [excalidraw collabration](./excalidraw/README.md):
Deploy excalidraw(<https://github.com/excalidraw/excalidraw>) to your cluster as communication tool.
