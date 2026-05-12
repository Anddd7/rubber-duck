# netshoot sidecar component

## Introduction

Add a `netshoot` sidecar for fast network diagnostics inside app pods.

## Use cases

- Run network diagnostics (`dig`, `tcpdump`, `curl`, `traceroute`) next to your app
- Investigate DNS/routing/TLS issues without rebuilding app images

## How to use

In your overlay kustomization:

```yaml
components:
- ../../_patches/netshoot
```
