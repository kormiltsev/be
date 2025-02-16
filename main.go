package main

import (
	"log/slog"
	"os"

	cli "github.com/urfave/cli/v2"

	"github.com/kormiltsev/be/cmd/run"
	"github.com/kormiltsev/be/config"
	"github.com/kormiltsev/be/version"
)

// RENAMEME =>

func main() {
	app := cli.NewApp()
	app.Name = "RENAMEME"
	app.Usage = "RENAMEME"
	app.Copyright = "(c) Artem Kormiltsev"
	app.Version = version.Version + " (" + version.GitCommit + ")"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug, d",
			Usage:   "Enable verbose logging, log level DEBUG",
			EnvVars: []string{"RENAMEME_DEBUG"},
		},
	}

	app.Before = func(c *cli.Context) error {
		config.DebugMode = c.Bool("debug")
		return nil
	}

	app.Commands = []*cli.Command{
		run.Command,
	}

	if err := app.Run(os.Args); err != nil {
		slog.Info("exit RENAMEME", slog.String("err", err.Error()))
	}
}
