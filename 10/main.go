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
	// цикл температуры от -30 до 40 с шагом 10
	for i := minimum; i <= maximum; i += step {
		minimum = i             // минимальная температура
		letterMinimum := i + 10 // следущая температура
		for j := 0; j < len(temp); j++ {
			// если температура минимальная температура меньше или равна температурным значением И следущая минимальная температура больше или равна температурным значение
			if minimum <= int(temp[j]) && letterMinimum >= int(temp[j]) {
				noName[minimum] = append(noName[minimum], temp[j]) // то записываем температуру в карту минимальную температуру и все температурные значения
			}
		}
	}
	fmt.Println(noName)
}
