package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	time.Sleep(1 * time.Second)
}

func work() {
	fmt.Println("work")
}
