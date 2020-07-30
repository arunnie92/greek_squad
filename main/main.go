package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type product struct {
	SkuID              string `json:"currentProductSku"`
	ProductDescription string `json:"productDescription"`
	TermBeginDate      string `json:"termBeginDate"`
	TermEndDate        string `json:"termEndDate"`
}

type products []product

func main() {
	jsonFile, jsonFilErr := os.Open("../data_2.json")
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return
	}

	fmt.Println("Successfully Opened data")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var items products

	err := json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("here")
	fmt.Println(items[0].SkuID)
}
