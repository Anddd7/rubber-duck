package main

import (
	"context"

	tui "github.com/Anddd7/rubber-duck/rubber-duck-tui"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type SidecarCmds struct {
	List     SidecarListCmd     `cmd:""`
	Patch    SidecarPatchCmd    `cmd:""`
	Terminal SidecarTerminalCmd `cmd:""`
	Restore  SidecarRestoreCmd  `cmd:""`
}

type SidecarListCmd struct {
}

func (cmd *SidecarListCmd) Run(g *GlobalSettings) error {
	// list pods installed with sidecar (by label)

	return nil
}

type SidecarPatchCmd struct {
}

func (cmd *SidecarPatchCmd) Run(g *GlobalSettings) error {
	// Load configs from system setting
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	// Get the current namespace from the context
	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return err
	}

	rawConfig, err := clientConfig.RawConfig()
	if err != nil {
		return err
	}
	log.Infof("current context: %s, namespace: %s", rawConfig.CurrentContext, namespace)

	// Get clientset
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	// list pods in current namespace
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	// interactive select a pod
	podNames := []string{}
	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}
	selection := tui.NewStringSelection("Select a Pod", podNames)
	err = selection.Start()
	if err != nil {
		return err
	}
	selected := selection.GetSelected()
	if len(selected) == 0 {
		println("No pod selected")
		return nil
	}

	podName := selected[0]

	log.Infof("selected pod %s", podName)

	// // Get the specified Pod
	// pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), selectedPodName, metav1.GetOptions{})
	// if err != nil {
	// 	return err
	// }

	// // Define the new container
	// newContainer := corev1.Container{
	// 	Name:  "sidecar-curl",
	// 	Image: "curlimages/curl:8.5.0",
	// 	Args:  []string{"sleep", "3600"},
	// }

	// // Add the new container to the Pod
	// pod.Spec.Containers = append(pod.Spec.Containers, newContainer)
	// // Add labels to the Pod
	// labels := pod.GetLabels()
	// labels["rubber-duck/component"] = "sidecar"                     // command name
	// labels["rubber-duck/managed-by"] = "rubber-duck-cli/sidecar"    // command path
	// labels["rubber-duck/overlays"] = "curl"                         // sub command name
	// labels["rubber-duck/sidecar"] = "sidecar-curl"                  // container name
	// labels["rubber-duck/started"] = time.Now().Format(time.RFC3339) // started at

	// pod.SetLabels(labels)

	// // Update the Pod
	// _, err = clientset.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	// if err != nil {
	// 	return err
	// }

	return nil
}

type SidecarTerminalCmd struct {
}

func (cmd *SidecarTerminalCmd) Run(g *GlobalSettings) error {
	// open terminal to sidecar container

	return nil
}

type SidecarRestoreCmd struct {
}

func (cmd *SidecarRestoreCmd) Run(g *GlobalSettings) error {
	// remove sidecar from pod, deployment

	return nil
}
