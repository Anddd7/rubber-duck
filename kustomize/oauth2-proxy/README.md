# oauth2-proxy

## Introduction

Deploy [oauth2-proxy](https://github.com/oauth2-proxy/oauth2-proxy) to protect services behind OAuth authentication.

## Use cases

- Add OAuth authn in front of internal HTTP services
- Reuse as a reference setup for ingress auth integration

## How to use

- Base deployment:

  ```sh
  kustomize build kustomize/oauth2-proxy/kustomization/base
  ```

- Ingress overlay:

  ```sh
  kustomize build kustomize/oauth2-proxy/kustomization/overlays/ingress
  ```

- End-to-end scenario:

  ```sh
  kustomize build kustomize/_scenarios/oauth-github
  ```

- Ingress overlay manifests: [kustomization/overlays/ingress](./kustomization/overlays/ingress)
