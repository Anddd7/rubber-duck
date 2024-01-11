package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var CLI struct {
	K struct {
		Commands []string `arg:"" name:"cmd" help:"Commands to execute."`
	} `cmd:"" help:"wrapper for kubectl"`
}

func main() {
	// TODO for development
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "k <cmd>":
		k()
	default:
		panic(ctx.Command())
	}
}
