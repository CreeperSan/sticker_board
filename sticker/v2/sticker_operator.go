package StickerV2

type StickerOperator struct {
	
}

func (operator *StickerOperator) Initialize () {
	//operator.FindAllTag("612b8ef853c5e6cc7ee457c7")
	//operator.CreateCategory("612b8ef853c5e6cc7ee457c7", "", "Category01", "", 0)
	//operator.FindAllCategory("612b8ef853c5e6cc7ee457c7")

	//operator.CreatePlainTextSticker("6126e612a9192b1b0c9628be", 0, false,
	//	StickerModuleConst.StickerStatusProcessing, "Test Plain Text 02", "",
	//	[]string{}, "", "xz564c56xz1cas4 5asd15a5s")
	operator.FindSticker("6126e612a9192b1b0c9628be", 1, 1)
}