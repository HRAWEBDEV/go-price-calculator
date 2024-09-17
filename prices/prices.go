package prices

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadDate()
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result

	file, err := os.Create(fmt.Sprintf("result_%.0f.json", job.TaxRate*100))
	if err != nil {
		fmt.Println("could not create file!")
		return
	}

	encode := json.NewEncoder(file)
	err = encode.Encode(job)

	if err != nil {
		fmt.Println("failed to convert data to json")
		return
	}
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadDate() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		file.Close()
		return
	}

	job.InputPrices = prices
}
