package main

import "fmt"

// Point Точки x и y
type Point struct {
	x int
	y int
}

func (p Point) searchDistance() int {
	var result int
	if p.x >= p.y { // Если х больше то вычитакем x из y тем самым находим расстояние
		result = p.x - p.y
		return result
	} else {
		result = p.y - p.x // Иначе
		return result
	}
}

// Функция-конструктор для указания точек(значений)
func newSetParameters() Point {
	return Point{23, 2}
}
func main() {
	fmt.Println(newSetParameters().searchDistance())
}
