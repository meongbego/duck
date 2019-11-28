package cmd

import (
	"github.com/sofyan48/duck/src/libs"
	"github.com/urfave/cli"
)

var app *cli.App

// ArgsMapping object mapping
type ArgsMapping struct {
	ConfigPath string
	TestArg    string
}

// Args Glabal Acces args command
var Args ArgsMapping

// Init Initialise a CLI app
func Init() *cli.App {
	app = cli.NewApp()
	app.Name = "duck"
	app.Usage = "duck scheduler task send"
	app.Author = "meong"
	app.Email = "meongbego@gmail.com"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "",
			Destination: &Args.ConfigPath,
			Usage:       "Path to a configuration file",
		},
	}
	return app
}

// AppCommands mapping command
func AppCommands() *cli.App {
	app := Init()
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch worker",
			Action: func(c *cli.Context) error {
				srv, err := InitServer()
				if err != nil {
					return err
				}
				libs.WorkerStart(srv)
				return nil
			},
		},
	}
	return app
}