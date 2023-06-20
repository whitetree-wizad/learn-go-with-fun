package main

type celsius float64
type fahreinheit float64
type kelvins float64

func (k celsius) ToFahrenheit() fahreinheit {
	return fahreinheit((k * 9.0 / 5.0) + 32.0)
}
func (k kelvins) ToFahrenheit() fahreinheit {
	return (k - 273.15).ToFahrenheit()
}

func (k celsius) ToKelvins() kelvins {
	return kelvins(k + 273.5)
}

func (k fahreinheit) ToKelvins() kelvins {
	return kelvins((k-32.0)*5/9 + 273.15)
}

func (k kelvins) ToCelsius() celsius {
	return celsius(k - 273.15)
}

func (k fahreinheit) ToCelsius() celsius {
	return celsius((k - 32.0) * 5 / 9)
}

func main() {

}
