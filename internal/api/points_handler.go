package api

import (
	"net/http"

	points_service "github.com/123MayankSharma/receipt-processor-challenge/internal/service"
	"github.com/gorilla/mux"
)

func PointsHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", points_service.PointsService).Methods("GET")
	http.Handle("/", router)
}
