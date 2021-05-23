package com.creepersan.sticker_board.common.utils

import java.text.SimpleDateFormat

object FormatUtils {
    private val datetimeDateFormatter = SimpleDateFormat("yyyy-MM-dd HH:mm:ss.SSS")

    fun toDateTime(timestamp: Long): String{
        return datetimeDateFormatter.format(timestamp)
    }

}