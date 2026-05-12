# ls

## Introduction

Run a minimal pod for filesystem checks inside the cluster.

## Use cases

- Verify mounted volumes and file permissions
- Inspect runtime filesystem state quickly

## How to use

- Base pod:

  ```sh
  kustomize build kustomize/ls/kustomization/base
  ```

- PVC overlay:

  ```sh
  kustomize build kustomize/ls/kustomization/overlays/pvc
  ```
