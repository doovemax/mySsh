package main

import (
	"fmt"
	"os"
	"path/filepath"

	"./core"
)

func main() {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	app := core.App{ServerPath: path + "/server.json"}
	app.Exec()
}
