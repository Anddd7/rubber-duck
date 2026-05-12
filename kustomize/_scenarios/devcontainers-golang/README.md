# scenario: devcontainers-golang

## Introduction

This scenario deploys a Golang dev container environment in Kubernetes.

## Use cases

- Cloud-native development workspace in cluster
- Quick onboarding for remote coding with VSCode attach

## How to use

```sh
kustomize build kustomize/_scenarios/devcontainers-golang | kubectl apply -f -
```
