package com.creepersan.sticker_board.common.initializer

/**
 * Initializer for each function module
 * @author CreeperSan
 */
interface StickerBoardInitializer {

    /**
     * Get current function version
     * @return
     */
    fun getVersion(): Int

    /**
     * Initializer for function
     * Will call this method when application starting
     */
    fun initialize()

    /**
     * Upgrade current version
     * @param version current version
     * @return version after upgrade
     */
    fun upgrade(version: Int): Int

}

/**
 *
 */
class StickerBoardInitializeResult(
        val isSuccess: Boolean,
        val message: String,
)

