package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.ReplaceAll(s, ",", ".")
	nums := strings.Split(s, ";")
	n1 := strings.ReplaceAll(nums[0], " ", "")
	n2 := strings.ReplaceAll(nums[1], " ", "")
	num1, _ := strconv.ParseFloat(n1, 64)
	num2, _ := strconv.ParseFloat(n2, 64)
	ans := num1 / num2
	fmt.Printf("%.4f", ans)
}
