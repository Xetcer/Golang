//Урок 29 ошибки в Go
// https://golangs.org/errors
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// List1 Листинг 1
func List1() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	fmt.Println("Ошибки в Go")
	List1()
}
