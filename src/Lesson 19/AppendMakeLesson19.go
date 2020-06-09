package main

import "fmt"

//List1 реализация функции append расширения массива
func List1() {
	fmt.Println("Листинг №1 расширяем срез, функция append")
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Println("Срез до увеличения:", dwarfs, "Планет всего", len(dwarfs))
	dwarfs = append(dwarfs, "Оркус")
	fmt.Println("Увеличиваем срез:", dwarfs, "Планет всего", len(dwarfs))
	dwarfs = append(dwarfs, "Сальция", "Квавар", "Седна")
	fmt.Println("Append - вариативная функция, принимает много аргументов")
	fmt.Println("Увеличиваем срез:", dwarfs, "Планет всего", len(dwarfs))
}

//List2dump длина, вместимость, содержимое среза
func List2dump(label string, slice []string) {
	fmt.Println("Листинг №2, реализация метода cap- вместимость среза")
	fmt.Printf("%v: длина %v, вместимость %v, %v\n", label, len(slice), cap(slice), slice)
}

// List3 разбор функции append
func List3() {
	fmt.Println("Листинг 3 разбор функции append")
	dwarfs1 := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	fmt.Println("Исходный срез:", dwarfs1, "длина:", len(dwarfs1), "вместимость:", cap(dwarfs1))
	dwarfs2 := append(dwarfs1, "оркус") // вместимость dwarfs2 увеличивается в 2 раза от dwarfs1
	fmt.Println("Расширим срез:", dwarfs2, "длина:", len(dwarfs2), "вместимость:", cap(dwarfs2))
	dwarfs3 := append(dwarfs2, "Салация", "Квавар", "Седна")
	fmt.Println("Расширим срез:", dwarfs3, "длина:", len(dwarfs3), "вместимость:", cap(dwarfs3))
}

// List4 Трех-индексный срез для ограничения вместимости среза
func List4() {
	fmt.Println("Листинг 4 трехиндексный срез, позволяет не переписывать базовый массив")
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}
	fmt.Println("Исходный срез:", planets)
	terrestrial := planets[0:4:4] // длина 4 вместимость 4
	fmt.Println("Трех-индексный срез:", terrestrial)
	worlds := append(terrestrial, "Церера") // перепишем юпитер
	fmt.Println("Расширяем трехиндексный срез:", worlds)
	fmt.Println("Исходный срез:", planets)
}

// List5 предварительное выделение срезов
func List5() {
	dwarfs := make([]string, 0, 10) // Объявляем срез типа стрнг, в 10 элеменвтов каждый из которых 0 длины
	fmt.Println("объявленный массив", dwarfs)
	dwarfs = append(dwarfs, "Церера", "Плутно", "Хаумеа", "Макемаке", "Эрида")
	fmt.Println("Добавляем элементы", dwarfs)
}

// List6Terraform вариативная функция
func List6Terraform(prefix string, worlds ...string) []string {

	newWorlds := make([]string, len(worlds)) // Создаем новый срез вместо прямого изменения worlds
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

// List6 функция для раскрытия функционала вариативных функций
func List6() {
	fmt.Println("Листинг 6")
	twoWorlds := List6Terraform("Нью", "Венера", "Марс")
	fmt.Println("Создаем новые планеты", twoWorlds)
	planets := []string{"Венера", "Марс", "Юпитер"}
	fmt.Println("Передаем срез в качестве аргумента", planets)
	newPlanets := List6Terraform("New", planets...)
	fmt.Println("новые миры", newPlanets)
}

//  controlWork создаем срез и добавляем в него maxCycles элементов
func controlWork(maxCycles int) {
	slice := []string{}    // пустой срез
	sliceCap := cap(slice) // Вместимость среза
	sliceLen := len(slice) // число элементов среза
	fmt.Println("Исходная вместимость:", sliceCap, "Исходная длина:", sliceLen)
	for i := 0; i < maxCycles; i++ {
		slice = append(slice, "+ Элемент")
		if sliceCap != cap(slice) {
			sliceCap = cap(slice)
			sliceLen = len(slice)
			fmt.Println("Вместимость увеличина до:", sliceCap, "Число элементов:", sliceLen)
		}
	}
}

func main() {
	fmt.Println("Lesson 19")
	List1()
	dwarfs := []string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	List2dump("dwarfs", dwarfs)
	List2dump("dwarfs [1:2]", dwarfs[1:2])
	List3()
	List4()
	List5()
	List6()
	controlWork(200)
}
