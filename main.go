package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _,taxRate := range taxRates{
		taxInclPrices := []float64{}
		for _,price := range prices{
			taxInclPrices= append(taxInclPrices, price * (1+ taxRate))
		}
		result[taxRate] = taxInclPrices
	}
fmt.Println(result)
}
