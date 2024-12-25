package bybit

import (
	"context"
	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

type V5APIExchangeEarning struct {
	CommonV5Response `json:",inline"`
	Result           V5APIExchangeEarningResult `json:"result"`
}

type V5APIExchangeEarningResult struct {
	TotalEarningCat TotalEarningCategory `json:"totalEarningCat"`
	Details         []EarningDetail      `json:"details"`
	NextPageCursor  string               `json:"nextPageCursor"`
}

type TotalEarningCategory struct {
	Spot        []Earning `json:"spot"`
	Derivatives []Earning `json:"derivatives"`
	Options     []Earning `json:"options"`
	Convert     []Earning `json:"convert"`
	Total       []Earning `json:"total"`
}

type Earning struct {
	Coin    string `json:"coin"`
	Earning string `json:"earning"`
}

type EarningDetail struct {
	UserID         string          `json:"userId"`
	BizType        string          `json:"bizType"`
	Symbol         string          `json:"symbol"`
	Coin           string          `json:"coin"`
	Earning        decimal.Decimal `json:"earning"`
	MarkupEarning  string          `json:"markupEarning"`
	BaseFeeEarning string          `json:"baseFeeEarning"`
	OrderID        string          `json:"orderId"`
	ExecTime       string          `json:"execTime"`
}

type ExchangeEarningRequest struct {
	BizType *string `url:"bizType,omitempty"`
	Begin   *string `url:"begin,omitempty"`
	End     *string `url:"end,omitempty"`
	UID     *string `url:"uid,omitempty"`
	Limit   *uint   `url:"limit,omitempty"`
	Cursor  *string `url:"cursor,omitempty"`
}

func (s *V5MarketService) GetExchangeEarning(ctx context.Context, param ExchangeEarningRequest) (res *V5APIExchangeEarning, err error) {
	q, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	err = s.client.getV5PrivatelyCtx(ctx, "/v5/broker/earnings-info", q, &res)

	return
}
