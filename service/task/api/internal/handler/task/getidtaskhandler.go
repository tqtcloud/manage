package task

import (
	"net/http"

	logic "github.com/tqtcloud/manage/service/task/api/internal/logic/task"
	"github.com/tqtcloud/manage/service/task/api/internal/svc"
	"github.com/tqtcloud/manage/service/task/api/internal/types"
	"github.com/tqtcloud/resp/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetIdTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetIdTaskLogic(r.Context(), svcCtx)
		resp, err := l.GetIdTask(&req)
		response.Response(w, resp, err)

	}
}