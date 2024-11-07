package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	utilPkg "github.com/123MayankSharma/receipt-processor-challenge/internal/utils"

	"github.com/google/uuid"
)

func UploadService(w http.ResponseWriter, r *http.Request) {
	//parse request body
	var RequestBody utilPkg.Receipt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&RequestBody)
	if err != nil {
		log.Println("Request Body Parsing Error...")
		log.Println(err)
		http.Error(w, "Failed to Parse Request Body", http.StatusBadRequest)
		return
	}
	//get global store for storing receipts
	store := utilPkg.GetStore()

	//generate id for receipt
	id := uuid.New()
	id_string := id.String()

	//store receipt in string
	store.Receipts[id_string] = RequestBody
	fmt.Println(RequestBody.PurchaseTime)

	//preparing response
	response_map := map[string]string{
		"id": id_string,
	}
	log.Printf("https://localhost:8000/receipts/%s/points", id_string)
	//Set response header
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//encode response map as json
	json.NewEncoder(w).Encode(response_map)

}
