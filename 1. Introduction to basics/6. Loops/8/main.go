package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	var sa []int
	var sb []int
	var res []int
	for a > 0 {
		sa = append(sa, a%10)
		a /= 10
	}

	for b > 0 {
		sb = append(sb, b%10)
		b /= 10
	}

	for i := 0; i < len(sa); i++ {
		for j := 0; j < len(sb); j++ {
			if sa[i] == sb[j] {
				res = append(res, sa[i])
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}
