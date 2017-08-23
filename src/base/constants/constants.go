package constants

import (
	"os"
	"path/filepath"
)

var CurrentDirectory string   // Parent Directory of src
var LogDirectory string
var LogFilePath string
var MsgRelayLogFilePath string
var ConfigDirectory string
var DoctorSSLKeyPath string
var DoctorSSLCrtPath string
var KeeperLogFilePath string
var KeeperSSHKeyFile string
var KeeperSSLKeyPath string
var KeeperSSLCrtPath string
var ScriptsDirectory string
var OIDCPublicKeyDirectory string

func init (){

	CurrentDirectory, _ = os.Getwd()

	LogDirectory = filepath.Join(CurrentDirectory, "logs")
	logDirPath, _ := os.Stat(LogDirectory)
	if (logDirPath == nil){
		// Create log directory it doesn't existed
		os.Mkdir(LogDirectory, os.ModePerm)
	}
	LogFilePath = filepath.Join(LogDirectory, "taishan.log")
	MsgRelayLogFilePath = filepath.Join(LogDirectory, "doctor_msg_relay.log")

	ConfigDirectory = filepath.Join(CurrentDirectory, "config")
}

const (
	DoctorVersion  = "4.11.26"
	KeeperVersion	=	"4.3.1"
)

// Service name
const (
	// Service
	ServiceMetadata = "metadata"
)

