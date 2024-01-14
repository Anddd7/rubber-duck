package main

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

var CLI struct {
	K struct {
		Commands []string `arg:"" name:"cmd" help:"Commands to execute."`
	} `cmd:"" help:"wrapper for kubectl"`
}

func main() {
	// TODO for development logging
	log.SetLevel(log.TraceLevel)

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "k <cmd>":
		k()
	default:
		panic(ctx.Command())
	}
}
