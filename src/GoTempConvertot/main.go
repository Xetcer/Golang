package main

import "fmt"

func main() {
	var c celsius = 127 //°C
	var k kelvin
	var f fahrenheit
	k = c.kelvin()
	f = c.fahrenheit()
	fmt.Printf("%v°C соответствует %v°K или %v°F", c, k, f)
}
