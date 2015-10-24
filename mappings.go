package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/kaneshin/esindex/elasticsearch"
)

type Mappings struct {
}

func NewMappings() *Mappings {
	return &Mappings{}
}

func (cmd *Mappings) Run(args []string) error {
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	indexName := args[0]
	cmdFlag.Parse(args[1:])

	client := elasticsearch.NewClient(
		elasticsearch.NewConfig().WithURL(*urlStr),
	)

	req := elasticsearch.NewRequest("GET", path.Join(indexName, "_mappings"), nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	b, _ := json.MarshalIndent(res[indexName], "", "  ")
	fmt.Println(string(b))

	return nil
}

func (cmd *Mappings) Name() string {
	return "mappings"
}

func (cmd *Mappings) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<index-name>", "[options]")
}
