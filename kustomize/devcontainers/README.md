# Dev Container in Kubernetes

Install a dev container inside your cluster, develop and test your application in a real environment.

## Description

Devcontainers is a remote development environment, you can use vscode to connect to a remote server, and use the server's resources to develop your code. When you run it in a kubernetes cluster, you can use the cluster's resources(compute, network, storage) and develop a true 'cloud-native' application.

VSCode (client) will interact with the dev container(pod) in the cluster, via kubectl and kube-proxy. So you don't need to open any port for the deployment.

## Usage

- install required extensions for vscode
  - Dev Container
  - Kubernetes
- deploy a dev container
- set kubeconfig
- select the pod and `Attach Visual Studio Code``

![Alt text](docs/attach-pod.png)

### For WSL user

:rotating_light: **Attach to Container** is not working in wsl, see [here](https://code.visualstudio.com/docs/devcontainers/containers#_docker-extension-limitations).

So, you have to install the tools (kubectl, vscode extensions) in your windows machine.

1. Install kubectl.exe and add into `PATH`
2. Set KUBECONFIG into env

  ```powershell
  $kubeconfig = "C:\Users\your\local\folder\config.yaml"

  [System.Environment]::SetEnvironmentVariable("KUBECONFIG", $kubeconfig, "User")

  ```

3. follow the steps [above](#usage)

### For storage(cache)

- [ ] workspace cache, for whole user path
- [ ] vscode cache, for extensions
