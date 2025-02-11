package prices

import (
	"errors"
	"fmt"
	"strconv"

	"com.sikora/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	fm                iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) Process(channel chan bool, errChannel chan error) {

	err := job.ReadInPrices()

	if err != nil {

		errChannel <- err
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	err = job.fm.WriteResult(job)

	if err != nil {
		errChannel <- err
		return
	}
	channel <- true
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

func NewTaxIncludedPriceJob(rate float64, ioManager iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     rate,
		InputPrices: []float64{},
		fm:          ioManager,
	}
}
