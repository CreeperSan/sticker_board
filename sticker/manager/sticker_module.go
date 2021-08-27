package StickerModule

var _instance StickerInterface

func InstallOperator(instance StickerInterface){
	_instance = instance
}

func UninstallOperator(){
	_instance = nil
}

func GetOperator() StickerInterface {
	if _instance != nil {
		return _instance
	}

	panic("Sticker Module has not been install yet!")
}

