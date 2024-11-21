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

func printNumber(index int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range ch {
		fmt.Printf("%v: I am proud to print: %d\n", index, i)
	}
}

func main() {
	var ch chan int = make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printNumber(i, ch, &wg)
	}

	numberGenerator(2000, ch, &wg)
	close(ch)

	fmt.Println("Waiting for goroutines...")
	wg.Wait()
	fmt.Println("...finished!")
}
