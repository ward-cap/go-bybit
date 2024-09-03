//go:build integrationtestfutureusdtperpetual

package integrationtestfutureusdtperpetual

import (
	"github.com/hirokisan/bybit/v2/src"
	"testing"
	"time"

	"github.com/hirokisan/bybit/v2"
	"github.com/hirokisan/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestAPIKeyInfo(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().APIKeyInfo()
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-api-key-info.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().APIKeyInfo()
		require.Error(t, err)
	})
}

func TestBalance(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().Balance(src.CoinUSDT)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().Balance(src.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().OrderBook(src.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListLinearKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().ListLinearKline(src.ListLinearKlineParam{
		Symbol:   src.SymbolFutureBTCUSDT,
		Interval: src.Interval120,
		From:     time.Now().AddDate(0, 0, -1).Unix(),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/public-linear-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().Tickers(src.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().Symbols()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().OpenInterest(src.OpenInterestParam{
		Symbol: src.SymbolFutureBTCUSD,
		Period: src.Period1h,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-open-interest.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestBigDeal(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().USDTPerpetual().BigDeal(src.BigDealParam{
		Symbol: src.SymbolFutureBTCUSD,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-big-deal.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestAccountRatio(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().USDTPerpetual().AccountRatio(src.AccountRatioParam{
		Symbol: src.SymbolFutureBTCUSD,
		Period: src.Period1h,
		Limit:  &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-account-ratio.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestCreateLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSDT,
			OrderType:   src.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-order-create.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSDT,
			OrderType:   src.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListLinearOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	symbol := src.SymbolFutureBTCUSDT

	var orderID string
	{
		price := 10000.0
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
			Side:        src.SideBuy,
			Symbol:      symbol,
			OrderType:   src.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.NoError(t, err)
		orderID = res.Result.OrderID
	}

	{
		orderStatus := src.OrderStatusNew
		res, err := client.Future().USDTPerpetual().ListLinearOrder(src.ListLinearOrderParam{
			Symbol:      symbol,
			OrderStatus: &orderStatus,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-order-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}

	{
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{
			Symbol:  symbol,
			OrderID: &orderID,
		})
		require.NoError(t, err)
	}
}

func TestListLinearPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().ListLinearPosition(src.SymbolFutureBTCUSDT)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ListLinearPosition(src.SymbolFutureBTCUSDT)
		require.Error(t, err)
	})
}

func TestListLinearPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().ListLinearPositions()
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ListLinearPositions()
		require.Error(t, err)
	})
}

func TestCancelLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		var orderID string
		{
			price := 47000.0
			res, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSDT,
				OrderType:   src.OrderTypeLimit,
				Qty:         0.001,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		res, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{
			Symbol:  src.SymbolFutureBTCUSDT,
			OrderID: &orderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestSaveLinearLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		_, err := client.Future().USDTPerpetual().SaveLinearLeverage(src.SaveLinearLeverageParam{
			Symbol:       src.SymbolFutureBTCUSDT,
			BuyLeverage:  2.0,
			SellLeverage: 2.0,
		})
		require.NoError(t, err)
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{})
		require.Error(t, err)
	})
}

func TestLinearExecutionList(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().LinearExecutionList(src.LinearExecutionListParam{
			Symbol: src.SymbolFutureBTCUSDT,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-trade-execution-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().LinearExecutionList(src.LinearExecutionListParam{})
		require.Error(t, err)
	})
}

func TestLinearCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().USDTPerpetual().LinearCancelAllOrder(src.LinearCancelAllParam{
			Symbol: src.SymbolFutureBTCUSDT,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-cancel-all-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().LinearCancelAllOrder(src.LinearCancelAllParam{})
		require.Error(t, err)
	})
}

func TestQueryLinearOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	symbol := src.SymbolFutureBTCUSDT

	var orderID string
	{
		price := 10000.0
		res, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
			Side:        src.SideBuy,
			Symbol:      symbol,
			OrderType:   src.OrderTypeLimit,
			Qty:         0.001,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.NoError(t, err)
		orderID = res.Result.OrderID
	}

	{
		res, err := client.Future().USDTPerpetual().QueryLinearOrder(src.QueryLinearOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-order-search.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}

	{
		_, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{
			Symbol:  symbol,
			OrderID: &orderID,
		})
		require.NoError(t, err)
	}
}

func TestCreateLinearStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := src.SymbolFutureBTCUSDT
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
				Side:           src.SideBuy,
				Symbol:         symbol,
				OrderType:      src.OrderTypeMarket,
				Qty:            0.001,
				BasePrice:      price,
				StopPx:         price + 200,
				TimeInForce:    src.TimeInForceGoodTillCancel,
				TriggerBy:      src.TriggerByFutureLastPrice,
				ReduceOnly:     true,
				CloseOnTrigger: true,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/private-linear-stop-order-create.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
			stopOrderID = res.Result.StopOrderID
		}
		{
			_, err := client.Future().USDTPerpetual().CancelLinearStopOrder(src.CancelLinearStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 19400.5
		_, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
			Side:           src.SideBuy,
			Symbol:         src.SymbolFutureBTCUSDT,
			OrderType:      src.OrderTypeMarket,
			Qty:            0.001,
			BasePrice:      price,
			StopPx:         price + 200,
			TimeInForce:    src.TimeInForceGoodTillCancel,
			TriggerBy:      src.TriggerByFutureLastPrice,
			ReduceOnly:     true,
			CloseOnTrigger: true,
		})
		require.Error(t, err)
	})
}

func TestListLinearStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := src.SymbolFutureBTCUSDT
		var stopOrderID string
		{
			price := 19800.5
			res, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
				Side:           src.SideBuy,
				Symbol:         symbol,
				OrderType:      src.OrderTypeMarket,
				Qty:            0.001,
				BasePrice:      price,
				StopPx:         price + 200,
				TimeInForce:    src.TimeInForceGoodTillCancel,
				TriggerBy:      src.TriggerByFutureLastPrice,
				ReduceOnly:     true,
				CloseOnTrigger: true,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}
		{
			status := src.OrderStatusUntriggered
			res, err := client.Future().USDTPerpetual().ListLinearStopOrder(src.ListLinearStopOrderParam{
				Symbol:          symbol,
				StopOrderStatus: &status,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/private-linear-stop-order-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		{
			_, err := client.Future().USDTPerpetual().CancelLinearStopOrder(src.CancelLinearStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ListLinearStopOrder(src.ListLinearStopOrderParam{})
		require.Error(t, err)
	})
}

func TestCancelLinearStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		symbol := src.SymbolFutureBTCUSDT
		var stopOrderID string
		{
			price := 19800.5
			res, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
				Side:           src.SideBuy,
				Symbol:         symbol,
				OrderType:      src.OrderTypeMarket,
				Qty:            0.001,
				BasePrice:      price,
				StopPx:         price + 200,
				TimeInForce:    src.TimeInForceGoodTillCancel,
				TriggerBy:      src.TriggerByFutureLastPrice,
				ReduceOnly:     true,
				CloseOnTrigger: true,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}
		res, err := client.Future().USDTPerpetual().CancelLinearStopOrder(src.CancelLinearStopOrderParam{
			Symbol:      symbol,
			StopOrderID: &stopOrderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-stop-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelLinearStopOrder(src.CancelLinearStopOrderParam{})
		require.Error(t, err)
	})
}

func TestCancelAllLinearStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		symbol := src.SymbolFutureBTCUSDT
		{
			price := 19800.5
			_, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
				Side:           src.SideBuy,
				Symbol:         symbol,
				OrderType:      src.OrderTypeMarket,
				Qty:            0.001,
				BasePrice:      price,
				StopPx:         price + 200,
				TimeInForce:    src.TimeInForceGoodTillCancel,
				TriggerBy:      src.TriggerByFutureLastPrice,
				ReduceOnly:     true,
				CloseOnTrigger: true,
			})
			require.NoError(t, err)
		}
		res, err := client.Future().USDTPerpetual().CancelAllLinearStopOrder(src.CancelAllLinearStopOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-stop-order-cancel-all.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().CancelAllLinearStopOrder(src.CancelAllLinearStopOrderParam{})
		require.Error(t, err)
	})
}

func TestQueryLinearStopOrder(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	symbol := src.SymbolFutureBTCUSDT

	var stopOrderID string
	{
		price := 19800.5
		res, err := client.Future().USDTPerpetual().CreateLinearStopOrder(src.CreateLinearStopOrderParam{
			Side:           src.SideBuy,
			Symbol:         symbol,
			OrderType:      src.OrderTypeMarket,
			Qty:            0.001,
			BasePrice:      price,
			StopPx:         price + 200,
			TimeInForce:    src.TimeInForceGoodTillCancel,
			TriggerBy:      src.TriggerByFutureLastPrice,
			ReduceOnly:     true,
			CloseOnTrigger: true,
		})
		require.NoError(t, err)
		stopOrderID = res.Result.StopOrderID
	}

	{
		res, err := client.Future().USDTPerpetual().QueryLinearStopOrder(src.QueryLinearStopOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/private-linear-stop-order-search.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	}

	{
		_, err := client.Future().USDTPerpetual().CancelLinearStopOrder(src.CancelLinearStopOrderParam{
			Symbol:      symbol,
			StopOrderID: &stopOrderID,
		})
		require.NoError(t, err)
	}
}

func TestLinearTradingStop(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()

	{
		_, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSDT,
			OrderType:   src.OrderTypeMarket,
			Qty:         0.001,
			TimeInForce: src.TimeInForceGoodTillCancel,
		})
		require.NoError(t, err)
	}

	{
		price := 20000.0
		_, err := client.Future().USDTPerpetual().LinearTradingStop(src.LinearTradingStopParam{
			Symbol:     src.SymbolFutureBTCUSDT,
			Side:       src.SideBuy,
			TakeProfit: &price,
		})
		require.NoError(t, err)
	}
}

func TestReplaceLinearOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()

		symbol := src.SymbolFutureBTCUSDT
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().USDTPerpetual().CreateLinearOrder(src.CreateLinearOrderParam{
				Side:        src.SideBuy,
				Symbol:      symbol,
				OrderType:   src.OrderTypeLimit,
				Qty:         0.001,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		{
			newPrice := 11000.0
			res, err := client.Future().USDTPerpetual().ReplaceLinearOrder(src.ReplaceLinearOrderParam{
				Symbol:   symbol,
				OrderID:  &orderID,
				NewPrice: &newPrice,
			})
			require.NoError(t, err)
			goldenFilename := "./testdata/private-linear-order-replace.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))

			orderID = res.Result.OrderID
		}
		{
			_, err := client.Future().USDTPerpetual().CancelLinearOrder(src.CancelLinearOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().USDTPerpetual().ReplaceLinearOrder(src.ReplaceLinearOrderParam{})
		require.Error(t, err)
	})
}
