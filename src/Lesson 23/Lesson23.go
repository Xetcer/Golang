// https://golangs.org/oop
package main

import "fmt"

// coordinate координаты в градусах, минутах, секундах для сферы N/S/E/W
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal переводит из координат DMS в градусы с десятичными значениями
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// List12 листинг 12
func List12() {
	fmt.Println("Листинг 1-2 метод для структурированного типа")
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println("в градусах широта:", lat.decimal(), "Долгота:", long.decimal())
}

func main() {
	fmt.Println("Структуры и методы Go")
	List12()
}
