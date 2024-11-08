package main

import (
	"fmt"
	"net/http"

	"github.com/123MayankSharma/receipt-processor-challenge/internal/api"
)

const PORT = 8000

func startUpMessage() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("               Welcome to the Receipt Processor API       ")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("The API is running and ready to process your receipts!")
	fmt.Println()
	fmt.Println("Available Endpoints:")
	fmt.Printf("  * Process a new receipt:           http://localhost:%d/receipts/process\n", PORT)
	fmt.Printf("  * Retrieve points for a receipt:   http://localhost:%d/receipts/{id}/points\n", PORT)
	fmt.Println()
	fmt.Println("Note: Replace {id} with the Receipt ID you receive after processing a receipt.")
	fmt.Println("---------------------------------------------------------")
}

func main() {
	startUpMessage()
	api.UploadHandler()
	api.PointsHandler()

	http.ListenAndServe(":8000", nil)

}
