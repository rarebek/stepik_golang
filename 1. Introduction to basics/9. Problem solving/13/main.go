package main

import "fmt"

func main() {
	var first, second, third, limit, iter, check int
	var mylist []int
	check = 0
	fmt.Scan(&limit)
	first = 1
	second = 1
	for second <= limit {
		iter += 1
		third = first + second
		first = second
		second = third
		mylist = append(mylist, second)

	}
	for i := 0; i < len(mylist); i++ {
		if mylist[i] == limit {
			check = 1
			break
		}
	}
	if check == 1 {
		fmt.Print(iter + 1)
	} else {
		fmt.Print(-1)
	}

}
