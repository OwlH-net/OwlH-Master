package master 

import (
    "github.com/astaxie/beego/logs"
    "github.com/google/gopacket/pcap"
    "owlhmaster/utils"
    "owlhmaster/validation"
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
    if err != nil {logs.Error("DeleteServiceMaster error: "+err.Error()); return err}

    //StopService

    // if _, err := os.Stat("/etc/suricata/bpf/"+anode["uuid"]+" - filter.bpf"); !os.IsNotExist(err) {
    //     err = os.Remove("/etc/suricata/bpf/"+anode["uuid"]+" - filter.bpf")
    //     if err != nil {logs.Error("plugin/SaveSuricataInterface error deleting a pid file: "+err.Error())}
    // }

    return err
}

func ModifyStapValuesMaster(anode map[string]string)(err error) {
    allPlugins,err := ndb.PingPlugins()
    if anode["type"] == "socket-network" || anode["type"] == "socket-pcap"{
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"name",anode["name"]); if err != nil {logs.Error("ModifyStapValuesMaster "+anode["type"]+" Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"port",anode["port"]) ; if err != nil {logs.Error("ModifyStapValuesMaster "+anode["type"]+" Error: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster(anode["uuid"],"cert",anode["cert"]) ; if err != nil {logs.Error("ModifyStapValuesMaster "+anode["type"]+" Error: "+err.Error()); return err}
        if anode["type"] == "socket-pcap"{
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pcap-path",anode["pcap-path"]) ; if err != nil {logs.Error("ModifyStapValuesMaster socket-pcap Error: "+err.Error()); return err}
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pcap-prefix",anode["pcap-prefix"]) ; if err != nil {logs.Error("ModifyStapValuesMaster socket-pcap Error: "+err.Error()); return err}
        }
        for x := range allPlugins{
            if ((allPlugins[x]["type"] == "socket-network" || allPlugins[x]["type"] == "socket-pcap") && (anode["uuid"] != x)){
                if allPlugins[x]["port"] == anode["port"] {
                    err = StopStapServiceMaster(anode); if err != nil {logs.Error("ModifyStapValues "+anode["type"]+" stopping error: "+err.Error()); return err}        
                    logs.Error("Can't deploy "+anode["type"]+" or "+anode["type"]+" with the same port")
                    return errors.New("Can't deploy "+anode["type"]+" or "+anode["type"]+" with the same port")
                }
            }
        }        
        //restart services after update
        if allPlugins[anode["uuid"]]["pid"] != "none"{
            err = StopStapServiceMaster(anode); if err != nil {logs.Error("ModifyStapValuesMaster "+anode["type"]+" Error stopping service: "+err.Error()); return err}
            err = DeployStapServiceMaster(anode); if err != nil {logs.Error("ModifyStapValuesMaster "+anode["type"]+" Error deploying service: "+err.Error()); return err}
            logs.Notice(allPlugins[anode["uuid"]]["name"]+" service updated!!!")
        }        
    }
    return nil
}

func CheckServicesStatus()(){
    allPlugins,err := ndb.PingPlugins()
    if err != nil {logs.Error("CheckServicesStatus error getting all master plugins: "+err.Error())}

    for w := range allPlugins {
        if allPlugins[w]["pid"] != "none"{
            if allPlugins[w]["type"] == "socket-network" || allPlugins[w]["type"] == "socket-pcap"{
                pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[w]["port"]+" | grep -v grep | awk '{print $2}'").Output()
                if err != nil {logs.Error("CheckServicesStatus Checking previous PID for socket-network: "+err.Error())}
                pidValue := strings.Split(string(pid), "\n")
                
                if pidValue[0] == ""{
                    if allPlugins[w]["type"] == "socket-network"{
                        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[w]["interface"]+" -\" &")
                        var errores bytes.Buffer
                        cmd.Stdout = &errores
                        err = cmd.Start()
                        if err != nil {logs.Error("CheckServicesStatus deploying Error socket-network: "+err.Error())}        
                    }else{
                        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[w]["pcap-path"]+allPlugins[w]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[w]["bpf"]+"\" &")
                        var errores bytes.Buffer
                        cmd.Stdout = &errores
                        err = cmd.Start()
                        if err != nil {logs.Error("CheckServicesStatus deploying Error socket-network: "+err.Error())}        
                    }                    

                    pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[w]["port"]+" | grep -v grep | awk '{print $2}'").Output()
                    if err != nil {logs.Error("CheckServicesStatus deploy socket-network Error: "+err.Error())}
                    pidValue = strings.Split(string(pid), "\n")
                    if pidValue[0] != "" {
                        err = ndb.UpdatePluginValueMaster(w,"pid",pidValue[0]); if err != nil {logs.Error("CheckServicesStatus change pid to value Error socket-network: "+err.Error())}
                    }
                    logs.Notice("Socket-network deploy after Master stops: PID: "+pidValue[0])
                }
            // }else if  allPlugins[w]["type"] == "socket-pcap"{
            //     pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[w]["port"]+" | grep -v grep | awk '{print $2}'").Output()
            //     if err != nil {logs.Error("CheckServicesStatus Checking previous PID for socket-pcap: "+err.Error())}
            //     pidValue := strings.Split(string(pid), "\n")
                
            //     if pidValue[0] == ""{
            //         cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[w]["pcap-path"]+allPlugins[w]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[w]["bpf"]+"\" &")
            //         var errores bytes.Buffer
            //         cmd.Stdout = &errores
            //         err = cmd.Start()
            //         if err != nil {logs.Error("CheckServicesStatus deploying Error socket-pcap: "+err.Error())}        

            //         pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[w]["port"]+" | grep -v grep | awk '{print $2}'").Output()
            //         if err != nil {logs.Error("CheckServicesStatus deploy socket-pcap Error: "+err.Error())}
            //         pidValue = strings.Split(string(pid), "\n")
            //         if pidValue[0] != "" {
            //             err = ndb.UpdatePluginValueMaster(w,"pid",pidValue[0]); if err != nil {logs.Error("CheckServicesStatus change pid to value Error socket-pcap: "+err.Error())}
            //         }
            //         logs.Notice("Socket-network deploy after Master stops: PID: "+pidValue[0])
            //     }
            }
        }
    }
}

