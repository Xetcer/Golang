//Урок 29 ошибки в Go
// https://golangs.org/errors
package main

import (
	"errors"
	"fmt"
	"io"
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

// defer ключевое слово, с помощью которого можно убедиться в том что файл закрыт правильно

// List4 Листинг 4
func List4(name string) error {
	fmt.Println("Листинг 4 ключевое слово defer")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")
	return err
}

// Листинг 5 обработка ошибок с сохранением в файле ошибок и пропуском последующей записи если была ошибка

type safeWriter struct {
	w   io.Writer
	err error // место для хранения первой ошибки
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // пропускаем ошибку если уже были ошибки
	}

	_, sw.err = fmt.Fprintln(sw.w, s) // записать строку и хранить ошибку
}

// proverbsList5 красивая обработка ошибок с пропуском повторов
func proverbsList5(name string) error {
	fmt.Println("Листинг 5")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	return sw.err // Возвращает ошибку в случае ее возникновения
}

// Новые ошибки в Golang
const rows, columns = 9, 9

//Grid сетка судоку
type Grid [rows][columns]int8

// Set проверка на ввод числа в сетку поля 9 на 9
func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of bounds")
	}
	g[row][column] = digit
	return nil
}

// inBounds позволяет убедиться что мы не выходим за пределы сетки 9 на 9
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func main() {
	fmt.Println("Ошибки в Go")
	List1()
	List2Proverbs("HelloWorld")
	List4("HelloWorld")
	proverbsList5("HelloWorld")
	fmt.Println("Листинг 7-10 новые ошибки в Go")
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		fmt.Printf("Ошибка: %v. \n", err)
		os.Exit(1)
	}
}
