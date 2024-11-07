package service

import (
	"encoding/json"
	"net/http"

	utilPkg "github.com/123MayankSharma/receipt-processor-challenge/internal/utils"
	"github.com/gorilla/mux"
)

func PointsService(w http.ResponseWriter, r *http.Request) {
	receipt_id := mux.Vars(r)["id"]
	print(receipt_id, "----")
	store := utilPkg.GetStore()

	receipt := store.Receipts[receipt_id]

	pointsValue := utilPkg.PointsCalculation(receipt)

	responseMap := map[string]int64{
		"points": pointsValue,
	}

	w.Header().Set("content", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(responseMap)

}
