package uconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

// using AbsoluteZeroC and FreezingK for conversion
const (
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)
const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CelsiusToKelvin(c Celsius) Kelvin {
	return Kelvin(c + AbsoluteZeroC)
}

func KelvinToCelsius(k Kelvin) Celsius {
	return Celsius(k + FreezingK)
}

func KelvinToFahrenheit(k Kelvin) Fahrenheit {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}

func FahrenheitToKelvin(f Fahrenheit) Kelvin {
	return CelsiusToKelvin(FahrenheitToCelsius(f))
}
