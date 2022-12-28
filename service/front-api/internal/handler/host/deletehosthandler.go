package host

import (
	"net/http"

	"github.com/tqtcloud/manage/common/result"
	"github.com/tqtcloud/manage/service/front-api/internal/logic/host"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteHostRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := host.NewDeleteHostLogic(r.Context(), svcCtx)
		resp, err := l.DeleteHost(&req)
		result.HttpResult(r, w, resp, err)

	}
}
