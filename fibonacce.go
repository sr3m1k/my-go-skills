package main

import "fmt"

func forwword(s []int) {
	for _, el := range s {
		fmt.Println(el)

	}

}

func main() {
	var value int
	fmt.Scan(&value)
	slice := []int{0, 1}
	for value > len(slice) {
		index := len(slice) - 1
		sum := slice[index] + slice[index-1]
		slice = append(slice, sum)
	}
	fmt.Println("________")
	forwword(slice)
}
