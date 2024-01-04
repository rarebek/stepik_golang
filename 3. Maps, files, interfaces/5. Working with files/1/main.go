package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	var (
		sum int
	)
	for buf.Scan() {
		num, err := strconv.Atoi(buf.Text())
		if err != nil {
			return
		}
		sum += num
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
}
