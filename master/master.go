package master 

import (
    "github.com/astaxie/beego/logs"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
    // "owlhmaster/database"
	// "owlhmaster/aboutme"
	"io/ioutil"
	"encoding/json"
)

func GetMasterTitle() (data string, err error) {
	confFilePath := "/etc/owlh/conf/main.conf"
	jsonPathBpf, err := ioutil.ReadFile(confFilePath)
    if err != nil {
        logs.Error("utils/GetConf -> can't open Conf file -> " + confFilePath)
        return "",err
    }
    var anode map[string]map[string]string
    json.Unmarshal(jsonPathBpf, &anode)

    return anode["master"]["name"], err
}