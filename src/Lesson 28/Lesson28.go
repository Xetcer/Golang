// Урок 28 значение nil в Golang
// https://golangs.org/nil
package main

import (
	"fmt"
)

// List1_2 Листинг 1-2 проверка указателей на занчение
func List1_2() {
	var nowhere *int
	fmt.Println("адрес указателя:", nowhere)
	// fmt.Println("значение указателя приведет к ошибке:", *nowhere)
	if nowhere != nil { // Проверка на наличие значения по указателю
		fmt.Println("значение указателя:", *nowhere)
	}
}

func main() {
	fmt.Println("Значение nil в golang")
	List1_2()

}
