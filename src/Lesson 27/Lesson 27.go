// Урок 27 Указатели в Go https://golangs.org/pointers
package main

import "fmt"

// List1 листинг 1
func List1() {
	fmt.Println("Листинг 1 указатели и адрес в памяти")
	variable := 42
	fmt.Println("Адрес в памяти переменной", &variable)
	adress := &variable
	fmt.Println("Значение переменной:", *adress)
	fmt.Println(*(&variable))
}

func main() {
	fmt.Println("Урок 27 указатели в Go")
	List1()
}
