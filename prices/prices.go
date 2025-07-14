package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/fileManager"
)

type TaxInclPriceJob struct {
	TaxRate       float64
	InputPrices   []float64
	TaxInclPrices map[string]string
}

func (job *TaxInclPriceJob) LoadData(){
	lines, err:=fileManager.ReadLines("prices.txt")
	if err != nil{
			fmt.Println(err)
			return
		}

	prices, err:= conversion.StringsToFloats(lines)
		if err != nil{
			fmt.Println(err)
			return
		}
	job.InputPrices = prices
	
}

func (job *TaxInclPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice :=price*(1+job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxInclPrice)
	}
	job.TaxInclPrices = result
	fileManager.WriteJSON(fmt.Sprintf("result_%.0f.json",job.TaxRate*100),job)
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}