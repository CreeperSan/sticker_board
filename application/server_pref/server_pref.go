package ServerPref

import (
	StickerBoard "sticker_board/application/const"
	SharedPreferences "sticker_board/lib/shared_preferences"
)

var domain string
var port string
func GetHostDomain() string {
	if len(domain) <= 0 {
		domain = SharedPreferences.GetString(StickerBoard.SPServeDomain, "")
	}
	return domain
}

func GetHostPort() string {
	if len(port) <= 0 {
		port = SharedPreferences.GetString(StickerBoard.SPServePort, "")
	}
	return port
}
