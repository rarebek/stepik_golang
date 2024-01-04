package main

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	prevValue := ""
	for v := range inputStream {
		if v != prevValue {
			outputStream <- v
			prevValue = v
		}
	}
	close(outputStream)
}

func main() {
	input := make(chan string)
	output := make(chan string)

	go removeDuplicates(input, output)

	data := []string{"a", "b", "b", "c", "c", "c", "a", "b"}

	go func() {
		for _, d := range data {
			input <- d
		}
		close(input)
	}()

	for result := range output {
		fmt.Println(result)
	}
}
