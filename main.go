package main

import "fmt"

func numberGenerator(total int) {
	for i := 0; i < total; i++ {
		fmt.Printf("Generating number %d\n", i)
	}
}

func printThreeNumber() {
	for i := 0; i < 4; i++ {
		fmt.Printf("I am poud to print: %d\n", i)
	}
}
func main() {
	printThreeNumber()
	numberGenerator(20)
}
