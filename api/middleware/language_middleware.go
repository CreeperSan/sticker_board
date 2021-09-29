package ApiMiddleware

import (
	"github.com/kataras/iris/v12"
	Localization "sticker_board/lib/localization"
)

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

func LanguageMiddlewareTrText(ctx iris.Context, key string, params ...interface{}) string {
	var lang = LanguageMiddleWareGetResponse(ctx).Language
	return Localization.TrText(lang, key, params...)
}
