package context

import (
	"context"
	"testing"

	"github.com/abuabdillatief/s16-research-ventures/config"
	"github.com/stretchr/testify/assert"
)

func TestIsAuthenticated(t *testing.T) {
	// initiate config value
	_, err := config.NewConfig("../config/test_config.yaml")
	assert.NoError(t, err)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success case",
			args: args{ctx: context.WithValue(context.TODO(), APIKey, "testsecretkey")},
			want: true,
		},
		{
			name: "Fail case",
			args: args{ctx: context.WithValue(context.TODO(), APIKey, "failcase")},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAuthenticated(tt.args.ctx); got != tt.want {
				t.Errorf("IsAuthenticated() = %v, want %v", got, tt.want)
			}
		})
	}
}
