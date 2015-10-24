package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/kaneshin/esindex/elasticsearch"
)

type Settings struct {
}

func NewSettings() *Settings {
	return &Settings{}
}

func (cmd *Settings) Run(args []string) error {
	if len(args) < 1 {
		cmd.Usage(os.Stdout)
		return nil
	}
	indexName := args[0]
	cmdFlag.Parse(args[1:])

	client := elasticsearch.NewClient(
		elasticsearch.NewConfig().WithURL(*urlStr),
	)

	req := elasticsearch.NewRequest("GET", path.Join(indexName, "_settings"), nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	b, _ := json.MarshalIndent(res[indexName], "", "  ")
	fmt.Println(string(b))

	return nil
}

func (cmd *Settings) Name() string {
	return "settings"
}

func (cmd *Settings) Usage(w io.Writer) {
	fmt.Fprintf(w, "%-8s%10s %s", cmd.Name(), "<index-name>", "[options]")
}
