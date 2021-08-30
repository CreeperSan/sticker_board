package OSSAlicloud

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	StickerBoard "sticker_board/application/const"
	LogService "sticker_board/lib/log_service"
	SharedPreferences "sticker_board/lib/shared_preferences"
	"strings"
)

var Endpoint string
var AccessKeyID string
var AccessKeySecret string
var BucketName string
var Client *oss.Client
var Bucket *oss.Bucket
func Initialize()  {
	if len(Endpoint) <= 0 {
		Endpoint = SharedPreferences.GetString(StickerBoard.SPOSSAlicloudEndpoint, "")
	}
	if len(AccessKeyID) <= 0 {
		AccessKeyID = SharedPreferences.GetString(StickerBoard.SPOSSAlicloudAccessKeyID, "")
	}
	if len(AccessKeySecret) <= 0 {
		AccessKeySecret = SharedPreferences.GetString(StickerBoard.SPOSSAlicloudAccessKeySecret, "")
	}
	if len(BucketName) <= 0 {
		BucketName = SharedPreferences.GetString(StickerBoard.SPOSSAlicloudBucket, "")
	}

	// Initialize oss Client
	tmpClient, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err!=nil || tmpClient==nil {
		LogService.Error("Initializing OSS Module Failed! Can not connect to alicloud. err =", err)
		os.Exit(StickerBoard.ExitCodeAlicloudOSSCollectionFailed)
		return
	}
	Client = tmpClient

	// Initialize oss Client Bucket
	Bucket, err = Client.Bucket(BucketName)
	if err!=nil || tmpClient==nil {
		LogService.Error("Initializing OSS Module Failed! Can not connect to alicloud oss Bucket. err =", err)
		os.Exit(StickerBoard.ExitCodeAlicloudOSSBucketCollectionFailed)
		return
	}


}

func UploadFile() {
	err := Bucket.PutObject("sticker_board/obj-key", strings.NewReader("56sc4x5czx35cxz5c4c456sa4d46ac32xzcs8scasv"))
	if err != nil {
		LogService.Warming("Upload file invoke fail, err=", err)
		return
	}
}

