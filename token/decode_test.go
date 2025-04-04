package token

import (
	"context"
	"testing"
)

func Test_Decode(t *testing.T) {
	claims, err := Decode(context.Background(), "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJTTDExMjAyMjEzIiwiUGhvbmVfbnVtYmVyIjoiKzgwMDEwMDE0MDAiLCJEZXZpY2VfaWQiOjE0MjAyLCJleHAiOjE2NjkyMTEzODh9.jj0gWi88W2fw6ciTuLJTFMYqpLhd2Wai7urVH1lswzM")
	t.Log(claims, err)
}
