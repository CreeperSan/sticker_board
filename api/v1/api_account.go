package ApiV1

import (
	"github.com/kataras/iris/v12"
	AccountModule "sticker_board/account/manager"
	ApiMiddleware "sticker_board/api/middleware"
	"sticker_board/application/formatter"
	Formatter "sticker_board/lib/formatter"
	"strings"
)

func InitializeAccount(app *iris.Application){
	accountApi := app.Party("/api/account/v1")

	accountApi.Use(ApiMiddleware.LanguageMiddleware)
	accountApi.Use(ApiMiddleware.AuthVersionMiddleware)

	accountApi.Post("/login", login)
	accountApi.Post("/register", register)
	accountApi.Post("/auth_token", ApiMiddleware.AuthAccountMiddleware, authToken)
}

func login(ctx iris.Context)  {
	type RequestParams struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Platform int `json:"platform"`
		Brand string `json:"brand"`
		DeviceName string `json:"device_name"`
		MachineCode string `json:"machine_code"`
	}
	type ResponseParamsData struct {
		Token       string `json:"token"`
		ExpiredTime int64  `json:"expired_time"`
		UID         string `json:"uid"`
	}
	type ResponseParams struct {
		Code    int                `json:"code"`
		Message string             `json:"msg"`
		Data    ResponseParamsData `json:"data"`
	}
	// parse params
	requestParams := RequestParams{}
	err := ctx.ReadJSON(&requestParams)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: "Params Error",
		})
		return
	}

	// Check whether the params is valid

	if  !ApplicationFormatter.IsPlatformValid(requestParams.Platform) ||
		!Formatter.CheckStringWithLength(strings.TrimSpace(requestParams.Brand), 1, 64) ||
		!Formatter.CheckStringWithLength(strings.TrimSpace(requestParams.DeviceName), 1, 64) ||
		len(strings.TrimSpace(requestParams.MachineCode))!= 18 {

		ctx.JSON(ResponseParams{
			Code: 406,
			Message: "Request params error",
		})
		return
	}


	// Login
	loginResponse := AccountModule.GetOperator().LoginAccount(requestParams.Account, requestParams.Password,
		requestParams.Platform, requestParams.Brand, requestParams.DeviceName, requestParams.MachineCode)

	// return data to client
	ctx.JSON(ResponseParams{
		Code:    loginResponse.Code,
		Message: loginResponse.Message,
		Data: ResponseParamsData{
			Token:       loginResponse.Token,
			ExpiredTime: loginResponse.ExpireTime,
			UID:         loginResponse.AccountID,
		},
	})
}

func register(ctx iris.Context)  {
	type RequestParams struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		UserName string `json:"username"`
		Email    string `json:"email"`
	}
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	// parse params
	requestParams := RequestParams{}
	err := ctx.ReadJSON(&requestParams)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 406,
			Message: "Params Error",
		})
		return
	}

	// Register
	registerResponse := AccountModule.GetOperator().RegisterAccount(requestParams.Account, requestParams.Password,
		requestParams.UserName, requestParams.Email)


	// return data to client
	ctx.JSON(ResponseParams{
		Code:    registerResponse.Code,
		Message: registerResponse.Message,
	})
}

func authToken(ctx iris.Context)  {
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	}

	// auth header
	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	// Auth Code
	authResponse := AccountModule.GetOperator().AuthToken(authResult.AccountID, authResult.Token, authResult.Platform,
		authResult.Brand, authResult.DeviceName, authResult.MachineCode)

	// return data to client
	ctx.JSON(ResponseParams{
		Code:    authResponse.Code,
		Message: authResponse.Message,
	})

}
