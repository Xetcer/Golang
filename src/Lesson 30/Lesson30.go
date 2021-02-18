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
	fmt.Println("... snore ...")
}

func main() {
	fmt.Println("Урок 30, горутины в Go")
	go listOneGopher()          // начало горутины
	time.Sleep(4 * time.Second) // Ожидаем когда гофер дас о себе знать.
}
