package main

import (
	"fmt"
	"sync"
)

func numberGenerator(total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < total; i++ {
		fmt.Printf("Generating number %d\n", i)
	}
}

func printThreeNumber(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		fmt.Printf("I am poud to print: %d\n", i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printThreeNumber(&wg)
	go numberGenerator(2000, &wg)

	fmt.Println("Waiting for goroutines...")
	wg.Wait()
	fmt.Println("...finished!")
}
