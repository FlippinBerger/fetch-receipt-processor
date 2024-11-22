package utils

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func GetAlphaNumericCount(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			count++
		}
	}
	return count
}

func GetCentsPoints(s string) int {
	priceParts := strings.Split(s, ".")
	if len(priceParts) != 2 {
		return 0
	}

	cents := priceParts[1]
	if cents == "00" {
		return 75
	} else if cents == "25" || cents == "50" || cents == "75" {
		return 25
	}

	return 0
}

func GetPairPoints(length int) int {
	return (length / 2) * 5
}

func IsOddDayInDateString(s string) bool {
	dateParts := strings.Split(s, "-")
	if len(dateParts) != 3 {
		return false
	}

	day := dateParts[2]
	num, _ := strconv.Atoi(day)

	return num%2 == 1
}

func IsBetweenTwoAndFourPM(s string) bool {
	timeParts := strings.Split(s, ":")
	if len(timeParts) != 2 {
		return false
	}

	hour, _ := strconv.Atoi(timeParts[0])
	minutes, _ := strconv.Atoi(timeParts[1])

	// 2PM is NOT after 2PM
	if hour == 14 && minutes == 0 {
		return false
	}

	// Strictly between 2:01PM and 4PM (not inclusive of 4PM)
	if hour >= 14 && hour < 16 {
		return true
	}

	return false
}

func GetItemDescriptionPoints(s string, priceString string) int {
	s = strings.TrimSpace(s)
	if len(s)%3 != 0 {
		return 0
	}

	pricePieces := strings.Split(priceString, ".")
	if len(pricePieces) != 2 {
		return 0
	}

	dollars, _ := strconv.Atoi(pricePieces[0])
	cents, _ := strconv.Atoi(pricePieces[1])

	price := float64(dollars) + (float64(cents) / 100)

	return int(math.Ceil(price * 0.2))
}
