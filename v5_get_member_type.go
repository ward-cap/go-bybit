package bybit

import (
	"context"
	"fmt"
	"net/url"
)

type V5APIGetMemberType struct {
	CommonV5Response `json:",inline"`
	Result           V5APIGetMemberTypeResult `json:"result"`
}

type V5APIGetMemberTypeResult struct {
	Accounts []V5MemberAccountInfo `json:"accounts"`
}

type V5MemberAccountInfo struct {
	UID         string   `json:"uid"`         // Унікальний ідентифікатор акаунта
	AccountType []string `json:"accountType"` // Типи акаунта: SPOT, UNIFIED, FUND, CONTRACT тощо
}

// GetMemberType отримує UID акаунтів і їх типи гаманців.
func (s *V5UserService) GetMemberType(ctx context.Context) (*V5APIGetMemberType, error) {
	var res V5APIGetMemberType

	// Виконуємо GET-запит до API
	if err := s.client.getV5PrivatelyCtx(ctx, "/v5/user/get-member-type", url.Values{}, &res); err != nil {
		return nil, fmt.Errorf("не вдалося отримати типи акаунтів: %w", err)
	}

	return &res, nil
}
