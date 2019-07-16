package master 

import (
    "github.com/astaxie/beego/logs"
	"owlhmaster/utils"
	"io/ioutil"
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

func GetFileContent(file string) (data map[string]string, err error) {
	sendBackArray := make(map[string]string)
    
    //create map and obtain file
    loadData := map[string]map[string]string{}
	loadData["files"] = map[string]string{}
	loadData["files"][file] = ""
	loadData,err = utils.GetConf(loadData)
	if err != nil {logs.Error("SendFile Error getting data from main.conf"); return nil,err}
	    
    //save url from file selected and open file
    fileConfPath := loadData["files"][file]
	fileReaded, err := ioutil.ReadFile(fileConfPath) // just pass the file name
    if err != nil {
		logs.Error("Error reading file for path: "+fileConfPath)
        return nil,err
    }
	
	sendBackArray["fileContent"] = string(fileReaded)
    sendBackArray["fileName"] = file

	return sendBackArray, nil
}

func SaveFileContent(file map[string]string) (err error) {
	    //Get full path
		loadData := map[string]map[string]string{}
		loadData["files"] = map[string]string{}
		loadData["files"][file["file"]] = ""
		loadData,err = utils.GetConf(loadData)
		if err != nil {
			logs.Error("SaveFile Error getting data from main.conf")
		}
	
		//make file backup before overwrite
		err = utils.BackupFullPath(loadData["files"][file["file"]])
		if err != nil {
			logs.Info("SaveFile. Error doing backup with function BackupFullPath: "+err.Error())
			return err
		}
	
		//make byte array for save the file modified
		bytearray := []byte(file["content"])
		err = utils.WriteNewDataOnFile(loadData["files"][file["file"]], bytearray)
		if err != nil {
			logs.Info("SaveFile. Error doing backup with function WriteNewDataOnFile: "+err.Error())
			return err
		}
		return nil

    return err
}