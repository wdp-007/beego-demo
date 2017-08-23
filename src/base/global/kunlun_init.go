package global

// Gene: This file is used to store the global variable for serivce
var RegistryEndpoint  *string
var ServiceArray []string
var ServicePort *string
var ShuttleToken *string

func init (){
	RegistryEndpoint = new(string)
	ServiceArray = *new([]string)
	ServicePort = new(string)
	ShuttleToken = new(string)
}
