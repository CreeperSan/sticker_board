package com.creepersan.sticker_board.common.manager

import com.alibaba.fastjson.JSON
import com.alibaba.fastjson.JSONObject
import java.io.File
import java.lang.Exception

abstract class VersionManager {
    private lateinit var versionFile: File
    private lateinit var versionObject: VersionManagerObject

    abstract fun getModuleName(): String

    fun init(){
        val tmpVersionFile = File("./data/${getModuleName()}/version/version.json")
        val tmpVersionParentFile = tmpVersionFile.parentFile

        if(tmpVersionParentFile.exists() && tmpVersionParentFile.isFile){
            if(!tmpVersionParentFile.delete()){
                throw VersionManagerException("Version initialize failed! Version directory is not a directory!")
            }
        }

        if(!tmpVersionParentFile.exists()){
            if(!tmpVersionParentFile.mkdirs()){
                throw VersionManagerException("Version initialize failed! Version directory create failed!")
            }
        }

        if(tmpVersionFile.exists() && tmpVersionFile.isDirectory){
            if(!tmpVersionFile.delete()){
                throw VersionManagerException("Version initialize failed! Version file is not a file!")
            }
        }

        if(!tmpVersionFile.exists()){
            if(!tmpVersionFile.createNewFile()){
                throw VersionManagerException("Version initialize failed! Version file create failed!")
            }
        }

        versionFile = tmpVersionFile

        var versionText = tmpVersionFile.readText()
        if(versionText.isEmpty()){
            versionText = "{}"
        }
        versionObject = JSON.parseObject(versionText, VersionManagerObject::class.java)

    }

    fun getVersionCode(): Int{
        return versionObject.versionCode
    }

    fun getVersionName(): String{
        return versionObject.versionName
    }

    fun getUpgradeTime(): Long{
        return versionObject.upgradeTime
    }

    fun upgrade(
            versionCode: Int,
            versionName: String,
    ){
        versionObject.versionCode = versionCode
        versionObject.versionName = versionName
        versionObject.upgradeTime = System.currentTimeMillis()

        val jsonString = JSON.toJSONString(versionObject)
        versionFile.writeText(jsonString)
    }

}

class VersionManagerException(msg: String) : Exception(msg)

class VersionManagerObject {
    var versionName: String = ""
    var versionCode: Int = 1
    var upgradeTime: Long = 0
}