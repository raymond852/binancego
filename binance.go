package binancego

const (
	baseAPIMainURL    = "https://api.binance.com"
	baseAPITestnetURL = "https://testnet.binance.vision"

	headerAPIKey = "X-MBX-APIKEY"
)

type Binance struct {
	apiKey       string
	apiSecretKey string
	binanceUrl   string
}

func NewBinanceClient(apiKey string, apiSecretKey string) Binance {
	return Binance{
		apiKey:       apiKey,
		apiSecretKey: apiSecretKey,
		binanceUrl:   baseAPIMainURL,
	}
}

func (b Binance) UsingTestNet(val bool) Binance {
	b.binanceUrl = baseAPITestnetURL
	return b
}

func (b Binance) GeneralEndpoint() EndPointGeneral {
	return EndPointGeneral{
		conf: b,
	}
}

func (b Binance) MarketDataEndpoint() EndPointMarketData {
	return EndPointMarketData{
		conf: b,
	}
}

func (b Binance) AccountDataEndpoint() EndPointAccount {
	return EndPointAccount{
		conf: b,
	}
}
