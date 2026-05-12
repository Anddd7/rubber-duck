# nginx

## Introduction

Deploy nginx as a lightweight web endpoint for connectivity and config tests.

## Use cases

- Validate ingress routing quickly
- Test custom nginx config and static content mounting

## How to use

- Base:

  ```sh
  kustomize build kustomize/nginx/kustomization/base
  ```

- Ingress overlay:

  ```sh
  kustomize build kustomize/nginx/kustomization/overlays/ingress
  ```

- External config overlay:

  ```sh
  kustomize build kustomize/nginx/kustomization/overlays/extconf
  ```
