# netshoot sidecar component

## Introduction

This component patches `apps/v1` Deployments by appending a `netshoot` sidecar container.

## Use cases

- Run network diagnostics (`dig`, `tcpdump`, `curl`, `traceroute`) next to your app
- Investigate DNS/routing/TLS issues without rebuilding app images

## How to use

In your overlay kustomization:

```yaml
components:
- ../../_patches/netshoot
```
