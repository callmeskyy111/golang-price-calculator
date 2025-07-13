package prices

type TaxInclPriceJob struct {
	TaxRate       float64
	InputPrices   []float64
	TaxInclPrices map[string]float64
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
	}
}