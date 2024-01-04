package main

import (
	"fmt"
	"time"
)

func main() {
	a := ""
	fmt.Scan(&a)
	mytime, _ := time.Parse(time.RFC3339, a)
	fmt.Println(mytime.Format(time.UnixDate))
}
