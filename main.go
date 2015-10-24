package main

import (
	"bytes"
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
			var buf bytes.Buffer
			cmd.Usage(&buf)
			fmt.Printf("    %s\n", buf.String())
		}
	}
	flag.Parse()
	os.Exit(run())
}

func run() int {
	const (
		success  = 0
		failure  = 1
		notfound = 127
	)

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return success
	}
	name := args[0]
	if name == "help" {
		flag.Usage()
		return success
	}

	for _, cmd := range commands {
		if cmd.Name() == name {
			if err := cmd.Run(args[1:]); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return failure
			}
			return success
		}
	}

	return notfound
}
