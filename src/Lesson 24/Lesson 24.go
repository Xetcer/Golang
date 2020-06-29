// Урок 24 композиция и встраивание методов
// https://golangs.org/composition-and-forwarding

package main

import "fmt"

type celsius float64

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
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

// List1 композиция структур Go
func List1() {
	fmt.Println("Листинг 1-2")

	bradbury := location{lat: -4.5895, long: 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report1 := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report1)
	fmt.Printf("a balmy %v °C\n", report1.temperature.high)
	fmt.Println("Среднее значение температуры:", report1.temperature.average())
	fmt.Println("Среднее значение температуры 2:", report1.average())
}

func main() {
	fmt.Println("Урок 24 композиция и встраивание методов")
	List1()
}