func PingPlugins()(data map[string]map[string]string, err error) {
    allPlugins,err := ndb.PingPlugins()
    for x,_ := range allPlugins {
        //check if process is running even though database status is enabled
        if (allPlugins[x]["type"] == "socket-pcap" || allPlugins[x]["type"] == "socket-network") && allPlugins[x]["pid"] != "none"{
            pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[x]["port"]+" | grep -v grep | awk '{print $2}'").Output()
            if err != nil {logs.Error("ping/PingPluginsNode Checking STAP PID: "+err.Error())}
            if strings.Split(string(pid), "\n")[0] == "" {
                allPlugins[x]["running"] = "false"
            }else{
                allPlugins[x]["running"] = "true"
            }
        }
    }
    return allPlugins,err
}

func DeployStapServiceMaster(anode map[string]string)(err error) { 
    allPlugins,err := ndb.PingPlugins()
    if anode["type"] == "socket-network" {
        pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapService deploy socket-network Error: "+err.Error()); return err}
        pidValue := strings.Split(string(pid), "\n")
        // logs.Error(pidValue)
        if pidValue[0] != "" {
            logs.Error("Socket to network deployed. Can't deploy more than one stap service at the same port")
            return errors.New("Can't deploy more than one socket at the same port")
        }
        // logs.Debug("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'")
        // logs.Debug("/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
        var errores bytes.Buffer
        cmd.Stdout = &errores
        err = cmd.Start()
        if err != nil {logs.Error("DeployStapService deploying Error: "+err.Error()); return err}        

        pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapService deploy socket-network Error: "+err.Error()); return err}
        pidValue = strings.Split(string(pid), "\n")
        if pidValue[0] != "" {
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid",pidValue[0]); if err != nil {logs.Error("DeployStapService change pid to value Error: "+err.Error()); return err}
        }
        logs.Notice("Deploy successful --> Type: "+allPlugins[anode["uuid"]]["type"]+" Description: "+allPlugins[anode["uuid"]]["name"]+"  --  SOCAT: "+pidValue[0])
    }else if anode["type"] == "socket-pcap" {
        pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapService deploy socket-network Error: "+err.Error()); return err}
        pidValue := strings.Split(string(pid), "\n")
        // logs.Error(pidValue)
        if pidValue[0] != "" {
            logs.Error("Socket to pcap deployed. Can't deploy more than one stap service at the same port")
            return errors.New("Can't deploy more than one socket at the same port")   
        }
        // logs.Debug("ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'")
        // logs.Debug("/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[anode["uuid"]]["pcap-path"]+allPlugins[anode["uuid"]]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[anode["uuid"]]["bpf"]+"\" &")
        cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[anode["uuid"]]["pcap-path"]+allPlugins[anode["uuid"]]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[anode["uuid"]]["bpf"]+"\" &")
        var errores bytes.Buffer
        cmd.Stdout = &errores
        err = cmd.Start()
        if err != nil {logs.Error("DeployStapService deploying Error: "+err.Error()); return err}        

        pid, err = exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapService deploy socket-network Error: "+err.Error()); return err}
        pidValue = strings.Split(string(pid), "\n")
        if pidValue[0] != "" {
            err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid",pidValue[0]); if err != nil {logs.Error("DeployStapService change pid to value Error: "+err.Error()); return err}
        }
        logs.Notice("Deploy successful --> Type: "+allPlugins[anode["uuid"]]["type"]+" Description: "+allPlugins[anode["uuid"]]["name"]+"  --  SOCAT: "+pidValue[0])
    }
    return nil
}

