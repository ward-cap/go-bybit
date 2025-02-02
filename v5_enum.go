package bybit

type MarginMode string

const (
	MarginModeRegular   MarginMode = "REGULAR_MARGIN"
	MarginModePortfolio MarginMode = "PORTFOLIO_MARGIN"
)

type CategoryV5 string

const (
	CategoryV5Spot    CategoryV5 = "spot"
	CategoryV5Linear  CategoryV5 = "linear"
	CategoryV5Inverse CategoryV5 = "inverse"
	CategoryV5Option  CategoryV5 = "option"
)

type TriggerDirection int

const (
	TriggerDirectionRise TriggerDirection = iota + 1
	TriggerDirectionFall
)

type IsLeverage int

const (
	IsLeverageFalse IsLeverage = iota
	IsLeverageTrue
)

type OrderFilter string

const (
	OrderFilterOrder     OrderFilter = "Order"
	OrderFilterStopOrder OrderFilter = "StopOrder"
	OrderFilterTpSlOrder OrderFilter = "tpslOrder"
)

// TriggerBy :
type TriggerBy string

const (
	TriggerByLastPrice  TriggerBy = "LastPrice"
	TriggerByIndexPrice TriggerBy = "IndexPrice"
	TriggerByMarkPrice  TriggerBy = "MarkPrice"
)

// PositionIdx :
type PositionIdx int

// PositionIdx :
const (
	PositionIdxOneWay PositionIdx = iota
	PositionIdxHedgeBuy
	PositionIdxHedgeSell
)

type ContractType string

// ContractType :
const (
	ContractTypeInversePerpetual ContractType = "InversePerpetual"
	ContractTypeLinearPerpetual  ContractType = "LinearPerpetual"
	ContractTypeInverseFutures   ContractType = "InverseFutures"
)

type InstrumentStatus string

const (
	// linear & inverse:
	InstrumentStatusPending  InstrumentStatus = "Pending"
	InstrumentStatusTrading  InstrumentStatus = "Trading"
	InstrumentStatusSettling InstrumentStatus = "Settling"
	InstrumentStatusClosed   InstrumentStatus = "Closed"

	// option
	InstrumentStatusWaitingOnline InstrumentStatus = "WAITING_ONLINE"
	InstrumentStatusOnline        InstrumentStatus = "ONLINE"
	InstrumentStatusDelivering    InstrumentStatus = "DELIVERING"
	InstrumentStatusOffline       InstrumentStatus = "OFFLINE"

	// spot
	InstrumentStatusAvailable InstrumentStatus = "1"
)

// OptionsType :
type OptionsType string

// OptionsType :
const (
	OptionsTypeCall OptionsType = "Call"
	OptionsTypePut  OptionsType = "Put"
)

type Innovation string

const (
	InnovationFalse Innovation = "0"
	InnovationTrue  Innovation = "1"
)

type PositionMode int

const (
	PositionModeMergedSingle PositionMode = 0
	PositionModeBothSides    PositionMode = 3
)

type PositionMarginMode int

const (
	PositionMarginCross PositionMarginMode = iota
	PositionMarginIsolated
)

// ExecTypeV5 :
type ExecTypeV5 string

const (
	ExecTypeV5Trade            ExecTypeV5 = "Trade"
	ExecTypeV5AdlTrade         ExecTypeV5 = "AdlTrade"
	ExecTypeV5Funding          ExecTypeV5 = "Funding"
	ExecTypeV5BustTrade        ExecTypeV5 = "BustTrade"
	ExecTypeV5Delivery         ExecTypeV5 = "Delivery"
	ExecTypeV5SessionSettlePnL ExecTypeV5 = "SessionSettlePnL"
	ExecTypeV5Settle           ExecTypeV5 = "Settle"
	ExecTypeV5BlockTrade       ExecTypeV5 = "BlockTrade"
	ExecTypeV5MovePosition     ExecTypeV5 = "MovePosition"
	ExecTypeV5UNKNOWN          ExecTypeV5 = "UNKNOWN"
)

// TransferStatusV5 :
type TransferStatusV5 string

const (
	TransferStatusV5SUCCESS TransferStatusV5 = "SUCCESS"
	TransferStatusV5PENDING TransferStatusV5 = "PENDING"
	TransferStatusV5FAILED  TransferStatusV5 = "FAILED"
)

