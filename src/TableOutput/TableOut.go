package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64
type cnvrtT func(t int) (string, string)

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32.0)
}

const (
	line     = "======================="
	dataline = "| %-8s | %-8s |\n"
)

//  drawTable отрисовка таблицы с конвертацией
func drawTable(s1 string, s2 string, convertor cnvrtT) {
	fmt.Println(line)
	fmt.Printf(dataline, s1, s2)
	fmt.Println(line)
	for i := 40; i <= 100; i += 5 {
		s1, s2 = convertor(i)
		fmt.Printf(dataline, s1, s2)
		fmt.Println(line)
	}
}

// cToF возвращает t°С и t°F в строке с форматом 0.2
func cToF(t int) (string, string) {
	c := celsius(t)
	f := c.fahrenheit()
	return fmt.Sprintf("%.2f", float64(c)), fmt.Sprintf("%.2f", float64(f))
}

//fToC возвращает t°F и t°С в строке с форматом 0.2
func fToC(t int) (string, string) {
	f := fahrenheit(t)
	c := f.celsius()
	return fmt.Sprintf("%.2f", float64(f)), fmt.Sprintf("%.2f", float64(c))
}

func main() {
	drawTable("°C", "°F", cToF)
	fmt.Println("*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*")
	drawTable("°F", "°C", fToC)
}
