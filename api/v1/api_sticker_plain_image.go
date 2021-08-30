package ApiV1

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
	ApiMiddleware "sticker_board/api/middleware"
	OSSAlicloud "sticker_board/application/oss_alicloud"
	Formatter "sticker_board/lib/formatter"
	LogService "sticker_board/lib/log_service"
	"strconv"
	"strings"
)

func InitializeStickerPlainImage(app *iris.Application)  {
	stickerPlainText := app.Party("/api/sticker/v1/plain_image")

	stickerPlainText.Use(ApiMiddleware.AuthVersionMiddleware)
	stickerPlainText.Use(ApiMiddleware.AuthAccountMiddleware)

	stickerPlainText.Post("/create", createPlainImageSticker)
}

func createPlainImageSticker(ctx iris.Context){
	type RequestParams struct {
		Star             int      `json:"star"`
		Status           int      `json:"status"`
		Title            string   `json:"title"`
		Background       string   `json:"background"`
		CategoryID       string   `json:"category_id"`
		TagID            []string `json:"tag_id"`
		IsPinned         bool     `json:"is_pinned"`
		ImageDescription string   `json:"description"`
	}
	type ResponseParams struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}

	ctx.SetMaxRequestBodySize(32 * 1024 * 1024) // Limit the size of image should be in 32 MB

	authResult := ApiMiddleware.AuthAccountMiddleWareGetResponse(ctx)

	// parse params
	pStar, err := strconv.Atoi(ctx.FormValue("star"))
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Params Error",
		})
		return
	}
	pStatus, err := strconv.Atoi(ctx.FormValue("status"))
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 401,
			Message: "Params Error",
		})
		return
	}
	requestParams := RequestParams{
		Star: pStar,
		Status: pStatus,
		Title: ctx.FormValue("title"),
		Background: ctx.FormValue("background"),
		CategoryID: ctx.FormValue("category_id"),
		TagID: strings.Split(ctx.FormValue("tag_id"), ","),
		IsPinned: ctx.FormValue("is_pinned") == "true",
		ImageDescription: ctx.FormValue("description"),
	}
	LogService.Info("Request Params : ", requestParams)

	// Extract file from form request
	uploadFile, uploadFileHeader, err := ctx.FormFile("image")
	if err!= nil {
		ctx.JSON(ResponseParams{
			Code: 400,
			Message: "No image uploaded",
		})
		return
	}


	dstFilePathDirectory := fmt.Sprint("./tmp/", authResult.AccountID, "/upload/") // tempotory file path
	err = os.MkdirAll(dstFilePathDirectory, os.ModePerm)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 500,
			Message: fmt.Sprintf("Error occur while initialzing tempotory directory,err=%s",err),
		})
		return
	}
	_, err = ctx.UploadFormFiles(dstFilePathDirectory)
	if err != nil {
		ctx.JSON(ResponseParams{
			Code: 500,
			Message: fmt.Sprintf("Error occur while decoding request,err=%s",err),
		})
		return
	}

	//dstFilePath := fmt.Sprintf(dstFilePathDirectory, "/", uploadFileHeader.Filename)
	dstFilePath := dstFilePathDirectory
	dstFilePath += "/"
	dstFilePath += uploadFileHeader.Filename

	defer func(){
		// 1. close file
		err := uploadFile.Close()
		if err != nil {
			LogService.Warming("Can not close uploaded file in request form")
			return
		}
		// 2. delete file
		err = os.Remove(dstFilePath)
		LogService.Warming("Can not close delete file copy from request form, path=", dstFilePath)
		if err != nil {
			return
		}
	}()

	file, err := os.Open(dstFilePath)
	if err != nil {
		LogService.Info(err)
		return
	}
	LogService.Info(file)

	// Upload file to OSS
	var fileData []byte
	_, err = uploadFile.Read(fileData)
	if err != nil {
		LogService.Warming("Error occur while read file data from form. err=", err)
		ctx.JSON(ResponseParams{
			Code: 500,
			Message: "Error occur while read file data from form",
		})
		return
	}
	err = OSSAlicloud.Bucket.PutObjectFromFile(fmt.Sprint("sticker_board/sticker/", Formatter.CurrentTimestampMillisecond(), ".jpg"), dstFilePath)
	//err = OSSAlicloud.Bucket.PutObject(fmt.Sprint("sticker_board/sticker/", Formatter.CurrentTimestampMillisecond(), ".jpg"), bytes.NewReader(fileData))
	if err != nil {
		LogService.Warming("Error occur while uploading data to oss. err=", err)
		ctx.JSON(ResponseParams{
			Code: 500,
			Message: "Service internal error, please try again later.",
		})
		return
	}
	OSSAlicloud.UploadFile()


	// TODO
	ctx.Writef("FileName=", uploadFileHeader.Filename, "  FileSize=", uploadFileHeader.Size)



	//createResult := StickerModule.GetOperator().CreatePlainImageSticker(
	//	authResult.AccountID,
	//	requestParams.Star,
	//	requestParams.IsPinned,
	//	requestParams.Status,
	//	requestParams.Title,
	//	requestParams.Background,
	//	requestParams.TagID,
	//	requestParams.CategoryID,
	//	"todo", // TODO
	//	requestParams.ImageDescription,
	//)
	//
	//if !createResult.IsSuccess() {
	//	ctx.JSON(ResponseParams{
	//		Code: createResult.Code,
	//		Message: createResult.Message,
	//	})
	//} else {
	//	ctx.JSON(ResponseParams{
	//		Code: 200,
	//		Message: createResult.Message,
	//		Data: createResult.Sticker,
	//	})
	//}

}


