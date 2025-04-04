package config

import (
	"context"
	"fmt"
	"testing"
)

func Test_New(t *testing.T) {
	config := New()
	value, err := config.Config(context.Background(), "sit-poc", WithDefault("default"))
	if err != nil {
		if err == ErrNotFound {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("Getting value: ", value)
}
