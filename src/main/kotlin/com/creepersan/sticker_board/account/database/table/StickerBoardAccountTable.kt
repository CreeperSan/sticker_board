package com.creepersan.sticker_board.account.database.table

import com.creepersan.sticker_board.account.const.AccountConst
import com.creepersan.sticker_board.common.utils.FileUtils
import java.io.File
import java.sql.DriverManager

object StickerBoardAccountTable {
    private val path = File("${FileUtils.getModuleDirectory(AccountConst.ModuleName).absolutePath}/database/account.sqlite")

    fun init(){
        // Load SQLite Drivers
        Class.forName("org.sqlite.JDBC")
        FileUtils.initDirectory(path.parentFile)
//        val connection = DriverManager.getConnection("jdbc:sqlite:${path.absolutePath}")
    }

}