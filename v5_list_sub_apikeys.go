package bybit

import (
	"context"
	"net/url"
	"strconv"
)

type V5APISubUIDAPIKeys struct {
	CommonV5Response `json:",inline"`
	Result           V5APISubUIDAPIKeysResult `json:"result"`
}

type V5APISubUIDAPIKeysResult struct {
	Result         []V5SubUIDAPIKey `json:"result"`
	NextPageCursor string           `json:"nextPageCursor"`
}

// V5SubUIDAPIKey is one of a sub account's api keys. Only the fields needed to
// identify a key and read its IP whitelist are mapped; the secret is masked by
// the exchange and the permission block is not used here.
type V5SubUIDAPIKey struct {
	ID     string `json:"id"`
	APIKey string `json:"apiKey"`
	Note   string `json:"note"`
	// Ips is the key's IP whitelist. A single "*" means the key is unrestricted.
	Ips []string `json:"ips"`
	// Status: 1 permanent, 2 expired, 3 active, 4 expires within 7 days.
	Status   int  `json:"status"`
	Type     int  `json:"type"`
	ReadOnly bool `json:"readOnly"`
}

type GetSubUIDAPIKeysRequest struct {
	SubMemberID string
	// Limit is capped at 20 by the exchange; zero leaves it to the default.
	Limit  int
	Cursor string
}

// GetSubUIDAPIKeys lists a sub account's api keys. Only a master account may call
// it, which is what makes it usable for checking a sub key's whitelist without
// holding that key's secret.
func (s *V5UserService) GetSubUIDAPIKeys(
	ctx context.Context,
	param GetSubUIDAPIKeysRequest,
) (res *V5APISubUIDAPIKeys, err error) {
	query := url.Values{}
	query.Set("subMemberId", param.SubMemberID)
	if param.Limit > 0 {
		query.Set("limit", strconv.Itoa(param.Limit))
	}
	if param.Cursor != "" {
		query.Set("cursor", param.Cursor)
	}

	err = s.client.getV5PrivatelyCtx(ctx, "/v5/user/sub-apikeys", query, "V5UserService", &res)

	return
}
