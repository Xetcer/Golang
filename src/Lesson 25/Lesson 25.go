// Урок 25 Интерфейсы Go
// https://golangs.org/interface
// Интерфейсы сконцентрированы на том, что тип делает, а не на сохраняемых им значениях.
// Методы выражают поведение предоставленного типа, поэтому интерфейсы объявляются с набором методов, которых тип должен удовлетворить
// Правило именования интерфейсов добавлять в конце название -er, например для метода talk - интерфейс talker
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
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

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate возвращает выдуманное измерение времени
func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

type sol int

// YearDay Получение числа дней в годах
func (s sol) YearDay() int {
	return int(s % 668) // Марсианский год состоит из 668 дней
}

// Hour часы на марсе для реализации интерфейса stardater
func (s sol) Hour() int {
	return 0 // Неизвестный час
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

// list8_11 Листинг с примерами интерфейсов Go
func list8_11() {
	fmt.Println("Листинг 8-11 примеры интерфейсов")
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1442)
	fmt.Printf("%.1f С Днем рождения\n", stardate(s))
}

// location с широтой и долготой в десятичных градусах
type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// List12 листинг 12
func List12() {
	fmt.Println("Листинг 12 удовлетворение требований интерфейса Go")
	fmt.Println("Используем интерфейс stringer пакета fmt в для нашего метода")
	curiosity := location{lat: -4.5895, long: 137.4417}
	fmt.Sprintln(curiosity)
}

// coordinate координаты в минутах
type coordinate struct {
	d, m, s float64
	h       rune
}

// String метод для интерфейса stringer пакета fmt
func (c coordinate) String() string {
	return fmt.Sprintf("%.1f°%.1f'%.1f\" %v", c.d, c.m, c.s, c.h)
}

// location1 координаты
type location1 struct {
	lat, long coordinate
}

// String метод для интерфейса stringer пакета fmt
func (l location1) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// list12CtlWork контрольное задание листинга 12
func list12CtlWork() {
	fmt.Println("Контрольное задание листинга 12")
	elisiym := location1{
		lat:  coordinate{d: 4, m: 30, s: 0.0, h: 'N'},
		long: coordinate{d: 135, m: 54, s: 0.0, h: 'E'},
	}
	fmt.Println("Elysium Planitia is at", elisiym)
}

// Итоговое контрольное заданиеъ
// decimal перевод в десятичные градусы из DMS
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// MarshalJSON метод для преобразования в формат json
func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

type locationCW struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

// mainCtrlWork контрольное задание
func mainCtrlWork() {
	fmt.Println("Итоговое задание")
	elisiym := locationCW{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	bytes, err := json.MarshalIndent(elisiym, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}

func main() {
	fmt.Println("Урок 25 интерфейсы Go")
	List123()
	fmt.Println("Листинг 4-5 ")
	shoutL456(martian{})
	shoutL456(laser(2))
	list7()
	ControlWork()
	list8_11()
	List12()
	list12CtlWork()
	mainCtrlWork()
}
