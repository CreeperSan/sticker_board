package v1

import (
	LogService "sticker_board/lib/log_service"
	"sticker_board/sticker/v1/database"
)

func Init()  {
	LogService.Info("Initializing Sticker Service Database")

	StickerDatabase.Initialize()

	LogService.Info("Sticker Service Database Initialized.")
}


func CreateCategory(){

}
