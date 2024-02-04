package main

import (
	"context"

	tui "github.com/Anddd7/rubber-duck/rubber-duck-tui"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	sidecarLabel   = "rubber-duck/sidecar"
	containerLabel = "rubber-duck/container"
)

type SidecarCmds struct {
	List    SidecarListCmd    `cmd:""`
	Patch   SidecarPatchCmd   `cmd:""`
	Restore SidecarRestoreCmd `cmd:""`
}

type SidecarListCmd struct {
}

func (cmd *SidecarListCmd) Run(g *GlobalSettings) error {
	clientset, namespace, err := connect()
	if err != nil {
		return err
	}

	// list with labels and status
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: sidecarLabel,
	})
	if err != nil {
		return err
	}

	var headers = []string{"Name", "Sidecar", "Container"}
	var rows = [][]string{}
	for _, deploy := range deployments.Items {
		sidecar := deploy.Labels[sidecarLabel]
		container := deploy.Labels[containerLabel]

		rows = append(rows, []string{
			deploy.Name,
			sidecar,
			container,
		})
	}

	ptable := tui.NewPrettyTable("Deployments running patches", headers, rows)
	ptable.Print()

	return nil
}

type SidecarPatchCmd struct {
	Name string `arg:"" optional:""`

	Curl bool `help:"add curl sidecar to pod"`
}

func (cmd *SidecarPatchCmd) Run(g *GlobalSettings) error {
	clientset, namespace, err := connect()
	if err != nil {
		return errors.Wrap(err, "failed to connect to k8s")
	}

	deployname := cmd.Name
	if deployname == "" {
		// list deployments in current namespace
		deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}

		// interactive select a deployment
		options := []string{}
		for _, deploy := range deployments.Items {
			options = append(options, deploy.Name)
		}

		selection := tui.NewStringSelection("Select a Deployment", options)
		err = selection.Start()
		if err != nil {
			return err
		}

		selected := selection.GetSelected()
		if len(selected) == 0 {
			println("No deployment selected")
			return nil
		}
		deployname = selected[0]
	}

	log.Infof("selected deployment %s", deployname)

	// Get the specified Deployment
	deploy, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployname, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// Add sidecar container to the Deployment
	var patchName string
	var container corev1.Container
	if cmd.Curl {
		// TODO load from kustomization-patches
		patchName = "curl"
		container = corev1.Container{
			Name:  "sidecar-curl",
			Image: "curlimages/curl:8.5.0",
			Args:  []string{"sleep", "3600"},
		}
	} else {
		return errors.New("no sidecar specified")
	}

	// Update the Deployment
	deploy.Spec.Template.Spec.Containers = append(deploy.Spec.Template.Spec.Containers, container)

	// annoations := deploy.GetAnnotations()
	// annoations["rubber-duck/component"] = "sidecar"
	// annoations["rubber-duck/managed-by"] = "rubber-duck-cli/sidecar"
	// deploy.SetAnnotations(annoations)

	labels := deploy.GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}
	labels[sidecarLabel] = patchName
	labels[containerLabel] = container.Name
	deploy.SetLabels(labels)

	_, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

type SidecarRestoreCmd struct {
	Name string `arg:"" optional:""`
}

func (cmd *SidecarRestoreCmd) Run(g *GlobalSettings) error {
	clientset, namespace, err := connect()
	if err != nil {
		return errors.Wrap(err, "failed to connect to k8s")
	}

	deployname := cmd.Name
	if deployname == "" {
		// list deployments in current namespace
		deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: sidecarLabel,
		})
		if err != nil {
			return err
		}

		// interactive select a deployment
		options := []string{}
		for _, deploy := range deployments.Items {
			options = append(options, deploy.Name)
		}

		selection := tui.NewStringSelection("Select a Deployment", options)
		err = selection.Start()
		if err != nil {
			return err
		}

		selected := selection.GetSelected()
		if len(selected) == 0 {
			println("No deployment selected")
			return nil
		}

		deployname = selected[0]
	}

	log.Infof("selected deployment %s", deployname)

	// Get the specified Deployment
	deploy, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deployname, metav1.GetOptions{})
	if err != nil {
		return err
	}

	// Remove labels
	labels := deploy.GetLabels()
	if labels == nil {
		return errors.New("no sidecar found")
	}
	containerName, ok := labels[containerLabel]
	if !ok {
		return errors.New("no sidecar found")
	}
	delete(labels, sidecarLabel)
	delete(labels, containerLabel)
	deploy.SetLabels(labels)

	// Remove sidecar container
	var newContainers []corev1.Container
	for _, container := range deploy.Spec.Template.Spec.Containers {
		if container.Name != containerName {
			newContainers = append(newContainers, container)
		}
	}
	deploy.Spec.Template.Spec.Containers = newContainers

	// Update the Deployment
	_, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func connect() (kubernetes.Interface, string, error) {
	// Load configs from system setting
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	// Get the current namespace from the context
	namespace, _, err := clientConfig.Namespace()
	if err != nil {
		return nil, "", err
	}

	rawConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, "", err
	}
	log.Infof("current context: %s, namespace: %s", rawConfig.CurrentContext, namespace)

	// Get clientset
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, "", err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, "", err
	}

	return clientset, namespace, nil
}
