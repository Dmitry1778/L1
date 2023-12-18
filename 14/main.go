package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Hello"
	num := 42
	float := 42.2
	boolean := true
	ch := make(chan int)

	fmt.Println(typeOf(str))
	fmt.Println(typeOf(num))
	fmt.Println(typeOf(boolean))
	fmt.Println(typeOf(ch))

	typeOfSwitch(float)
	typeOfSwitch(num)
	typeOfSwitch(boolean)
}

func typeOf(v interface{}) interface{} {
	return reflect.TypeOf(v)
}

func typeOfSwitch(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
	return
}
