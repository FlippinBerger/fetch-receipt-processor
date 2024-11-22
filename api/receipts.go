package api

import (
	"net/http"

	"github.com/flippinberger/fetch-receipt-processor/model"
	"github.com/flippinberger/fetch-receipt-processor/store"
	"github.com/labstack/echo/v4"
)

type ProcessResponse struct {
	ID string `json:"id"`
}

func ProcessReceipt(c echo.Context) error {
	receiptStore := store.GetStore()
	id := receiptStore.GenerateID()

	var receipt model.Receipt
	err := c.Bind(&receipt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorDecodingBody{error: err.Error()})
	}

	receiptStore.StoreReceipt(id, receipt)

	return c.JSON(http.StatusOK, ProcessResponse{ID: id})
}

type PointsResponse struct {
	Points int `json:"points"`
}

func GetPoints(c echo.Context) error {
	id := c.Param("id")

	receiptStore := store.GetStore()
	receipt, err := receiptStore.GetReceipt(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "{}")
	}

	return c.JSON(http.StatusOK, PointsResponse{Points: receipt.Rewards()})
}
