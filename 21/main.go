package main

import "fmt"

type Target interface {
	Operation()
}

// ConcreteAdapter класс конкретного адаптера
type ConcreteAdapter struct {
	*Adaptee
}

// Operation реализация метода интерфейса, реализующего вызов адаптируемого класса
func (adapter *ConcreteAdapter) Operation() {
	adapter.AdaptedOperation()
}

// Adaptee адаптируемый класс
type Adaptee struct {
}

// AdaptedOperation Метод адаптируемого класса, который нужно вызвать где-то
func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}

// NewAdapter конструктор адаптера
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
