package bybit

import (
	"context"

	"github.com/google/go-querystring/query"
)

// V5APIGetSubDepositAddress описує структуру відповіді від API для отримання депозитних адрес.
type V5APIGetSubDepositAddress struct {
	CommonV5Response `json:",inline"`
	Result           V5APIGetSubDepositAddressResult `json:"result"`
}

// V5APIGetSubDepositAddressResult містить інформацію про депозитні адреси.
type V5APIGetSubDepositAddressResult struct {
	//Coin   string `json:"coin"`
	Chains struct {
		//ChainType         string `json:"chainType"`
		AddressDeposit string `json:"addressDeposit"`
		TagDeposit     string `json:"tagDeposit"`
		//Chain             string `json:"chain"`
		BatchReleaseLimit string `json:"batchReleaseLimit"`
	} `json:"chains"`
}

// GetSubDepositAddressRequest описує параметри запиту для отримання депозитних адрес.
type GetSubDepositAddressRequest struct {
	Coin        string `url:"coin,omitempty"`        // Монета, наприклад, "USDT"
	ChainType   string `url:"chainType,omitempty"`   // Блокчейн, наприклад, "TRX"
	SubMemberID string `url:"subMemberId,omitempty"` // ID субакаунта
}

// GetSubDepositAddress отримує депозитну адресу для субакаунта.
func (s *V5UserService) GetSubDepositAddress(ctx context.Context, param GetSubDepositAddressRequest) (*V5APIGetSubDepositAddress, error) {
	// Ініціалізуємо змінну для збереження відповіді
	var res *V5APIGetSubDepositAddress

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	// Виконуємо POST-запит до API
	if err := s.client.getV5PrivatelyCtx(ctx, "/v5/asset/deposit/query-sub-member-address", queryString, &res); err != nil {
		return nil, err
	}

	return res, nil
}
