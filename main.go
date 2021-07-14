package main

import (
	"fmt"
	StickerBoardAccount "sticker_board/account"
	ApiGeneral "sticker_board/api"
)


func main() {
	fmt.Println("========= Sticker Board =========")
	StickerBoardAccount.Init()
	ApiGeneral.Initialize()
}
