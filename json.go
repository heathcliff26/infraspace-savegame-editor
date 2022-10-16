package main

func GetJSONObject(parent map[string]interface{}, path ...string) (map[string]interface{}, bool) {
	return getJSONObject(parent, path)
}

func GetJSONArray(parent map[string]interface{}, path ...string) ([]interface{}, bool) {
	i := len(path) - 1
	var ok bool
	if i > 0 {
		parent, ok = getJSONObject(parent, path[:i])
		if !ok {
			return nil, false
		}
	}
	resultInterface, ok := parent[path[i]]
	if !ok || resultInterface == nil {
		return nil, false
	}
	result, ok := resultInterface.([]interface{})
	if !ok {
		return nil, false
	}
	return result, true
}

func getJSONObject(parent map[string]interface{}, path []string) (map[string]interface{}, bool) {
	if parent == nil || len(path) <= 0 {
		return nil, false
	}
	var result map[string]interface{}
	for i := 0; len(path) > i; i++ {
		resultInterface, ok := parent[path[i]]
		if !ok || resultInterface == nil {
			return nil, false
		}
		result, ok = resultInterface.(map[string]interface{})
		if !ok {
			return nil, false
		} else {
			parent = result
		}
	}
	return result, true
}
