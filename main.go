package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	cmdFlag = flag.NewFlagSet("esindex", flag.ExitOnError)
	urlStr  = cmdFlag.String("url", "http://127.0.0.1:9200", "")
)

type Command interface {
	Run([]string) error
	Name() string
	Usage(io.Writer)
}

var commands = []Command{
	NewCreate(),
	NewAlias(),
	NewList(),
}

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println(`Usage of esindex:
  Commands:`)
		for _, cmd := range commands {
			fmt.Print("    ")
			cmd.Usage(os.Stdout)
		}
	}
	flag.Parse()
	os.Exit(run())
}

func run() int {
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return 0
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			if err := cmd.Run(args[1:]); err != nil {
				return 1
			}
			return 0
		}
	}

	return 127
}
