package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var (
	ostype   = os.Getenv("GOOS")
	listFile []string
)

func ConfigPath(conf string) (servers []Server, err error) {
	file, err := os.Stat(conf)
	if err != nil {
		Printer.Errorln(err)
		os.Exit(1)
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
