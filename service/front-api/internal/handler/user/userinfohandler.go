package user

import (
	"net/http"

	"github.com/tqtcloud/manage/common/response"
	logic "github.com/tqtcloud/manage/service/front-api/internal/logic/user"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)

	}
}
