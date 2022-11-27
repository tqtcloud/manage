package secret

import (
	"net/http"

	logic "github.com/tqtcloud/manage/service/secret/api/internal/logic/secret"
	"github.com/tqtcloud/manage/service/secret/api/internal/svc"
	"github.com/tqtcloud/manage/service/secret/api/internal/types"
	"github.com/tqtcloud/resp/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetIdSecretHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetIdSecretLogic(r.Context(), svcCtx)
		resp, err := l.GetIdSecret(&req)
		response.Response(w, resp, err)

	}
}
