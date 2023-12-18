package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Human cтруктура с произвольным набором полей
type Human struct {
	age    int
	wight  int
	height int
	name   string
	gender string
}

// Action - это структура для реализации встраивания методов от структуры Human
// Она имеет все те же поля что и Human и все теже методы
type Action struct {
	human Human
}

// selectHeight метод определенный для типа структуры Human.
func (h *Human) selectHeight() *int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("how tall are you?")
	scanner.Scan()
	input := h.height
	newInput := strconv.Itoa(input)
	newInput = scanner.Text()
	newInputInt, err := strconv.Atoi(newInput)
	if err != nil {
		panic(err.Error())
	}
	return &newInputInt
}

// selectGender метод определенный для типа структуры Human.
func (h *Human) selectGender() *string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You male or female?")
	scanner.Scan()
	input := h.gender
	input = scanner.Text()
	return &input
}

// Функция-конструктор которая создает объект
func newAction(human Human) *Action {
	return &Action{human: human}
}

// Action реализация c произвольными данными
func (a Action) Action() *Action {
	a.human.age = 22
	a.human.name = "Dmitry"
	a.human.wight = 70
	a.human.gender = *a.human.selectGender()
	a.human.height = *a.human.selectHeight()
	return &a
}

func main() {
	out := newAction(Human{}).Action()
	fmt.Println(*out)
}

// https://www.youtube.com/watch?v=VVbqgRWDgzk&t=152s
