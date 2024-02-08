package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
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
	// for debugging
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &opts)))

	ctx := kong.Parse(&CLI)
	slog.Debug(fmt.Sprintf("cmd: %s", ctx.Command()))

	err := ctx.Run(&GlobalSettings{
		Version: "0.0.1",
	})
	if err != nil {
		slog.Debug(fmt.Sprintf("Got error: %+v", err))
	}
	ctx.FatalIfErrorf(err)
}
