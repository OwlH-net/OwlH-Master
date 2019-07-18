package master 

import (
	"github.com/astaxie/beego/logs"
    "github.com/google/gopacket/pcap"
	"owlhmaster/utils"
	"io/ioutil"
	"os/exec"
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

func GetNetworkInterface()(values map[string]string, err error) {
	data := make(map[string]string)
	interfaces, err := pcap.FindAllDevs()
	if err != nil {logs.Error("GetNetworkInterface Master pcap.FindAllDevs error: "+err.Error()); return nil,err}

	for _, localInt := range interfaces {
		data[localInt.Name] = localInt.Name
	}

	logs.Notice(data)

	return data, err
}

func DeployMaster(anode map[string]string)(err error) {
	loadData := map[string]map[string]string{}
	loadData["deploy"] = map[string]string{}
	loadData["deploy"][anode["value"]] = ""
	loadData,err = utils.GetConf(loadData)
	if err != nil { logs.Error("DeployMaster Error getting data from main.conf"); return err}
	
	_,err = exec.Command("bash", "-c", loadData["deploy"][anode["value"]]).Output()
	if err != nil{logs.Error("DeployMaster Error deploying "+loadData["deploy"][anode["value"]]+": "+err.Error()); return err}
	
	return nil
}