package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(num uint) uint {
		n := strconv.Itoa(int(num))
		nw := ""
		for _, v := range n {
			a, _ := strconv.Atoi(string(v))
			if a%2 == 0 && a > 0 {
				nw += string(v)
			}
		}
		ans, _ := strconv.Atoi(nw)
		if ans == 0 {
			return 100
		} else {
			return uint(ans)
		}
	}
	fmt.Println(fn(10))

}
