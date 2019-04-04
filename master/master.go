package master 

import (
    "github.com/astaxie/beego/logs"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
    // "owlhmaster/database"
	// "owlhmaster/aboutme"
	"owlhmaster/utils"
	// "io/ioutil"
	// "encoding/json"
)

//Obtain title for webpage from main.conf
func GetMasterTitle() (data string, err error) {
	loadData := map[string]map[string]string{}
	loadData["master"] = map[string]string{}
	loadData["master"]["name"] = ""
	loadData,err = utils.GetConf(loadData)
	if err != nil {
		logs.Error("Error getting Master title from main.conf")
		return "-",err
	}
	return loadData["master"]["name"], nil
}