package main

import (
	"fmt"
	"sync"
)

func numberGenerator(total int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < total; i++ {
		fmt.Printf("Generating number %d\n", i*i)
		ch <- i * i
	}
}

func printNumber(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range ch {
		fmt.Printf("I am proud to print: %d\n", i)
	}
}

func main() {
	var ch chan int = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumber(ch, &wg)
	numberGenerator(20, ch, &wg)
	close(ch)

	fmt.Println("Waiting for goroutines...")
	wg.Wait()
	fmt.Println("...finished!")
}
