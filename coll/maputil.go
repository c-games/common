package coll

import (
	"errors"
	"sort"
)

func GetKeysFromMap(ori_map map[string]interface{}) []string {
	var keys []string
	for key, _ := range ori_map {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func CheckMapNotOuterKeys(allow_keys []string, ori_map map[string]interface{}) error {
	for _, key := range allow_keys {
		if _, ok := ori_map[key] ; !ok {
			return errors.New("Unexpect keys: " + key)
		}
	}
	return nil
}
