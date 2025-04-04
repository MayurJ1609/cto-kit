package token

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Decode(ctx context.Context, token string) (map[string]interface{}, error) {
	token = strings.ReplaceAll(token, "Bearer ", "")

	claims := map[string]interface{}{}
	t, _, err := jwt.NewParser().ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return claims, err
	}

	data, _ := json.Marshal(t.Claims)
	json.Unmarshal(data, &claims)

	return claims, nil
}
