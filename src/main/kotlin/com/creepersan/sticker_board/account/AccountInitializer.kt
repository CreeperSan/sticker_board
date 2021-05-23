package com.creepersan.sticker_board.account

import com.creepersan.sticker_board.common.initializer.StickerBoardInitializer
import com.creepersan.sticker_board.common.manager.VersionManager

class AccountInitializer : StickerBoardInitializer {

    override fun getVersion(): Int {
        return 1
    }

    override fun initialize() {
        val accountVersionManager = AccountVersionManager()
        accountVersionManager.init()
        accountVersionManager.upgrade(1, "0.0.1")
    }

    override fun upgrade(version: Int): Int {
        return 1
    }


}