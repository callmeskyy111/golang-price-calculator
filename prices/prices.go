package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxInclPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate       float64 `json:"tax_rate"`
	InputPrices   []float64 `json:"input_prices"`
	TaxInclPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxInclPriceJob) LoadData() error{
	lines, err:=job.IOManager.ReadLines()
	if err != nil{
			//fmt.Println(err)
			return err
		}

	prices, err:= conversion.StringsToFloats(lines)
		if err != nil{
			//fmt.Println(err)
			return err
		}
	job.InputPrices = prices
	return  nil
	
}

func (job *TaxInclPriceJob) Process() error {
	 err:= job.LoadData()
	 if err != nil{
		return err
	 }
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice :=price*(1+job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxInclPrice)
	}
	job.TaxInclPrices = result
	return job.IOManager.WriteResult(job)
	
}

func NewTaxInclPriceJob(iom iomanager.IOManager, taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		IOManager: iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}