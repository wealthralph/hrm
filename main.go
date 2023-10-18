package main

import (
	"fmt"
	"math"
)

func add(x, y, z float64) float64 {
	r := x + y
	return math.Pow(r, z)
}

func main() {

	ralph := Address{"ralph", 12, "development", "engineer", "nigeria"}
	ralph.GetStatus()
	sum := add(2, 3, 7)
	fmt.Println(sum)
}

type Address struct {
	Name    string
	Age     int
	Hobby   string
	Job     string
	country string
}

func (a Address) GetStatus() {
	fmt.Println(`My status is married`, a.Hobby)
}
