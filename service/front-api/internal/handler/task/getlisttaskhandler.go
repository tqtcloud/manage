package task

import (
	"net/http"

	"github.com/tqtcloud/manage/common/result"
	"github.com/tqtcloud/manage/service/front-api/internal/logic/task"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetListTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTaskListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := task.NewGetListTaskLogic(r.Context(), svcCtx)
		resp, err := l.GetListTask(&req)
		result.HttpResult(r, w, resp, err)

	}
}
