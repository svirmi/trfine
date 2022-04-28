package initdata

type TradeInfo struct {
	ExternalID       uint64
	SellFilledOrders string
	SellOpenOrders   string
}

func GetTradeInfo() TradeInfo {
	tradeInfo := TradeInfo{
		ExternalID:       1,
		SellFilledOrders: "0",
		SellOpenOrders:   "0",
	}
	return tradeInfo
}

type TradeParams struct {
	ExternalID      uint64
	NameList        string
	MinBalance      string
	MinOrder        string
	MinPrice        string
	DailyPercent    string
	SellUpp         string
	BuyDown         string
	MaxTradePairs   string
	AutoTradePairs  bool
	DeltaPercent    bool
	NumAver         bool
	StepAver        string
	MaxAver         string
	QuantityAver    string
	TrailingStop    bool
	TrailingPercent string
	TrailingPart    string
	TrailingPrice   string
	UserOrder       bool
	FiatCurrencies  string
	QuoteAsset      string
	DoubleAsset     bool
	Active          bool
}

// TODO: делать списки параметров с привязкой квот ассет
func GetTradeParams() TradeParams {
	return TradeParams{
		ExternalID:      1,
		NameList:        "native",
		MinBalance:      "25",
		MinOrder:        "1.1",
		MinPrice:        "0.005",
		DailyPercent:    "-5",
		SellUpp:         "1.75",
		BuyDown:         "-5",
		MaxTradePairs:   "20",
		AutoTradePairs:  true,
		DeltaPercent:    true,
		NumAver:         true,
		StepAver:        "1.15",
		MaxAver:         "4",
		QuantityAver:    "2",
		TrailingStop:    false,
		TrailingPercent: "0.35",
		TrailingPart:    "10",
		TrailingPrice:   "0.15",
		UserOrder:       true,
		FiatCurrencies:  "RUB",
		QuoteAsset:      "USDT BTC",
		DoubleAsset:     false,
		Active:          true,
	}
}

// TODO: возможно не понадобится
type SystemKeys struct {
	Keys []string `json:"keys"`
}

func GetSystemKeys() SystemKeys {
	return SystemKeys{
		Keys: []string{
			"symbols",
			"trade_info",
			"api_key",
			"trade_params",
			"trade_params_list",
			"white_list",
			"trade_pairs",
			"trailing_orders",
			"daily_profit",
			"bnb_burn",
			"average_percent",
		},
	}
}
