package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64

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

func drawTable(s1 string, s2 string, isHeader bool, isEnd bool) {
	if isHeader == true {
		fmt.Println(line)
		fmt.Printf(dataline, s1, s2)
		fmt.Println(line)
	} else {
		if isEnd == true {
			fmt.Println(line)
		} else {
			fmt.Printf(dataline, s1, s2)
		}
	}
}

func convertTemp() {
	var f fahrenheit
	var c celsius
	drawTable("°C", "°F", true, false)
	for i := 40; i <= 100; i += 5 {
		c = celsius(i)
		f = c.fahrenheit()
		drawTable(fmt.Sprintf("%.2f", float64(c)), fmt.Sprintf("%.2f", float64(f)), false, false)
	}
	drawTable("", "", false, true)
	fmt.Println("*+*+*+*+*+*+*+*+*+*+*+*+*+*+*+*")
	drawTable("°F", "°C", true, false)
	for i := 40; i <= 100; i += 5 {
		f = fahrenheit(i)
		c = f.celsius()
		drawTable(fmt.Sprintf("%.2f", float64(f)), fmt.Sprintf("%.2f", float64(c)), false, false)
	}
	drawTable("", "", false, true)
}

func main() {
	convertTemp()
}
