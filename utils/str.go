package utils

import "strings"

func StrFirstToUpper(str string) string {
	strArr := strings.Split(str, "_")
	var upperStr string

	for _, v := range strArr {
		item := []rune(v)
		if item[0] >= 97 && item[0] <= 122 {
			item[0] -= 32
		}
		upperStr += string(item)
	}

	return upperStr
}
