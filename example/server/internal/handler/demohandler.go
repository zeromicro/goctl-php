package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/goctl-php/example/internal/logic"
	"github.com/zeromicro/goctl-php/example/internal/svc"
	"github.com/zeromicro/goctl-php/example/internal/types"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
