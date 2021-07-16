package ApiMiddleware

import "github.com/kataras/iris/v12"

type AuthAccountMiddlewareResult struct {
	AccountID uint
}

func AuthAccountMiddleware(ctx iris.Context) {
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	//token := ctx.GetHeader("token")

	//ctx.Params().Set("_AccountInfo", "123")
	//ctx.Params().Set("_AccountInfo", &interface{
	//
	//})

	ctx.Values().Set("_AccountInfo", AuthAccountMiddlewareResult{
		AccountID: 16,
	})

	//if len(token) <= 0 {
	//	ctx.JSON(ResponseParams{
	//		Code: 401,
	//		Message: "Please Login, token = "+token,
	//	})
	//	return
	//}

	ctx.Next()
}


