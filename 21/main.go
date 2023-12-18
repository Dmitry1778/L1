package main

import "fmt"

type Target interface {
	Operation()
}

// класс конкретного адаптера
type ConcreteAdapter struct {
	*Adaptee
}

// реализация метода интерфейса, реализующего вызов адаптируемого класса
func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

// адаптируемый класс
type Adaptee struct {
}

// Метод адаптируемого класса, который нужно вызвать где-то
func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}

// конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &ConcreteAdapter{adaptee}
}

// основной метод для демонстрации
func main() {
	fmt.Println("\nAdapter demo:\n")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}

//https://habr.com/ru/articles/765468/
//https://www.youtube.com/watch?v=6xDBbYe11HQ
//https://www.youtube.com/watch?v=o9sCFOv-uKE
