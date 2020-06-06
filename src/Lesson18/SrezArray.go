package main

import (
	"fmt"
	"sort"
	"strings"
)

// Planets тип среза  для домашнего задания
type Planets []string

// Terraform метод типа Planets, который добавляет ко всем элементам среза p строку S в начало элемента
func (p Planets) Terraform(s string) {
	for i := range p {
		p[i] = s + p[i]
	}
}

// List1 реализация листинга 1 урока
func List1(planets []string) {
	fmt.Println("листинг 1, в качестве параметра получен срез массива размерностью", len(planets), "элементов")
	terrestrial := planets[0:4] // Срез земная группа
	gasGiants := planets[4:6]   // Срез газовые гиганты
	iceGiants := planets[6:8]   // Срез ледяные гиганты
	fmt.Println(terrestrial, gasGiants, iceGiants)
	// Присвоение элементу среза нового значения вызовет изменение в основном массиве
	gasGiants[0] = "Изменено"
	fmt.Println(planets)
}

// List2 срезы со значениями индекса по умолчанию
func List2(planets []string) {
	// Пропуск первого индекса - равнозначен указанию первого элемента
	// Пропуск второго индекса - равнозначен указанию последнего элемента массива
	fmt.Println("листинг 2")
	terrestrial := planets[:4] // аналогично 0:4
	gasGiants := planets[4:6]
	iceGiants := planets[6:] // аналогично 6:8
	allPlanets := planets[:] // аналогично 0:8
	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(allPlanets)
}

// List3 композитные литералы для срезов
func List3() {
	fmt.Println("листинг 3")
	// базовый массив
	dwarfArray := [...]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	// Срез массива
	dwarfSlice := dwarfArray[:]
	// Объявление среза напрямую через композитный литерал
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Printf("тип базового массива: %T,\n тип его среза %T,\n тип среза композитного литерала %T\n", dwarfArray, dwarfSlice, dwarfs)
}

// List4Hyperspace приемущества среза массива
func List4Hyperspace(worlds []string) {
	fmt.Println("Листинг 4 работа со срезами и удаление пробелов в названиях элементов исходного массива")
	fmt.Println("Исходный срез:", worlds)
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
	fmt.Println("срез после удаления пробелов", worlds)
	fmt.Println("Объединим миры:", strings.Join(worlds, ""))
}

// List5 срезы с методами
func List5() {
	fmt.Println("Листинг 5 используем метод sort среза композитного литерала")
	planets := []string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	sort.StringSlice(planets).Sort() // сортирует срез в алфавитном порядке
	fmt.Println(planets)
}

//controlWork контрольное задание
func controlWork() {
	fmt.Println("Домашнее задание")
	planets := [...]string{ // Базовый массив
		"Марс",
		"Уран",
		"Нептун",
	}
	const New = "Новый "
	planetSlice := planets[:] // Срез массива
	fmt.Println("До терраформирования были планеты:", planetSlice)
	Planets(planetSlice).Terraform(New)
	fmt.Println("После терраформирования получаем планеты:", planetSlice)
}

func main() {
	planets := [...]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}
	list4Planets := []string{" Венера   ", "Земля  ", " Марс"} // срез композитного литерала, названия разделены пробелами
	planetSlice := planets[:]                                  // передавать срез в качестве аргумента проще, чем массив, потому что иначе в параметре объявляемой функции придется указывать размерность массива
	List1(planetSlice)
	List2(planetSlice)
	List3()
	List4Hyperspace(list4Planets)
	List5()
	controlWork()
}
