package metadata

import (
	"base/service"
	"github.com/astaxie/beego"
)

type MetadataService struct {
	service.Service
}

func (serv *MetadataService) Init(name string) {

	beego.Informational("Init service: " + name)
	serv.Service.InitService(name)
	MetadataRouter()
}
