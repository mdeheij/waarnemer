package utils

//StringInSlice : There is no built-in operator to do it in Go. You need to iterate over the array.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