// AccountTypeV5 :
type AccountTypeV5 string

const (
	AccountTypeV5CONTRACT   AccountTypeV5 = "CONTRACT"
	AccountTypeV5SPOT       AccountTypeV5 = "SPOT"
	AccountTypeV5INVESTMENT AccountTypeV5 = "INVESTMENT"
	AccountTypeV5OPTION     AccountTypeV5 = "OPTION"
	AccountTypeV5UNIFIED    AccountTypeV5 = "UNIFIED"
	AccountTypeV5FUND       AccountTypeV5 = "FUND"
)

// UnifiedMarginStatus :
type UnifiedMarginStatus int

const (
	UnifiedMarginStatusRegular UnifiedMarginStatus = iota + 1
	UnifiedMarginStatusUnifiedMargin
	UnifiedMarginStatusUnifiedTrade
	UnifiedMarginStatusUTAPro
)

// TransactionLogTypeV5 :
type TransactionLogTypeV5 string

const (
	TransactionLogTypeV5TRANSFERIN   TransactionLogTypeV5 = "TRANSFER_IN"
	TransactionLogTypeV5TRANSFEROUT  TransactionLogTypeV5 = "TRANSFER_OUT"
	TransactionLogTypeV5TRADE        TransactionLogTypeV5 = "TRADE"
	TransactionLogTypeV5SETTLEMENT   TransactionLogTypeV5 = "SETTLEMENT"
	TransactionLogTypeV5DELIVERY     TransactionLogTypeV5 = "DELIVERY"
	TransactionLogTypeV5LIQUIDATION  TransactionLogTypeV5 = "LIQUIDATION"
	TransactionLogTypeV5BONUS        TransactionLogTypeV5 = "BONUS"
	TransactionLogTypeV5FEEREFUND    TransactionLogTypeV5 = "FEE_REFUND"
	TransactionLogTypeV5INTEREST     TransactionLogTypeV5 = "INTEREST"
	TransactionLogTypeV5CURRENCYBUY  TransactionLogTypeV5 = "CURRENCY_BUY"
	TransactionLogTypeV5CURRENCYSELL TransactionLogTypeV5 = "CURRENCY_SELL"
)

// InternalDepositStatusV5 :
type InternalDepositStatusV5 int

const (
	InternalDepositStatusV5Processing InternalDepositStatusV5 = iota + 1
	InternalDepositStatusV5Success
	InternalDepositStatusV5Failed
)

// DepositStatusV5 :
type DepositStatusV5 int

const (
	DepositStatusV5Unknown DepositStatusV5 = iota
	DepositStatusV5ToBeConfirmed
	DepositStatusV5Processing
	DepositStatusV5Success
	DepositStatusV5Failed
)

type WithdrawTypeV5 int

const (
	WithdrawTypeOnChain WithdrawTypeV5 = iota
	WithdrawTypeOffChain
	WithdrawTypeAll
)

type WithdrawStatusV5 string

const (
	WithdrawStatusV5SecurityCheck       WithdrawStatusV5 = "SecurityCheck"
	WithdrawStatusV5Pending             WithdrawStatusV5 = "Pending"
	WithdrawStatusV5Success             WithdrawStatusV5 = "success"
	WithdrawStatusV5CancelByUser        WithdrawStatusV5 = "CancelByUser"
	WithdrawStatusV5Reject              WithdrawStatusV5 = "Reject"
	WithdrawStatusV5Fail                WithdrawStatusV5 = "Fail"
	WithdrawStatusV5BlockchainConfirmed WithdrawStatusV5 = "BlockchainConfirmed"
)

type IsLowestRisk int

const (
	IsLowestRiskFalse IsLowestRisk = iota
	IsLowestRiskTrue
)

type CollateralSwitchV5 string

const (
	CollateralSwitchV5On  CollateralSwitchV5 = "ON"
	CollateralSwitchV5Off CollateralSwitchV5 = "OFF"
)

// AdlRankIndicator : Auto-deleverage rank indicator
type AdlRankIndicator int

const (
	AdlRankIndicator0 AdlRankIndicator = iota
	AdlRankIndicator1
	AdlRankIndicator2
	AdlRankIndicator3
	AdlRankIndicator4
	AdlRankIndicator5
)
