package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Human struct {
	age    int
	wight  int
	height int
	name   string
	gender string
}

type Action struct {
	human Human
}

func (h Human) selectHeight() *int {
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

func (h Human) selectGender() *string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You male or female?")
	scanner.Scan()
	input := h.gender
	input = scanner.Text()
	return &input
}

func newAction(human Human) *Action {
	return &Action{human: human}
}

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
// Язык golang (GO) за 1 час. ООП - полное руководство.
