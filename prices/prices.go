package prices

import (
	"bufio"
	"fmt"
	"os"
	"example.com/price-calculator/conversion"
)

type TaxInclPriceJob struct {
	TaxRate       float64
	InputPrices   []float64
	TaxInclPrices map[string]float64
}

func (job *TaxInclPriceJob) LoadData(){
	file,err:=os.Open("prices.txt")
	if err!= nil{
		fmt.Println("❌ Err. Could Not Open file:",err)
		return
	}
	scanner:=bufio.NewScanner(file)

	var lines []string //helper variab;e

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	err=scanner.Err()
	if err != nil{
		fmt.Println("❌ Err. Could not read file-content:",err)
		file.Close()
		return
	}

	prices, err:= conversion.StringsToFloats(lines)
		if err != nil{
			fmt.Println("❌ Err. Could not covert price to float:",err)
			file.Close()
			return
		}
	job.InputPrices = prices
	file.Close()
}

func (job *TaxInclPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxInclPrice :=price*(1+job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxInclPrice)
	}
	fmt.Println(result)
}

func NewTaxInclPriceJob(taxRate float64) *TaxInclPriceJob {
	return &TaxInclPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}