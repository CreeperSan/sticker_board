package com.creepersan.sticker_board.account.controller.v1

import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.*

@Controller
@RequestMapping(value = ["/account/v1"])
class AccountControllerV1{

    @ResponseBody
    @RequestMapping(value = ["/login"], method = [RequestMethod.POST, RequestMethod.GET])
    fun login(@RequestBody requestBodyString: String ) : String{
        return requestBodyString
    }

    @ResponseBody
    @RequestMapping(value = ["/register"], method = [RequestMethod.PUT, RequestMethod.GET])
    fun register() : String{
        return "Not Supported yet"
    }

    @ResponseBody
    @RequestMapping(value = ["/auth"], method = [RequestMethod.POST, RequestMethod.GET])
    fun auth() : String{
        return "Learning & Developing"
    }


    @ResponseBody
    @RequestMapping(value = ["/logout"], method = [RequestMethod.POST, RequestMethod.GET])
    fun logout() : String{
        return "Learning & Developing"
    }

}