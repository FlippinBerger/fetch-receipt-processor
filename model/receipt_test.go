package model

import (
	"testing"
)

func TestTargetReceiptRewards(t *testing.T) {
	targetReceipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
	}

	rewardPoints := targetReceipt.Rewards()
	if rewardPoints != 28 {
		t.Fatalf("Reward points should be 28, is actually %d", rewardPoints)
	}
}

func TestCornerMarketReceiptRewards(t *testing.T) {
	receipt := Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}

	rewardPoints := receipt.Rewards()
	if rewardPoints != 109 {
		t.Fatalf("Reward points should be 109, is actually %d", rewardPoints)
	}
}
