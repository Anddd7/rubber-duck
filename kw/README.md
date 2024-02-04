# kw: kubectl wrapper

A simple wrapper for kubectl to make it easier to use.

## Overview

quick export yaml file

## Index

* [kw](#kw)

### kw

provide some shortcuts for your frequently used kubectl command, support kubectl_completion (for zsh)

#### Example

```bash
  # export resource to a temp yaml file
  kw get pod your-pod -oy
  > kubectl get pod your-pod -o yaml > pod_your-pod_112422.yaml

@arg $@ kubectl command
```

```bash
  # exec into container
  kw -x your-pod 
  > kubectl exec -it your-pod -- bash
```