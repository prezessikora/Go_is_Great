package main

import (
	"fmt"

	"com.sikora/price-calculator/filemanager"
	"com.sikora/price-calculator/prices"
)

func main() {

	taxRates := []float64{0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		ioManager := filemanager.New("prices.txt", fmt.Sprintf("prices_%v.json", taxRate))
		job := prices.NewTaxIncludedPriceJob(taxRate, ioManager)

		go job.Process(doneChans[index], errorChans[index])

	}
	//wait for the go routines to finish by emitting a bool or error
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
