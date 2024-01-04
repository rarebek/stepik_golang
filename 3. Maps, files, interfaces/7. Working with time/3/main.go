package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	a = strings.TrimSuffix(a, "\n")
	b := strings.Split(a, ",")
	ftime := b[0]
	stime := b[1]
	var ans time.Duration
	firstTime, _ := time.Parse("02.01.2006 15:04:05", ftime)
	secondTime, _ := time.Parse("02.01.2006 15:04:05", stime)
	if firstTime.After(secondTime) {
		ans = time.Since(secondTime) - time.Since(firstTime)
	} else {
		ans = time.Since(firstTime) - time.Since(secondTime)
	}
	fmt.Println(ans.Round(time.Second))

}
