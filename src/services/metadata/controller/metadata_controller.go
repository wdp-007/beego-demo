package controller

import (
	"github.com/astaxie/beego"
)

type MetadataController struct {
	beego.Controller
}

func (this *MetadataController) Test()  {
	this.Ctx.WriteString("Metadata Work")
	beego.Info("this is test api")
}

func (this *MetadataController) Healthz()  {
	this.Ctx.WriteString("ok")
}