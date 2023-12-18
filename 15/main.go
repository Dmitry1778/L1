package main

import (
	"fmt"
	"strconv"
	"strings"
)

var justString string

func someFunc() {
	v, err := createHugeStrings(1 << 10) // строка будет иметь длину 1024 (2^ 10)
	if err != nil {
		panic(err)
	}
	//str := "123456789"
	//fmt.Println(str[:10])
	if len(v) >= 100 {
		justString = v[:100]
		fmt.Println(justString)
	} else {
		panic("<100")
	}
}

func createHugeString2(str int) (string string, err error) {
	var s strings.Builder
	for i := 0; i < 10000; i++ {
		_, err = s.WriteString(" Hello World!")
		if err != nil {
			panic(err)
		}
	}
	return s.String(), nil
}

// При создании большой строки может возникнуть ошибка времени выполнения: границы среза выходят за пределы диапазона [:100] длиной 50
func createHugeStrings(str int) (string string, err error) {
	newStr := strconv.Itoa(str)
	newStr = strings.Repeat("Hello", 20)
	return newStr, nil
}

func main() {
	someFunc()
}

//1. Возможна паника при выходе за границы массива при попытке присвоить
//justString значение v[:100]. Если длина v менее 100 элементов, то возникнет паника. index out of range
//
//2. В основной функции main нет обработки ошибок,
//которая может возникнуть в функции someFunc() при вызове createHugeString.
