package com.creepersan.sticker_board.account

import com.creepersan.sticker_board.common.manager.VersionManager

class AccountVersionManager : VersionManager() {
    override fun getModuleName(): String {
        return "Account"
    }
}