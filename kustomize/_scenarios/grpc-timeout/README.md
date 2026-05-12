# scenario: grpc-timeout

## Introduction

This scenario is intended to validate gRPC timeout behavior through ingress by composing timeout-related variants.

## Use cases

- Compare ingress timeout and upstream gRPC timeout behavior
- Reproduce timeout edge cases for troubleshooting

## How to use

1. Add concrete resources/patches in `kustomization.yaml` for the timeout variants you want to test.
2. Build and apply:

```sh
kustomize build kustomize/_scenarios/grpc-timeout | kubectl apply -f -
```
