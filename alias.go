package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kaneshin/esindex/elasticsearch"
)

type Alias struct {
}

func NewAlias() *Alias {
	return &Alias{}
}

func (cmd *Alias) Run(args []string) error {
	only := cmdFlag.Bool("only", true, "")
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	indexName := args[0]
	aliasName := aliasNameFromIndexName(indexName)
	cmdFlag.Parse(args[1:])

	client := elasticsearch.NewClient(
		elasticsearch.NewConfig().WithURL(*urlStr),
	)

	_, aliased := indexNamesAndAliasedIndexByAliasName(client, aliasName)
	if len(aliased) == 0 {
		if err := makeAlias(client, aliasName, indexName); err != nil {
			return err
		}
		fmt.Printf("%s <- %s\n", indexName, aliasName)
		return nil
	}

	if *only {
		addActions := []Action{
			Action{aliasName, indexName},
		}
		removeActions := []Action{}
		for name, _ := range aliased {
			if name != indexName {
				removeActions = append(removeActions, Action{aliasName, name})
			}
		}
		if err := updateAliases(client, addActions, removeActions); err != nil {
			return err
		}
	} else {
		if err := makeAlias(client, aliasName, indexName); err != nil {
			return err
		}
	}

	fmt.Printf("%s <- %s\n", indexName, aliasName)
	return nil
}

func (cmd *Alias) Name() string {
	return "alias"
}

func (cmd *Alias) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<index-name>", "[options]")
}
