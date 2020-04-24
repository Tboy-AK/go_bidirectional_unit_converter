package main

import (
	"fmt"
	"math"
)

// Converter measures distance
type Converter struct{}

// Feet measures distance
type Feet float64

// Centimeter measures distance
type Centimeter float64

// Minutes measures time
type Minutes float64

// Seconds measures time
type Seconds float64

// Celsius measures temperature
type Celsius float64

// Farenheit measures temperature
type Farenheit float64

// Radian measures angle
type Radian float64

// Degree measures angle
type Degree float64

// Kilogram measures mass
type Kilogram float64

// Pounds measures mass
type Pounds float64

// FeetToCentimeter converts feet to centimeter
func (cvr Converter) FeetToCentimeter(c Feet) Centimeter {
	return (c / 0.0328084)
}

// CentimeterToFeet converts Centimeter to Feet
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return (c * 0.0328084)
}

// MinutesToSeconds converts Minutes to Seconds
func (cvr Converter) MinutesToSeconds(c Minutes) Seconds {
	return (c * 60)
}

// SecondsToMinutes converts Seconds to Minutes
func (cvr Converter) SecondsToMinutes(c Seconds) Minutes {
	return (c / 60)
}

// CelsiusToFarenheit converts Celsius to Farenheit
func (cvr Converter) CelsiusToFarenheit(c Celsius) Farenheit {
	return ((c * 9 / 5) + 32)
}

// FarenheitToCelsius converts Farenheit to Celsius
func (cvr Converter) FarenheitToCelsius(c Farenheit) Celsius {
	return ((c - 32) * 5 / 9)
}

// RadianToDegree converts Radian to Degree
func (cvr Converter) RadianToDegree(c Radian) Degree {
	return (c * math.Pi * 180)
}

// DegreeToRadian converts Degree to Radian
func (cvr Converter) DegreeToRadian(c Degree) Radian {
	return (c / (math.Pi * 180))
}

// KilogramToPounds converts Kilogram to Pounds
func (cvr Converter) KilogramToPounds(c Kilogram) Pounds {
	return (c * 2.205)
}

// PoundsToKilogram converts Pounds to Kilogram
func (cvr Converter) PoundsToKilogram(c Pounds) Kilogram {
	return (c / 2.205)
}

func main() {
	fmt.Println(Converter.FeetToCentimeter())
	fmt.Println(Converter.CentimeterToFeet())
	fmt.Println(Converter.MinutesToSeconds())
	fmt.Println(Converter.SecondsToMinutes())
	fmt.Println(Converter.CelsiusToFarenheit())
	fmt.Println(Converter.FarenheitToCelsius())
	fmt.Println(Converter.RadianToDegree())
	fmt.Println(Converter.DegreeToRadian())
	fmt.Println(Converter.KilogramToPounds())
	fmt.Println(Converter.PoundsToKilogram())
}
