package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kaneshin/esindex/elasticsearch"
)

type Create struct {
	mappings *string
}

func NewCreate() *Create {
	return &Create{
		mappings: cmdFlag.String("mappings", "", ""),
	}
}

func (cmd *Create) Run(args []string) error {
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	aliasName := args[0]
	indexName := indexNameFromAliasName(aliasName)
	cmdFlag.Parse(args[1:])

	client := elasticsearch.NewClient(
		elasticsearch.NewConfig().WithURL(*urlStr),
	)

	req := elasticsearch.NewRequest("PUT", indexName, strings.NewReader(*cmd.mappings))
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if v, ok := res["acknowledged"].(bool); ok && v {
		_, aliased := indexNamesAndAliasedIndexByAliasName(client, aliasName)
		if len(aliased) == 0 {
			if err := makeAlias(client, aliasName, indexName); err != nil {
				return err
			}
			fmt.Printf("%s <- %s\n", indexName, aliasName)
			return nil
		}
		fmt.Printf("%s\n", indexName)
	}

	if v, ok := res["error"].(string); ok && len(v) > 0 {
		return errors.New(v)
	}

	return nil
}

func (cmd *Create) Name() string {
	return "create"
}

func (cmd *Create) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<alias-name>", "[options]")
}
