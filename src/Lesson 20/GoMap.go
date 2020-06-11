// Ссылка на урок https://golangs.org/map
package main

import "fmt"

// List1 реализация листинга один
func List1() {
	fmt.Println("")
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

func main() {
	fmt.Println("Карты в Go")
	List1()
}
