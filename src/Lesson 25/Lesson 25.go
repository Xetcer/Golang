// Урок 25 Интерфейсы Go
// https://golangs.org/interface
// Интерфейсы сконцентрированы на том, что тип делает, а не на сохраняемых им значениях.
// Методы выражают поведение предоставленного типа, поэтому интерфейсы объявляются с набором методов, которых тип должен удовлетворить
// Правило именования интерфейсов добавлять в конце название -er, например для метода talk - интерфейс talker
package main

import (
	"fmt"
	"strings"
)

// Переменная t может хранить значения любого типа если у него есть метод talk без аргументов и возвращает строку
var t interface {
	talk() string
}

// типизированный интерфейс talker
type talker interface {
	talk() string
}

type martian struct{}

// talk метод типа martian возвращает строку
func (m martian) talk() string {
	return "тук тук"
}

type laser int

type starship struct {
	laser
}

type rover struct {
	roverName string
	rptCnt    int
}

// talk метод типа laser возвращает строку
func (l laser) talk() string {
	return strings.Repeat("пиу ", int(l))
}

func (r rover) talk() string {
	return strings.Repeat("Врр ", r.rptCnt)
}

// List123 реализация листинга 1
func List123() {
	fmt.Println("Листинг 1-2-3")
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())

}

// List456 Листинг 4-5-6
func shoutL456(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

// list7 интерфейс и встраиваемые структуры
func list7() {
	fmt.Println("Листинг 7 интерфейсы и встраиваемые структуры")
	s := starship{laser: 3}
	fmt.Println(s.talk())
	shoutL456(s)
}

// ControlWork контрольное задание
func ControlWork() {
	fmt.Println("контрольное задание")
	myRover := rover{roverName: "Xetzer rover", rptCnt: 10}
	fmt.Printf("Создаем новый ровер %+v\n", myRover)
	fmt.Println("Звук моего ровера", myRover.talk())
	fmt.Println("Звук ровера через интерфейс:")
	shoutL456(myRover)
}

func main() {
	fmt.Println("Урок 25 интерфейсы Go")
	List123()
	fmt.Println("Листинг 4-5 ")
	shoutL456(martian{})
	shoutL456(laser(2))
	list7()
	ControlWork()
}
