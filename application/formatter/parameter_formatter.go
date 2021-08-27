package ApplicationFormatter

import (
	ApplicationConst "sticker_board/application/const"
)

func IsPlatformValid(platform int) bool {
	switch platform {
		case
			ApplicationConst.PlatformAndroid,
			ApplicationConst.PlatformIOS,
			ApplicationConst.PlatformWindows,
			ApplicationConst.PlatformLinux,
			ApplicationConst.PlatformMac,
			ApplicationConst.PlatformBrowser: {
			return true
		}
	}
	return false
}
