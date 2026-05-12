# scenario: oauth-github

## Introduction

This scenario composes `httpbin` with `oauth2-proxy` ingress to demonstrate GitHub OAuth-protected ingress access.

## Use cases

- Validate ingress auth flow with oauth2-proxy
- Reuse as a starting point for protecting other internal services

## How to use

1. Update host/tls values in scenario patch files.
2. Build and apply:

```sh
kustomize build kustomize/_scenarios/oauth-github | kubectl apply -f -
```
