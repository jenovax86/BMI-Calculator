package main

import (
	"fmt"
	"math"
)

func bmiCalculater() float64 {
	var weight float64
	var height float64
	fmt.Println("Enter your weight and your height to calculate your BMI(body mass index)")

	fmt.Scan(&weight, &height)
	finalHeight := math.Pow(height/100, 2)
	result := weight / finalHeight
	return result
}

func main() {
	fmt.Println(bmiCalculater())
}
