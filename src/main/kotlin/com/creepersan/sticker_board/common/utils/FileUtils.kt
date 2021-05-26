package com.creepersan.sticker_board.common.utils

import java.io.File

object FileUtils {

    // User data root directory
    private val rootDirectory = File("./data")

    fun getRootDirectory(): File{
        return rootDirectory
    }

    fun getModuleDirectory(moduleName: String): File{
        return File("${rootDirectory.absolutePath}/$moduleName")
    }


    fun initDirectory(directory: File): File?{
        val isExist = directory.exists()
        // if directory is exist and it is a directory
        if(isExist && directory.isDirectory){
            return directory
        }
        return if(!isExist){
            // if not exist, create
            if(directory.mkdirs()) directory else null
        } else {
            null
        }
    }

}