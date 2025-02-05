package bybit

import (
	"context"
	"encoding/json"
)

type V5SetMarginModeParam struct {
	Mode string `json:"setMarginMode"` // ISOLATED_MARGIN, REGULAR_MARGIN(i.e. Cross margin), PORTFOLIO_MARGIN
}

type V5SetMarginModeResponse struct {
	CommonV5Response `json:",inline"`
	Result           struct {
		Reasons []struct {
			ReasonCode string `json:"reasonCode"`
			ReasonMsg  string `json:"reasonMsg"`
		} `json:"reasons"`
	} `json:"result"`
}

func (s *V5AccountService) SetMarginMode(
	ctx context.Context,
	param V5SetMarginModeParam,
) (res []V5SetMarginModeResponse, err error) {

	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	err = s.client.postV5JSON(ctx, "/v5/account/set-margin-mode", body, &res)

	return
}
