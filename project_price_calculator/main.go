package main

import (
	"fmt"

	"com.sikora/price-calculator/filemanager"
	"com.sikora/price-calculator/prices"
)

func main() {

	taxRates := []float64{0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		ioManager := filemanager.New("prices.txt", fmt.Sprintf("prices_%v.json", taxRate))
		job := prices.NewTaxIncludedPriceJob(taxRate, ioManager)
		err := job.ReadInPrices()

		if err != nil {
			fmt.Println(err)
			continue
		}
		job.Process()
	}

}
