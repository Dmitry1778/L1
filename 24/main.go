package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) searchDistance() int {
	var result int
	if p.x >= p.y {
		result = p.x - p.y
		return result
	} else {
		result = p.y - p.x
		return result
	}
}

func newSetParameters() Point {
	return Point{23, 2}
}
func main() {
	fmt.Println(newSetParameters().searchDistance())
}
