package i18n

type BundleSimplifyChinese struct {

}

func (bundle *BundleSimplifyChinese) BundleName() string {
	return "zh_CN"
}

func (bundle *BundleSimplifyChinese) TextMap() map[string]string {
	return map[string]string{
		// Account
		"account_expire_need_login" : "账号信息已过期，请重新登录",

		// Sticker
		"sticker_plain_image_url_empty" : "图片不能为空",
		"sticker_plain_sound_url_empty" : "音频不能为空",

		// Version
		"version_out_of_date" : "当前版本太旧，请升级至新版本",

		// Common
		"common_params_error" : "参数错误",
		"common_operate_success" : "操作成功",
	}
}
