package utils

func Contains(array []string, key string) bool {
	for _, val := range array {
		if val == key {
			return true
		}
	}
	return false
}
