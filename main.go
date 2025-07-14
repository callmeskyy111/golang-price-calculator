package main

import (
	// "example.com/price-calculator/cmdmanager"
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}


	for _,taxRate := range taxRates{
		fm:=fileManager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		//cmdm := cmdmanager.New() // Swappable Struct - Optnl. ✔️✔️✔️
		//priceJob:=prices.NewTaxInclPriceJob(cmdm,taxRate)
		priceJob:=prices.NewTaxInclPriceJob(fm,taxRate)
		err:=priceJob.Process()
		if err != nil{
			fmt.Println("🔴 Could not process job:",err)
		}
	}
}
