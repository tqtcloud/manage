// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	secret "github.com/tqtcloud/manage/service/secret/api/internal/handler/secret"
	"github.com/tqtcloud/manage/service/secret/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: secret.CreateSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete/:id",
				Handler: secret.DeleteSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: secret.UpdateSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getlist",
				Handler: secret.GetListSecretHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: secret.GetIdSecretHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/secret/v1"),
	)
}
