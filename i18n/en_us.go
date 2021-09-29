package i18n

type BundleEnglish struct {

}

func (bundle *BundleEnglish) BundleName() string {
	return "en_US"
}

func (bundle *BundleEnglish) TextMap() map[string]string {
	return map[string]string{
		// Account
		"account_expire_need_login" : "Account information expired, please login",

		// Sticker
		"sticker_plain_image_url_empty" : "Image path can not be empty",
		"sticker_plain_sound_url_empty" : "Sound path can not be empty",

		// Version
		"version_out_of_date" : "Current version is out of date, please update to latest version",

		// Common
		"common_params_error" : "Params error",
		"common_operate_success" : "operate_success",
	}
}