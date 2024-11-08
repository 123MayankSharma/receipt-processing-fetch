package utils

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	models "github.com/123MayankSharma/receipt-processor-challenge/models"
)

func isAlphaNumericScore(text string) int64 {
	var score int64 = 0
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			score += 1
		}
	}
	return score
}

func totalValueScore(total string) int64 {
	total_values := strings.Split(total, ".")

	decimal_value, _ := strconv.ParseFloat(total_values[1], 64)

	if decimal_value == 0.0 {
		return 50
	}
	return 0

}

func totalMultipleScore(total string) int64 {
	total_value, _ := strconv.ParseFloat(total, 64)

	var remainder float64 = math.Mod(total_value, 0.25)

	if remainder > 0.0 {
		return 0
	}

	return 25

}

func itemsLengthScore(ItemList []models.Item) int64 {
	var score int64 = 0

	score += int64(len(ItemList) / 2)

	return score * 5
}

func itemDescriptionLengthScore(ItemList []models.Item) int64 {
	var score int64 = 0

	for _, item := range ItemList {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			score_addition := int64(math.Ceil(price * 0.2))
			score += score_addition

		}
	}

	return score
}

func purchaseDayScore(purchaseDate string) int64 {
	date_values := strings.Split(purchaseDate, "-")
	day_value, _ := strconv.ParseInt(date_values[2], 10, 64)

	if day_value%2 != 0 {
		return 6
	}

	return 0
}

func purchaseTimeScore(purchaseTime string) int64 {
	startTime := 14 * 60
	endTime := 16 * 60

	purchaseTimeValues := strings.Split(purchaseTime, ":")
	hourValue, _ := strconv.ParseInt(purchaseTimeValues[0], 10, 64)
	minuteValue, _ := (strconv.ParseInt(purchaseTimeValues[1], 10, 64))
	purchaseTimeMin := (hourValue)*60 + minuteValue

	if purchaseTimeMin > int64(startTime) && purchaseTimeMin < int64(endTime) {
		return 10
	}

	return 0

}

func PointsCalculation(receipt models.Receipt) int64 {
	var points int64 = 0
	//calculating points according to given rules
	points += isAlphaNumericScore(receipt.Retailer)

	points += totalValueScore(receipt.Total)

	points += totalMultipleScore(receipt.Total)

	points += itemsLengthScore(receipt.Items)

	points += itemDescriptionLengthScore(receipt.Items)

	points += purchaseDayScore(receipt.PurchaseDate)

	points += purchaseTimeScore(receipt.PurchaseTime)

	return int64(points)

}
