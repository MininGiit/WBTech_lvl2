package sort

import (
	"flag"
	"os"
)

type Config struct {
	SortedColumn int
	NumSort      bool
	Reversed     bool
	Unique       bool
	InputFile  string
	OutputFile string
}

func (cfg *Config) ParseArgs() {

	flag.IntVar(&cfg.SortedColumn, "k", 0, "column for sort")
	flag.BoolVar(&cfg.NumSort, "n", false, "numeric sort")
	flag.BoolVar(&cfg.Reversed, "r", false, "desc sort")
	flag.BoolVar(&cfg.Unique, "u", false, "only unique rows")
	flag.Parse()

	if cfg.SortedColumn < 0 {
		flag.Usage()
		os.Exit(2)
	}
	args := flag.Args()
	if len(args) > 2 { 
		flag.Usage()
		os.Exit(2)
	}
	if len(args) >= 1 {
		cfg.InputFile = args[0]
	}
	if len(args) == 2 {
		cfg.OutputFile = args[1]
	}
	check := func() func(bool) {
		var f bool
		return func(b bool) {
			if f && b {
				flag.Usage()
				os.Exit(2)
			}
			f = f || b
		}
	}()
	check(cfg.NumSort)
}
