package main

import (
	"log"
	"os"
	"task10/wget"
)

func main() {
	var cfg wget.Config
	cfg.ParseConfig()
	logger := log.New(os.Stderr, "error: ", log.Ldate|log.Ltime|log.Lshortfile)
	wget.Wget(&cfg, logger)
}