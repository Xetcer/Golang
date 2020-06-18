// Ссылка на урок https://golangs.org/struct
package main

import "fmt"

// List1 описание функций листинга 1
func List1() {
	fmt.Println("Листинг 1 объявление структуры")
	var curiosity struct {
		lat  float64 // широта
		long float64 // долгота
	}
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}

// List2 описание функций листинга 2
func List2() {
	fmt.Println("Листинг 2 объявляем типизированную структуру")
	type location struct { // Типизированная структура
		lat  float64
		long float64
	}

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734
	fmt.Println(spirit, opportunity)
}

// List345 инициализация структур через композитный литерал
func List345() {
	fmt.Println("Листинг 3-4-5 инициализация структуру через строковые литералы")
	type location struct {
		lat  float64
		long float64
	}
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)
	// композитный литерал задает значения переменны по их порядку в структуре если не указаны сами переменные
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)

	// вывод названий полей в консоли с использованием %+v
	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%+v\n", curiosity)
}

// List6 копирование структуру листинг 6
func List6() {
	fmt.Println("Листинг 6 Копирование стурктур")
	type location struct {
		lat  float64
		long float64
	}
	bradbary := location{lat: -4.8595, long: 137.4414}
	curiosity := bradbary
	curiosity.long += 0.0106

	fmt.Printf("Координты кратера бредбери %+v, координаты марсохода %+v\n", bradbary, curiosity)
}

func main() {
	fmt.Println("Структуры Go и экспорт структур Json")
	List1()
	List2()
	List345()
	List6()
}
