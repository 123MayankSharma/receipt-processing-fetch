package main

import (
	"fmt"
	"net/http"

	"github.com/123MayankSharma/receipt-processor-challenge/internal/api"
)

func main() {
	fmt.Println("------------Welcome to Receipt Processor API ------------")
	fmt.Println("Please Refer Below for the available endpoints")
	const PORT = 8000

	fmt.Printf("* http://localhost:%d/receipts/process-------->\n", PORT)
	fmt.Printf("* http://localhost:%d/receips/{id}/points-------->\n", PORT)
	fmt.Println("{id} denotes the Receipt id received after uploading the receipt")

	api.UploadHandler()
	api.PointsHandler()

	http.ListenAndServe(":8000", nil)

}
