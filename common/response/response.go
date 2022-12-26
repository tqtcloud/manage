package response

import (
	"net/http"

	"github.com/tqtcloud/resp/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		if e, ok := err.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			body.Code = e.Code
			body.Msg = e.Msg
		} else {
			body.Code = 9999
			body.Msg = err.Error()
		}
	} else {
		body.Msg = "Success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
