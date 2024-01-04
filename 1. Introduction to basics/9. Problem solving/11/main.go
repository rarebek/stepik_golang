package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if (a == 1 || a%10 == 1) && a != 11 {
		fmt.Print(a, " korova")
	} else if (a == 3 || a == 4 || a == 5 || a%10 == 3 || a%10 == 4 || a%10 == 2) && a != 12 && a != 13 && a != 14 {
		fmt.Print(a, " korovy")
	} else {
		fmt.Print(a, " korov")
	}
}
