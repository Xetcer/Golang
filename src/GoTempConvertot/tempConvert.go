package main

type celsius float64
type kelvin float64
type fahrenheit float64

// kelvinToCelcius конвертирует °K в °С
func kelvinToCelcius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// celsiusToKelvin конвертирует °С в °K
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

//celsiusToFahrenheit конвертирует °C в °F
func celsiusToFahrenheit(c celsius) fahrenheit {
	c = (c * 9.0 / 5.0) + 32.0
	return fahrenheit(c)
}

//  kelvinToFahrenheit конвертирует °К в °F
func kelvinToFahrenheit(k kelvin) fahrenheit {
	return celsiusToFahrenheit(kelvinToCelcius(k))
}

//  fahrenheitTocelsius конвертирует °F в °С
func fahrenheitTocelsius(f fahrenheit) celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// методы для типов
// метод конвертирует °C
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// метод конвертирует в °K
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) kelvin() kelvin {
	return (f.celsius()).kelvin()
}

// конвертируем в °F
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (k kelvin) fahrenheit() fahrenheit {
	return (k.celsius()).fahrenheit()
}
