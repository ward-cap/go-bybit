package src

import (
	"context"
	"encoding/json"
	"fmt"
)

type V5APIModifySubAPIKey struct {
	CommonV5Response `json:",inline"`
	Result           V5APIModifySubAPIKeyResult `json:"result"`
}

type V5APIModifySubAPIKeyResult struct {
	ID          string                          `json:"id"`
	Note        string                          `json:"note"`
	APIKey      string                          `json:"apiKey"`
	ReadOnly    int                             `json:"readOnly"`
	Secret      string                          `json:"secret"`
	Permissions V5APIModifySubAPIKeyPermissions `json:"permissions"`
	IPs         []string                        `json:"ips"`
}

type V5APIModifySubAPIKeyPermissions struct {
	ContractTrade []string `json:"ContractTrade,omitempty"`
	Spot          []string `json:"Spot,omitempty"`
	Wallet        []string `json:"Wallet,omitempty"`
	Options       []string `json:"Options,omitempty"`
	Exchange      []string `json:"Exchange,omitempty"`
	CopyTrading   []string `json:"CopyTrading,omitempty"`
}

type ModifySubAPIKeyRequest struct {
	APIKey      *string                          `json:"apikey,omitempty"`      // API-ключ субакаунта (обов'язковий, якщо викликається з головного акаунта)
	ReadOnly    *int                             `json:"readOnly,omitempty"`    // 0: Читання та запис (за замовчуванням), 1: Тільки читання
	IPs         *string                          `json:"ips,omitempty"`         // Прив'язка до IP, наприклад: "192.168.0.1,192.168.0.2"
	Permissions *V5APIModifySubAPIKeyPermissions `json:"permissions,omitempty"` // Дозволи для API-ключа
}

// ModifySubAPIKey модифікує налаштування API-ключа субакаунта.
func (s *V5UserService) ModifySubAPIKey(ctx context.Context, param ModifySubAPIKeyRequest) (*V5APIModifySubAPIKey, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for ModifySubAPIKey: %w", err)
	}

	var res *V5APIModifySubAPIKey

	if err := s.client.postV5JSON(ctx, "/v5/user/update-sub-api", body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
