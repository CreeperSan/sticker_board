package ApiMiddleware

import (
	"github.com/kataras/iris/v12"
	Account "sticker_board/account/database"
	"strconv"
)

type AuthAccountMiddlewareResult struct {
	Token                 string
	AccountID             uint
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

	token := ctx.GetHeader("token")
	brand := ctx.GetHeader("brand")
	deviceName := ctx.GetHeader("device_name")
	machineCode := ctx.GetHeader("machine_code")
	platform, convertPlatformErr := strconv.Atoi(ctx.GetHeader("platform"))
	if convertPlatformErr != nil {
		platform = 0
	}
	uid, convertUIDErr := strconv.Atoi(ctx.GetHeader("uid"))
	if convertUIDErr != nil {
		uid = -1
	}

	authTokenResult := Account.AuthToken(uint(uint64(uid)), token, platform, brand, deviceName, machineCode)

	// if token auth error or token account id not the same as request header's uid
	if authTokenResult.Code != 200 {
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
		ExpireTimeMilliSecond: authTokenResult.ExpireTimeMilliSecond,
		UpdateTime:            authTokenResult.UpdateTime,
	})

	ctx.Next()
}

func AuthAccountMiddleWareGetResponse(ctx iris.Context) AuthAccountMiddlewareResult {
	return ctx.Values().Get("_AccountInfo").(AuthAccountMiddlewareResult)
}








type AuthVersionMiddlewareResult struct {
	Platform    int
	VersionCode int
}

func AuthVersionMiddleware(ctx iris.Context) {
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	platform, convertPlatformErr := strconv.Atoi(ctx.GetHeader("platform"))
	if convertPlatformErr != nil {
		platform = 0
	}
	versionCode, convertVersionCodeErr := strconv.Atoi(ctx.GetHeader("version_code"))
	if convertVersionCodeErr != nil {
		platform = 0
	}

	if versionCode <= 0 || platform <= 0 {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Current application version is out of date, please update to newest version.",
		})
		return
	}

	// auth pass
	ctx.Values().Set("_VersionInfo", AuthVersionMiddlewareResult{
		Platform:    platform,
		VersionCode: versionCode,
	})

	ctx.Next()
}

func AuthVersionMiddleWareGetResponse(ctx iris.Context) AuthVersionMiddlewareResult {
	return ctx.Values().Get("_VersionInfo").(AuthVersionMiddlewareResult)
}
