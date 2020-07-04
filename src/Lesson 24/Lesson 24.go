// Урок 24 композиция и встраивание методов
// https://golangs.org/composition-and-forwarding

package main

import (
	"fmt"
	"math"
)

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.lat, l.long)
}

type report struct {
	sol         int
	temperature temperature
	location    location
}

// average метод который расчитывает среднее значение температур
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// average метод который вызывает метод из встроенной структуры
func (r report) average() celsius {
	return r.temperature.average()
}

type sol int

type reportl5 struct {
	sol
	location
	temperature
}

type world struct {
	wName  string
	radius float64
}

type gps struct {
	cLoc, nLoc location //текущее/следующее положение
	cWorld     world    // текущий мир
}

// rad переводит градусы в радианы
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance метод, который прикрепляется к типу world считает расстояние
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// distance возвращает расстояение между точками
func (g gps) distance() float64 {
	return g.cWorld.distance(g.cLoc, g.nLoc)
}

// message возвращает строку с расстоянием до конечной точки
func (g gps) message() string {
	return fmt.Sprintf("Расстояние до точки %v составляет %.2f км", g.nLoc.description(), g.distance())
}

type rover struct {
	gps
}

// days метод типа sol для определения разницы дней
func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

// List123 композиция структур Go
func List123() {
	fmt.Println("Листинг 1-2-3")

	bradbury := location{lat: -4.5895, long: 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report1 := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report1)
	fmt.Printf("a balmy %v °C\n", report1.temperature.high)
	fmt.Println("Среднее значение температуры:", report1.temperature.average())
	fmt.Println("Среднее значение температуры 2:", report1.average())
}

// List45 Встраивание методов Go
func List45() {
	fmt.Println("Листинг 4-5")
	report1 := report{sol: 15,
		location:    location{name: "Name", lat: -4.5895, long: 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Printf("average %v °C\n", report1.average())
	// можно напрямую вызывать методы или поля интегрированных типов структур
	report2 := reportl5{sol: 15}
	fmt.Println("Дни по полному пути до метода:", report2.sol.days(1446)) // полный путь до метода
	fmt.Println("Дни по интегрированному пути:", report2.days(1446))      // или интегрированный вариант пути
}

// controlWork контрольное задание
func controlWork() {
	fmt.Println("Контрольное задание")
	mars := world{wName: "Mars", radius: 3389.5}
	bradbary := location{name: "Bradbary Landing", lat: -4.5895, long: 137.4417}
	elysium := location{name: "Elysium Planitia", lat: 4.5, long: 135.9}
	curiosity := rover{gps{cLoc: bradbary, cWorld: mars, nLoc: elysium}}
	fmt.Println(curiosity.message())
}

func main() {
	fmt.Println("Урок 24 композиция и встраивание методов")
	List123()
	List45()
	controlWork()
}
