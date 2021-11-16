package binancego

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type EndPointAccount struct {
	conf Binance
}

type EnumOrderSide string

const (
	ENUM_ORDER_SIDE_SELL EnumOrderSide = "SELL"
	ENUM_ORDER_SIDE_BUY  EnumOrderSide = "BUY"
)

type EnumOrderType string

const (
	ENUM_ORDER_TYPE_LIMIT             EnumOrderType = "LIMIT"
	ENUM_ORDER_TYPE_MARKET            EnumOrderType = "MARKET"
	ENUM_ORDER_TYPE_STOP_LOSS         EnumOrderType = "STOP_LOSS"
	ENUM_ORDER_TYPE_STOP_LOSS_LIMIT   EnumOrderType = "STOP_LOSS_LIMIT"
	ENUM_ORDER_TYPE_TAKE_PROFIT       EnumOrderType = "TAKE_PROFIT"
	ENUM_ORDER_TYPE_TAKE_PROFIT_LIMIT EnumOrderType = "TAKE_PROFIT_LIMIT"
	ENUM_ORDER_TYPE_LIMIT_MAKER       EnumOrderType = "LIMIT_MAKER"
)

type EnumTimeInForce string

const (
	ENUM_ORDER_TIME_IN_FORCE_GTC EnumTimeInForce = "GTC"
	ENUM_ORDER_TIME_IN_FORCE_IOC EnumTimeInForce = "IOC"
	ENUM_ORDER_TIME_IN_FORCE_FOK EnumTimeInForce = "FOK"
)

type EnumOrderNewOrderRespType string

const (
	ENUM_ORDER_NEW_ENUM_ORDER_RESP_TYPE_ACK    EnumOrderNewOrderRespType = "ACK"
	ENUM_ORDER_NEW_ENUM_ORDER_RESP_TYPE_RESULT EnumOrderNewOrderRespType = "RESULT"
)

type CreateOrderRequest struct {
	Symbol           string                     `url:"symbol"`
	Side             EnumOrderSide              `url:"side"`
	Type             EnumOrderType              `url:"type"`
	TimeInForce      *EnumTimeInForce           `url:"timeInForce,omitempty"`
	Quantity         *float64                   `url:"quantity,omitempty"`
	QuoteOrderQty    *float64                   `url:"quoteOrderQty,omitempty"`
	Price            *float64                   `url:"price,omitempty"`
	NewClientOrderId *string                    `url:"newClientOrderId,omitempty"`
	StopPrice        *float64                   `url:"stopPrice,omitempty"`
	IcebergQty       *float64                   `url:"icebergQty,omitempty"`
	NewOrderRespType *EnumOrderNewOrderRespType `url:"newOrderRespType,omitempty"`
	RecvWindow       *int64                     `url:"recvWindow,omitempty"`
	Timestamp        int64                      `url:"timestamp,omitempty"`
}

func (tr CreateOrderRequest) Validate() error {
	missingFields := make([]string, 0, 3)
	switch tr.Type {
	case ENUM_ORDER_TYPE_LIMIT:
		if tr.TimeInForce == nil {
			missingFields = append(missingFields, "timeInForce")
		}
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if tr.Price == nil {
			missingFields = append(missingFields, "price")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_MARKET:
		if tr.Quantity == nil && tr.QuoteOrderQty == nil {
			missingFields = append(missingFields, "quantity or quoteOrderQty")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_STOP_LOSS:
		if tr.StopPrice == nil {
			missingFields = append(missingFields, "stopPrice")
		}
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_STOP_LOSS_LIMIT:
		if tr.TimeInForce == nil {
			missingFields = append(missingFields, "timeInForce")
		}
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if tr.Price == nil {
			missingFields = append(missingFields, "price")
		}
		if tr.StopPrice == nil {
			missingFields = append(missingFields, "stopPrice")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_TAKE_PROFIT:
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if tr.StopPrice == nil {
			missingFields = append(missingFields, "stopPrice")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_TAKE_PROFIT_LIMIT:
		if tr.TimeInForce == nil {
			missingFields = append(missingFields, "timeInForce")
		}
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if tr.Price == nil {
			missingFields = append(missingFields, "price")
		}
		if tr.StopPrice == nil {
			missingFields = append(missingFields, "stopPrice")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	case ENUM_ORDER_TYPE_LIMIT_MAKER:
		if tr.Quantity == nil {
			missingFields = append(missingFields, "quantity")
		}
		if tr.Price == nil {
			missingFields = append(missingFields, "price")
		}
		if len(missingFields) > 0 {
			return BadRequestError{
				MissingFields: missingFields,
			}
		}
	default:
		return BadRequestError{
			IncorrectFields: []string{"type"},
		}
	}
	return nil
}

type CreateOrderFullResponse struct {
	Symbol              string                          `json:"symbol"`
	OrderId             int                             `json:"orderId"`
	OrderListId         int                             `json:"orderListId"`
	ClientOrderId       string                          `json:"clientOrderId"`
	TransactTime        int                             `json:"transactTime"`
	Price               string                          `json:"price"`
	OrigQty             string                          `json:"origQty"`
	ExecutedQty         string                          `json:"executedQty"`
	CummulativeQuoteQty string                          `json:"cummulativeQuoteQty"`
	Status              string                          `json:"status"`
	TimeInForce         string                          `json:"timeInForce"`
	Type                string                          `json:"type"`
	Side                string                          `json:"side"`
	Fills               []CreateOrderFullResponse_Fills `json:"fills"`
}

type CreateOrderFullResponse_Fills struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

type CreateOrderResultResponse struct {
	Symbol              string `json:"symbol"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int    `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

type CreateOrderAckResponse struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	OrderListId   int    `json:"orderListId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int    `json:"transactTime"`
}

func (e EndPointAccount) CreateOrder(ctx context.Context, req CreateOrderRequest, ret interface{}) error {
	if err := req.Validate(); err != nil {
		return err
	}

	path := "/api/v3/order"
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodPost, e.conf.binanceUrl+path, nil)
	v, _ := query.Values(req)
	h := hmac.New(sha256.New, []byte(e.conf.apiSecretKey))
	h.Write([]byte(v.Encode()))
	httpReq.URL.RawQuery = v.Encode() + "&signature=" + hex.EncodeToString(h.Sum(nil))
	fmt.Println(httpReq.URL.RawQuery)
	httpReq.Header.Set(headerAPIKey, e.conf.apiKey)

	if resp, err := http.DefaultClient.Do(httpReq); err != nil {
		return err
	} else if bodyByte, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return ResponseError{
			Body:       bodyByte,
			StatusCode: resp.StatusCode,
		}
	} else {
		if ret != nil {
			if err := json.Unmarshal(bodyByte, ret); err != nil {
				return err
			}
		}
		return nil
	}
}
