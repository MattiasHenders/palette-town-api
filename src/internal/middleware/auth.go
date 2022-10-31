package middleware

import (
	"context"

	"github.com/MattiasHenders/palette-town-api/src/models"
)

var (
	userCtxKey = &contextKey{"User"}
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "constants context value " + k.name
}

func GetAuthUser(ctx context.Context) *models.User {
	return ctx.Value(userCtxKey).(*models.User)
}