func StopStapServiceMaster(anode map[string]string)(err error) {
    allPlugins,err := ndb.PingPlugins()
    if err != nil {logs.Error("StopStapServiceMaster error getting plugins from database: "+err.Error()); return err}
    
    pidToInt,_ := strconv.Atoi(allPlugins[anode["uuid"]]["pid"])
    process, _ := os.FindProcess(pidToInt)
    _ = process.Kill()
    
    err = ndb.UpdatePluginValueMaster(anode["uuid"],"pid","none") ; if err != nil {logs.Error("StopStapServiceMaster change pid to none Error: "+err.Error()); return err}

    return nil
}

func SaveStapInterface(anode map[string]string) (err error) {    
    err = ndb.UpdatePluginValueMaster(anode["uuid"], anode["param"], anode["value"])
    if err != nil {logs.Error("SaveStapInterface Change STAP Interface error: "+err.Error()); return err}
    return nil
}

func SetBPF(anode map[string]string)(err error){
    err = ndb.UpdatePluginValueMaster(anode["uuid"],anode["param"],anode["value"])
    if err != nil { logs.Error("master/UpdateMasterStapInterface Error updating interface selected: "+err.Error()); return err}

    return nil
}

func GetIncidents()(data map[string]map[string]string, err error){
    data,err = ndb.GetIncidents()
    if err!=nil{logs.Error("Error getting incidents from database: "+err.Error()); return nil,err}

    return data,err
}

func PutIncident(anode map[string]string)(err error){
    err = ndb.PutIncident(anode["uuid"], anode["param"], anode["value"])
    if err!=nil{logs.Error("Error putting incidents into database: "+err.Error()); return err}

    return nil
}

