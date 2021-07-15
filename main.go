package main

import (
	"fmt"
	StickerAccount "sticker_board/account"
	ApiGeneral "sticker_board/api"
	Sticker "sticker_board/sticker"
)


func main() {
	fmt.Println("========= Sticker Board =========")
	StickerAccount.Init()
	Sticker.Init()
	ApiGeneral.Initialize()
}
