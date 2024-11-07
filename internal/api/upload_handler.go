package api

import (
	"net/http"

	upload_service "github.com/123MayankSharma/receipt-processor-challenge/internal/service"
)

func UploadHandler() {
	http.HandleFunc("/receipts/process", upload_service.UploadService)
}
