package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success case",
			args: args{
				configPath: "./test_config.yaml",
			},
			wantErr: false,
		},
		{
			name: "Fail case",
			args: args{
				configPath: "./non_existent.yaml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.configPath)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
				assert.NoError(t, err)
			}
		})
	}
}
