package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	r, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r = strings.TrimSuffix(r, "\n")
	mytime, _ := time.Parse("2006-01-02 15:04:05", r)
	if mytime.Hour() >= 13 {
		mytime = mytime.AddDate(0, 0, 1)
	}
	fmt.Println(mytime.Format("2006-01-02 15:04:05"))
}
