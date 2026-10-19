package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Exchange broker per-UID API rate limits. Only the broker's master account may
// call these.

type V5APIBrokerAPILimitSet struct {
	CommonV5Response `json:",inline"`
	Result           V5APIBrokerAPILimitSetResult `json:"result"`
}

type V5APIBrokerAPILimitSetResult struct {
	Result []V5BrokerAPILimitSetItem `json:"result"`
}

// V5BrokerAPILimitSetItem is the outcome for one entry of the request: a
// retCode of 0 does not mean every entry was applied.
type V5BrokerAPILimitSetItem struct {
	UIDs    string `json:"uids"`
	BizType string `json:"bizType"`
	Rate    int    `json:"rate"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type V5BrokerAPILimit struct {
	// UIDs is comma-separated: one row covers every uid that shares a rate.
	UIDs    string `json:"uids"`
	BizType string `json:"bizType"`
	Rate    int    `json:"rate"`
}

type SetBrokerAPILimitRequest struct {
	List []V5BrokerAPILimit `json:"list"`
}

// SetBrokerAPILimit is limited by the exchange to 1 request per second.
func (s *V5UserService) SetBrokerAPILimit(ctx context.Context, param SetBrokerAPILimitRequest) (*V5APIBrokerAPILimitSet, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for SetBrokerAPILimit: %w", err)
	}

	var res *V5APIBrokerAPILimitSet

	if err := s.client.postV5JSON(ctx, "/v5/broker/apilimit/set", body, "V5UserService", &res); err != nil {
		return nil, err
	}

	return res, nil
}

type V5APIBrokerAPILimits struct {
	CommonV5Response `json:",inline"`
	Result           V5APIBrokerAPILimitsResult `json:"result"`
}

type V5APIBrokerAPILimitsResult struct {
	List           []V5BrokerAPILimit `json:"list"`
	NextPageCursor string             `json:"nextPageCursor"`
}

type GetBrokerAPILimitsRequest struct {
	// UIDs is comma-separated; empty lists every configured uid.
	UIDs string
	// Limit is capped at 1000 by the exchange; zero leaves it to the default.
	Limit  int
	Cursor string
}

// GetBrokerAPILimits lists configured per-UID limits. A uid that was never
// configured is absent and runs on the exchange default. Limited by the
// exchange to 1 request per second.
func (s *V5UserService) GetBrokerAPILimits(ctx context.Context, param GetBrokerAPILimitsRequest) (res *V5APIBrokerAPILimits, err error) {
	query := url.Values{}
	if param.UIDs != "" {
		query.Set("uids", param.UIDs)
	}
	if param.Limit > 0 {
		query.Set("limit", strconv.Itoa(param.Limit))
	}
	if param.Cursor != "" {
		query.Set("cursor", param.Cursor)
	}

	err = s.client.getV5PrivatelyCtx(ctx, "/v5/broker/apilimit/query-all", query, "V5UserService", &res)

	return
}

type V5APIBrokerAPILimitCap struct {
	CommonV5Response `json:",inline"`
	Result           V5APIBrokerAPILimitCapResult `json:"result"`
}

type V5APIBrokerAPILimitCapResult struct {
	List []V5BrokerAPILimitCap `json:"list"`
}

type V5BrokerAPILimitCap struct {
	BizType string `json:"bizType"`
	// TotalRate is what is already handed out across the master and every sub account.
	TotalRate string `json:"totalRate"`
	EbCap     string `json:"ebCap"`
	UIDCap    string `json:"uidCap"`
}

// GetBrokerAPILimitCap returns the broker-level and per-UID ceilings. The list
// is empty when the account manager never configured custom caps.
func (s *V5UserService) GetBrokerAPILimitCap(ctx context.Context) (res *V5APIBrokerAPILimitCap, err error) {
	err = s.client.getV5PrivatelyCtx(ctx, "/v5/broker/apilimit/query-cap", url.Values{}, "V5UserService", &res)

	return
}