func SaveZeekValues(anode map[string]string) (err error) {
    if anode["param"] == "nodeConfig"{
        if _, err := os.Stat(anode["nodeConfig"]); os.IsNotExist(err) {
            logs.Error("master/SaveZeekValues node config path error: "+err.Error()); return err
        }
        err = ndb.UpdatePluginValueMaster("zeek", anode["param"], anode["nodeConfig"])
        if err != nil{logs.Error("master/SaveZeekValues Error writting node content: "+err.Error()); return err}
    }
    if anode["param"] == "networksConfig"{
        if _, err := os.Stat(anode["networksConfig"]); os.IsNotExist(err) {
            logs.Error("master/SaveZeekValues networks config path error: "+err.Error()); return err
        }
        err = ndb.UpdatePluginValueMaster("zeek", anode["param"], anode["networksConfig"])
        if err != nil{logs.Error("master/SaveZeekValues Error writting networks content: "+err.Error()); return err}
        }            
    if anode["param"] == "policies"{
        if _, err := os.Stat(anode["policiesMaster"]); os.IsNotExist(err) {
            logs.Error("master/SaveZeekValues policies path error: "+err.Error()); return err
        }
        err = ndb.UpdatePluginValueMaster("zeek", "policiesMaster", anode["policiesMaster"]); if err != nil{logs.Error("master/SaveZeekValues Error writting policies content: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster("zeek", "policiesNode", anode["policiesNode"]); if err != nil{logs.Error("master/SaveZeekValues Error writting policies content: "+err.Error()); return err}
    }            
    if anode["param"] == "variables"{
        if _, err := os.Stat(anode["variables1"]); os.IsNotExist(err) {
            logs.Error("master/SaveZeekValues variables path error: "+err.Error()); return err
        }
        err = ndb.UpdatePluginValueMaster("zeek", "variables1", anode["variables1"]); if err != nil{logs.Error("master/SaveZeekValues Error writting variables content: "+err.Error()); return err}
        err = ndb.UpdatePluginValueMaster("zeek", "variables2", anode["variables2"]); if err != nil{logs.Error("master/SaveZeekValues Error writting variables content: "+err.Error()); return err}
    }            

    return err
}

func PingPluginsMaster()(data map[string]map[string]string, err error){
    data,err = ndb.GetPlugins()
    if err != nil{logs.Error("master/PingPluginsMaster Error Getting Master plugins: "+err.Error()); return nil,err}    
    return data,nil
}

func GetPathFileContent(param string) (file map[string]string, err error) {
    data,err := ndb.GetPlugins()
    sendBackArray := make(map[string]string)

    fileReaded, err := ioutil.ReadFile(data["zeek"][param]) // just pass the file name
    if err != nil {logs.Error("GetPathFileContent Error reading file for path: "+data["zeek"][param]); return nil,err}
    
    sendBackArray["fileContent"] = string(fileReaded)
    sendBackArray["fileName"] = param

    return sendBackArray, nil
}

func SaveFilePathContent(file map[string]string) (err error) {
    data,err := ndb.GetPlugins()

    //make file backup before overwrite
    err = utils.BackupFullPath(data["zeek"][file["file"]])
    if err != nil {
        logs.Info("SaveFilePathContent. Error doing backup with function BackupFullPath: "+err.Error())
        return err
    }

    //make byte array for save the file modified
    bytearray := []byte(file["content"])
    err = utils.WriteNewDataOnFile(data["zeek"][file["file"]], bytearray)
    if err != nil {
        logs.Info("SaveFilePathContent. Error doing backup with function WriteNewDataOnFile: "+err.Error())
        return err
    }
    
    return nil
}

func Login(data map[string]string)(newToken string, err error){ 
    users,err := ndb.GetLoginData()
    if err != nil{logs.Error("master/Login Error Getting user values: "+err.Error()); return "",err}    
    
    //check values
    for x := range users{
        if users[x]["user"] == data["user"]{
            check, err := validation.CheckPasswordHash(data["password"], users[x]["pass"])
            if err != nil{return "", err}
            if check{
                // userExists = true
                token, err := validation.Encode(x, data["user"], users[x]["secret"])
                if err != nil {return "",err}
                return token,nil
            }
        }
    }

    return "", errors.New("There are not token. Error creating Token")
}