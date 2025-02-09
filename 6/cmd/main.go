package main

import (
	"task6/grep"
)

func main() {
	var cfg grep.Config
	cfg.ParseArgs()
	grep.FindRows(cfg)
}