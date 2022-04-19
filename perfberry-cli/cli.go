package main

import (
	"github.com/perfberry/cli/api"
	"github.com/perfberry/cli/cmd"
	"github.com/perfberry/cli/helpers"
	"github.com/perfberry/cli/ui"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	appName = "Perfberry CLI"
	version = "1.0.0"
)

type App struct {
	CLI *cli.App
	API *api.Client
	UI  *ui.Client
}

func (a *App) Run() {
	helpers.ClearStatus()

	if err := a.CLI.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

	if code, err := helpers.ReadStatus(); err == nil {
		os.Exit(code)
	}
}

var app App

func init() {
	app.API = api.NewClient().SetUserAgent(appName + "/" + version)
	app.UI = ui.NewClient(app.API)
	app.CLI = cmd.NewCmd(appName, version, app.API, app.UI)
}

func main() {
	app.Run()
}
