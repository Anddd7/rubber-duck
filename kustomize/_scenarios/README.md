# kustomization scenarios

## Introduction

Scenarios compose multiple app overlays into runnable end-to-end examples.

## Use cases

- Validate cross-component integration quickly
- Keep reproducible demo/testing environments

## How to use

```sh
kustomize build kustomize/_scenarios/<scenario> | kubectl apply -f -
```
