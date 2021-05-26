package com.creepersan.sticker_board.account

import com.creepersan.sticker_board.account.database.table.StickerBoardAccountTable
import com.creepersan.sticker_board.common.initializer.StickerBoardInitializer
import com.creepersan.sticker_board.common.manager.VersionManager

class AccountInitializer : StickerBoardInitializer {

    override fun getVersion(): Int {
        return 1
    }

    override fun initialize() {
        val accountVersionManager = AccountVersionManager()
        accountVersionManager.init()

        StickerBoardAccountTable.init()
    }

    override fun upgrade(version: Int): Int {
        return when(version){
            1 -> {
                1
            }
            else -> {
                version
            }
        }

    }


}