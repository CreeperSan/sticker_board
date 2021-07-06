// @Title log_service.go
// @Description 
// @Author creepersan, 2021-07-06 17:13:43
// @Update creepersan, 2021-07-06 17:13:43

package LogService

import "log"

func Debug(message ...interface{}) {
	log.Println("【 Debug 】", message)
}

func Info(message ...interface{})  {
	log.Println("【 Info  】", message)
}

func Warming(message ...interface{}) {
	log.Println("【Warming】", message)
}

func Error(message ...interface{}) {
	log.Println("【 Error 】", message)
}



