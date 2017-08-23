package main

import (
	"flag"
	"fmt"
	"strings"
	"runtime"
	"strconv"
	"encoding/json"
	"github.com/astaxie/beego"

	"base/initialize"
	"base/constants"
	"base/config"
	"base/global"
)

var Config *string
var Services *string
var Port *string

func main()  {

	runtime.GOMAXPROCS(runtime.NumCPU()); // Use all processors core to handle request. Default is just use one core.

	/**** Check parameter ****/
	Config = flag.String("c", "", "Configuration File")
	Services = flag.String("s", "", "Services Name")
	Port = flag.String("p", "", "Service Port")	// If needed. For most case, we can get port from config file
	flag.Parse()

	if (*Config == "" || *Services==""){
		panic("Parameter error!")
	}

	/**** Load Yaml File ****/
	config.LoadYamlFile(*Config)

	/**** Enable log ****/
	beego.SetLogger("console", fmt.Sprintf("{\"level\":%d}", config.GetLogLevel()))  // Output to console

	loggerConfig := make(map[string]interface{})
	loggerConfig["filename"] = constants.LogFilePath
	loggerConfig["level"] = config.GetLogLevel()
	loggerConfig["maxlines"] = 10000
	loggerConfigJson, _ := json.Marshal(loggerConfig)

	beego.SetLogger("file", string(loggerConfigJson))	// Output to file

	ServiceArray := strings.Split(*Services, ",")

	var servicePort string
	if (*Port == ""){
		port := config.GetPort()
		servicePort = strconv.Itoa(port)
	} else {
		servicePort = *Port
	}
	*global.ServicePort = servicePort

	_port, _:=strconv.Atoi(servicePort)
	//initialize.InitAPPSSL(_port)
	initialize.InitAPP(_port)
	beego.Informational("Using config file: ", constants.ConfigDirectory + "/" + *Config)

	/**** Init Services ****/
	initialize.InitService(ServiceArray)

	beego.Informational("Code name: Kunlun	Version:", constants.DoctorVersion)
	beego.Informational(beego.BConfig.AppName)
	beego.Run()

}