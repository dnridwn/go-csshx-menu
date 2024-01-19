package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const baseCommand = "csshx --login %s %s"

func OpenCSSHX(server SSHServer) {
	cmd := exec.Command(
		"osascript",
		"-s", "h",
		"-e", `tell application "Terminal" to activate do script "`+
			fmt.Sprintf(baseCommand, server.User, strings.Join(server.IPs, " "))+`"`,
	)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}

func OpenCSSHXSpecificIP(server SSHServer, ip string) {
	cmd := exec.Command(
		"osascript",
		"-s", "h",
		"-e", `tell application "Terminal" to activate do script "`+
			fmt.Sprintf(baseCommand, server.User, ip)+`"`,
	)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
