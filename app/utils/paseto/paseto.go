package paseto

import (
	"os"
	"time"
	"wordora/app/modules/users/model"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type TokenHelper struct {
	pasetoKey []byte
}

func NewTokenHelper() *TokenHelper {
	secretKey := os.Getenv("PASETO_SECRET_KEY")
	if len(secretKey) != chacha20poly1305.KeySize {
		panic("Invalid PASETO key size")
	}

	return &TokenHelper{
		pasetoKey: []byte(secretKey),
	}
}

func (t *TokenHelper) GenerateToken(user *model.User) (string, error) {
	now := time.Now()
	expiration := now.Add(24 * time.Hour)
	token := paseto.NewV2()

	payload := map[string]interface{}{
		"sub": user.ID,
		"role": user.Role,
		"exp": expiration.Unix(),
	}

	return token.Encrypt(t.pasetoKey, payload, nil)
}

func (t *TokenHelper) ValidateToken(tokenStr string) (map[string]interface{}, error) {
	token := paseto.NewV2()
	var payload map[string]interface{}

	err := token.Decrypt(tokenStr, t.pasetoKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
