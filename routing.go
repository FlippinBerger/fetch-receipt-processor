package main

import (
	"github.com/flippinberger/fetch-receipt-processor/api"
	"github.com/labstack/echo/v4"
)

func registerRoutes(e *echo.Echo) {
	e.POST("/receipts/process", api.ProcessReceipt)
	e.GET("/receipts/:id/points", api.GetPoints)
}
