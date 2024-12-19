package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	// необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			work()
		}(wg)
	}
	wg.Wait()
}
