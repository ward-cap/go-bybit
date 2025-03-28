package bybit

import (
	"context"
	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

// V5ExecutionServiceI :
type V5ExecutionServiceI interface {
	GetExecutionList(context.Context, V5GetExecutionParam) (*V5GetExecutionListResponse, error)
}

// V5ExecutionService :
type V5ExecutionService struct {
	client *Client
}

// V5GetExecutionParam :
type V5GetExecutionParam struct {
	Category CategoryV5 `url:"category"`

	Symbol      *string     `url:"symbol,omitempty"`
	OrderID     *string     `url:"orderId,omitempty"`
	OrderLinkID *string     `url:"orderLinkId,omitempty"`
	BaseCoin    *string     `url:"baseCoin,omitempty"`
	StartTime   *int64      `url:"startTime,omitempty"`
	EndTime     *int64      `url:"endTime,omitempty"`
	ExecType    *ExecTypeV5 `url:"execType,omitempty"`
	Limit       *int        `url:"limit,omitempty"`
	Cursor      *string     `url:"cursor,omitempty"`
}

// V5GetExecutionListResponse :
type V5GetExecutionListResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetExecutionListResult `json:"result"`
}

// V5GetExecutionListResult :
type V5GetExecutionListResult struct {
	NextPageCursor string                   `json:"nextPageCursor"`
	Category       CategoryV5               `json:"category"`
	List           []V5GetExecutionListItem `json:"list"`
}

// V5GetExecutionListItem :
type V5GetExecutionListItem struct {
	Symbol          string          `json:"symbol"`
	OrderID         string          `json:"orderId"`
	OrderLinkID     string          `json:"orderLinkId"`
	Side            Side            `json:"side"`
	OrderPrice      string          `json:"orderPrice"`
	OrderQty        string          `json:"orderQty"`
	LeavesQty       string          `json:"leavesQty"`
	OrderType       OrderType       `json:"orderType"`
	StopOrderType   string          `json:"stopOrderType"`
	ExecFee         decimal.Decimal `json:"execFee"`
	FeeCurrency     string          `json:"feeCurrency"`
	ExecID          string          `json:"execId"`
	ExecPrice       decimal.Decimal `json:"execPrice"`
	ExecQty         decimal.Decimal `json:"execQty"`
	ExecType        ExecTypeV5      `json:"execType"`
	ExecValue       string          `json:"execValue"`
	ExecTime        string          `json:"execTime"`
	IsMaker         bool            `json:"isMaker"`
	FeeRate         string          `json:"feeRate"`
	TradeIv         string          `json:"tradeIv"`
	MarkIv          string          `json:"markIv"`
	MarkPrice       string          `json:"markPrice"`
	IndexPrice      string          `json:"indexPrice"`
	UnderlyingPrice string          `json:"underlyingPrice"`
	BlockTradeID    string          `json:"blockTradeId"`
	ClosedSize      string          `json:"closedSize"`
}

func (s *V5ExecutionService) GetExecutionList(ctx context.Context, param V5GetExecutionParam) (res *V5GetExecutionListResponse, err error) {
	queryString, err := query.Values(param)
	if err != nil {
		return
	}

	err = s.client.getV5PrivatelyCtx(ctx, "/v5/execution/list", queryString, &res)

	return
}
