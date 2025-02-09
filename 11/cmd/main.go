package main

import (
	"log"
	"task11/telnet"
)

func main() {
	var cfg telnet.Config
	if err := telnet.Telnet(&cfg); err != nil {
		log.Fatalln(err)
	}
}
