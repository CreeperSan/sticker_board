package ApiV1

import (
	"github.com/kataras/iris/v12"
	StickerBoardConst "sticker_board/account/const"
	StickerBoardAccount "sticker_board/account/database"
	ApiMiddleware "sticker_board/api/middleware"
	Formatter "sticker_board/lib/formatter"
	LogService "sticker_board/lib/log_service"
	"strings"
)

func InitializeAccount(app *iris.Application){
	accountApi := app.Party("/api/account/v1")

	accountApi.Post("/login", login)
	accountApi.Post("/register", register)
	accountApi.Post("/auth_token", ApiMiddleware.AuthAccountMiddleware, authToken)
}

func login(ctx iris.Context)  {
	type RequestParams struct {
		Account  string `json:"account"`
		Password string `json:"password"`
		Platform int8 `json:"platform"`
		Brand string `json:"brand"`
		DeviceName string `json:"device_name"`
		MachineCode string `json:"machine_code"`
	}
	type ResponseParamsData struct {
		Token         string `json:"token"`
		EffectiveTime int64  `json:"effective_time"`
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
	if requestParams.Platform <= StickerBoardConst.PlatformUndefined ||
		requestParams.Platform > StickerBoardConst.PlatformBrowser ||
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
	loginResponse := StickerBoardAccount.LoginAccount(requestParams.Account, requestParams.Password,
		requestParams.Platform, requestParams.Brand, requestParams.DeviceName, requestParams.MachineCode)
	LogService.Info(loginResponse)

	// return data to client
	ctx.JSON(ResponseParams{
		Code:    loginResponse.Code,
		Message: loginResponse.Message,
		Data: ResponseParamsData{
			Token:         loginResponse.Token,
			EffectiveTime: loginResponse.EffectiveTime,
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
	registerResponse := StickerBoardAccount.RegisterAccount(requestParams.Account, requestParams.Password,
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
	authResponse := StickerBoardAccount.AuthToken(authResult.AccountID, authResult.Token, authResult.Platform,
		authResult.Brand, authResult.DeviceName, authResult.MachineCode)

	// return data to client
	ctx.JSON(ResponseParams{
		Code:    authResponse.Code,
		Message: authResponse.Message,
	})

}
