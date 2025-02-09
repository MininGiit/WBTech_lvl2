package main

import (
	"task7/cut"
)

func main() {
	var cfg cut.Config
	cfg.ParseArgs()
	cut.ReadSTDIN(cfg)
}