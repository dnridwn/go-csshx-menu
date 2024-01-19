package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

const (
	appName = "GO CSSHX Menu"
	version = "v1.0"
	author  = "Den Ridwan Saputra (https://github.com/dnridwn)"
)

var (
	sshConf     SSHConf
	serverNames []string
	pages       *tview.Pages
)

func main() {
	app := tview.NewApplication()
	defer handlePanic(app)

	sshConf = loadConfiguration()
	serverNames = sshConf.ParseServerNames()

	pages = tview.NewPages()
	pages.AddPage("Home", homeModal(), true, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func handlePanic(app *tview.Application) {
	if err := recover(); err != nil {
		app.Stop()
		log.Print(err)
	}
}

func loadConfiguration() SSHConf {
	conf, err := ReadConfFile(GetFileConfPath())
	if err != nil {
		panic(err)
	}

	return conf
}

func homeModal() *tview.Modal {
	m := tview.NewModal()
	m.SetText(fmt.Sprintf("%s (%s)\nby %s\nChoose Server", appName, version, author))
	m.AddButtons(serverNames)
	m.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if server, found := sshConf.FindServerByName(buttonLabel); found {
			pages.AddAndSwitchToPage(server.Name, ipListModal(server), true)
		}
	})
	return m
}

func ipListModal(server SSHServer) *tview.Modal {
	var buttons = []string{"All", "Back"}
	buttons = append(buttons[:1], append(server.IPs, buttons[1:]...)...)

	m := tview.NewModal()
	m.SetText(fmt.Sprintf("%s (%s)\nby %s\nChoose IP", appName, version, author))
	m.AddButtons(buttons)
	m.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "All" {
			OpenCSSHX(server)
		} else if buttonLabel == "Back" {
			pages.RemovePage(server.Name)
		} else if ip, found := server.FindIP(buttonLabel); found {
			OpenCSSHXSpecificIP(server, ip)
		}
	})
	return m
}
