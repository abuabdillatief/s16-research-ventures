package context

import (
	"context"

	"github.com/abuabdillatief/s16-research-ventures/config"
)

type ContextKey string

const (
	APIKey     ContextKey = "api_key"
)


func IsAuthenticated(ctx context.Context) bool {
	key, ok := ctx.Value(APIKey).(string)
	if !ok {
		return false
	}
	return config.ConfigInUse.Server.APIKey == key
}