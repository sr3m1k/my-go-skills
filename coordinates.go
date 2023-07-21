package main

import (
	"fmt"
	"math"
)

func print(len float64) {
	fmt.Print("Distance between these coordinates: ", len)
}

func math_coor(XY float64) {
	len := math.Sqrt(XY)
	print(len)
}

func coord(x1, y1, x2, y2 float64) {
	XY := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	math_coor(XY)
}

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64

	fmt.Println("Write your coordinates")
	fmt.Print("Write a lat: ")
	fmt.Scan(&x1)
	fmt.Print("Write a long: ")
	fmt.Scan(&y1)

	fmt.Println("Write the coordinates of another location")
	fmt.Print("Write a lat: ")
	fmt.Scan(&x2)
	fmt.Print("Write a long: ")
	fmt.Scan(&y2)

	coord(x1, y1, x2, y2)

}
