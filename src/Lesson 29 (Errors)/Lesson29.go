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

// List2Proverbs функция для обработки и возвращения ошибок при работе с файлом
func List2Proverbs(name string) error {
	fmt.Println("Листинг 2 создание файла и обработка ошибок")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just chek errors, handle them gracefully")
	f.Close()
	return err
}

func main() {
	fmt.Println("Ошибки в Go")
	List1()
	List2Proverbs("HelloWorld")
}
