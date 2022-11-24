package handler

import (
	"net/http"

	"github.com/tqtcloud/manage/service/user/api/internal/logic"
	"github.com/tqtcloud/manage/service/user/api/internal/svc"
	"github.com/tqtcloud/manage/service/user/api/internal/types"
	"github.com/tqtcloud/resp/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(w, resp, err)

	}
}
