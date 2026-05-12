# netshoot

## Introduction

Deploy `nicolaka/netshoot` for in-cluster network diagnostics.

## Use cases

- Run a dedicated troubleshooting deployment
- Run a standalone pod for short-lived debugging sessions
- Pair with sidecar patch when debugging existing apps

## How to use

- Deployment base:

  ```sh
  kustomize build kustomize/netshoot/kustomization/base
  ```

- Pod overlay:

  ```sh
  kustomize build kustomize/netshoot/kustomization/overlays/pod
  ```
