package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"path/filepath"

	"errors"
)

var (
	ostype         = os.Getenv("GOOS")
	listFile       []string
	configTemplete string = `[
  {
    "name": "vagrant",
    "ip": "192.168.10.19",
    "port": 22,
    "user": "root",
    "password": "root",
    "method": "password"
  },
  {
    "name": "ssh-pem",
    "ip": "192.168.33.11",
    "port": 22,
    "user": "root",
    "password": "your pem file password or empty",
    "method": "pem",
    "key": "your pem file path"
  }
]`
)

func ConfigPath(conf string) (servers []Server, err error) {
	file, err := os.Stat(conf)
	if err != nil {
		Printer.Errorln(err)
		return nil, err
	}

	//config path  is a dir
	if file.IsDir() {
		servers, err = configDir(conf)
	} else { //当个文件处理
		servers, err = configFile(conf)
	}
	return
}

func configDir(conf string) ([]Server, error) {
	var servers []Server
	err := filepath.Walk(conf, listFunc)
	if err != nil {
		Printer.Errorln(err)
		return nil, err
	}
	if len(listFile) < 1 {
		err := errors.New("No config file")
		return nil, err
	}
	for _, file := range listFile {
		serversTmp, err := configFile(file)
		if err != nil {
			Printer.Errorln(err)
			break
		}
		servers = append(servers, serversTmp...)
	}
	return servers, nil
}

func configFile(conf string) ([]Server, error) {
	var server []Server
	b, err := ioutil.ReadFile(conf)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func listFunc(path string, f os.FileInfo, err error) error {
	var strRet string
	if ostype == "windows" {
		strRet += "\\"
	} else {
		strRet += "/"
	}

	if f == nil {
		return err
	}

	if f.IsDir() {
		return nil
	}

	strRet += path
	ok := strings.HasSuffix(strRet, ".json")
	if ok {
		listFile = append(listFile, strRet)
	}
	return nil
}

func CreatConfig(file string) error {
	err := os.Mkdir(file, 0755)
	if err != nil {
		return err
	}

	configByte := []byte(configTemplete)
	err = ioutil.WriteFile(file+"/server.json", configByte, 0655)
	if err != nil {
		return err
	}
	return nil
}
