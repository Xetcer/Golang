// https://golangs.org/oop
package main

import (
	"fmt"
	"math"
)

// coordinate координаты в градусах, минутах, секундах для сферы N/S/E/W
type coordinate struct {
	d, m, s float64
	h       rune
}

// location координаты в десятичных градусах
type location struct {
	pointName string
	lat, long float64
}

// world типизированная структура для хранения радиуса планет.
type world struct {
	wName  string
	radius float64
}

// distance типизированная структура, хрянящая расстояния между точками
type pointsDist struct {
	point1, point2 location // точка 1, точка 2, между которыми надо найти дистанцию
	cworld         world
	distance       float64
}

// calcDist метод типа poitDist заполняет структуру расстоянием между двуямя ее точками и ее миром
func (pd pointsDist) calcDist() {
	pd.distance = pd.cworld.distance(pd.point1, pd.point2)
	fmt.Printf("Расстояние между точкой %v и точкой %v на поверхности мира %v составляет %.2f км\n", pd.point1.pointName, pd.point2.pointName, pd.cworld.wName, pd.distance)
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

// newLocation конструктор типа location
func newLocation(pName string, lat, long coordinate) location {
	return location{pointName: pName, lat: lat.decimal(), long: long.decimal()}
}

// distance метод, который прикрепляется к типу world считает расстояние
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad переводит градусы в радианы
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// List12 листинг 12
func List12() {
	fmt.Println("Листинг 1-2 метод для структурированного типа")
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println("в градусах широта:", lat.decimal(), "Долгота:", long.decimal())
}

// List3 Листинг 3 использование конструктора типа
func List3() {
	fmt.Println("Листинг 3 конструктор типа")
	/*Функции в форме newType или NewType используются для создания значения указанного типа. */
	curiosity := newLocation("Curiosity", coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}

// List4 классы в goLang
func List4() {
	fmt.Println("Листинг 4 имитация классов в Go")
	var mars = world{radius: 3389.5}
	spirit := location{lat: -14.5684, long: 175.472636}
	opportunity := location{lat: -1.9462, long: 354.4734}
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("Расстояние составляет: %.2f km\n", dist)
}

// ControlWork12 контрольное задание #12
func ControlWork12() {
	fmt.Println("контрольное задание №1 создаем координаты марсоходов и выводим их в тесятичной системе")
	spirit := newLocation("Spirit", coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation("Opportunity", coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation("Curiosity", coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	inSight := newLocation("InSight", coordinate{4, 30, 0., 'N'}, coordinate{135, 54, 0.0, 'E'})
	fmt.Println(spirit)
	fmt.Println(opportunity)
	fmt.Println(curiosity)
	fmt.Println(inSight)
	fmt.Println("контрольное задание 2 расчитывает расстояния между марсоходами")
	var mars = world{wName: "Mars", radius: 3389.5}
	var spirOppDist = pointsDist{point1: spirit, point2: opportunity, cworld: mars}
	var spirCurDist = pointsDist{point1: spirit, point2: curiosity, cworld: mars}
	var spirInsDist = pointsDist{point1: spirit, point2: inSight, cworld: mars}
	var OppCurDist = pointsDist{point1: opportunity, point2: curiosity, cworld: mars}
	var OppInsDist = pointsDist{point1: opportunity, point2: inSight, cworld: mars}
	var CurInsDist = pointsDist{point1: curiosity, point2: inSight, cworld: mars}
	spirOppDist.calcDist()
	spirCurDist.calcDist()
	spirInsDist.calcDist()
	OppCurDist.calcDist()
	OppInsDist.calcDist()
	CurInsDist.calcDist()
	london := newLocation("London", coordinate{51, 30, 0.0, 'N'}, coordinate{0, 39, 0.0, 'W'})
	paris := newLocation("Paris", coordinate{48, 51, 0.0, 'N'}, coordinate{2, 21, 0.0, 'E'})
	var earth = world{wName: "Earth", radius: 6371.0}
	var lonParisDist = pointsDist{point1: london, point2: paris, cworld: earth}
	lonParisDist.calcDist()
	mountSharp := newLocation("mountSharp", coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0.0, 'E'})
	OlympusMons := newLocation("OlympusMons", coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0.0, 'E'})
	var sharpOlympusDist = pointsDist{point1: mountSharp, point2: OlympusMons, cworld: mars}
	sharpOlympusDist.calcDist()
}

func main() {
	fmt.Println("Структуры и методы Go")
	List12()
	List3()
	List4()
	ControlWork12()
}
