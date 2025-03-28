package bybit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/shopspring/decimal"
)

type V5OrderServiceI interface {
	CreateOrder(context.Context, V5CreateOrderParam) (*V5CreateOrderResponse, error)
	AmendOrder(context.Context, V5AmendOrderParam) (*V5AmendOrderResponse, error)
	CancelOrder(context.Context, V5CancelOrderParam) (*V5CancelOrderResponse, error)
	GetOpenOrders(context.Context, V5GetOpenOrdersParam) (*V5GetOrdersResponse, error)
	CancelAllOrders(context.Context, V5CancelAllOrdersParam) (*V5CancelAllOrdersResponse, error)
	GetHistoryOrders(context.Context, V5GetHistoryOrdersParam) (*V5GetOrdersResponse, error)
}

type V5OrderService struct {
	client *Client
}

type V5CreateOrderParam struct {
	Category  CategoryV5 `json:"category"`
	Symbol    string     `json:"symbol"`
	Side      Side       `json:"side"`
	OrderType OrderType  `json:"orderType"`
	Qty       string     `json:"qty"`

	IsLeverage            *IsLeverage       `json:"isLeverage,omitempty"`
	Price                 *string           `json:"price,omitempty"`
	TriggerDirection      *TriggerDirection `json:"triggerDirection,omitempty"`
	OrderFilter           *OrderFilter      `json:"orderFilter,omitempty"` // If not passed, Order by default
	TriggerPrice          *string           `json:"triggerPrice,omitempty"`
	TriggerBy             *TriggerBy        `json:"triggerBy,omitempty"`
	OrderIv               *string           `json:"orderIv,omitempty"`     // option only.
	TimeInForce           *TimeInForce      `json:"timeInForce,omitempty"` // If not passed, GTC is used by default
	PositionIdx           *PositionIdx      `json:"positionIdx,omitempty"` // Under hedge-mode, this param is required
	OrderLinkID           *string           `json:"orderLinkId,omitempty"`
	TakeProfit            *string           `json:"takeProfit,omitempty"`
	StopLoss              *string           `json:"stopLoss,omitempty"`
	TpTriggerBy           *TriggerBy        `json:"tpTriggerBy,omitempty"`
	SlTriggerBy           *TriggerBy        `json:"slTriggerBy,omitempty"`
	ReduceOnly            *bool             `json:"reduce_only,omitempty"`
	CloseOnTrigger        *bool             `json:"closeOnTrigger,omitempty"`
	SmpType               *string           `json:"smpType,omitempty"`
	MarketMakerProtection *bool             `json:"mmp,omitempty"` // option only
	TpSlMode              *TpSlMode         `json:"tpslMode,omitempty"`
	TpLimitPrice          *string           `json:"tpLimitPrice,omitempty"`
	SlLimitPrice          *string           `json:"slLimitPrice,omitempty"`
	TpOrderType           *OrderType        `json:"tpOrderType,omitempty"`
	SlOrderType           *OrderType        `json:"slOrderType,omitempty"`
	MarketUnit            *MarketUnit       `json:"marketUnit,omitempty"` // The unit for qty when create Spot market orders for UTA account.
}

type V5CreateOrderResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CreateOrderResult `json:"result"`
}

type V5CreateOrderResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

func (s *V5OrderService) CreateOrder(ctx context.Context, param V5CreateOrderParam) (*V5CreateOrderResponse, error) {
	var res V5CreateOrderResponse

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/order/create", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

type V5AmendOrderParam struct {
	Category CategoryV5 `json:"category"`
	Symbol   string     `json:"symbol"`

	OrderID      *string    `json:"orderId,omitempty"`
	OrderLinkID  *string    `json:"orderLinkId,omitempty"`
	OrderIv      *string    `json:"orderIv,omitempty"`
	TriggerPrice *string    `json:"triggerPrice,omitempty"`
	Qty          *string    `json:"qty,omitempty"`
	Price        *string    `json:"price,omitempty"`
	TakeProfit   *string    `json:"takeProfit,omitempty"`
	StopLoss     *string    `json:"stopLoss,omitempty"`
	TpTriggerBy  *TriggerBy `json:"tpTriggerBy,omitempty"`
	SlTriggerBy  *TriggerBy `json:"slTriggerBy,omitempty"`
	TriggerBy    *TriggerBy `json:"triggerBy,omitempty"`
}

func (p V5AmendOrderParam) validate() error {
	if p.OrderID == nil && p.OrderLinkID == nil {
		return fmt.Errorf("orderId or orderLinkId must be passed")
	}
	if p.Category != CategoryV5Option && p.OrderIv != nil {
		return fmt.Errorf("orderIv is for option only")
	}
	return nil
}

// V5AmendOrderResponse :
type V5AmendOrderResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5AmendOrderResult `json:"result"`
}

