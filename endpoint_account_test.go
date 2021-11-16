package binancego

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestEndPointTrade_Trade(t *testing.T) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_API_SECRET")
	symbol := "CELRUSDT"
	tif := ENUM_ORDER_TIME_IN_FORCE_GTC
	err := NewBinanceClient(apiKey, secretKey).AccountDataEndpoint().CreateOrder(context.Background(), CreateOrderRequest{
		Symbol:      symbol,
		Side:        ENUM_ORDER_SIDE_SELL,
		Type:        ENUM_ORDER_TYPE_LIMIT,
		TimeInForce: &tif,
		Quantity:    Float64Pointer(192),
		Price:       Float64Pointer(0.49),
		Timestamp:   time.Now().UnixMilli(),
	}, nil)
	if err != nil {
		fmt.Println(err)
	}
}
