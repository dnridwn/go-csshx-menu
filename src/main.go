package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	appName = "GO CSSHX Menu"
	version = "v2.0"
	author  = "Den Ridwan Saputra (https://github.com/dnridwn)"
)

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

func main() {
	app := tview.NewApplication()
	defer handlePanic(app)

	ipView := tview.NewFlex()
	serverView := tview.NewFlex()
	homeView := tview.NewFlex()

	sshConf := loadConfiguration()
	serverNames := sshConf.ParseServerNames()

	ipView.SetBorder(true)
	ipView.SetTitle("CHOOSE IP")
	ipView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			ipView.Clear()
			app.SetFocus(serverView)
		}
		return event
	})

	serverList := tview.NewList()
	serverList.ShowSecondaryText(false)
	serverList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		if server, found := sshConf.FindServerByName(s1); found {
			ipList := tview.NewList()
			ipList.ShowSecondaryText(false)
			ipList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
				switch s1 {
				case "All":
					OpenCSSHX(server)
				default:
					OpenCSSHXSpecificIP(server, s1)
				}
			})
			for _, ip := range server.IPs {
				ipList.AddItem(ip, "", '-', nil)
			}

			ipView.Clear()
			ipView.AddItem(ipList, 0, 1, true)
			app.SetFocus(ipView)
		}
	})
	for _, serverName := range serverNames {
		serverList.AddItem(serverName, "", '-', nil)
	}

	serverView.SetBorder(true)
	serverView.SetTitle("CHOOSE SERVER")
	serverView.AddItem(serverList, 0, 1, true)

	homeView.SetBorder(true)
	homeView.SetTitle(fmt.Sprintf("%s %s by %s", appName, version, author))
	homeView.AddItem(serverView, 0, 1, true)
	homeView.AddItem(ipView, 0, 1, false)

	if err := app.SetRoot(homeView, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
