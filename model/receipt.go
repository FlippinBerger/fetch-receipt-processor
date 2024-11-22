package model

import (
	"github.com/flippinberger/fetch-receipt-processor/utils"
)

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

type Item struct {
	ShortDescription string
	Price            string
}

func (r Receipt) Rewards() int {
	points := utils.GetAlphaNumericCount(r.Retailer)

	points += utils.GetCentsPoints(r.Total)
	points += utils.GetPairPoints(len(r.Items))

	for _, item := range r.Items {
		points += utils.GetItemDescriptionPoints(item.ShortDescription, item.Price)
	}

	if utils.IsOddDayInDateString(r.PurchaseDate) {
		points += 6
	}

	if utils.IsBetweenTwoAndFourPM(r.PurchaseTime) {
		points += 10
	}

	return points
}
