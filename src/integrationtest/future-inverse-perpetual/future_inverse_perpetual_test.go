//go:build integrationtestfutureinverseperpetual

package integrationtestfutureinverseperpetual

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
		res, err := client.Future().InversePerpetual().APIKeyInfo()
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-api-key-info.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().APIKeyInfo()
		require.Error(t, err)
	})
}

func TestBalance(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual().Balance(src.CoinUSDT)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-wallet-balance.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().Balance(src.CoinBTC)
		require.Error(t, err)
	})
}

func TestOrderBook(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().OrderBook(src.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-order-book-l2.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestListKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().ListKline(src.ListKlineParam{
		Symbol:   src.SymbolFutureBTCUSD,
		Interval: src.Interval120,
		From:     time.Now().AddDate(0, 0, -1).Unix(),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-kline-list.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTickers(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().Tickers(src.SymbolFutureBTCUSD)
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-tickers.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestTradingRecords(t *testing.T) {
	client := bybit.NewTestClient()
	limit := 10
	res, err := client.Future().InversePerpetual().TradingRecords(src.TradingRecordsParam{
		Symbol: src.SymbolFutureBTCUSD,
		Limit:  &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-trading-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestSymbols(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().Symbols()
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-symbols.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestMarkPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InverseFuture().MarkPriceKline(src.MarkPriceKlineParam{
		Symbol:   src.SymbolFutureBTCUSD,
		Interval: src.IntervalD,
		From:     time.Now().AddDate(0, 0, -1).Unix(),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-mark-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestIndexPriceKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().IndexPriceKline(src.IndexPriceKlineParam{
		Symbol:   src.SymbolFutureBTCUSD,
		Interval: src.IntervalD,
		From:     time.Now().AddDate(0, 0, -1).Unix(),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-index-price-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestOpenInterest(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().OpenInterest(src.OpenInterestParam{
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
	res, err := client.Future().InversePerpetual().BigDeal(src.BigDealParam{
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
	res, err := client.Future().InversePerpetual().AccountRatio(src.AccountRatioParam{
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

func TestPremiumIndexKline(t *testing.T) {
	client := bybit.NewTestClient()
	res, err := client.Future().InversePerpetual().PremiumIndexKline(src.PremiumIndexKlineParam{
		Symbol:   src.SymbolFutureBTCUSD,
		Interval: src.Interval120,
		From:     time.Now().AddDate(0, 0, -1).Unix(),
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v2-public-premium-index-kline.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestCreateOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		price := 28383.5
		res, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSD,
			OrderType:   src.OrderTypeLimit,
			Qty:         1,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-create.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
		// clean
		{
			orderID := res.Result.OrderID
			_, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{
				Symbol:  src.SymbolFutureBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 28383.5
		_, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSD,
			OrderType:   src.OrderTypeLimit,
			Qty:         1,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeLimit,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}
		{
			res, err := client.Future().InversePerpetual().ListOrder(src.ListOrderParam{
				Symbol: src.SymbolFutureBTCUSD,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-order-list.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{
				Symbol:  src.SymbolFutureBTCUSD,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListOrder(src.ListOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestListPosition(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual().ListPosition(src.SymbolFutureBTCUSD)
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-position-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListPosition(src.SymbolFutureBTCUSD)
		require.Error(t, err)
	})
}

func TestListPositions(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		res, err := client.Future().InversePerpetual().ListPositions()
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-position-lists.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListPositions()
		require.Error(t, err)
	})
}

func TestCancelOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var orderID string
		{
			price := 28383.5
			res, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeLimit,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}

		res, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{
			Symbol:  src.SymbolFutureBTCUSD,
			OrderID: &orderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestCancelAllOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			price := 10000.0
			_, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeLimit,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
		}
		res, err := client.Future().InversePerpetual().CancelAllOrder(src.CancelAllOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order-cancel-all.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelAllOrder(src.CancelAllOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := src.SymbolFutureBTCUSD
		var orderID string
		{
			price := 10000.0
			res, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideBuy,
				Symbol:      symbol,
				OrderType:   src.OrderTypeLimit,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
				Price:       &price,
			})
			require.NoError(t, err)
			orderID = res.Result.OrderID
		}

		res, err := client.Future().InversePerpetual().QueryOrder(src.QueryOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		{
			_, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{
				Symbol:  symbol,
				OrderID: &orderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().QueryOrder(src.QueryOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestTradingStop(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := src.SymbolFutureBTCUSD
		{
			_, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideBuy,
				Symbol:      symbol,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}

		{
			price := 20000.0
			res, err := client.Future().InversePerpetual().TradingStop(src.TradingStopParam{
				Symbol:     symbol,
				TakeProfit: &price,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-position-trading-stop.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CreateOrder(src.CreateOrderParam{
				Side:        src.SideSell,
				Symbol:      symbol,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().TradingStop(src.TradingStopParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestSaveLeverage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			res, err := client.Future().InversePerpetual().SaveLeverage(src.SaveLeverageParam{
				Symbol:   src.SymbolFutureBTCUSD,
				Leverage: 2.0,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-position-leverage-save.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
		}
	})
	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelOrder(src.CancelOrderParam{})
		require.Error(t, err)
	})
}

func TestCreateStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			{
				goldenFilename := "./testdata/v2-private-stop-order-create.json"
				testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
				testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			}
			stopOrderID = res.Result.StopOrderID
		}
		// clean
		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(src.CancelStopOrderParam{
				Symbol:      src.SymbolFutureBTCUSD,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		price := 10000.0
		_, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
			Side:        src.SideBuy,
			Symbol:      src.SymbolFutureBTCUSD,
			OrderType:   src.OrderTypeLimit,
			Qty:         1,
			TimeInForce: src.TimeInForceGoodTillCancel,
			Price:       &price,
		})
		require.Error(t, err)
	})
}

func TestListStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		symbol := src.SymbolFutureBTCUSD
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
				Side:        src.SideBuy,
				Symbol:      symbol,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}

		// need to wait until the order status becode untriggered
		time.Sleep(10 * time.Second)

		status := src.OrderStatusUntriggered
		res, err := client.Future().InversePerpetual().ListStopOrder(src.ListStopOrderParam{
			Symbol:          symbol,
			StopOrderStatus: &status,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-list.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		// clean
		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(src.CancelStopOrderParam{
				Symbol:      src.SymbolFutureBTCUSD,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().ListStopOrder(src.ListStopOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}
		res, err := client.Future().InversePerpetual().CancelStopOrder(src.CancelStopOrderParam{
			Symbol:      src.SymbolFutureBTCUSD,
			StopOrderID: &stopOrderID,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-cancel.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelStopOrder(src.CancelStopOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestCancelAllStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		{
			price := 19400.5
			_, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
				Side:        src.SideBuy,
				Symbol:      src.SymbolFutureBTCUSD,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
		}
		res, err := client.Future().InversePerpetual().CancelAllStopOrder(src.CancelAllStopOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order-cancel-all.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().CancelAllStopOrder(src.CancelAllStopOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}

func TestQueryStopOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		client := bybit.NewTestClient().WithAuthFromEnv()
		symbol := src.SymbolFutureBTCUSD
		var stopOrderID string
		{
			price := 19400.5
			res, err := client.Future().InversePerpetual().CreateStopOrder(src.CreateStopOrderParam{
				Side:        src.SideBuy,
				Symbol:      symbol,
				OrderType:   src.OrderTypeMarket,
				Qty:         1,
				BasePrice:   price,
				StopPx:      price + 200,
				TimeInForce: src.TimeInForceGoodTillCancel,
			})
			require.NoError(t, err)
			stopOrderID = res.Result.StopOrderID
		}

		res, err := client.Future().InversePerpetual().QueryStopOrder(src.QueryStopOrderParam{
			Symbol: symbol,
		})
		require.NoError(t, err)
		{
			goldenFilename := "./testdata/v2-private-stop-order.json"
			testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
			testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		}

		{
			_, err := client.Future().InversePerpetual().CancelStopOrder(src.CancelStopOrderParam{
				Symbol:      symbol,
				StopOrderID: &stopOrderID,
			})
			require.NoError(t, err)
		}
	})

	t.Run("auth error", func(t *testing.T) {
		client := bybit.NewTestClient()
		_, err := client.Future().InversePerpetual().QueryStopOrder(src.QueryStopOrderParam{
			Symbol: src.SymbolFutureBTCUSD,
		})
		require.Error(t, err)
	})
}
