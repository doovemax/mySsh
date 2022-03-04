package main

import (
	"fmt"
	"os"

	"os/user"

	"flag"

	"github.com/doovemax/mySsh/core"
)

func main() {
	//获取用户$HOME
	who, err := user.Current()
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	prefix := who.HomeDir
	flag.Parse()
	app := core.App{ServerPath: prefix + "/.mySsh"}
	app.Exec()
	
}
