package main

func merge(base map[string]interface{}, over map[string]interface{}) map[string]interface{} {
	for key, value := range over {
		if _, ok := base[key]; !ok {
			// Key does not exist in base, simply assign
			base[key] = over[key]
		} else {
			switch v := value.(type) {
			case map[string]interface{}:
				switch b := base[key].(type) {
				case map[string]interface{}:
					base[key] = merge(b, v)
				default:
					base[key] = v
				}
			default:
				base[key] = value
			}
		}
	}
	return base
}
