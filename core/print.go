package core

import "fmt"

type Print struct {
}

var Printer Print

func (print Print) Logln(a ...interface{}) {
	fmt.Println(a...)
}

func (print Print) Log(a ...interface{}) {
	fmt.Print(a...)
}

func (print Print) Infoln(a ...interface{}) {
	fmt.Print("\033[32m")
	fmt.Print(a...)
	fmt.Println("\033[0m")
}

func (print Print) Info(a ...interface{}) {
	fmt.Print("\033[32m")
	fmt.Print(a...)
	fmt.Print("\033[0m")
}

func (print Print) Errorln(a ...interface{}) {
	fmt.Print("\033[31m")
	fmt.Print(a...)
	fmt.Println("\033[0m")
}

func (print Print) Error(a ...interface{}) {
	fmt.Print("\033[31m")
	fmt.Print(a...)
	fmt.Print("\033[0m")
}
