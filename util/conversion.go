package util

import "strconv"

// string 转 int 默认为0
func StringToIntDefault(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return result
}
