package handler

import (
	"net/http"

	"github.com/tqtcloud/manage/service/user/api/internal/logic"
	"github.com/tqtcloud/manage/service/user/api/internal/svc"
	"github.com/tqtcloud/resp/response"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)

	}
}
