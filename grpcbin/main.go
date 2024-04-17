package main

import (
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
)

type Globals struct {
	Debug bool   `short:"d"`
	Host  string `default:"localhost"`
	Port  int    `default:"50051"`
}

type CLI struct {
	Globals

	Serve                  ServeCmd                  `cmd:""`
	Unary                  UnaryCmd                  `cmd:""`
	ServerStreaming        ServerStreamingCmd        `cmd:"" aliases:"srvstr"`
	ClientStreaming        ClientStreamingCmd        `cmd:"" aliases:"clistr"`
	BidirectionalStreaming BidirectionalStreamingCmd `cmd:"" aliases:"bistr"`
}

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("grpcbin"),
		kong.Description("A gRPC server and client for testing"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		},
	)

	if cli.Globals.Debug {
		opts := slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &opts)))
	}

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
}
