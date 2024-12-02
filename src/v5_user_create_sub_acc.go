package src

import (
	"encoding/json"
	"fmt"
)

type V5APICreateSubAcc struct {
	CommonV5Response `json:",inline"`
	Result           V5APICreateSubAccResult `json:"result"`
}

type V5APICreateSubAccResult struct {
	Uid        string `json:"uid"`
	Username   string `json:"username"`
	MemberType int    `json:"memberType"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
}

func (s *V5UserService) CreateSubAcc(param CancelFuturesStopOrderParam) (res *V5APICreateSubAcc, _ error) {

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateSubAcc: %w", err)
	}

	if err := s.client.postJSON("/v5/user/create-sub-member", body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
