package src

import (
	"context"
)

type V5APIBrokerAccountInfo struct {
	CommonV5Response `json:",inline"`
	Result           V5APIBrokerAccountInfoResult `json:"result"`
}

type V5APIBrokerAccountInfoResult struct {
	BrokerID        string  `json:"brokerId"`
	Rebate          string  `json:"rebate"`
	TotalSubMembers int     `json:"totalSubMembers"`
	TotalBalance    float64 `json:"totalBalance"`
	CreatedAt       int64   `json:"createdAt"`
}

func (s *V5UserService) GetBrokerAccountInfo(ctx context.Context) (res *V5APIBrokerAccountInfo, err error) {
	err = s.client.getV5PrivatelyCtx(ctx, "/v5/broker/account-info", nil, &res)
	return
}
