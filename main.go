package main

import "fmt"

var K float64 = 480.0

func main() {
	C := K - 273

	fmt.Printf("The Kelvin temperature is %.2f and its equivalent in Celsius is %.2f", K, C)
}
