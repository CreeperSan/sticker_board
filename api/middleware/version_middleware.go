package ApiMiddleware

import (
	"github.com/kataras/iris/v12"
	"strconv"
)

type AuthVersionMiddlewareResult struct {
	Platform    int
	VersionCode int
}

func AuthVersionMiddleware(ctx iris.Context) {
	type ResponseParams struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	platform, convertPlatformErr := strconv.Atoi(ctx.GetHeader("sticker-board-platform"))
	if convertPlatformErr != nil {
		platform = 0
	}
	versionCode, convertVersionCodeErr := strconv.Atoi(ctx.GetHeader("sticker-board-version-code"))
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
