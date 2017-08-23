package initialize

import (
	"base/constants"
	"services/metadata"
)

func init()  {

}

func InitService(services []string) {

	for _, service := range services {

		if(service == constants.ServiceMetadata){

			service := new(metadata.MetadataService)
			service.Init(constants.ServiceMetadata)
		} else{

		}
	}


}