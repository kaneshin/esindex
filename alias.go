package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Alias struct {
}

func NewAlias() *Alias {
	return &Alias{}
}

func (cmd *Alias) Run(args []string) error {
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	indexName := args[0]
	aliasName := aliasNameFromIndexName(indexName)
	cmdFlag.Parse(args[1:])

	associatedIndexName, has := func(aliasName string) (string, bool) {
		req := NewRequest("GET", "_aliases", nil)
		result, err := DefaultClient.Do(req)
		if err != nil {
			return "", false
		}
		for idxName, idx := range result {
			if idx, ok := idx.(map[string]interface{}); ok {
				if aliases, ok := idx["aliases"].(map[string]interface{}); ok {
					for alias := range aliases {
						if alias == aliasName {
							return idxName, true
						}
					}
				}
			}
		}
		return "", false
	}(aliasName)

	var data string
	if has {
		data = `{
            "actions": [{
                "remove": {
                    "alias": "` + aliasName + `",
                    "index": "` + associatedIndexName + `"
                }
            }, {
                "add": {
                    "alias": "` + aliasName + `",
                    "index": "` + indexName + `"
                }
            }]
        }`

	} else {
		data = `{
            "actions": [{
                "add": {
                    "alias": "` + aliasName + `",
                    "index": "` + indexName + `"
                }}
                ]
            }`
	}
	req := NewRequest("POST", "_aliases", strings.NewReader(data))
	result, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if v, ok := result["acknowledged"].(bool); ok && v {
		fmt.Printf("%s <- %s\n", indexName, aliasName)
		return nil
	}

	return nil
}

func (cmd *Alias) Name() string {
	return "alias"
}

func (cmd *Alias) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<index-name>", "[options]")
}
