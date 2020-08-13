// Урок 27 Указатели в Go https://golangs.org/pointers
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

// List5 листинг 5
func List5() {
	fmt.Println("Листинг 5")
	var administrator *string
	scolese := "Кристофер Сколезе"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Чарльз Болден"
	administrator = &bolden
	fmt.Println(*administrator)
}

// List67 листинг 6
func List67() {
	fmt.Println("Листинг 6-7 структуры и указатели")
	type person struct {
		name, superpower string
		age              int
	}
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}

// List8 листинг 8 указатели и массивы
func List8() {
	fmt.Println("Листинг 8 указатели и массивы")
	superpowers := &[3]string{"Полет", "невидимость", "супер сила"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
}

// Указатели в функции
type person struct {
	name, superpower string
	age              int
}

// List9birthday принимает в аргумента указатель типа person
func List9birthday(p *person) {
	p.age++
}

// List10 указатели в функциях
func List10() {
	fmt.Println("Указатели в функциях")
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	fmt.Printf("До вызова функции с указателем: %+v\n", rebecca)
	List9birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)
}

// birthday метод приемник указателя типа person
func (p *person) birthday() {
	p.age++
}

// List11_14 Листинг 11-14
func List11_14() {
	fmt.Println("Листинг 11-14 приемники указателя в методе")
	terry := &person{
		name:       "Terry",
		superpower: "fear",
		age:        15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)
	nathan := person{
		name:       "Nathan",
		superpower: "eat",
		age:        17,
	}
	nathan.birthday()
	fmt.Printf("%+v\n", nathan)
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)
	fmt.Println(day.Format(layout))
	fmt.Println(tomorrow.Format(layout))
}

type stats struct {
	level             int
	endurance, health int
}

// List15LevelUp увеличение характеристик
func List15LevelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

// List16 листинг 16 указатели на внутреннюю структуру
func List16() {
	type character struct {
		name  string
		stats stats
	}
	fmt.Println("Листинг 15-16 внутренние указатели на структуры")
	player := character{name: "Gumboldt"}
	List15LevelUp(&player.stats)
	fmt.Printf("%+v\n", player)
}

// List17Reset сброс значений массива по указателю
func List17Reset(board *[8][8]rune) {
	fmt.Println("Листинг 17 управление элементами массива по указателю")
	board[0][0] = 'r'
	//...
}

// List18 срезы - указатели на массив
func List18(planets *[]string) {
	*planets = (*planets)[0:8]
}

// интерфейс в Golang
type talker interface {
	talk() string
}

// shout перевод в верхний регистр интерфейса
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

// talk метод типа martian возвращающий строку
func (m martian) talk() string {
	return "Тук тук"
}

// List19 листинг 19
func List19() {
	fmt.Println("Листинг 19 интерфейсы и указатели")
	shout(martian{})
	shout(&martian{})
}

type laser int

// talk метод типа laser
func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

// List20 листинг 20
func List20() {
	fmt.Println("Листинг 20")
	pew := laser(2)
	shout(&pew)
	fmt.Printf("тип переменной pew %T", pew)
}

// Итоговое задание

type coord struct {
	x, y int
}

// up вверх на 1
func (c *coord) up() {
	fmt.Println("вверх")
	c.y++
}

// down вниз на 1
func (c *coord) down() {
	fmt.Println("вниз")
	c.y--
}

// left влево на 1
func (c *coord) left() {
	fmt.Println("влево")
	c.x--
}

// right вправо на 1
func (c *coord) right() {
	fmt.Println("вправо")
	c.x++
}

// getCoord получить координаты черепахи
func (c *coord) getCoord() {
	fmt.Println("координаты", c)
}

// ctrlWork контрольное задание, движение черепахи
func ctrlWork() {
	fmt.Println("контрольное задание")
	turtleCoord := coord{x: 0, y: 0}
	turtleCoord.getCoord()
	for i := 0; i < 20; i++ {
		switch rand.Intn(4) {
		case 0:
			turtleCoord.down()
		case 1:
			turtleCoord.up()
		case 2:
			turtleCoord.left()
		case 3:
			turtleCoord.right()
		}
	}
	turtleCoord.getCoord()
}

func main() {
	fmt.Println("Урок 27 указатели в Go")
	List12()
	List34()
	List5()
	List67()
	List8()
	List10()
	List11_14()
	List16()
	var board [8][8]rune
	List17Reset(&board)
	fmt.Printf("%c", board[0][0])
	fmt.Println("Листинг 18")
	planets := []string{"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	List18(&planets)
	fmt.Println(planets)
	List19()
	List20()
	ctrlWork()
}
