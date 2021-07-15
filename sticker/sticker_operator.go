package Sticker

import (
	LogService "sticker_board/lib/log_service"
	StickerDatabase "sticker_board/sticker/database"
)

func Init()  {
	LogService.Info("Initializing Sticker Service Database")

	StickerDatabase.Initialize()

	LogService.Info("Sticker Service Database Initialized.")
}


func CreateCategory(){

}
