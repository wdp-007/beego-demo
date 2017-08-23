package metadata

import (
	"services/metadata/controller"
	"github.com/astaxie/beego"
)

func MetadataRouter()  {
	beego.Router("/metadata/test", &controller.MetadataController{}, "get:Test")
	beego.Router("/metadata/healthz", &controller.MetadataController{}, "get:Healthz")
}