// @Title shared_preferences.go
// @Description provide an easy way to to save and read some simple value into local files
// @Author CreeperSan, 2021-07-06 15:48:12
// @Update CreeperSan, 2021-07-06 15:48:12

package SharedPreferences

import (
	"encoding/json"
	"fmt"
	"os"
)

// config file path
const _ConfigFilePath = "config.json"

// config data
var _data map[string]interface{}


// @title 		: isConfigFileExist
// @description : like the name said
// @auth 		: CreeperSan
// @param		: <empty>
// @return 		: bool, if the config file is exist
func isConfigFileExist() bool {
	_, err := os.Stat(_ConfigFilePath)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}


// @title 		: getMap
// @description : get shared preferences data in memory
// @auth 		: CreeperSan
// @param		: empty
// @return 		: map[string] interface{}
func getMap() map[string] interface{} {
	if _data == nil {
		// check the config file is exist
		if isConfigFileExist() {
			// if exist, read configs from file
			fileContent, errReadFile := os.ReadFile(_ConfigFilePath)
			if errReadFile != nil {
				fmt.Println(" Error occurred while reading application config")
				panic(errReadFile)
			}
			// parse to map from the data read from config file
			errParseJson := json.Unmarshal(fileContent, &_data)
			if errParseJson != nil {
				fmt.Println(" Error occurred while parsing application config file's data")
				panic(errParseJson)
			}
		} else {
			// or just create an empty config file
			_data = make(map[string]interface{})
		}
	}
	return _data
}



// get
// @description : get value by key
// @createDate  : 2021-07-06 16:15:12
// @auth 		: creepersan
// @param		: [key, defaultValue]
// @return 		:
func get(key string, defaultValue interface{}) interface{} {
	var resultValue = getMap()[key]
	if resultValue != nil {
		return resultValue
	}
	return defaultValue
}

// GetString
// @description : Like the name said
// @createDate  : 2021-07-06 16:18:13
// @auth 		: creepersan
// @param		:
// @return 		:
func GetString(key string, defaultValue string) string {
	return get(key, defaultValue).(string)
}

// GetInt
// @description : Like the name said
// @createDate  : 2021-07-06 16:16:26
// @auth 		: creepersan
// @param		:
// @return 		:
func GetInt(key string, defaultValue int) int {
	return int(get(key, float64(defaultValue)).(float64))
}

// GetBool
// @description : Like the name said
// @createDate  : 2021-07-06 16:16:34
// @auth 		: creepersan
// @param		:
// @return 		:
func GetBool(key string, defaultValue bool) bool {
	return get(key, defaultValue).(bool)
}

// GetFloat
// @description : Like the name said
// @createDate  : 2021-07-06 16:16:42
// @auth 		: creepersan
// @param		:
// @return 		:
func GetFloat(key string, defaultValue float64) float64 {
	return get(key, defaultValue).(float64)
}



// set
// @description : Like the name said
// @createDate  : 2021-07-06 16:18:47
// @auth 		: creepersan
// @param		:
// @return 		:
func set(key string, value interface{}) {
	getMap()[key] = value

	result, err := json.Marshal(getMap())

	if err != nil {
		fmt.Println("error ", err)
		return
	}
	// fmt.Println("Updating Map -> ", string(result))

	err1 := os.WriteFile(_ConfigFilePath, result, os.ModePerm)
	if err1 != nil {
		fmt.Println("Error occurred while updating config -> ", err1)
		return
	}
}

// SetString
// @description : Like the name said
// @createDate  : 2021-07-06 16:18:57
// @auth 		: creepersan
// @param		:
// @return 		:
func SetString(key string, value string) {
	set(key, value)
}

// SetInt
// @description : Like the name said
// @createDate  : 2021-07-06 16:19:02
// @auth 		: creepersan
// @param		:
// @return 		:
func SetInt(key string, value int){
	set(key, value)
}

// SetBool
// @description : Like the name said
// @createDate  : 2021-07-06 16:19:39
// @auth 		: creepersan
// @param		:
// @return 		:
func SetBool(key string, value bool){
	set(key, value)
}

// SetFloat
// @description : Like the name said
// @createDate  : 2021-07-06 16:19:48
// @auth 		: creepersan
// @param		:
// @return 		:
func SetFloat(key string, value float64){
	set(key, value)
}