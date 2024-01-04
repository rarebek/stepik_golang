package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(adding("10", "40"))
}
func adding(s1, s2 string) int64 {
	var cleanStr1 strings.Builder
	for _, char := range s1 {
		if unicode.IsDigit(char) {
			cleanStr1.WriteRune(char)
		}
	}

	var cleanStr2 strings.Builder
	for _, char := range s2 {
		if unicode.IsDigit(char) {
			cleanStr2.WriteRune(char)
		}
	}
	num1, _ := strconv.ParseInt(cleanStr1.String(), 10, 64)
	num2, _ := strconv.ParseInt(cleanStr2.String(), 10, 64)
	return num1 + num2
}
