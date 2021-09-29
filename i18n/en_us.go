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

		// Version
		"version_out_of_date" : "Current version is out of date, please update to latest version",
	}
}