package main

import (
	"fmt"
	StickerAccount "sticker_board/account/v1"
	ApiGeneral "sticker_board/api"
	Sticker "sticker_board/sticker"
)


func main() {
	fmt.Println("========= Sticker Board =========")
	StickerAccount.Init()
	Sticker.Init()
	ApiGeneral.Initialize()
}
