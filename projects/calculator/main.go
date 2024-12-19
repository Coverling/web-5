package main

import "fmt"

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	var val int

	go func() {
		for {
			select {
			case val = <-firstChan:
				ch <- val * val
			case val = <-secondChan:
				ch <- val * 3
			case <-stopChan:
				close(ch)
				return
			}
		}
	}()
	return ch
}
func main() {
	// здесь должен быть код для проверки правильности работы функции calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
	//type Str struct{}

	firstChan, secondChan := make(chan int), make(chan int)
	stopChan := make(chan struct{})

	resultChan := calculator(firstChan, secondChan, stopChan)

	firstChan <- 10
	fmt.Println(<-resultChan)

	secondChan <- 15
	fmt.Println(<-resultChan)

	close(stopChan)
	fmt.Println(<-resultChan)
}
