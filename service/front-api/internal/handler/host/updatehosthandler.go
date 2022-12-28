package host

import (
	"net/http"

	"github.com/tqtcloud/manage/common/result"
	"github.com/tqtcloud/manage/service/front-api/internal/logic/host"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
	"github.com/tqtcloud/manage/service/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHostRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := host.NewUpdateHostLogic(r.Context(), svcCtx)
		resp, err := l.UpdateHost(&req)
		result.HttpResult(r, w, resp, err)

	}
}
