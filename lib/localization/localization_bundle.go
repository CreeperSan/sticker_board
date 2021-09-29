package Localization

type LocalizationBunble interface {

	BundleName() string

	TextMap() map[string]string

}
