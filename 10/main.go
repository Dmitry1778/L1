package main

import (
	"fmt"
)

func main() {
	var minimum = -30
	var maximum = 40
	step := 10
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var noName = make(map[int][]float32)

	for i := minimum; i <= maximum; i += step {
		minimum = i
		letterMinimum := i + 10
		for j := 0; j < len(temp); j++ {
			if minimum <= int(temp[j]) && letterMinimum >= int(temp[j]) {
				noName[minimum] = append(noName[minimum], temp[j])
			}
		}
	}
	fmt.Println(noName)
}

//20-29
//10-19
