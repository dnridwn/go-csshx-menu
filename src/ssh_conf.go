package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SSHServer struct {
	Name string   `yaml:"name"`
	User string   `yaml:"user"`
	IPs  []string `yaml:"ips"`
}

func (s *SSHServer) FindIP(target string) (string, bool) {
	for _, ip := range s.IPs {
		if ip == target {
			return ip, true
		}
	}
	return "", false
}

type SSHConf struct {
	SSHServers []SSHServer `yaml:"ssh_servers"`
}

func (sc *SSHConf) ParseServerNames() (serverNames []string) {
	for _, server := range sc.SSHServers {
		serverNames = append(serverNames, server.Name)
	}

	return
}

func (sc *SSHConf) FindServerByName(target string) (SSHServer, bool) {
	for _, server := range sc.SSHServers {
		if server.Name == target {
			return server, true
		}
	}
	return SSHServer{}, false
}

func GetFileConfPath() string {
	return os.Getenv("CSSHX_SERVER_CONF_FILE_PATH")
}

func ReadConfFile(path string) (sshConf SSHConf, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return
	}

	if fileInfo.IsDir() {
		err = errors.New("invalid path: directory provided instead of a file")
		return
	}

	if ext := filepath.Ext(path); ext != ".yml" {
		err = fmt.Errorf("invalid conf file: not a YAML file (got %s)", ext)
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
