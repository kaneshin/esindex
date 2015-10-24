package main

import (
	"fmt"
	"io"
	"os"
)

type List struct {
}

func NewList() *List {
	return &List{}
}

func (cmd *List) Run(args []string) error {
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	aliasName := args[0]
	cmdFlag.Parse(args[1:])

	req := NewRequest("GET", "_aliases", nil)
	result, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	for idxName, idx := range result {
		if aliasNameFromIndexName(idxName) != aliasName {
			continue
		}
		fmt.Printf("%s", idxName)
		if idx, ok := idx.(map[string]interface{}); ok {
			if aliases, ok := idx["aliases"].(map[string]interface{}); ok {
				for alias := range aliases {
					if alias == aliasName {
						fmt.Printf(" <- %s", aliasName)
					}
				}
			}
		}
		fmt.Print("\n")
	}

	return nil
}

func (cmd *List) Name() string {
	return "list"
}

func (cmd *List) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<alias-name>", "[options]")
}
