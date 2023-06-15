package main

import (
	"encoding/json"
	"flag"
	"fmt"
	U "github.com/manfioLP/go-cli-tax-calculator/utils"
)

func main() {
	// Define command-line flags
	inputFlag := flag.String("input", "", "Input data in JSON format please")
	flag.Parse()

	// Check for input
	if *inputFlag == "" {
		fmt.Println("Please provide the input data using the -input flag")
		return
	}

	//todo: add read from file?
	var input []U.Order
	err := json.Unmarshal([]byte(*inputFlag), &input)
	if err != nil {
		fmt.Println("Parsing input data failed, err: ", err)
		return
	}

	taxResults := U.CalculateTax(input)

	taxResultsJSON, err := json.Marshal(taxResults)
	if err != nil {
		fmt.Println("marshalling tax results failed, err:", err)
		return
	}
	fmt.Println(string(taxResultsJSON))
}
