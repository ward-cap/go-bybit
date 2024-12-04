package src

import (
	"context"
	"encoding/json"
	"fmt"
)

// V5APIDeleteSubAPIKey описує структуру відповіді від API при видаленні API-ключа субакаунта.
type V5APIDeleteSubAPIKey struct {
	CommonV5Response `json:",inline"`
}

// DeleteSubAPIKeyRequest описує параметри запиту для видалення API-ключа субакаунта.
type DeleteSubAPIKeyRequest struct {
	APIKey *string `json:"apikey,omitempty"` // API-ключ субакаунта (обов'язковий, якщо викликається з головного акаунта)
}

// DeleteSubAPIKey видаляє API-ключ субакаунта.
func (s *V5UserService) DeleteSubAPIKey(ctx context.Context, param DeleteSubAPIKeyRequest) (*V5APIDeleteSubAPIKey, error) {
	// Серіалізуємо параметри в JSON
	body, err := json.Marshal(param)
	if err != nil {
		return nil, fmt.Errorf("json marshal for DeleteSubAPIKey: %w", err)
	}

	// Ініціалізуємо змінну для збереження відповіді
	var res *V5APIDeleteSubAPIKey

	// Виконуємо POST-запит до API
	if err := s.client.postV5JSON(ctx, "/v5/user/delete-sub-api", body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
