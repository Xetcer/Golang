// Ссылка на урок https://golangs.org/map
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// List1 реализация листинга один
func List1() {
	fmt.Println("Листинг 1")
	temperature := map[string]int{
		"Земля": 15,
		"Марс":  -65,
		"Луна":  0,
	}
	temp := temperature["Земля"]
	fmt.Printf("Средняя температура на поверхности Земли составляет %v °C.\n", temp)
	temperature["Земля"] = 16
	temperature["Венера"] = 464 // Добавили новй ключ/значение
	fmt.Println(temperature)

	if saturn, ok := temperature["Сатурн"]; ok { // comma, ok
		fmt.Printf("Средняя температура на поверхности сатурна = %v °C.\n", saturn)
	} else {
		fmt.Println("Планеты с ключом Сатурн в карте нет.")
	}
}

// List2 копирование карт
func List2() {
	fmt.Println("листинг 2")
	planets := map[string]string{
		"Земля": "Сектор ZZ9",
		"Марс":  "Сектор ZZ9",
	}
	fmt.Println("Исходная карта", planets)
	planetsMkII := planets
	planets["Земля"] = "Упс"
	fmt.Println("Изменим сектор земли во второй карте")
	fmt.Println("Карта 1", planets)
	fmt.Println("Карта 2", planetsMkII)
	fmt.Println("Удалим землю")
	delete(planets, "Земля")
	fmt.Println("Карта 2", planetsMkII)
}

// List3 использование карты для подсчета частоты использования элементов
func List3() {
	fmt.Println("Листинг 3")
	temperature := []float64{ // массив температур
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int) // карта
	fmt.Println(frequency)
	for _, t := range temperature { // Заполним карту ключом будут значения элемента массива
		frequency[t]++ // Значениями карты будет число повторений температуры в массиве
	}
	fmt.Println(frequency)
	for t, num := range frequency {
		fmt.Printf("%+.2f встречается %d раз(а) \n", t, num)
	}
}

// List4 группировка данных с картами и срезами в GoLang
func List4() {
	fmt.Println("Листинг 4")
	temperatures := []float64{ // исходный массив температур
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64) //карта с ключами типа float64 и значениями типа []float64

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10       //скругляем температуры до 10°С вниз
		groups[g] = append(groups[g], t) // в карту по ключу с температурой g добавляем
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}

// List5 Множество
func List5() {
	fmt.Println("Листинг 5, множества")
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool) // Создаем карту с ключом типа float64 и значениями типа bool
	for _, t := range temperatures {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("часть множества")
	}

	fmt.Println(set)
	unique := make([]float64, 0, len(set)) // создаем пустой срез вместимостью
	// заполняем его значениями карты
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique) // сортируем по возрастанию

	fmt.Println(unique)
}

// ControlWork считаем частоту упоминания слов в строке, строку переводим в нижный регистр и обрезаем знаки препинания
func ControlWork(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text)) // получаем срезы слов
	frequency := make(map[string]int, len(words))  // создаем карту с ключами типа string значениями типа int длиной words
	for _, word := range words {                   // бежим по срезу слов
		word = strings.Trim(word, `.,"-;`) // удаляем его знаки препинания
		frequency[word]++                  // увеличиваем число повторений
	}
	return frequency
}

func main() {
	fmt.Println("Карты в Go")
	List1()
	List2()
	List3()
	List4()
	List5()
	text := `As far as eye could reach he saw nothing but the stems of the
	great plants about him receding in the violet shade, and far overhead
	the multiple transparency of huge leaves filtering the sunshine to the
	solemn splendour of twilight in which he walked. Whenever he felt able
	he ran again; the ground continued soft and springy, covered with the
	same resilient weed which was the first thing his hands had touched in
	Malacandra. Once or twice a small red creature scuttled across his
	path, but otherwise there seemed to be no life stirring in the wood;
	nothing to fear -- except the fact of wandering unprovisioned and alone
	in a forest of unknown vegetation thousands or millions of miles
	beyond the reach or knowledge of man.`
	fmt.Println("Контрольное задание")
	frequency := ControlWork(text)
	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("Слово %v встречается %v раз(аит) \n", word, count)
		}
	}

}
