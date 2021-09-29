package ApiMiddleware

import "github.com/kataras/iris/v12"

type LanguageMiddlewareResult struct {
	Platform    int
	VersionCode int
}

func LanguageMiddleware(ctx iris.Context){


	ctx.Next()
}
