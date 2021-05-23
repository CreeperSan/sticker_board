package com.creepersan.sticker_board.common.utils

object LogUtils {
    private object Level{
        val VERBOSE = "Verbose"
        val INFO    = "Info   "
        val WARMING = "Warming"
        val ERROR   = "Error  "
    }

    private fun log(message: String , level: String){
        println("【${level}】[${FormatUtils.toDateTime(System.currentTimeMillis())}] $message")
    }

    fun v(message: String){
        log(message, Level.VERBOSE)
    }

    fun i(message: String){
        log(message, Level.INFO)
    }

    fun w(message: String){
        log(message, Level.WARMING)
    }

    fun e(message: String){
        log(message, Level.ERROR)
    }

}