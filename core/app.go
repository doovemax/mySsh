package core

import (
	"fmt"
	"os"
	"strconv"
)

const VERSION = "0.2"

type App struct {
	ServerPath string
	servers    []Server
}

//func (app *App) version() {
//	fmt.Println("Autossh", VERSION, "。")
//	fmt.Println("由 Lenbo 编写，项目地址：https://github.com/islenbo/autossh。")
//}
//
//func (app *App) help() {
//	fmt.Println("go写的一个ssh远程客户端。可一键登录远程服务器，主要用来弥补Mac/Linux Terminal ssh无法保存密码的不足。")
//	fmt.Println("基本用法：")
//	fmt.Println("  直接输入autossh不带任何参数，列出所有服务器，输入对应编号登录。")
//	fmt.Println("参数：")
//	fmt.Println("  -v, --version", "\t", "显示 autossh 的版本信息。")
//	fmt.Println("  -h, --help   ", "\t", "显示帮助信息。")
//	fmt.Println("操作：")
//	fmt.Println("  list         ", "\t", "显示所有server。")
//}
func (app *App) inputsh() Server {
	var input string
	Printer.Info("Enter number: ")
	fmt.Scanln(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		Printer.Errorln("Input error,Please again")
		return app.inputsh()
	}

	if num <= 0 || num > len(app.servers) {
		Printer.Errorln("Input errors,Please again")
		return app.inputsh()
	}

	return app.servers[num-1]
}

func (app *App) start() {
	Printer.Infoln("=========Welcome Auto SSH===========")
	for i, server := range app.servers {
		Printer.Infoln("\033[31m["+strconv.Itoa(i+1)+"]\033[0m", server.Name)
	}
	Printer.Infoln("====================================")

	server := app.inputsh()
	server.Connection()
}
func (app *App) list() {
	for _, server := range app.servers {
		Printer.Logln(server.Name)
	}
}

func (app *App) Exec() {
	var err error
	if len(os.Args) == 0 {
		app.servers, err = ConfigPath(app.ServerPath)
		if os.IsNotExist(err) {
			var FLAG string
			Printer.Info("Creating config path ?(Yes/Quit): ")
			fmt.Scan(&FLAG)
			switch FLAG {
			case "Y", "yes", "y":
				err = CreatConfig(app.ServerPath)
				if err != nil {
					Printer.Errorln(err)
					os.Exit(2)

				}
			case "Q", "q", "quit":
				os.Exit(0)
			}

		} else if err != nil {
			Printer.Errorln(err)
		}
	}
	if len(os.Args) > 1 {
		err = Args(app)
		if err != nil {
			Printer.Errorln(err)
			os.Exit(2)
		}
	}
	app.start()
}
