package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/sirupsen/logrus"
)

func main() {
	app := cli.App("cli", "CLI App")

	// Declare our first command, which is invocable with "uman list"
	app.Command("hello-world", "a simple cli command", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			logrus.Info("Hello World!")
		}
	})

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Panic("Failed to run CLI app")
	}
}
