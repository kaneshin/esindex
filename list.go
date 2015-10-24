package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kaneshin/esindex/elasticsearch"
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

	client := elasticsearch.NewClient(
		elasticsearch.NewConfig().WithURL(*urlStr),
	)

	indices, aliased := indexNamesAndAliasedIndexByAliasName(client, aliasName)
	for _, index := range indices {
		fmt.Printf("%s", index)
		if aliased[index] {
			fmt.Printf(" <- %s", aliasName)
		}
		fmt.Println("")
	}

	return nil
}

func (cmd *List) Name() string {
	return "list"
}

func (cmd *List) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<alias-name>", "[options]")
}
