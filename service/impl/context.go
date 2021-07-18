package impl

import (
	"context"
	"internetbar_echo/model"
)

const contextKey = "context key"

type Context struct {
	Token *model.Token
	Admin *model.Admin
}

func WithContext(parent context.Context, ctx *Context) context.Context {
	i, ok := parent.(interface {
		Set(key string, value interface{})
	})
	if ok {
		i.Set(contextKey, ctx)
		return parent
	}
	return context.WithValue(parent, contextKey, ctx)
}
func GetContext(ctx context.Context) *Context {
	return ctx.Value(contextKey).(*Context)
}
