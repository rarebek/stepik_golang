package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	s1 := []rune(s)
	if len(s1) >= 5 {
		check := true
		for i := 0; i < len(s1); i++ {
			if !unicode.Is(unicode.Latin, s1[i]) && !unicode.IsDigit(s1[i]) {
				check = false
				break
			}
		}
		if check {
			fmt.Println("Ok")
		} else {
			fmt.Println("Wrong password")
		}

	} else {
		fmt.Println("Wrong password")
	}
}
