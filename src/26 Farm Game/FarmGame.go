// https://golangs.org/silmulator-ferma урок 26 симулятор ферма
package main

import (
	"fmt"
	"math/rand"
)

type cow struct {
	name string
}

// Respond Возвращает название коровы
func (c cow) respond() string {
	return c.name
}

// eat возвращает то что сейчас ест корова
func (c cow) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "траву"
	default:
		return "сено"
	}
}

// move текущий процесс движения коровы
func (c cow) move() string {
	switch rand.Intn(3) {
	case 0:
		return "стоит"
	case 1:
		return "идет"
	default:
		return "лежит"
	}
}

type dog struct {
	name string
}

// respond возвращаем имя собаки
func (d dog) respond() string {
	return d.name
}

// eat возвращает то что сейчас ест собака
func (d dog) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "кость"
	case 1:
		return "мясо"
	default:
		return "сухой корм"
	}
}

// move возвращает текущее действие собаки
func (d dog) move() string {
	switch rand.Intn(6) {
	case 0:
		return "прыгает"
	case 1:
		return "машет хвостом"
	case 2:
		return "бежит"
	case 3:
		return "лежит"
	case 4:
		return "спит"
	case 5:
		return "воет на солнце"
	}
}

// animal Интерфейс, для отклика животного. Возвращает название животного
type animal interface {
	respond() string
	eat() string
	move() string
}

func main() {
	fmt.Println("Марсианская ферма")
}
