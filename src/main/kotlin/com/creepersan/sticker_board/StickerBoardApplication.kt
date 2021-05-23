package com.creepersan.sticker_board

import com.creepersan.sticker_board.account.AccountInitializer
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class StickerBoardApplication

fun main(args: Array<String>) {
	AccountInitializer().initialize()

	runApplication<StickerBoardApplication>(*args)
}