// V5AmendOrderResult :
type V5AmendOrderResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

func (s *V5OrderService) AmendOrder(ctx context.Context, param V5AmendOrderParam) (*V5AmendOrderResponse, error) {
	var res V5AmendOrderResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/order/amend", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

type V5CancelOrderParam struct {
	Category CategoryV5 `json:"category"`
	Symbol   string     `json:"symbol"`

	OrderID     *string      `json:"orderId,omitempty"`
	OrderLinkID *string      `json:"orderLinkId,omitempty"`
	OrderFilter *OrderFilter `json:"orderFilter,omitempty"` // If not passed, Order by default
}

type V5CancelOrderResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CancelOrderResult `json:"result"`
}

type V5CancelOrderResult struct {
	OrderID     string `json:"orderId"`
	OrderLinkID string `json:"orderLinkId"`
}

func (s *V5OrderService) CancelOrder(ctx context.Context, param V5CancelOrderParam) (*V5CancelOrderResponse, error) {
	var res V5CancelOrderResponse

	if param.OrderID == nil && param.OrderLinkID == nil {
		return nil, fmt.Errorf("either OrderID or OrderLinkID needed")
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/order/cancel", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}

type V5GetOpenOrdersParam struct {
	Category CategoryV5 `url:"category"`

	Symbol      *string      `url:"symbol,omitempty"`
	BaseCoin    *string      `url:"baseCoin,omitempty"`
	SettleCoin  *string      `url:"settleCoin,omitempty"`
	OrderID     *string      `url:"orderId,omitempty"`
	OrderLinkID *string      `url:"orderLinkId,omitempty"`
	OpenOnly    *int         `url:"openOnly,omitempty"`
	OrderFilter *OrderFilter `url:"orderFilter,omitempty"` // If not passed, Order by default
	Limit       *int         `url:"limit,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
}

type V5GetHistoryOrdersParam struct {
	Category CategoryV5 `url:"category"`

	Symbol      *string      `url:"symbol,omitempty"`
	BaseCoin    *string      `url:"baseCoin,omitempty"`
	OrderID     *string      `url:"orderId,omitempty"`
	OrderLinkID *string      `url:"orderLinkId,omitempty"`
	OrderFilter *OrderFilter `url:"orderFilter,omitempty"` // If not passed, Order by default
	OrderStatus *OrderStatus `url:"orderStatus,omitempty"`
	StartTime   *int64       `url:"startTime,omitempty"`
	EndTime     *int64       `url:"endTime,omitempty"`
	Limit       *int         `url:"limit,omitempty"`
	Cursor      *string      `url:"cursor,omitempty"`
}

// V5GetOrdersResponse :
type V5GetOrdersResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetOrdersResult `json:"result"`
}

// V5GetOrdersResult :
type V5GetOrdersResult struct {
	Category       CategoryV5   `json:"category"`
	NextPageCursor string       `json:"nextPageCursor"`
	List           []V5GetOrder `json:"list"`
}

type V5GetOrder struct {
	Symbol             string              `json:"symbol"`
	OrderType          OrderType           `json:"orderType"`
	OrderLinkID        string              `json:"orderLinkId"`
	OrderID            string              `json:"orderId"`
	CancelType         string              `json:"cancelType"`
	AvgPrice           string              `json:"avgPrice"`
	StopOrderType      string              `json:"stopOrderType"`
	LastPriceOnCreated string              `json:"lastPriceOnCreated"`
	OrderStatus        OrderStatus         `json:"orderStatus"`
	TakeProfit         string              `json:"takeProfit"`
	CumExecValue       string              `json:"cumExecValue"`
	TriggerDirection   int                 `json:"triggerDirection"`
	IsLeverage         string              `json:"isLeverage"`
	RejectReason       string              `json:"rejectReason"`
	Price              decimal.Decimal     `json:"price"`
	OrderIv            string              `json:"orderIv"`
	CreatedTime        string              `json:"createdTime"`
	TpTriggerBy        string              `json:"tpTriggerBy"`
	PositionIdx        int                 `json:"positionIdx"`
	TimeInForce        TimeInForce         `json:"timeInForce"`
	LeavesValue        string              `json:"leavesValue"`
	UpdatedTime        string              `json:"updatedTime"`
	Side               Side                `json:"side"`
	TriggerPrice       NullDecimalV2       `json:"triggerPrice"`
	CumExecFee         string              `json:"cumExecFee"`
	LeavesQty          string              `json:"leavesQty"`
	SlTriggerBy        string              `json:"slTriggerBy"`
	CloseOnTrigger     bool                `json:"closeOnTrigger"`
	CumExecQty         decimal.NullDecimal `json:"cumExecQty"`
	ReduceOnly         bool                `json:"reduceOnly"`
	Qty                decimal.Decimal     `json:"qty"`
	StopLoss           string              `json:"stopLoss"`
	TriggerBy          TriggerBy           `json:"triggerBy"`
}

func (s *V5OrderService) GetOpenOrders(ctx context.Context, param V5GetOpenOrdersParam) (*V5GetOrdersResponse, error) {
	var res V5GetOrdersResponse

	if param.Category == "" {
		return nil, fmt.Errorf("category needed")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5PrivatelyCtx(ctx, "/v5/order/realtime", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *V5OrderService) GetHistoryOrders(ctx context.Context, param V5GetHistoryOrdersParam) (*V5GetOrdersResponse, error) {
	var res V5GetOrdersResponse

	if param.Category == "" {
		return nil, fmt.Errorf("category needed")
	}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getV5PrivatelyCtx(ctx, "/v5/order/history", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type V5CancelAllOrdersParam struct {
	Category CategoryV5 `json:"category"`

	Symbol      *string      `json:"symbol,omitempty"`
	BaseCoin    *string      `json:"baseCoin,omitempty"`
	SettleCoin  *string      `json:"settleCoin,omitempty"`
	OrderFilter *OrderFilter `json:"orderFilter,omitempty"` // If not passed, Order by default
}

func (p V5CancelAllOrdersParam) validate() error {
	if p.Category == CategoryV5Linear || p.Category == CategoryV5Inverse {
		if p.Symbol == nil && p.BaseCoin == nil && p.SettleCoin == nil {
			return fmt.Errorf("symbol or baseCoin or settleCoin is needed for linear and inverse")
		}
	}
	if p.Category != CategoryV5Spot && p.OrderFilter != nil {
		return fmt.Errorf("orderFilter is for spot only")
	}
	return nil
}

type V5CancelAllOrdersResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5CancelAllOrdersResult `json:"result"`
}

type V5CancelAllOrdersResult struct {
	LinearInverseOption *V5CancelAllOrdersLinearInverseOptionResult
	Spot                *V5CancelAllOrdersSpotResult
}

func (r *V5CancelAllOrdersResult) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &r.LinearInverseOption); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &r.Spot); err != nil {
		return err
	}
	return nil
}

type V5CancelAllOrdersLinearInverseOptionResult struct {
	List []struct {
		OrderID     string `json:"orderId"`
		OrderLinkID string `json:"orderLinkId"`
	} `json:"list"`
}

type V5CancelAllOrdersSpotResult struct {
	Success string `json:"success"` // 1: success, 0: fail
}

func (s *V5OrderService) CancelAllOrders(ctx context.Context, param V5CancelAllOrdersParam) (*V5CancelAllOrdersResponse, error) {
	var res V5CancelAllOrdersResponse

	if err := param.validate(); err != nil {
		return nil, fmt.Errorf("validate param: %w", err)
	}

	body, err := json.Marshal(param)
	if err != nil {
		return &res, fmt.Errorf("json marshal: %w", err)
	}

	if err := s.client.postV5JSON(ctx, "/v5/order/cancel-all", body, &res); err != nil {
		return &res, err
	}

	return &res, nil
}
