package main

type DevcontainerCmds struct {
	Create  DevcontainerCreateCmd  `cmd:""`
	Delete  DevcontainerDeleteCmd  `cmd:""`
	Pause   DevcontainerPauseCmd   `cmd:""`
	Connect DevcontainerConnectCmd `cmd:""`
}

type DevcontainerCreateCmd struct {
}

func (cmd *DevcontainerCreateCmd) Run(g *GlobalSettings) error {
	// create a devcontainer

	return nil
}

type DevcontainerDeleteCmd struct {
}

func (cmd *DevcontainerDeleteCmd) Run(g *GlobalSettings) error {
	// stop and delete devcontainer

	return nil
}

type DevcontainerPauseCmd struct {
}

func (cmd *DevcontainerPauseCmd) Run(g *GlobalSettings) error {
	// delete pod but keep volumes

	return nil
}

type DevcontainerConnectCmd struct {
}

func (cmd *DevcontainerConnectCmd) Run(g *GlobalSettings) error {
	// show connection status to export the kubeconfig

	return nil
}
