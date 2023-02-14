package main

import (
	"fmt"
	"sync"
)

// chan 读写
// chan <- 只写
// <-chan 只读
func worker(wg *sync.WaitGroup) {

	for {
		select {
		case num, ok := <-ch:
			if !ok {
				wg.Done()
				return
			}
			sum = sum + num

		}

	}

}

func producer() {
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

var ch = make(chan int)
var sum = 0

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)

	go worker(&wg)

	go producer()

	wg.Wait()

	fmt.Printf("sum is %d", sum)

}
