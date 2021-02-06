package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const active = "Active"



type product struct {
	ID                      int64  `json:"created"`
	Sku                     string `json:"currentProductSku"`
	Description             string `json:"productDescription"`
	Model                   string `json:"model"`
	PurchaseDate            string `json:"purchasedDate"`
	ProtectionPlanTitle     string `json:"title"`
	ProtectionPlanType      string `json:"contractClassDescription"`
	ProtectionPlanSku       string `json:"originalPlanSku"`
	ProtectionPlanStatus    string `json:"contractStatusDescription"`
	ProtectionPlanStartDate string `json:"termBeginDate"`
	ProtectionPlanEndDate   string `json:"termEndDate"`
	//
	BbyOrderNumber     string `json:"orderNumber"`
	ReceiptTransaction string `json:"fourPartKey"`
}

func ingestData() map[int64]product {
	fmt.Println("Reading Geek Squad GSP Products JSON file... ")
	jsonFile, jsonFilErr := os.Open("../data.json")
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return nil
	}
	fmt.Println("Successfully Ingested Geek Squad GSP Products JSON file...")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var productsArr []product

	err := json.Unmarshal(byteValue, &productsArr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	productsMap := make(map[int64]product)

	for _, product := range productsArr {
		if product.ProtectionPlanStatus == active {
			productsMap[product.ID] = product
		}
	}

	return productsMap
}

func main() {
	products := ingestData()
	fmt.Println(len(products))

	for _, product := range products {
		fmt.Println(product.Sku)
	}
}
