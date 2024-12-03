package src

import (
	"context"
	"encoding/json"
	"fmt"
)

type V5APICreateSubAPIKey struct {
	CommonV5Response `json:",inline"`
	Result           V5APICreateSubAPIKeyResult `json:"result"`
}

type V5APICreateSubAPIKeyResult struct {
	ID          string                          `json:"id"`
	Note        string                          `json:"note"`
	APIKey      string                          `json:"apiKey"`
	ReadOnly    int                             `json:"readOnly"`
	Secret      string                          `json:"secret"`
	Permissions V5APICreateSubAPIKeyPermissions `json:"permissions"`
}

type V5APICreateSubAPIKeyPermissions struct {
	ContractTrade []string `json:"ContractTrade,omitempty"`
	Spot          []string `json:"Spot,omitempty"`
	Wallet        []string `json:"Wallet,omitempty"`
	Options       []string `json:"Options,omitempty"`
	Derivatives   []string `json:"Derivatives,omitempty"`
	Exchange      []string `json:"Exchange,omitempty"`
	CopyTrading   []string `json:"CopyTrading,omitempty"`
}

type CreateSubAPIKeyRequest struct {
	SubUID      int64                           `json:"subuid"`         // Обов'язково: ID субакаунта
	Note        *string                         `json:"note,omitempty"` // Примітка (опціонально)
	ReadOnly    int                             `json:"readOnly"`       // 0: Читання та запис, 1: Тільки читання
	IPs         *string                         `json:"ips,omitempty"`  // Прив'язка до IP (опціонально)
	Permissions V5APICreateSubAPIKeyPermissions `json:"permissions"`    // Дозволи для API-ключа
}

func (s *V5UserService) CreateSubAPIKey(ctx context.Context, param CreateSubAPIKeyRequest) (*V5APICreateSubAPIKey, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for CreateSubAPIKey: %w", err)
	}

	var res *V5APICreateSubAPIKey

	if err := s.client.postV5JSON(ctx, "/v5/user/create-sub-api", body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
