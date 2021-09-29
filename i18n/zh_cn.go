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

		// Version
		"version_out_of_date" : "当前版本太旧，请升级至新版本",
	}
}
