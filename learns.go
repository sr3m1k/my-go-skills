package main

import "fmt"

func test() (a, b, c string) {
	a = "My"
	b = "name is"
	fmt.Scan(&c)
	return a, b, c
}

func main() {
	s, s2, s3 := test()
	fmt.Println(s, s2, s3)
}
