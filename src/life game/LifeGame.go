package main

import (
	"fmt"
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

func main() {
	fmt.Println("Здесь скоро будет игра имитирующая жизнь клеток")
	myUniverse := NewUniverse()
	myUniverse.Show()
}
