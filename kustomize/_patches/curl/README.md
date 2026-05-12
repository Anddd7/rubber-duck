# curl sidecar component

## Introduction

This component patches `apps/v1` Deployments by appending a `curl` sidecar container.

## Use cases

- Debug service-to-service connectivity from the same Pod network namespace
- Run quick HTTP checks without modifying the main container image

## How to use

In your overlay kustomization:

```yaml
components:
- ../../_patches/curl
```
