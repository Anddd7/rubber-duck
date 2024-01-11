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

## TODO

- [ ] makefile, utilize the envsubs with kustomize
- [ ] find a way to reduce the kustomization.yaml (or auto generate)
- [ ] tui engine for daily operations

### useful tools

- TUI
  - charmbracelet / bubbletea
  - hairyhenderson / gomplate
  - ~~jedib0t / go-pretty~~