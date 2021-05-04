// урок 30 горутины в Go
// https://golangs.org/goroutines
package main

/* Горутины это своего рода потоки в Go, позволяют работать с разными задачами в одно время,
* связь между горутинами осуществляется через каналы.
 */
import (
	"fmt"
	"time"
)

// listOne листинг 1 запуск горутины
func listOneGopher() {
	time.Sleep(3 * time.Second) // гофер спит
	fmt.Println("гофер листинг 1 закончил работу")
}

// litTwoGopher запуск нескольких горутин
func litTwoGopher() {
	fmt.Println("Листинг 2 - 3 запуск 5 гоферов с параметрами") // в результате видим что гофера работают не по порядку
	for i := 0; i < 5; i++ {
		go sleepyGopher(i + 1)
	}
}

// sleepyGopher тело одного гофера для листинга 2
func sleepyGopher(number int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("... гофер ... %d", number)
}

func main() {
	fmt.Println("Урок 30, горутины в Go")
	fmt.Println("Листинг 1 создание гофера")
	go listOneGopher() // начало горутины\
	litTwoGopher()
	time.Sleep(4 * time.Second) // Ожидаем когда гофер дас о себе знать.
}
