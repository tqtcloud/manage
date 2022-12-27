package user

import (
	"net/http"

	"github.com/tqtcloud/manage/common/result"
	"github.com/tqtcloud/manage/service/front-api/internal/logic/user"
	"github.com/tqtcloud/manage/service/front-api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		result.HttpResult(r, w, resp, err)

	}
}
