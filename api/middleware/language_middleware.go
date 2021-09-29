package ApiMiddleware

import "github.com/kataras/iris/v12"

type LanguageMiddlewareResult struct {
	Language string
}

func LanguageMiddleware(ctx iris.Context){

	lang := ctx.GetHeader("sticker-board-lang")

	ctx.Values().Set("_Lang", LanguageMiddlewareResult{
		Language: lang,
	})

	ctx.Next()
}

func LanguageMiddleWareGetResponse(ctx iris.Context) LanguageMiddlewareResult{
	return ctx.Values().Get("_Lang").(LanguageMiddlewareResult)
}
