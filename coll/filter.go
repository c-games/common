package coll

func FilterMapByKey(updatableKeys []string, data map[string]interface{}) map[string]interface{} {
	rlt := make(map[string]interface{})
	for _, k := range updatableKeys {
		if val, ok := data[k]; ok {
			rlt[k] = val
		}
	}
	return rlt
}

