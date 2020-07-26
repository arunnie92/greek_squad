package main

import (
	"fmt"
	"encoding/json"
)

type Product struct {
	skuId   string `json:"currentProductSku"`
	productDescription string `json:"productDescription"`
	purchasedDate string `json:"currentProductSku"`
	// expiredDate
}
