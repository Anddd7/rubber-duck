package main

type SidecarCmds struct {
	List    SidecarListCmd    `cmd:""`
	Patch   SidecarPatchCmd   `cmd:""`
	Restore SidecarRestoreCmd `cmd:""`
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
	// select one of pod, deployment to add sidecar

	return nil
}

type SidecarRestoreCmd struct {
}

func (cmd *SidecarRestoreCmd) Run(g *GlobalSettings) error {
	// remove sidecar from pod, deployment

	return nil
}
