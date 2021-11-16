package binancego

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func BenchmarkEndPointMarketData_GetOldTradeLookupRequest(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		apiKey := os.Getenv("BINANCE_API_KEY")
		secretKey := os.Getenv("BINANCE_API_SECRET")
		symbol := "BTCUSDT"
		client := NewBinanceClient(apiKey, secretKey)
		if _, err := client.MarketDataEndpoint().GetOldTradeList(context.Background(), GetOldTradeListRequest{
			Symbol: symbol,
			FromId: IntPointer(0),
			Limit:  nil,
		}); err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkEndPointMarketData_GetRecentTradeLookupRequest(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		apiKey := os.Getenv("BINANCE_API_KEY")
		secretKey := os.Getenv("BINANCE_API_SECRET")
		symbol := "BTCUSDT"
		client := NewBinanceClient(apiKey, secretKey)
		if _, err := client.MarketDataEndpoint().GetOldTradeList(context.Background(), GetOldTradeListRequest{
			Symbol: symbol,
			FromId: IntPointer(0),
			Limit:  nil,
		}); err != nil {
			fmt.Println(err)
		}
	}
}
