package token

import (
	"context"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	sign := New("test")
	token, err := sign.Generate(
		context.Background(),
		Claims{
			UserID:   "SL20220601",
			Phone:    "9844976308",
			DeviceID: "SJ13131PGW",
		},
		WithDuration(2*time.Second),
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "Expected to verify the token",
			token:   token,
			wantErr: false,
		},
		{
			name:    "Expected provided token to return invalid",
			token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQaG9uZV9udW1iZXIiOiIrODAwMTAwMTE1MyIsIkRldmljZV9pZCI6OTEzMiwiZXhwIjoxNjY0NjE2OTIyfQ.uP4n6B_OjE6iq39NcU786Ai1SxQ7Z3wPuOkh350pRx8",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sign.Verify(context.Background(), tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("sign.Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		claims Claims
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Expected to generate the token",
			args: args{
				claims: Claims{
					UserID:   "SL20220601",
					Phone:    "8001001050",
					DeviceID: "SJ13131PGW",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := New("test").Generate(context.Background(), tt.args.claims, WithDuration(24*time.Second))
			if (err != nil) != tt.wantErr {
				t.Errorf("sign.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(token)
		})
	}
}
