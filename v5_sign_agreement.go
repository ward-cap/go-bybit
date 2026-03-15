package bybit

import (
	"context"
	"encoding/json"
)

type V5APISignAgreementResponse struct {
	CommonV5Response `json:",inline"`
}

type V5APISignAgreementRequest struct {
	Category int  `json:"category"`
	Agree    bool `json:"agree"`
}

func (s *V5UserService) SignAgreement(ctx context.Context, category int) (res *V5APISignAgreementResponse, err error) {
	var rq = V5APISignAgreementRequest{Category: category, Agree: true}

	rqJson, err := json.Marshal(rq)
	if err != nil {
		return nil, err
	}

	err = s.client.postV5JSON(ctx, "/v5/user/agreement", rqJson, &res)

	return res, err
}
