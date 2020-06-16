package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80 // ширина мира в клетках
	height = 15 // Высота мира в клетках
)

// Universe тип двумерного мира клеток
type Universe [][]bool

// NewUniverse создает новую вселенную
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Show метод типа Unverse для вывода вселенной на экран
func (u Universe) Show() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch u[y][x] {
			case true:
				fmt.Print("*")
			case false:
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

// Seed во вселенной живые клетки - 25% от ее вместимости
func (u Universe) Seed() {
	for i := 0; i < (width*height)/4; i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

//Alive возвращает жива ли клета по координатам
func (u Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

//Neighbors возвращает число соседей у данной клетки
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

// Next состояние вселенной на следующем шаге
func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

// Step определяет состояние новой вселенной по старой
func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.Next(x, y)
		}
	}
}

func main() {
	fmt.Println("Здесь скоро будет игра имитирующая жизнь клеток")
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 10; i++ {
		Step(a, b)
		fmt.Print("\x0c") // очищаем экран
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}
}
