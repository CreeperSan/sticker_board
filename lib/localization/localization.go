package Localization

import "fmt"

var bundles = map[string]LocalizationBunble {}

var defaultBundle = ""

func RegisterBundle(bundle LocalizationBunble){
	bundles[bundle.BundleName()] = bundle
}

func UnregisterBundle(bundle LocalizationBunble){
	delete(bundles, bundle.BundleName())
}

func SetDefaultBundleName(bundleName string){
	defaultBundle = bundleName
}

func GetText(bundleName string, key string) string {
	// 1. find specific bundle
	if value,isContain := bundles[bundleName]; isContain {
		// Bundle exist
		if text,isContainText := value.TextMap()[key]; isContainText {
			return text
		}
	}

	// 2. find default bundle
	if value,isContain := bundles[defaultBundle]; isContain {
		// Bundle exist
		// Bundle exist
		if text,isContainText := value.TextMap()[key]; isContainText {
			return text
		}
	}

	// 3. missing data
	return key
}

func TrText(bundleName string, key string, params ...interface{}) string {
	var text = GetText(bundleName, key)
	return fmt.Sprintf(text, params)
}

