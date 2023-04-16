package main

import (
	"flag"

	"github.com/jeffcail/go-im/boot"
)

var (
	s = flag.String("serve", "http", "👉选择运行的服务👈")
)

func init() {
	flag.Parse()
	boot.Run()
}

func main() {
	switch *s {
	case "http":
		boot.ApiServer()
	default:
		boot.ApiServer()
	}
}
