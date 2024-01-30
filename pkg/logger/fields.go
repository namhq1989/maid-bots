package logger

func mergeFields(fields ...Fields) Fields {
	result := make(map[string]interface{})
	for _, m := range fields {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
