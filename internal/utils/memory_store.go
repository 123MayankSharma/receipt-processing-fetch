package utils

import (
	"fmt"
	"sync"

	models "github.com/123MayankSharma/receipt-processor-challenge/models"
)

var lock = &sync.Mutex{}

var store *models.Memory_store

func GetStore() *models.Memory_store {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			fmt.Println("Instantiating Receipt Store...\n")
			store = &models.Memory_store{
				Receipts: make(map[string]models.Receipt),
			}
		} else {
			fmt.Println("Fetching Existing Store....\n")
		}
	} else {
		fmt.Println("Fetching Existing Store....\n")
	}

	return store
}
