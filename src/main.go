package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/rivo/tview"
)

const (
	version   string = "v1.0"
	author    string = "Den Ridwan Saputra (https://github.com/dnridwn)"
	modalText string = "GO CSSHX MENU %s\nby %s\n\nCHOOSE SERVER"
)

func main() {
	sshConf, err := ReadConfFile(
		GetFileConfPath(),
	)

	if err != nil {
		log.Print(err.Error())
		return
	}

	serverNames := ParseServerNames(sshConf)
	if len(serverNames) == 0 {
		log.Print(errors.New("Server empty"))
		return
	}

	app := tview.NewApplication()
	modal := CreateModal(
		fmt.Sprintf(modalText, version, author),
		serverNames,
		func(buttonIndex int, buttonLabel string) {
			for _, server := range sshConf.SSHServers {
				if buttonLabel == server.Name {
					OpenCSSHX(server)
					app.Stop()
				}
			}
		},
	)

	if err := app.SetRoot(modal, true).EnableMouse(true).Run(); err != nil {
		log.Panic(err)
	}
}
