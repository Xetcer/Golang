//Урок 29 ошибки в Go
// https://golangs.org/errors
package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
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

// Новые ошибки в Golang листинг 7-10
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

// Причины каждой ошибки в Go листинг 11-13
// Объявление ошибки
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// validDigit проверка что число не выходит за диапазон
func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

// Set1 метод который возвращает объявленную заранее ошибку
func (g *Grid) Set1(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrBounds
	}
	g[row][column] = digit
	return nil
}

// Листинг 14 настраиваемые типы ошибок в Go
// интерфейс который можно использовать для описания любых ошибок
// type error interface {
// 	Error() string
// }

// Листинг 15 Множество ошибок в Go

// SudokuError массив ошибок
type SudokuError []error

//Error возвращает одну или несколько ошибок через запятые
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error()) // конвертирует ошибки в строки
	}
	return strings.Join(s, ", ")
}

// Set2 проверка на множество ошибок
func (g *Grid) Set2(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

// ctrlWork контрольное задание
func ctrlWork() {
	fmt.Println("Контрольное задание, парсинг URL")
	u, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err) // Выводит ошибку
		fmt.Printf("%#v\n", err)
		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}
		// os.Exit(1)
	}
	fmt.Println(u)
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
	}
	// Объявление ошибок пример
	fmt.Println("Листинг 11- 13 объявление своих ошибок в Go")
	err1 := g.Set1(0, 0, 15)
	if err1 != nil {
		switch err1 {
		case ErrBounds, ErrDigit:
			fmt.Println("Возникли ошибки которые были объявлены")
		default:
			fmt.Println(err1)
		}
	}

	fmt.Println("Листинг 15-16 множество ошибок в Go")
	err2 := g.Set2(0, 0, 15)
	if err2 != nil {
		switch err2 {
		case ErrBounds, ErrDigit:
			fmt.Println("Возникли ошибки которые были объявлены")
		default:
			fmt.Println(err2)
		}
	}

	fmt.Println("Листинг 17 утверждение ошибок в Go")
	err3 := g.Set2(10, 0, 15)
	if err != nil {
		if errs, ok := err3.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	// os.Exit(1)
	ctrlWork()
	//func анонимная функция
	fmt.Println("Листинг 18 тонкости работы panic в Go")
	defer func() {
		if e := recover(); e != nil { // Приходим в себя после panic
			fmt.Println(e) // Выводит: Я забыл свое полотенце
		}
	}()
	panic("Я забыл свое полотенце") // приводит к panic

}
