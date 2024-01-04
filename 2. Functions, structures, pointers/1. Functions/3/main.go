package main

import "fmt"

func main() {
	fmt.Println(vote(10, 20, 30))
}
func vote(x int, y int, z int) int {
	if x == 0 && y == 0 && z == 0 {
		return 0
	} else if (x == 0 && y == 0 && z == 1) || (x == 0 && y == 1 && z == 0) || (x == 1 && y == 0 && z == 0) {
		return 0
	} else if (x == 0 && y == 1 && z == 1) || (x == 1 && y == 1 && z == 0) || (x == 1 && y == 0 && z == 1) {
		return 1
	} else {
		return 1
	}
}
