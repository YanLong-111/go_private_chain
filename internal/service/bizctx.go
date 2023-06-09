package service

import (
	"context"
	"go_private_chain/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	ITesDBCtx interface {
		Init(req *ghttp.Request, ctx *model.Context)
		Get(ctx context.Context) *model.Context
		SetCiphertext(ctx context.Context, ctxGoTestDb *string)
	}
)

var (
	localTesDBCtx ITesDBCtx
)

func TesDBCtx() ITesDBCtx {
	if localTesDBCtx == nil {
		panic("implement not found for interface ITesDBCtx, forgot register?")
	}
	return localTesDBCtx
}

func RegisterTesDBCtx(i ITesDBCtx) {
	localTesDBCtx = i
}
