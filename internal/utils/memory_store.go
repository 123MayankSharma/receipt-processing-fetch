package utils

import (
	"fmt"
	"sync"
)

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []item `json:"items"`
	Total        string `json:"total"`
}

type Memory_store struct {
	Receipts map[string]Receipt
}

var lock = &sync.Mutex{}

var store *Memory_store

func GetStore() *Memory_store {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			fmt.Println("Instantiating Receipt Store...\n")
			store = &Memory_store{
				Receipts: make(map[string]Receipt),
			}
		} else {
			fmt.Println("Store already Exists")
		}
	} else {
		fmt.Println("Store already Exists")
	}

	return store
}
