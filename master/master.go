package master 

import (
	"github.com/astaxie/beego/logs"
    "github.com/google/gopacket/pcap"
	"owlhmaster/utils"
	"owlhmaster/database"
	"io/ioutil"
	"os"
	"errors"
	"strings"
	"bytes"
	"os/exec"
	"strconv"
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

func UpdateMasterNetworkInterface(anode map[string]string)(err error) {
	err = ndb.UpdateMasterNetworkInterface(anode)
	if err != nil { logs.Error("master/UpdateMasterNetworkInterface Error updating interface selected: "+err.Error()); return err}
    return nil
}

func UpdateMasterStapInterface(anode map[string]string)(err error) {
	err = ndb.UpdatePluginValueMaster(anode["uuid"],anode["param"],anode["value"])
	if err != nil { logs.Error("master/UpdateMasterStapInterface Error updating interface selected: "+err.Error()); return err}
    return nil
}

func LoadMasterNetworkValuesSelected()(data map[string]map[string]string ,err error) {
    data,err = ndb.LoadMasterNetworkValuesSelected()
	if err != nil { logs.Error("master/LoadMasterNetworkValuesSelected Error getting interface selected by user for Master: "+err.Error()); return nil,err}
    return data,err
}

func PingServiceMaster()(err error) {
	masterService := map[string]map[string]string{}
    masterService["service"] = map[string]string{}
    masterService["service"]["dstPath"] = ""
    masterService["service"]["file"] = ""
	masterService,err = utils.GetConf(masterService)
	if err != nil {logs.Error("master/PingServiceMaster -- Error GetConf service data: "+err.Error()); return err}
	dstPath := masterService["service"]["dstPath"]
	file := masterService["service"]["file"]

	if _, err := os.Stat(dstPath+file); os.IsNotExist(err) {
		return errors.New("Service don't exists")
	}else{
		logs.Info("OwlHmaster service already exists")
		return nil
	}
}

func DeployServiceMaster()(err error) {
	masterService := map[string]map[string]string{}
    masterService["service"] = map[string]string{}
    masterService["service"]["dstPath"] = ""
    masterService["service"]["file"] = ""
    masterService["service"]["origPath"] = ""
    masterService["service"]["reload"] = ""
    masterService["service"]["enable"] = ""
	masterService,err = utils.GetConf(masterService)
	if err != nil {logs.Error("master/DeployServiceMaster -- Error GetConf service data: "+err.Error()); return err}
	dstPath := masterService["service"]["dstPath"]
	file := masterService["service"]["file"]
	origPath := masterService["service"]["origPath"]
	reload := masterService["service"]["reload"]
	enable := masterService["service"]["enable"]

	if _, err := os.Stat(dstPath+file); os.IsNotExist(err) {
		// //copy file
		err = utils.CopyFile(dstPath, origPath, file, 0)
		if err != nil {logs.Error("master/Copy Error Copying file: "+err.Error()); return err}
	
		// //exec reload
		_,err = exec.Command("bash", "-c", reload).Output()
		if err != nil{logs.Error("DeployServiceMaster Error reload service: "+err.Error()); return err}

		// //exec enable
		_,err = exec.Command("bash", "-c", enable).Output()
		if err != nil{logs.Error("DeployServiceMaster Error enabling service: "+err.Error()); return err}

		// //return nil
		logs.Info("OwlHmaster service deployed successfully!")
		return nil
	}else{
		logs.Info("OwlHmaster service already exists")
		return nil
	}
}

func AddPluginServiceMaster(anode map[string]string) (err error) {
    uuid := utils.Generate()
    err = ndb.InsertPluginService(uuid, "name", anode["name"]); if err != nil {logs.Error("InsertPluginService name Error: "+err.Error()); return err}
	err = ndb.InsertPluginService(uuid, "type", anode["type"]); if err != nil {logs.Error("InsertPluginService type Error: "+err.Error()); return err}
	err = ndb.InsertPluginService(uuid, "interface", anode["interface"]); if err != nil {logs.Error("InsertPluginService interface Error: "+err.Error()); return err}
	err = ndb.InsertPluginService(uuid, "port", anode["port"]); if err != nil {logs.Error("InsertPluginService port Error: "+err.Error()); return err}
	err = ndb.InsertPluginService(uuid, "cert", anode["cert"]); if err != nil {logs.Error("InsertPluginService certtificate Error: "+err.Error()); return err}
	err = ndb.InsertPluginService(uuid, "pid", "none"); if err != nil {logs.Error("InsertPluginService pid Error: "+err.Error()); return err}

    if anode["type"] == "socket-pcap"{        
        err = ndb.InsertPluginService(uuid, "pcap-path", anode["pcap-path"]); if err != nil {logs.Error("InsertPluginService pcap-path Error: "+err.Error()); return err}
        err = ndb.InsertPluginService(uuid, "pcap-prefix", anode["pcap-prefix"]); if err != nil {logs.Error("InsertPluginService pcap-prefix Error: "+err.Error()); return err}
        err = ndb.InsertPluginService(uuid, "bpf", anode["bpf"]); if err != nil {logs.Error("InsertPluginService bpf Error: "+err.Error()); return err}
	}

    return nil
}

func DeleteServiceMaster(anode map[string]string)(err error) {
	err = ndb.DeleteServiceMaster(anode["uuid"])
    if err != nil {logs.Error("plugin/DeleteServiceMaster error: "+err.Error()); return err}

    if _, err := os.Stat("/etc/suricata/bpf/"+anode["uuid"]+" - filter.bpf"); !os.IsNotExist(err) {
        err = os.Remove("/etc/suricata/bpf/"+anode["uuid"]+" - filter.bpf")
        if err != nil {logs.Error("plugin/SaveSuricataInterface error deleting a pid file: "+err.Error())}
    }

    return err
}

func ModifyStapValuesMaster(anode map[string]string)(err error) {
    if anode["type"] == "socket-network"{
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"name",anode["name"]); if err != nil {logs.Error("UpdatePluginValueMaster socket-network Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"port",anode["port"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-network Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"cert",anode["cert"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-network Error: "+err.Error()); return err}
    }else if anode["type"] == "socket-pcap"{
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"name",anode["name"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-pcap Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"port",anode["port"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-pcap Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"cert",anode["cert"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-pcap Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"pcap-path",anode["pcap-path"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-pcap Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"pcap-prefix",anode["pcap-prefix"]) ; if err != nil {logs.Error("UpdatePluginValueMaster socket-pcap Error: "+err.Error()); return err}
    }
    return nil
}

func SetBPF(anode map[string]string)(err error){
    err = ndb.UpdatePluginValueMaster(anode["uuid"],anode["param"],anode["value"])
	if err != nil { logs.Error("master/UpdateMasterStapInterface Error updating interface selected: "+err.Error()); return err}
    return nil
}

func DeployStapServiceMaster(anode map[string]string)(err error) { 
	allPlugins,err := ndb.GetPlugins()
	if err != nil {logs.Error("DeployStapServiceMaster error getting all plugins from DB: "+err.Error()); return err}

    if anode["type"] == "socket-network" {
		logs.Debug("DEPLOY DEPLOY DEPLOY DEPLOY DEPLOY DEPLOY DEPLOY DEPLOY DEPLOY ")
		logs.Warn("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'")
        pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapServiceMaster deploy socket-network Error: "+err.Error()); return err}
        pidValue := strings.Split(string(pid), "\n")
		logs.Notice(pidValue[0])
		if pidValue[0] != "" {
            return nil
        }

		logs.Debug("/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
        var errores bytes.Buffer
        cmd.Stdout = &errores
        err = cmd.Start()
        if err != nil {logs.Error("DeployStapServiceMaster deploying Error: "+err.Error()); return err}        

		logs.Warn("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'")
        pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapServiceMaster deploy socket-network Error: "+err.Error()); return err}
		pidValue = strings.Split(string(pid), "\n")
		logs.Notice(pidValue[0])
        if pidValue[0] != "" {
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid",pidValue[0]); if err != nil {logs.Error("DeployStapServiceMaster change pid to value Error: "+err.Error()); return err}
        }else{
            return nil
        }
    }else if anode["type"] == "socket-pcap" {
		logs.Warn("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'")
		pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapServiceMaster deploy socket-network Error: "+err.Error()); return err}
        pidValue := strings.Split(string(pid), "\n")
		logs.Notice(pidValue[0])
        if pidValue[0] != "" {
            return nil
        }

        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[anode["uuid"]]["pcap-path"]+allPlugins[anode["uuid"]]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[anode["uuid"]]["bpf"]+"\" &")
        var errores bytes.Buffer
        cmd.Stdout = &errores
        err = cmd.Start()
        if err != nil {logs.Error("DeployStapServiceMaster deploying Error: "+err.Error()); return err}        

		logs.Warn("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'")
        pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v bash | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapServiceMaster deploy socket-network Error: "+err.Error()); return err}
		pidValue = strings.Split(string(pid), "\n")
		logs.Notice(pidValue[0])
        if pidValue[0] != "" {
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid",pidValue[0]); if err != nil {logs.Error("DeployStapServiceMaster change pid to value Error: "+err.Error()); return err}
        }else{
            return nil
        }
    }
    
    return nil
}

func StopStapServiceMaster(anode map[string]string)(err error) {
	logs.Debug("STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP STOP ")
    allPlugins,err := ndb.GetPlugins()
    pidToInt,err := strconv.Atoi(allPlugins[anode["uuid"]]["pid"])
    if err != nil {logs.Error("StopStapServiceMaster pid to int error: "+err.Error()); return err}
    process, err := os.FindProcess(pidToInt)
    if err != nil {logs.Error("StopStapServiceMaster pid process not found: "+err.Error()); return err}
    err = process.Kill()
    if err != nil {logs.Error("StopStapServiceMaster Kill pid process Error: "+err.Error()); return err}
    err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid","none") ; if err != nil {logs.Error("StopStapServiceMaster change pid to none Error: "+err.Error()); return err}

    return nil
}