package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	defaultNamespace = "default"
	defaultDebugPod  = "rubber-duck-netshoot"
)

type AppCmds struct {
	Connect AppConnectCmd `cmd:""`
	Sh      AppShCmd      `cmd:""`
}

type DebugCmds struct {
	Connect DebugConnectCmd `cmd:""`
	Sh      DebugShCmd      `cmd:""`
}

type AppConnectCmd struct {
	Svc        string `help:"k8s service name" required:""`
	Namespace  string `help:"namespace" default:"default"`
	LocalPort  int    `help:"local port" default:"8080"`
	RemotePort int    `help:"remote service port" default:"80"`
}

func (cmd *AppConnectCmd) Run(g *GlobalSettings) error {
	st, err := loadState()
	if err != nil {
		return err
	}

	st.App = &ConnectionState{
		Namespace:  cmd.Namespace,
		Name:       cmd.Svc,
		LocalPort:  cmd.LocalPort,
		RemotePort: cmd.RemotePort,
	}
	if err := saveState(st); err != nil {
		return err
	}

	fmt.Printf("Current app connection: ns=%s svc=%s\n", cmd.Namespace, cmd.Svc)
	fmt.Printf("Local URL: http://127.0.0.1:%d\n", cmd.LocalPort)

	return streamKubectl(
		"-n", cmd.Namespace,
		"port-forward", fmt.Sprintf("svc/%s", cmd.Svc),
		fmt.Sprintf("%d:%d", cmd.LocalPort, cmd.RemotePort),
	)
}

type AppShCmd struct {
	Pod       string `help:"target pod name"`
	Namespace string `help:"namespace (defaults to connected app namespace)"`
	TTY       bool   `help:"allocate terminal"`
	Shell     string `help:"shell path" default:"/bin/bash"`
}

func (cmd *AppShCmd) Run(g *GlobalSettings) error {
	st, err := loadState()
	if err != nil {
		return err
	}

	ns := cmd.Namespace
	if ns == "" {
		ns = defaultNamespace
		if st.App != nil && st.App.Namespace != "" {
			ns = st.App.Namespace
		}
	}

	pod := cmd.Pod
	if pod == "" {
		if st.App == nil || st.App.Name == "" {
			return fmt.Errorf("pod is required, or run app connect first")
		}
		pod, err = getServiceTargetPod(ns, st.App.Name)
		if err != nil {
			return err
		}
	}

	args := []string{"-n", ns, "exec", "-i"}
	if cmd.TTY {
		args = append(args, "-t")
	}
	args = append(args, pod, "--", cmd.Shell)

	return streamKubectl(args...)
}

type DebugConnectCmd struct {
	Namespace  string `help:"namespace" default:"default"`
	PodName    string `help:"debug pod name" default:"rubber-duck-netshoot"`
	LocalPort  int    `help:"optional local port (with remote port)"`
	RemotePort int    `help:"optional remote port (with local port)"`
}

func (cmd *DebugConnectCmd) Run(g *GlobalSettings) error {
	if err := ensureDebugPod(cmd.Namespace, cmd.PodName); err != nil {
		return err
	}

	st, err := loadState()
	if err != nil {
		return err
	}
	st.Debug = &ConnectionState{
		Namespace:  cmd.Namespace,
		Name:       cmd.PodName,
		LocalPort:  cmd.LocalPort,
		RemotePort: cmd.RemotePort,
	}
	if err := saveState(st); err != nil {
		return err
	}

	if cmd.LocalPort > 0 && cmd.RemotePort > 0 {
		fmt.Printf("Local URL: http://127.0.0.1:%d\n", cmd.LocalPort)
		return streamKubectl(
			"-n", cmd.Namespace,
			"port-forward", fmt.Sprintf("pod/%s", cmd.PodName),
			fmt.Sprintf("%d:%d", cmd.LocalPort, cmd.RemotePort),
		)
	}

	fmt.Printf("Debug pod ready: %s/%s\n", cmd.Namespace, cmd.PodName)
	return nil
}

type DebugShCmd struct {
	Namespace string `help:"namespace (defaults to connected debug namespace)"`
	Pod       string `help:"debug pod name (defaults to connected debug pod)"`
	TTY       bool   `help:"allocate terminal"`
	Cmd       string `help:"non-interactive command" default:"echo attached; uname -a"`
	Shell     string `help:"shell path" default:"/bin/bash"`
}

func (cmd *DebugShCmd) Run(g *GlobalSettings) error {
	st, err := loadState()
	if err != nil {
		return err
	}

	ns := cmd.Namespace
	if ns == "" {
		ns = defaultNamespace
		if st.Debug != nil && st.Debug.Namespace != "" {
			ns = st.Debug.Namespace
		}
	}

	pod := cmd.Pod
	if pod == "" {
		pod = defaultDebugPod
		if st.Debug != nil && st.Debug.Name != "" {
			pod = st.Debug.Name
		}
	}

	if err := ensureDebugPod(ns, pod); err != nil {
		return err
	}

	if cmd.TTY {
		return streamKubectl("-n", ns, "exec", "-it", pod, "--", cmd.Shell)
	}

	return streamKubectl("-n", ns, "exec", pod, "--", "bash", "-lc", cmd.Cmd)
}

type State struct {
	App   *ConnectionState `json:"app,omitempty"`
	Debug *ConnectionState `json:"debug,omitempty"`
}

type ConnectionState struct {
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
	LocalPort  int    `json:"localPort,omitempty"`
	RemotePort int    `json:"remotePort,omitempty"`
}

func stateFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "rubber-duck", "state.json"), nil
}

func loadState() (*State, error) {
	p, err := stateFilePath()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return &State{}, nil
		}
		return nil, err
	}

	st := &State{}
	if err := json.Unmarshal(b, st); err != nil {
		return &State{}, nil
	}
	return st, nil
}

func saveState(st *State) error {
	p, err := stateFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o644)
}

func ensureDebugPod(namespace, pod string) error {
	if err := runKubectl("-n", namespace, "get", "pod", pod); err != nil {
		manifest := fmt.Sprintf(`apiVersion: v1
kind: Pod
metadata:
  name: %s
  namespace: %s
  labels:
    app: %s
spec:
  containers:
  - name: netshoot
    image: nicolaka/netshoot:latest
    command: ["/bin/bash","-c","sleep infinity"]
    tty: true
`, pod, namespace, pod)
		apply := exec.Command("kubectl", "apply", "-f", "-")
		apply.Stdin = strings.NewReader(manifest)
		apply.Stdout = os.Stdout
		apply.Stderr = os.Stderr
		if err := apply.Run(); err != nil {
			return err
		}
	}

	return runKubectl(
		"-n", namespace,
		"wait", "--for=condition=Ready", fmt.Sprintf("pod/%s", pod), "--timeout=60s",
	)
}

func getServiceTargetPod(namespace, service string) (string, error) {
	jsonpath := "jsonpath={.subsets[0].addresses[0].targetRef.name}"
	out, err := captureKubectl(
		"-n", namespace,
		"get", "endpoints", service,
		"-o", jsonpath,
	)
	if err != nil {
		return "", err
	}
	if out == "" {
		return "", fmt.Errorf("no endpoint addresses found for service %s/%s", namespace, service)
	}
	return out, nil
}

func runKubectl(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func streamKubectl(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func captureKubectl(args ...string) (string, error) {
	out, err := exec.Command("kubectl", args...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
