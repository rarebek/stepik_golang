package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a/100 == a/10%10 || a/10%10 == a%10 || a/100 == a%10 {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
