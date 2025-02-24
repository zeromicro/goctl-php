package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/goctl-php/example/internal/svc"
	"github.com/zeromicro/goctl-php/example/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}

	return
}
