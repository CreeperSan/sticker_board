package StickerModule

type StickerInterface interface {

	Initialize()


	// Sticker Tag

	CreateTag()

	DeleteTag()

	FindAllTag()


	// Sticker Category

	CreateCategory()

	DeleteCategory()

	FindAllCategory()


	// Sticker

	CreatePlainTextSticker()

	DeleteSticker()

	FindSticker()

}


