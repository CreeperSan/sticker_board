package ApiMiddleware

import (
	"github.com/kataras/iris/v12"
	AccountModule "sticker_board/account/manager"
	"strconv"
)

type AuthAccountMiddlewareResult struct {
	Token                 string
	AccountID             string
	UpdateTime            int64
	ExpireTimeMilliSecond int64
	Platform              int
	Brand                 string
	DeviceName            string
	MachineCode           string
}

func AuthAccountMiddleware(ctx iris.Context) {
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	uid := ctx.GetHeader("sticker-board-uid")
	token := ctx.GetHeader("sticker-board-token")
	brand := ctx.GetHeader("sticker-board-brand")
	deviceName := ctx.GetHeader("sticker-board-device-name")
	machineCode := ctx.GetHeader("sticker-board-machine-code")
	platform, convertPlatformErr := strconv.Atoi(ctx.GetHeader("sticker-board-platform"))
	if convertPlatformErr != nil {
		platform = 0
	}

	accountOperator := AccountModule.GetOperator()
	authTokenResult := accountOperator.AuthToken(uid, token, platform, brand, deviceName, machineCode)

	// if token auth error or token account id not the same as request header's uid
	if !authTokenResult.IsSuccess() {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Login expired, please login in again",
		})
		return
	}

	// auth pass
	ctx.Values().Set("_AccountInfo", AuthAccountMiddlewareResult{
		Token:                 token,
		AccountID:             authTokenResult.AccountID,
		Platform:              platform,
		Brand:                 brand,
		DeviceName:            deviceName,
		MachineCode:           machineCode,
		ExpireTimeMilliSecond: authTokenResult.ExpireTime,
		UpdateTime:            authTokenResult.UpdateTime,
	})

	ctx.Next()
}

func AuthAccountMiddleWareGetResponse(ctx iris.Context) AuthAccountMiddlewareResult {
	return ctx.Values().Get("_AccountInfo").(AuthAccountMiddlewareResult)
}

