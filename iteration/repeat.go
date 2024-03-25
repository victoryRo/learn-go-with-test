package iteration

import "strings"

func Repeat(symbol string, repeat int) string {
	var result string

	for i := 0; i < repeat; i++ {
		result += symbol
	}
	return result
}

func Comparision(first, second string) int {
	num := strings.Compare(first, second)
	return num
}
