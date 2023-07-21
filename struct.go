package main

import "fmt"

type user struct {
	name  string
	age   int
	pass  string
	score []int
}

func (u user) highScore() int {
	high := 0
	for _, sc := range u.score {
		if high < sc {
			high = sc
		}
	}
	return high
}

func main() {
	user := user{"artem", 18, "Ki+1350P1f", []int{23, 6, 75, 12, 11, 9}}

	fmt.Println(user.highScore())
}
