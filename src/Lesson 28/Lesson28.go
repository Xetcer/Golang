// Урок 28 значение nil в Golang
// https://golangs.org/nil
package main

import (
	"fmt"
	"sort"
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

// Защита методов в Go
type person struct {
	age int
}

// birthday метод, приемником является указатель типа person
func (p *person) birthday() {
	if p != nil { // защита метода.
		p.age++ // разыменование указателя nil
	}
}

// List3_4 Листинг 3-4 Защита методов в Golang
func List3_4() {
	var nobody *person
	fmt.Println("Листинг 3-4 защита метода в Go", nobody)
	nobody.birthday()
}

// Листинг 5-6 значение функции nil
// sortStrings
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

// Срезы nil Golang

// List7_8 листинг 7-8
func List7_8() {
	fmt.Println("Листинг 7-8")
	var suop []string
	fmt.Println(suop == nil)
	for _, ingredient := range suop {
		fmt.Println(ingredient)
	}

	fmt.Println(len(suop))
	suop = append(suop, "чеснок", "морковь", "капуста")
	fmt.Println(suop)

	soup := mirepoix(nil)
	fmt.Println(soup)
}

// mirepoix
func mirepoix(ingredients []string) []string {
	fmt.Println("Функция добавления ингредиентов")
	return append(ingredients, "чеснок", "морковь", "капуста")
}

// Карты в nil в Golang

// List9 листинг 9
func List9() {
	fmt.Println("листинг 9")
	var soup map[string]int
	fmt.Println(soup == nil)
	measurement, ok := soup["чеснок"]
	measurement = 1
	if ok {
		fmt.Println(measurement)
	} else {
		fmt.Println("Значения с ключом чеснок нет")
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}
}

// Интерфейсы nil

// List10_11_12 интерфейсы и nil
func List10_11_12() {
	fmt.Println("Интерфейсы nil")
	var v interface{}
	fmt.Printf("Выводим пустой интерфейс типа %T значением %v равен nil %v\n", v, v, v == nil)
	var p *int
	v = p
	fmt.Printf("выводим указатель на int в интерфейс типа %T значением %v равен nil %v\n", v, v, v == nil)
	fmt.Printf("Вывод типа и значения с помощью специального символа %#v\n", v)
}

// Альтернатива nil

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

// String выводит не установлено или число
func (n number) String() string {
	if !n.valid {
		return "не установлено"
	}
	return fmt.Sprintf("%d", n.value)
}

// List13 листинг 13
func List13() {
	fmt.Println("Листинг 13 альтернатива Nil")
	n := newNumber(42)
	fmt.Printf("Число %v\n", n)
	e := number{}
	fmt.Printf("Число %v\n", e)
}

//Контрольное задание

// тип предмет
type item struct {
	itemName string // название предмета
}

// тип рука
type hand struct {
	name     string // название руки
	itemName item   // предмет в руке
	equiped  bool   // занята или свободна
}

// тип персонаж
type character struct {
	name      string // имя персонажа
	leftHand  hand   // левая рука персонажа
	rightHand hand   // правая рука
}

//
func (c *character) init(cName string) {
	if c != nil {
		c.name = cName
		c.leftHand.name = "LEFT"
		c.rightHand.name = "RIGHT"
	}
}

// pickup метод поднять предмет
func (c *character) pickup(i *item, h *hand) {
	if c != nil && i != nil && h != nil {
		if h.equiped == false {
			h.equiped = true
			h.itemName = *i
			fmt.Printf("%v взял в %v руку %v\n", c.name, h.name, i.itemName)
		} else {
			fmt.Printf("%v не смог взять в %v руку %v потому что рука занята\n", c.name, h.name, i.itemName)
		}
	} else {
		fmt.Println("Не удалось найти предмет чтобы его поднять")
	}
}

// give метод передать предмет
func (c *character) give(to *character, fh *hand, toh *hand) {
	if c != nil && to != nil && fh != nil && toh != nil {
		if fh.equiped == true && toh.equiped == false {
			toh.itemName = fh.itemName
			toh.equiped = true
			fh.equiped = false
			fmt.Printf("%v передал из %v руки %v в %v руку %v\n", c.name, fh.name, fh.itemName, toh.name, to.name)
		} else if fh.equiped == false {
			fmt.Printf("%v рука %v пуста\n", fh.name, c.name)
		} else if toh.equiped == true {
			fmt.Printf("%v рука %v занята\n", toh.name, to.name)
		}
	} else {
		fmt.Println("Не произошло никаких действий")
	}
}

// ctrlWork контрольное задание
func ctrlWork() {
	fmt.Println("Контрольное задание!")
	var artur character
	artur.init("Артур")
	var knight character
	knight.init("рыцарь Иван")
	sword := &item{itemName: "Меч"}
	shield := &item{itemName: "Щит"}
	mg := &item{itemName: "Пулемет"}
	fmt.Printf("Поединок %v против %v\n", artur.name, knight.name)
	fmt.Printf("%v выбирает оружие для соперника\n", artur.name)
	artur.pickup(sword, &artur.rightHand)
	artur.pickup(shield, &artur.leftHand)
	artur.give(&knight, &artur.leftHand, &knight.leftHand)
	artur.give(&knight, &artur.leftHand, &knight.leftHand)
	artur.give(&knight, &artur.rightHand, &knight.leftHand)
	artur.give(&knight, &artur.rightHand, &knight.rightHand)
	fmt.Printf("%v выбирает оружие для себя\n", artur.name)
	artur.pickup(mg, &artur.rightHand)
}

func main() {
	fmt.Println("Значение nil в golang")
	List1_2()
	List3_4()
	food := []string{"чеснок", "морковь", "капуста"}
	sortStrings(food, nil)
	fmt.Println(food)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) }) // сортировка по возрастанию длины массива передаем указатель на функцию первого класса
	fmt.Println(food)
	List7_8()
	List9()
	List10_11_12()
	List13()
	ctrlWork()
}
