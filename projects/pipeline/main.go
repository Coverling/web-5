package main

import (
	"fmt"
)

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	var lastValue string

	for currentValue := range inputStream {
		if currentValue != lastValue {
			outputStream <- currentValue
			lastValue = currentValue
		}
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		inputStream <- "hello"
		inputStream <- "hello"
		inputStream <- "world"
		inputStream <- "world"
		inputStream <- "!"
		inputStream <- "!"

		close(inputStream)
	}()

	for result := range outputStream {
		fmt.Println(result)
	}
}
