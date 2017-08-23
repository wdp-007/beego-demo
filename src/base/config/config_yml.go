package config

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/yaml"
	"path/filepath"
	"base/constants"
	"github.com/astaxie/beego/logs"
)

var ConfigInstance config.Configer
const (
	package_name         = "base_config"
)
func LoadYamlFile(filename string)  {

	configFile := filepath.Join(constants.ConfigDirectory, filename + ".yml")
	configAdapter, err := config.NewConfig("yaml", configFile)
	if err != nil {
		panic("Config file error " + err.Error())
	}

	ConfigInstance = configAdapter
}

func GetServerVersion () string  {

	return constants.DoctorVersion
//	serverInfoFile := filepath.Join(constants.ConfigDirectory, "server_info" + ".yml")
//	serverInfoAdapter, err := config.NewConfig("yaml", serverInfoFile)
//	if err != nil {
//		panic("Config file error " + err.Error())
//	}
//
//	_version, err := serverInfoAdapter.DIY("version")
//	if err != nil {
//		panic("Config file error, can not find verson in server_info.yml:" + err.Error())
//	}
//
//	var version map[string]interface{}
//	version = _version.(map[string]interface{})
//	majorVersion := strconv.FormatInt(version["major_version"].(int64) , 10)
//
//	var subVersion string
//
//	switch version["sub_version"].(type){
//		case string:
//			subVersion = version["sub_version"].(string)
//		case int64:
//			subVersion = strconv.FormatInt(version["sub_version"].(int64) , 10)
//	}
//	return majorVersion+"."+subVersion
}

func GetLogLevel() int {
	// info, debug, warning, error
	log_level := ConfigInstance.String("log_level")
	if (log_level == ""){
		return logs.LevelInformational
	} else if(log_level=="info"){
		return logs.LevelInformational
	} else if(log_level=="debug"){
		return logs.LevelDebug
	} else if(log_level=="warning"){
		return logs.LevelWarning
	} else if(log_level=="error"){
		return logs.LevelError
	} else{
		return logs.LevelInformational
	}
}

func GetPort() int {

	port, err := ConfigInstance.Int("port")
	if (err != nil){
		panic("Get server port error")
	}
	return port
}