package prices

import (
	"errors"
	"fmt"
	"strconv"

	"com.sikora/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
	fm                filemanager.FileManager
}

func (job *TaxIncludedPriceJob) Process() {

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	err := job.fm.WriteJson(job)

	if err != nil {
		fmt.Println("error saving json")
	}
}

func (job *TaxIncludedPriceJob) ReadInPrices() error {
	lines, err := job.fm.ReadLines()
	if err != nil {
		return errors.New("error reading in file")
	}
	for _, line := range *lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return errors.New("error reading in float price value")
		}
		job.InputPrices = append(job.InputPrices, floatPrice)
	}
	return nil
}

func NewTaxIncludedPriceJob(rate float64, ioManager filemanager.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     rate,
		InputPrices: []float64{},
		fm:          ioManager,
	}
}
