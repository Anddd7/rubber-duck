package main

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

var CLI struct {
	Cidr         CidrCmds         `cmd:""`
	Sidecar      SidecarCmds      `cmd:""`
	Devcontainer DevcontainerCmds `cmd:""`
}

type GlobalSettings struct {
	Version string
}

func main() {
	ctx := kong.Parse(&CLI)

	// TODO for development
	log.SetLevel(log.DebugLevel)
	log.Debugf("cmd: %s", ctx.Command())

	err := ctx.Run(&GlobalSettings{
		Version: "0.0.1",
	})
	if err != nil && log.GetLevel() == log.DebugLevel {
		log.Errorf("Got error: %+v", err)
	}
	ctx.FatalIfErrorf(err)
}
