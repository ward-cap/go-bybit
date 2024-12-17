package src

import (
	"context"
)

type V5APIBrokerAccountInfo struct {
	CommonV5Response `json:",inline"`
	Result           V5APIBrokerAccountInfoResult `json:"result"`
}

type V5APIBrokerAccountInfoResult struct {
	SubAcctQty        string `json:"subAcctQty"`
	MaxSubAcctQty     string `json:"maxSubAcctQty"`
	BaseFeeRebateRate struct {
		Spot        string `json:"spot"`
		Derivatives string `json:"derivatives"`
	} `json:"baseFeeRebateRate"`
	MarkupFeeRebateRate struct {
		Spot        string `json:"spot"`
		Derivatives string `json:"derivatives"`
		Convert     string `json:"convert"`
	} `json:"markupFeeRebateRate"`
}

func (s *V5UserService) GetBrokerAccountInfo(ctx context.Context) (res *V5APIBrokerAccountInfo, err error) {
	err = s.client.getV5PrivatelyCtx(ctx, "/v5/broker/account-info", nil, &res)
	return
}
