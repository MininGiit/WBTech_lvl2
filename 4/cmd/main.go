package main

import (
	"task4/sort"
)

func main() {
	var cfg sort.Config
	cfg.ParseArgs()
	sort.Sort(cfg)
}
