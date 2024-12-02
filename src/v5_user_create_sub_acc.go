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

type CreateSubUserRequest struct {
	Username   string  `json:"username"`
	Password   *string `json:"password,omitempty"`
	MemberType int     `json:"memberType"` // 1: normal sub account, 6: custodial sub account
	Switch     int     `json:"switch"`     // 0: turn off quick login (default),
	//IsUTA      bool    `json:"isUta,omitempty"`
	//Note       string  `json:"note,omitempty"`
}

func (s *V5UserService) CreateSubAcc(param CreateSubUserRequest) (res *V5APICreateSubAcc, _ error) {

	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateSubAcc: %w", err)
	}

	if err := s.client.postJSON("/v5/user/create-sub-member", body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
