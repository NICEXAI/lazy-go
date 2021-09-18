package util

// Include Is in the array
func Include(name string, list []string) bool {
	for _, v := range list {
		if name == v {
			return true
		}
	}
	return false
}
