package core

import (
	"flag"
)

var (
	f    string
	host string
	port int
	list bool
)

func init() {
	flag.StringVar(&f, "f", "", "specify config file")
	flag.StringVar(&host, "host", "", "specity remote host")
	flag.IntVar(&port, "port", 22, "specity remote port")
	flag.BoolVar(&list, "list", flase, "list remote hosts")

}
