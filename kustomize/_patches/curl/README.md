# curl sidecar component

## Introduction

Add a `curl` sidecar for quick in-pod HTTP/API connectivity checks.

## Use cases

- Debug service-to-service connectivity from the same Pod network namespace
- Run quick HTTP checks without modifying the main container image

## How to use

In your overlay kustomization:

```yaml
components:
- ../../_patches/curl
```
