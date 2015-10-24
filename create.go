package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Create struct {
}

func NewCreate() *Create {
	return &Create{}
}

const (
	versionFormat = "2006150405"
)

func (cmd *Create) Run(args []string) error {
	mappings := cmdFlag.String("mappings", "", "")
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	aliasName := args[0]
	indexName := indexNameFromAliasName(aliasName)
	cmdFlag.Parse(args[1:])

	req := NewRequest("PUT", indexName, strings.NewReader(*mappings))
	result, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if v, ok := result["acknowledged"].(bool); ok && v {
		has := func(aliasName string) bool {
			req = NewRequest("GET", "_aliases", nil)
			result, err := DefaultClient.Do(req)
			if err != nil {
				return true
			}
			for _, idx := range result {
				if idx, ok := idx.(map[string]interface{}); ok {
					if aliases, ok := idx["aliases"].(map[string]interface{}); ok {
						for alias := range aliases {
							if alias == aliasName {
								return true
							}
						}
					}
				}
			}
			return false
		}(aliasName)
		if !has {
			var d = `{
                "actions": [{
                    "add": {
                        "alias": "` + aliasName + `",
                        "index": "` + indexName + `"
                    }}
                    ]
                }`
			req = NewRequest("POST", "_aliases", strings.NewReader(d))
			result, err := DefaultClient.Do(req)
			if err != nil {
				return err
			}
			if v, ok := result["acknowledged"].(bool); ok && v {
				fmt.Printf("%s <- %s\n", indexName, aliasName)
				return nil
			}
		}
		fmt.Printf("%s\n", indexName)
	}

	return nil
}

func (cmd *Create) Name() string {
	return "create"
}

func (cmd *Create) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s\n", cmd.Name(), "<alias-name>", "[options]")
}
