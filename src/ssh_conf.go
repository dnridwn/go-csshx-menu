package main

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SSHServer struct {
	Name string   `yaml:"name"`
	User string   `yaml:"user"`
	IPs  []string `yaml:"ips"`
}

type SSHConf struct {
	SSHServers []SSHServer `yaml:"ssh_servers"`
}

func GetFileConfPath() string {
	return os.Getenv("CSSHX_SERVER_CONF_FILE_PATH")
}

func ReadConfFile(path string) (sshConf SSHConf, err error) {
	if filepath.Ext(path) != ".yml" {
		err = errors.New("Invalid conf file")
		return
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return
	}

	if fileInfo.IsDir() {
		err = errors.New("Invalid path")
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(content, &sshConf); err != nil {
		return
	}

	return
}

func ParseServerNames(sshConf SSHConf) (serverNames []string) {
	for _, server := range sshConf.SSHServers {
		serverNames = append(serverNames, server.Name)
	}

	return
}
