package initialize

import (
	"github.com/astaxie/beego"
	"base/constants"
)

func InitErrorHandler()  {
	// Redefine 404,401,503 error page
	//Error404()
	
}

func InitAPP (port int)  {
	beego.BConfig.AppName = "app-wang"
	beego.BConfig.ServerName = "app-wang Server"
	beego.BConfig.RecoverPanic = true
	beego.BConfig.MaxMemory = 1 << 30  // 1GB
	beego.BConfig.CopyRequestBody = true // Need this config, this will help us to call beego API to get http body. Otherwise we need to get data from http request by ourselves
	beego.BConfig.Listen.HTTPPort = port
	beego.BConfig.Listen.ServerTimeOut = 45 // 40 is timeout, and 5 is buffer

	InitErrorHandler()
}

func InitAPPSSL (port int)   {
	beego.BConfig.AppName = "app-wang"
	beego.BConfig.ServerName = "app-wang Server"
	beego.BConfig.RecoverPanic = true
	beego.BConfig.MaxMemory = 1 << 30  // 1GB
	beego.BConfig.CopyRequestBody = true
	//beego.BConfig.RunMode = "prod"	// default is "dev"

	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSPort = port
	beego.BConfig.Listen.ServerTimeOut = 45 // 40 is timeout, and 5 is buffer
	beego.BConfig.Listen.HTTPSCertFile = constants.DoctorSSLCrtPath
	beego.BConfig.Listen.HTTPSKeyFile = constants.DoctorSSLKeyPath

	InitErrorHandler()
}
