package master 

import (
    "github.com/astaxie/beego/logs"
	"owlhmaster/utils"
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