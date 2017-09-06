package database
import (
	"strconv"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego"

	"base/constants"
	"utils/encrypt"
)

type Database struct  {
	configObject map[string]interface{}
}

func (dbInstance *Database) InitDatabase(obj interface{})  {
	dbInstance.configObject = obj.(map[string]interface{})
}

func (dbInstance *Database) GetDB() (*sql.DB, error) {
	databaseType := dbInstance.getDBType()
	var db *sql.DB

	if (databaseType == constants.PostgreSQL) {
		sslMode := dbInstance.getDBSSLMode()
		if (sslMode) {
			db, err := sql.Open("postgres", "user="+dbInstance.getDBUser()+" password="+dbInstance.getDBPassword()+" dbname="+dbInstance.getDBName()+" host="+dbInstance.getDBHost()+" port="+dbInstance.getDBPort())
			if err != nil {
				beego.Error("Can not open database")
				return db, err
			}
			return db, nil
		} else {
			db, err := sql.Open("postgres", "user="+dbInstance.getDBUser()+" password="+dbInstance.getDBPassword()+" dbname="+dbInstance.getDBName()+" host="+dbInstance.getDBHost()+" port="+dbInstance.getDBPort()+" sslmode=disable")
			if err != nil {
				beego.Error("Can not open database")
				return db, err
			}
			return db, nil
		}

	}

	return db, nil
}

func (dbInstance *Database) getDBType() string {
	if (dbInstance.configObject["dbtype"] != nil) {
		return dbInstance.configObject["dbtype"].(string)
	} else {
		return ""
	}
}

func (dbInstance *Database) getDBHost() string {
	if (dbInstance.configObject["dbhost"] != nil) {
		return dbInstance.configObject["dbhost"].(string)
	} else {
		return ""
	}
}

func (dbInstance *Database) getDBPort() string {
	if (dbInstance.configObject["dbport"] != nil) {
		return strconv.FormatInt(dbInstance.configObject["dbport"].(int64),10)
	} else {
		return ""
	}
}

func (dbInstance *Database) getDBUser() string {
	if (dbInstance.configObject["dbuser"] != nil) {
		return dbInstance.configObject["dbuser"].(string)
	} else {
		return ""
	}
}

func (dbInstance *Database) getDBPassword() string {
	// Database password must be encrypted by AES256
	if (dbInstance.configObject["dbpwd"] != nil) {
		encryptedPassword := dbInstance.configObject["dbpwd"].(string)
		encryptedPasswordHex, err := encrypt.StringToHex(encryptedPassword)
		if (err != nil) {
			beego.Error("String to hex error during decrypt")
		}
		AESKeyHex, err := encrypt.StringToHex(constants.AES_Key)
		if (err != nil) {
			beego.Error("String to hex error during decrypt")
		}
		AESIvHex, err := encrypt.StringToHex(constants.AES_IV)
		if (err != nil) {
			beego.Error("String to hex error during decrypt")
		}
		_password, err := encrypt.AESDecrypt(encryptedPasswordHex, AESKeyHex, AESIvHex)
		if (err != nil) {
			beego.Error("AES decrypt password error!", err.Error())
		}
		password := encrypt.StringChop(_password)
		return password

	} else {
		return ""
	}
}

func (dbInstance *Database) getDBName() string {
	if (dbInstance.configObject["dbname"] != nil) {
		return dbInstance.configObject["dbname"].(string)
	} else {
		return ""
	}
}

func (dbInstance *Database) getDBSSLMode() bool {
	if (dbInstance.configObject["sslmode"] != nil) {
		result := dbInstance.configObject["sslmode"].(bool)
		return result
	} else {
		return false
	}
}