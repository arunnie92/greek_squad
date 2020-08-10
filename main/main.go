package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/speps/go-hashids"
)

type hashID string

func newHash() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = "id"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)
	// e, _ := *h.Encode([]int{45, 434, 1313, 99})
	// fmt.Println(e)
	// d, _ := h.DecodeWithError(e)
	// fmt.Println(d)

	return h
}

type product struct {
	SkuID              string `json:"currentProductSku"`
	ProductDescription string `json:"productDescription"`
	TermBeginDate      string `json:"termBeginDate"`
	TermEndDate        string `json:"termEndDate"`
}

type userInfo struct {
	Email string
	// Password string // gmail password
}

type products []product

func main() {
	reader := bufio.NewReader(os.Stdin)
	var user userInfo

	fmt.Print("Please enter your emaill-> ")
	email, _ := reader.ReadString('\n')
	user.Email = email

	// TODO: validate email

	fmt.Println("Uploading Geek Squad GSP JSON file... ")
	jsonFile, jsonFilErr := os.Open("../data_2.json")
	if jsonFilErr != nil {
		fmt.Println(jsonFilErr)
		return
	}
	fmt.Println("Successfully Geek Squad GSP JSON file...")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var items products

	err := json.Unmarshal(byteValue, &items)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(items[0].SkuID)

	for index, item := range items {
		fmt.Println(index)

		fmt.Println(item.SkuID)
	}
	// fmt.Println(toHash(items[0].SkuID))
}
