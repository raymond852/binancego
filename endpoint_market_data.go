package binancego

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EndPointMarketData struct {
	conf Binance
}

type GetOrderBookRequest struct {
	Symbol string
	Limit  *int
}

type GetOrderBookResponse struct {
	LastUpdateId int64       `json:"lastUpdateId"`
	Bids         [][]float64 `json:"bids"`
	Asks         [][]float64 `json:"asks"`
}

func (e EndPointMarketData) GetOrderBook(ctx context.Context, req GetOrderBookRequest) ([]GetOrderBookResponse, error) {
	path := "/api/v3/depth"
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, e.conf.binanceUrl+path, nil)
	q := httpReq.URL.Query()
	q.Set("symbol", req.Symbol)
	if req.Limit != nil {
		q.Set("limit", strconv.Itoa(*req.Limit))
	}
	httpReq.URL.RawQuery = q.Encode()
	httpReq.Header.Set(headerAPIKey, e.conf.apiKey)
	var ret []GetOrderBookResponse
	if resp, err := http.DefaultClient.Do(httpReq); err != nil {
		return nil, err
	} else if bodyByte, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ResponseError{
			Body:       bodyByte,
			StatusCode: resp.StatusCode,
		}
	} else if err := json.Unmarshal(bodyByte, &ret); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}

type GetRecentTradeListRequest struct {
	Symbol string
	Limit  *int
}

type GetRecentTradeListResponse GetOldTradeListResponse

func (e EndPointMarketData) GetRecentTradeList(ctx context.Context, req GetRecentTradeListRequest) ([]GetRecentTradeListResponse, error) {
	path := "/api/v3/trades"
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, e.conf.binanceUrl+path, nil)
	q := httpReq.URL.Query()
	q.Set("symbol", req.Symbol)
	if req.Limit != nil {
		q.Set("limit", strconv.Itoa(*req.Limit))
	}
	httpReq.URL.RawQuery = q.Encode()
	httpReq.Header.Set(headerAPIKey, e.conf.apiKey)
	var ret []GetRecentTradeListResponse
	if resp, err := http.DefaultClient.Do(httpReq); err != nil {
		return nil, err
	} else if bodyByte, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ResponseError{
			Body:       bodyByte,
			StatusCode: resp.StatusCode,
		}
	} else if err := json.Unmarshal(bodyByte, &ret); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}


type GetOldTradeListRequest struct {
	Symbol string
	FromId *int
	Limit  *int
}

type GetOldTradeListResponse struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int    `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func (e EndPointMarketData) GetOldTradeList(ctx context.Context, req GetOldTradeListRequest) ([]GetOldTradeListResponse, error) {
	path := "/api/v3/historicalTrades"
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, e.conf.binanceUrl+path, nil)
	q := httpReq.URL.Query()
	q.Set("symbol", req.Symbol)
	if req.Limit != nil {
		q.Set("limit", strconv.Itoa(*req.Limit))
	}
	if req.FromId != nil {
		q.Set("fromId", strconv.Itoa(*req.FromId))
	}
	httpReq.URL.RawQuery = q.Encode()
	httpReq.Header.Set(headerAPIKey, e.conf.apiKey)
	var ret []GetOldTradeListResponse
	if resp, err := http.DefaultClient.Do(httpReq); err != nil {
		return nil, err
	} else if bodyByte, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, ResponseError{
			Body:       bodyByte,
			StatusCode: resp.StatusCode,
		}
	} else if err := json.Unmarshal(bodyByte, &ret); err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}
