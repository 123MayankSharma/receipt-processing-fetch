package service

import (
	"encoding/json"
	"net/http"

	utilPkg "github.com/123MayankSharma/receipt-processor-challenge/internal/utils"
	"github.com/gorilla/mux"
)

func PointsService(w http.ResponseWriter, r *http.Request) {
	//fetching id from url
	receipt_id := mux.Vars(r)["id"]

	//Getting Global receipt store
	store := utilPkg.GetStore()

	//fetching the receipt with receipt id from store
	receipt := store.Receipts[receipt_id]

	//calculating the points based on rules for given receipt
	pointsValue := utilPkg.PointsCalculation(receipt)

	//preparing json response
	responseMap := map[string]int64{
		"points": pointsValue,
	}

	//setting content type and status code
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)

	//sending response
	json.NewEncoder(w).Encode(responseMap)

}
