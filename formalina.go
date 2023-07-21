package main

import "fmt"

func main() {
	sum := 0
	for sum < 11 {
		switch i := isTest(sum); i {
		case 1:
			fmt.Println("sum == 2")
		case 2:
			fmt.Println("sum == 2")
		default:
			fmt.Println("you are pidr")

		}
		sum++
		fmt.Println(sum)
	}
}
func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 10 {
		return 10
	}
	return 0
}
