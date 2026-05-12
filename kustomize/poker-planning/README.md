# poker-planning

## Introduction

Deploy [poker-planning](https://github.com/ModPhoenix/poker-planning) to Kubernetes for planning sessions.

## Use cases

- Lightweight team estimation sessions in cluster
- Internal workshop/demo environment

## How to use

- Base:

  ```sh
  kustomize build kustomize/poker-planning/kustomization/base
  ```

- Ingress overlay:

  ```sh
  kustomize build kustomize/poker-planning/kustomization/overlays/ingress
  ```

## Notes

It doesn't provide docker image directly, so I build it from source code.

- Forked Repo with Dockerfiles: <https://github.com/Anddd7/poker-planning>
- Docker Images:
  - [anddd9527/poker-planning:v1.0.0](https://hub.docker.com/repository/docker/anddd9527/poker-planning)
  - [anddd9527/poker-planning-server:v1.0.0](https://hub.docker.com/repository/docker/anddd9527/poker-planning-server)
