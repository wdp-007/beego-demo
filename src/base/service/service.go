package service

import (
	"base/config"
	"fmt"
	"errors"
	"database/sql"
)

type Service  struct{
	ServiceName string
	ServiceDB *sql.DB
	UserDB *sql.DB	// Gene: For some service such as fabric/auth, it need read the user database
	ServiceConfig map[string] interface{}
}
var serviceMap map[string] *Service

func init()  {
	serviceMap = make(map[string] *Service)
}

func (ser *Service) InitService(serviceName string){
	// Gene: This method is used to get service config section from yml file
	ser.ServiceName = serviceName
	serviceMap[serviceName] = ser // Gene: Add service instance into service map

	section, err := config.ConfigInstance.DIY(serviceName)
	if(err != nil){
		panic("Get service config section error " + err.Error())
	}

	// Get database config and create a db object
	object := section.(map[string]interface{})
	//if (object["db"] != nil){
	//	db := new (database.Database)
	//	db.InitDatabase(object["db"])
	//	ser.ServiceDB, err = db.GetDB()
	//	if (err != nil){
	//		beego.Error("Get service database error:", serviceName)
	//	}
	//}


	ser.ServiceConfig = object
}


func (ser *Service) HelloWorld() {
	// Gene: Hello world
	fmt.Println("Hello world")
}

func GetServiceInstance(serviceName string) (*Service, error)  {
	instance, existed := serviceMap[serviceName]
	if(existed){
		return instance, nil
	} else {
		return nil, errors.New("Can not find service instance")
	}
}

//func GetNetworkInterface() string {
//
//}