package main

import "fmt"

func main() {
	defer fmt.Println("  pidr")
	defer fmt.Println("  are")
	defer fmt.Println("  You")
	defer fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")

}
