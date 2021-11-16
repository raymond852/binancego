package binancego

import (
	"context"
	"io/ioutil"
	"net/http"
)

type EndPointGeneral struct {
	conf Binance
}

func (e EndPointGeneral) TestConnectivity(ctx context.Context) error {
	path := "/api/v3/ping"
	httpReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, e.conf.binanceUrl+path, nil)
	if resp, err := http.DefaultClient.Do(httpReq); err != nil {
		return err
	} else if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		return nil
	} else if bodyByte, err:= ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		return ResponseError{
			Body:       bodyByte,
			StatusCode: resp.StatusCode,
		}
	}
}
