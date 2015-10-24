package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/kaneshin/esindex/elasticsearch"
)

const (
	versionFormat = "20060102150405"
)

func indexNameFromAliasName(aliasName string) string {
	return fmt.Sprintf("%s_v%s", aliasName, time.Now().Format(versionFormat))
}

func aliasNameFromIndexName(indexName string) string {
	indexLen := len(indexName)
	versionLen := len("_v" + versionFormat)
	if versionLen < indexLen {
		return indexName[:indexLen-versionLen]
	}
	return ""
}

func indexNamesAndAliasedIndexByAliasName(client *elasticsearch.Client, aliasName string) (names []string, aliased map[string]bool) {
	req := elasticsearch.NewRequest("GET", "_aliases", nil)
	res, err := client.Do(req)
	if err != nil {
		return
	}

	aliased = map[string]bool{}
	for idxName, idx := range res {
		if aliasNameFromIndexName(idxName) != aliasName {
			continue
		}
		names = append(names, idxName)
		if idx, ok := idx.(map[string]interface{}); ok {
			if aliases, ok := idx["aliases"].(map[string]interface{}); ok {
				for alias := range aliases {
					if alias == aliasName {
						aliased[idxName] = true
					}
				}
			}
		}
	}
	sort.Sort(sort.StringSlice(names))
	return
}

type Action struct {
	Alias string `json:"alias"`
	Index string `json:"index"`
}

func updateAliases(client *elasticsearch.Client, addActions, removeActions []Action) error {
	actions := []map[string]Action{}
	for _, action := range addActions {
		actions = append(actions, map[string]Action{
			"add": action,
		})
	}
	for _, action := range removeActions {
		actions = append(actions, map[string]Action{
			"remove": action,
		})
	}

	body, err := json.Marshal(map[string]interface{}{"actions": actions})
	if err != nil {
		return err
	}

	req := elasticsearch.NewRequest("POST", "_aliases", bytes.NewReader(body))
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if v, ok := res["error"].(string); ok && len(v) > 0 {
		return errors.New(v)
	}
	return nil
}

func makeAlias(client *elasticsearch.Client, aliasName, indexName string) error {
	return updateAliases(client, []Action{
		Action{aliasName, indexName},
	}, nil)
}
