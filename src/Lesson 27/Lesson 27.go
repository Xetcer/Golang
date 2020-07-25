// Урок 27 Указатели в Go https://golangs.org/pointers
package main

import "fmt"

// List12 листинг 1-2
func List12() {
	fmt.Println("Листинг 1 указатели и адрес в памяти")
	variable := 42
	fmt.Println("Адрес в памяти переменной", &variable)
	adress := &variable
	fmt.Println("Значение переменной:", *adress)
	fmt.Println(*(&variable))
}

// List34 листинг 3-4
func List34() {
	fmt.Println("Листинг 3-4")
	answer := 42
	address := &answer
	fmt.Printf("Тип %T\n", address)
	canada := "Canada"
	var home *string
	fmt.Printf("home is a %T\n", home)
	home = &canada
	fmt.Println(*home)
}

func main() {
	fmt.Println("Урок 27 указатели в Go")
	List12()
	List34()
}
