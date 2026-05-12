# Grpcbin

## Introduction

Deploy grpcbin for gRPC connectivity and timeout behavior testing.

## Use cases

- Validate gRPC ingress routing
- Reproduce timeout and TLS behavior

## How to use

Build one of the manifests:

```sh
kustomize build kustomize/grpcbin/kustomization/base
kustomize build kustomize/grpcbin/kustomization/overlays/ingress
kustomize build kustomize/grpcbin/kustomization/overlays/tls
```

Then test through ingress-nginx:

- Install ingress-nginx
- Deploy grpcbin ingress overlay
- Port-forward ingress-nginx-controller to localhost

```sh
kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 8080:80
```

- Test grpcbin with host

```sh
grpcbin unary --message hello --host=grpcbin.example.com --port=8080 
```

### Reference

[More usage of grpcbin ...](https://github.com/Anddd7/grpcbin/blob/main/README.md#usage)
