package bybit

type Coin string

type Side string

const (
	SideNone Side = "None"
	SideBuy  Side = "Buy"
	SideSell Side = "Sell"
)

type OrderType string

const (
	OrderTypeLimit  OrderType = "Limit"
	OrderTypeMarket OrderType = "Market"
)

type MarketUnit string

const (
	MarketUnitBaseCoin  MarketUnit = "baseCoin"
	MarketUnitQuoteCoin MarketUnit = "quoteCoin"
)

// OrderStatus :
type OrderStatus string

const (
	OrderStatusCreated         OrderStatus = "Created"
	OrderStatusRejected        OrderStatus = "Rejected"
	OrderStatusNew             OrderStatus = "New"
	OrderStatusPartiallyFilled OrderStatus = "PartiallyFilled"
	OrderStatusFilled          OrderStatus = "Filled"
	OrderStatusCancelled       OrderStatus = "Cancelled"
	OrderStatusPendingCancel   OrderStatus = "PendingCancel"
	OrderStatusDeactivated     OrderStatus = "Deactivated"
	OrderStatusTriggered       OrderStatus = "Triggered"
	OrderStatusActive          OrderStatus = "Active"
)

type Order string

const (
	OrderDesc Order = "desc"
	OrderAsc  Order = "asc"
)

type TimeInForce string

const (
	TimeInForceGoodTillCancel    TimeInForce = "GTC"
	TimeInForceImmediateOrCancel TimeInForce = "IOC"
	TimeInForceFillOrKill        TimeInForce = "FOK"
)

// Interval :
type Interval string

var Intervals = []Interval{
	Interval1,
	Interval3,
	Interval5,
	Interval15,
	Interval30,
	Interval60,
	Interval120,
	Interval240,
	Interval360,
	Interval720,
	IntervalD,
	IntervalW,
	IntervalM,
}

const (
	Interval1   = Interval("1")
	Interval3   = Interval("3")
	Interval5   = Interval("5")
	Interval15  = Interval("15")
	Interval30  = Interval("30")
	Interval60  = Interval("60")
	Interval120 = Interval("120")
	Interval240 = Interval("240")
	Interval360 = Interval("360")
	Interval720 = Interval("720")
	IntervalD   = Interval("D")
	IntervalW   = Interval("W")
	IntervalM   = Interval("M")
)

// TickDirection :
type TickDirection string

const (
	// TickDirectionPlusTick :
	TickDirectionPlusTick = TickDirection("PlusTick")
	// TickDirectionZeroPlusTick :
	TickDirectionZeroPlusTick = TickDirection("ZeroPlusTick")
	// TickDirectionMinusTick :
	TickDirectionMinusTick = TickDirection("MinusTick")
	// TickDirectionZeroMinusTick :
	TickDirectionZeroMinusTick = TickDirection("ZeroMinusTick")
)

// Period :
type Period string

const (
	// Period5min :
	Period5min = Period("5min")
	// Period15min :
	Period15min = Period("15min")
	// Period30min :
	Period30min = Period("30min")
	// Period1h :
	Period1h = Period("1h")
	// Period4h :
	Period4h = Period("4h")
	// Period1d :
	Period1d = Period("1d")
)

// TpSlMode :
type TpSlMode string

const (
	// TpSlModeFull :
	TpSlModeFull = TpSlMode("Full")
	// TpSlModePartial :
	TpSlModePartial = TpSlMode("Partial")
)

// ExecType :
type ExecType string

const (
	// ExecTypeTrade :
	ExecTypeTrade = ExecType("Trade")
	// ExecTypeAdlTrade :
	ExecTypeAdlTrade = ExecType("AdlTrade")
	// ExecTypeFunding :
	ExecTypeFunding = ExecType("Funding")
	// ExecTypeBustTrade :
	ExecTypeBustTrade = ExecType("BustTrade")
)

// Direction :
type Direction string

const (
	// DirectionPrev :
	DirectionPrev = Direction("prev")
	// DirectionNext :
	DirectionNext = Direction("next")
)
