// урок 30 горутины в Go
// https://golangs.org/goroutines
package main

/* Горутины это своего рода потоки в Go, позволяют работать с разными задачами в одно время,
* связь между горутинами осуществляется через каналы.
 */
import (
	"fmt"
	"math/rand"
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
	fmt.Printf("... гофер ... %d\r\n", number)
}

// list4Gopher гофер с каналом связи
func list4Gopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
	c <- id // отправляем значение обратно в Main
}

// list6SleepyGopher
func list6SleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

func main() {
	fmt.Println("Урок 30, горутины в Go")
	fmt.Println("Листинг 1 создание гофера")
	go listOneGopher() // начало горутины\
	litTwoGopher()
	time.Sleep(4 * time.Second) // Ожидаем когда гофер дас о себе знать.
	fmt.Println("Листинг 4 каналы связи между гоферами")
	c := make(chan int) // Создаем канал связи между гоферами.
	for i := 0; i < 5; i++ {
		go list4Gopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // получаем значение из канала
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
	fmt.Println("Листинг 5 использование select в канале горутин")
	timeout := time.After(4 * time.Second)
	for i := 0; i < 5; i++ {
		go list4Gopher(i, c)
	}
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c: // ждет когда проснется гофер
			fmt.Println("gopher", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("Мое терпение закончислось")
			return
		}
	}
	fmt.Println("Листинг 6 ограничение выполнения опреации старт в ", time.Now())
	for i := 0; i < 5; i++ {
		go list6SleepyGopher(i, c) // запускаем гофер с рандомным временем выполнения
		for i := 0; i < 5; i++ {
			select {
			case gopherID := <-c: // ждет когда проснется гофер
				fmt.Println("гофер ", gopherID, " завершил выполнение в ", time.Now())
				// fmt.Println("gopher", gopherID, " has finished sleeping")
			case <-timeout: // Ждет окончания времени
				fmt.Println("Мое терпение закончислось")
				return
			}
		}
	}
}
