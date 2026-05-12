# httpbin

## Introduction

Deploy [httpbin](https://httpbin.org/) for HTTP connectivity and behavior testing.

## Use cases

- Verify service/network routing in cluster
- Reproduce client/server request behavior quickly

## How to use

- Base only:

  ```sh
  kustomize build kustomize/httpbin/kustomization/base
  ```

- With ingress overlay:

  ```sh
  kustomize build kustomize/httpbin/kustomization/overlays/ingress
  ```
