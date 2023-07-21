package main

import "fmt"

func testi(newFunc func(int) int) {
	fmt.Println(25)
}

func main() {
	kvadr := func(x int) int {
		return x * x
	}
	testi(kvadr)
}
