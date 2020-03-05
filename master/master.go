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
    name, err := utils.GetKeyValueString("master", "name")
    if err != nil {
        logs.Error("Error getting Master title from main.conf")
        return "-",err
    }
    return name, nil
}

func GetFileContent(file string) (data map[string]string, err error) {
    sendBackArray := make(map[string]string)

    fileConfPath, err := utils.GetKeyValueString("files", file)
    if err != nil {logs.Error("SendFile Error getting data from main.conf"); return nil,err}
        
    //save url from file selected and open file
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
        filePath, err := utils.GetKeyValueString("files", file["file"])
        if err != nil {logs.Error("SaveFile Error getting data from main.conf"); return err}
    
        //make file backup before overwrite
        err = utils.BackupFullPath(filePath)
        if err != nil {
            logs.Info("SaveFile. Error doing backup with function BackupFullPath: "+err.Error())
            return err
        }
    
        //make byte array for save the file modified
        bytearray := []byte(file["content"])
        err = utils.WriteNewDataOnFile(filePath, bytearray)
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
    value, err := utils.GetKeyValueString("deploy", anode["value"])
    if err != nil { logs.Error("DeployMaster Error getting data from main.conf"); return err}
    
    _,err = exec.Command("bash", "-c", value).Output()
    if err != nil{logs.Error("DeployMaster Error deploying "+value+": "+err.Error()); return err}
    
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
    dstPath, err := utils.GetKeyValueString("service", "dstPath")
    if err != nil {logs.Error("master/PingServiceMaster -- Error main.conf service data: "+err.Error()); return err}
    file, err := utils.GetKeyValueString("service", "file")
    if err != nil {logs.Error("master/PingServiceMaster -- Error main.conf service data: "+err.Error()); return err}

    if _, err := os.Stat(dstPath+file); os.IsNotExist(err) {
        return errors.New("Service don't exists")
    }else{
        logs.Info("OwlHmaster service already exists")
        return nil
    }
}

func DeployServiceMaster()(err error) {
    dstPath, err := utils.GetKeyValueString("service", "dstPath")
    if err != nil {logs.Error("master/DeployServiceMaster -- Error main.conf service data: "+err.Error()); return err}
    file, err := utils.GetKeyValueString("service", "file")
    if err != nil {logs.Error("master/DeployServiceMaster -- Error main.conf service data: "+err.Error()); return err}
    origPath, err := utils.GetKeyValueString("service", "origPath")
    if err != nil {logs.Error("master/DeployServiceMaster -- Error main.conf service data: "+err.Error()); return err}
    reload, err := utils.GetKeyValueString("service", "reload")
    if err != nil {logs.Error("master/DeployServiceMaster -- Error main.conf service data: "+err.Error()); return err}
    enable, err := utils.GetKeyValueString("service", "enable")
    if err != nil {logs.Error("master/DeployServiceMaster -- Error main.conf service data: "+err.Error()); return err}

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
        }        
    }
    return nil
}

func CheckServicesStatus()(){
    socat, err := utils.GetKeyValueString("plugins", "socat")
    if err != nil {logs.Error("GetNodeFile error getting path from main.conf")}

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
                        // cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[w]["interface"]+" -\" &")
                        cmd := exec.Command("bash","-c",socat+" -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[w]["interface"]+" -\" &")
                        var errores bytes.Buffer
                        cmd.Stdout = &errores
                        err = cmd.Start()
                        if err != nil {logs.Error("CheckServicesStatus deploying Error socket-network: "+err.Error())}        
                    }else{
                        // cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[w]["pcap-path"]+allPlugins[w]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[w]["bpf"]+"\" &")
                        cmd := exec.Command("bash","-c",socat+" -d OPENSSL-LISTEN:"+allPlugins[w]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[w]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[w]["pcap-path"]+allPlugins[w]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[w]["bpf"]+"\" &")
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
    socat, err := utils.GetKeyValueString("plugins", "socat")
    if err != nil {logs.Error("GetNodeFile error getting path from main.conf"); return err}

    allPlugins,err := ndb.PingPlugins()
    if anode["type"] == "socket-network" {
        pid, err := exec.Command("bash","-c","ps -ef | grep socat | grep OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+" | grep -v grep | awk '{print $2}'").Output()
        if err != nil {logs.Error("DeployStapService deploy socket-network Error: "+err.Error()); return err}
        pidValue := strings.Split(string(pid), "\n")

        if pidValue[0] != "" {
            logs.Error("Socket to network deployed. Can't deploy more than one stap service at the same port")
            return errors.New("Can't deploy more than one socket at the same port")
        }

        // cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
        cmd := exec.Command("bash","-c",socat+" -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpreplay -t -i "+allPlugins[anode["uuid"]]["interface"]+" -\" &")
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
        
        if pidValue[0] != "" {
            logs.Error("Socket to pcap deployed. Can't deploy more than one stap service at the same port")
            return errors.New("Can't deploy more than one socket at the same port")   
        }
        
        // cmd := exec.Command("bash","-c","/usr/bin/socat -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[anode["uuid"]]["pcap-path"]+allPlugins[anode["uuid"]]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[anode["uuid"]]["bpf"]+"\" &")
        cmd := exec.Command("bash","-c",socat+" -d OPENSSL-LISTEN:"+allPlugins[anode["uuid"]]["port"]+",reuseaddr,pf=ip4,fork,cert="+allPlugins[anode["uuid"]]["cert"]+",verify=0 SYSTEM:\"tcpdump -n -r - -s 0 -G 50 -W 100 -w "+allPlugins[anode["uuid"]]["pcap-path"]+allPlugins[anode["uuid"]]["pcap-prefix"]+"%d%m%Y%H%M%S.pcap "+allPlugins[anode["uuid"]]["bpf"]+"\" &")
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
            if users[x]["ldap"] == "enabled" {
                check, err := validation.CheckLdap(data["password"], users[x]["pass"])
                if err != nil {return "",err}
                if check{
                    token, err := validation.Encode(x, data["user"], users[x]["secret"])
                    if err != nil {return "",err}
                    return token,nil
                }
            }else{
                check, err := validation.CheckPasswordHash(data["password"], users[x]["pass"])
                if err != nil {return "",err}
                if check{
                    token, err := validation.Encode(x, data["user"], users[x]["secret"])
                    if err != nil {return "",err}
                    return token,nil
                }
            }
        }
    }

    return "", errors.New("There are not token. Error creating Token")
}

func AddUser(data map[string]string)(err error){ 
    //check if this username already exists
    users,err := ndb.GetLoginData()
    for id := range users{
        if users[id]["user"] == data["user"]{
            return errors.New("This username is already in use")
        }
    }

    passHashed, err := validation.HashPassword(data["pass"])
    if err != nil{logs.Error("master/AddUser Error creating hash password: "+err.Error()); return err}    
    
    //insert username into db
    uuid := utils.Generate()
    secret := utils.Generate()
    //user
    err = ndb.InsertUser(uuid, "user", data["user"])
    if err != nil{logs.Error("master/AddUser Error inserting user into db: "+err.Error()); return err}    
    err = ndb.InsertUser(uuid, "pass", passHashed)
    if err != nil{logs.Error("master/AddUser Error inserting pass into db: "+err.Error()); return err}    
    err = ndb.InsertUser(uuid, "secret", secret)
    if err != nil{logs.Error("master/AddUser Error inserting secret into db: "+err.Error()); return err}
    err = ndb.InsertUser(uuid, "ldap", data["ldap"])
    if err != nil{logs.Error("master/AddUser Error inserting ldap into db: "+err.Error()); return err}
    
    return nil
}

func GetAllUsers()(data map[string]map[string]string, err error) {
    users,err := ndb.GetLoginData()
    if err != nil{logs.Error("master/GetAllUsers Error getting users: "+err.Error()); return nil, err}    

    //delete private data
    for id := range users{
        delete(users[id],"pass")
        delete(users[id],"secret")
    }

    //get user info
    usersGroupRoles,err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetAllUsers Error getting usersGroupRoles: "+err.Error()); return nil, err}    
    for user := range users{
        var groups []string 
        var roles []string
        for priv := range usersGroupRoles{
            if user == usersGroupRoles[priv]["user"]{
                if usersGroupRoles[priv]["group"] != "" { 
                    //get group name
                    group,err := ndb.GetuserGroupByUUID(usersGroupRoles[priv]["group"])
                    if err != nil{logs.Error("master/GetAllUsers Error getting groups: "+err.Error()); return nil, err}    
                    groups = append(groups, group[usersGroupRoles[priv]["group"]]["group"]) 
                }
                if usersGroupRoles[priv]["role"] != "" { 
                    //get role name
                    role,err := ndb.GetuserRoleByUUID(usersGroupRoles[priv]["role"])
                    if err != nil{logs.Error("master/GetAllUsers Error getting roles: "+err.Error()); return nil, err}    
                    roles = append(roles, role[usersGroupRoles[priv]["role"]]["role"]) 
                }                
            }
        }

        users[user]["roles"] = strings.Join(roles, ",")
        users[user]["groups"] = strings.Join(groups, ",") 
    }

    return users, err
}

func DeleteUser(anode map[string]string)(err error){    
    err = ndb.DeleteUser(anode["id"])
    if err != nil{logs.Error("master/DeleteUser Error deleting user: "+err.Error()); return err}    

    return nil
}

func AddGroupUsers(anode map[string]string) (err error) {
    uuid := utils.Generate()
    err = ndb.InsertGroupUsers(uuid, "group", anode["group"])
    // err = ndb.InsertGroupUsers(uuid, "permissions", anode["permissions"])
    if err != nil{logs.Error("master/AddGroupUsers Error inserting user into db: "+err.Error()); return err}    

    return nil
}

func AddRole(anode map[string]string) (err error) {
    uuid := utils.Generate()
    err = ndb.InsertRoleUsers(uuid, "role", anode["role"])
    err = ndb.InsertRoleUsers(uuid, "permissions", anode["permissions"])
    if err != nil{logs.Error("master/AddRole Error inserting user into db: "+err.Error()); return err}    
   
    return nil
}

func GetRolesForUser(roleID string)(data map[string]map[string]string, err error) {
    roles,err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/GetRolesForUser Error getting roles: "+err.Error()); return nil, err}    
    all,err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetRolesForUser Error getting all: "+err.Error()); return nil, err}    
    for x := range all{
        for y := range roles{
            if all[x]["role"] == y && all[x]["user"] == roleID{ delete(roles,y) }
        }
    }

    if err != nil{logs.Error("master/GetRolesForUser Error getting roles: "+err.Error()); return nil, err}    
    return roles, err
}

func GetGroupsForUser(groupID string)(data map[string]map[string]string, err error) {
    groups,err := ndb.GetUserGroups()
    if err != nil{logs.Error("master/GetGroupsForUser Error getting groups: "+err.Error()); return nil, err}    
    all,err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetGroupsForUser Error getting all: "+err.Error()); return nil, err}    
    for x := range all{
        for y := range groups{
            if all[x]["group"] == y && all[x]["user"] == groupID{ delete(groups,y) }
        }
    }

    return groups, err
}

func AddUsersTo(anode map[string]string) (err error) {
    values := strings.Split(anode["values"], ",")
        
    if anode["type"] == "role"{
        for _,x := range values{
            uuid := utils.Generate()
            err = ndb.InsertUserGroupRole(uuid,"user", anode["user"])
            if err != nil{logs.Error("master/AddUsersTo Error inserting user into InsertUserGroupRole: "+err.Error()); return err} 
            err = ndb.InsertUserGroupRole(uuid,"role", x)
            if err != nil{logs.Error("master/AddUsersTo Error inserting role into InsertUserGroupRole: "+err.Error()); return err}    
        }
    }
    if anode["type"] == "group"{
        for _,x := range values{
            uuid := utils.Generate()
            err = ndb.InsertUserGroupRole(uuid,"user", anode["user"])
            if err != nil{logs.Error("master/AddUsersTo Error inserting user into InsertUserGroupRole: "+err.Error()); return err}
            err = ndb.InsertUserGroupRole(uuid,"group", x)
            if err != nil{logs.Error("master/AddUsersTo Error inserting group into InsertUserGroupRole: "+err.Error()); return err}    
        }
    }
    return nil
}

func ChangePassword(anode map[string]string) (err error) {
    hashed,err := validation.HashPassword(anode["pass"])
    if err != nil{logs.Error("master/ChangePassword Error hashing new password: "+err.Error()); return err}

    err = ndb.UpdateUser(anode["user"], "pass", hashed)
    if err != nil{logs.Error("master/ChangePassword Error updating password: "+err.Error()); return err}

    return nil
}

func DeleteUserRole(anode map[string]string) (err error) {
    allRoles, err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/DeleteUserRole Error getting groups: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteUserRole Error getting userGroupRoles: "+err.Error()); return err}
    for x := range allRoles{
        if allRoles[x]["role"] == anode["role"]{
            for y := range userGroupRole{
                if userGroupRole[y]["role"] == x && userGroupRole[y]["user"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                }
            }
        }        
    }

    return nil
}

func DeleteRoleUser(anode map[string]string) (err error) {
    allUsers, err := ndb.GetLoginData()
    if err != nil{logs.Error("master/DeleteRoleUser Error getting users: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteRoleUser Error getting userGroupRoles: "+err.Error()); return err}

    for x := range allUsers{
        if allUsers[x]["user"] == anode["user"]{
            for y := range userGroupRole{
                if userGroupRole[y]["user"] == x && userGroupRole[y]["role"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                    if err != nil{logs.Error("master/DeleteRoleUser Error deleting userGroupRole user: "+err.Error()); return err}
                }
            }
        }        
    }
    return nil
}

func DeleteRoleGroup(anode map[string]string) (err error) {
    allGroups, err := ndb.GetUserGroups()
    if err != nil{logs.Error("master/DeleteRoleGroup Error getting users: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteRoleGroup Error getting userGroupRoles: "+err.Error()); return err}

    for x := range allGroups{
        if allGroups[x]["group"] == anode["group"]{
            for y := range userGroupRole{
                if userGroupRole[y]["group"] == x && userGroupRole[y]["role"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                    if err != nil{logs.Error("master/DeleteRoleGroup Error deleting userGroupRole group: "+err.Error()); return err}
                }
            }
        }        
    }
    return nil
}

func DeleteGroupUser(anode map[string]string) (err error) {
    allUsers, err := ndb.GetLoginData()
    if err != nil{logs.Error("master/DeleteGroupUser Error getting users: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteGroupUser Error getting userGroupRoles: "+err.Error()); return err}

    for x := range allUsers{
        if allUsers[x]["user"] == anode["user"]{
            for y := range userGroupRole{
                if userGroupRole[y]["user"] == x && userGroupRole[y]["group"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                    if err != nil{logs.Error("master/DeleteGroupUser Error deleting userGroupRole user group: "+err.Error()); return err}
                }
            }
        }        
    }
    return nil
}
//todo
func DeleteGroupRole(anode map[string]string) (err error) {
    allRoles, err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/DeleteGroupRole Error getting users: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteGroupRole Error getting userGroupRoles: "+err.Error()); return err}

    for x := range allRoles{
        if allRoles[x]["role"] == anode["role"]{
            for y := range userGroupRole{
                if userGroupRole[y]["role"] == x && userGroupRole[y]["group"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                    if err != nil{logs.Error("master/DeleteGroupRole Error deleting userGroupRole user role: "+err.Error()); return err}
                }
            }
        }        
    }
    return nil
}

func GetAllRoles() (data map[string]map[string]string, err error) {
    allRoles, err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/GetAllRoles Error getting roles data: "+err.Error()); return nil,err}
    allUsers, err := ndb.GetLoginData()
    if err != nil{logs.Error("master/GetAllRoles Error getting users data: "+err.Error()); return nil,err}
    allGroups, err := ndb.GetUserGroups()
    if err != nil{logs.Error("master/GetAllRoles Error getting groups data: "+err.Error()); return nil,err}
    
    allElements, err :=ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetAllRoles Error getting usergrouprole data: "+err.Error()); return nil,err}
    for x := range allRoles{
        var userNames []string 
        var groupNames []string 
        for y := range allElements{
            if x == allElements[y]["role"]{
                if allElements[y]["user"] != "" {
                    userNames = append(userNames, allUsers[allElements[y]["user"]]["user"])                                        
                }
                if allElements[y]["group"] != "" {
                    groupNames = append(groupNames, allGroups[allElements[y]["group"]]["group"])                                        
                }
            }
        }
        allRoles[x]["users"] = strings.Join(userNames, ",")
        allRoles[x]["groups"] = strings.Join(groupNames, ",")
    }

    return allRoles, err
}

//delete roles for role management
func DeleteRole(anode map[string]string) (err error) {
    //delete role
    err = ndb.DeleteUserRole(anode["id"])
    if err != nil{logs.Error("master/DeleteRole Error deleting roles: "+err.Error()); return err}
    
    //delete usergrouproles
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteRole Error getting userGroupRoles: "+err.Error()); return err}
    for y := range userGroupRole{
        if userGroupRole[y]["role"] == anode["id"] { 
            err = ndb.DeleteUserGroupRole(y)
            if err != nil{logs.Error("master/DeleteRole Error deleting a role: "+err.Error()); return err}
        }
    }

    return nil
}

//delete groups for group management
func DeleteUserGroup(anode map[string]string) (err error) {
    err = ndb.DeleteUserGroup(anode["id"])
    if err != nil{logs.Error("master/DeleteUserGroup Error getting groups: "+err.Error()); return err}
    userGroupRole, err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/DeleteUserRole Error getting userGroupRoles: "+err.Error()); return err}

    for x := range userGroupRole{
        if userGroupRole[x]["group"] == anode["group"]{
            for y := range userGroupRole{
                if userGroupRole[y]["group"] == x && userGroupRole[y]["user"] == anode["id"] { 
                    err = ndb.DeleteUserGroupRole(y)
                    if err != nil{logs.Error("master/DeleteUserGroup Error deleting a role: "+err.Error()); return err}
                }
            }
        }        
    }

    return nil
}

//edit roles for role management
func EditRole(anode map[string]string) (err error) {
    //update name
    err = ndb.UpdateUserRole(anode["id"], "role", anode["role"])
    if err != nil{logs.Error("master/EditRole Error updating role name: "+err.Error()); return err}
    //update permissions
    err = ndb.UpdateUserRole(anode["id"], "permissions", anode["permissions"])
    if err != nil{logs.Error("master/EditRole Error updating role permissions: "+err.Error()); return err}

    return err
}

//edit groups for group management
func EditUserGroup(anode map[string]string) (err error) {
    //update name
    err = ndb.UpdateUserGroup(anode["id"], "group", anode["group"])
    if err != nil{logs.Error("master/EditUserGroup Error updating group name: "+err.Error()); return err}

    return err
}

func GetAllUserGroups() (data map[string]map[string]string, err error) {
    allRoles, err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/GetAllUserGroups Error getting roles data: "+err.Error()); return nil,err}
    allGroups, err := ndb.GetUserGroups()   
    if err != nil{logs.Error("master/GetAllUserGroups Error getting groups data: "+err.Error()); return nil,err}
    allUsers, err := ndb.GetLoginData()
    if err != nil{logs.Error("master/GetAllUserGroups Error getting users data: "+err.Error()); return nil,err}
    allElements, err :=ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetAllUserGroups Error getting usergrouprole data: "+err.Error()); return nil,err}

    for y := range allGroups{
        var groupUsers []string 
        var groupRoles []string 
            for x := range allElements{
            if allElements[x]["group"] != "" && allElements[x]["group"] == y {
                if allElements[x]["user"] != ""{
                    groupUsers = append(groupUsers, allUsers[allElements[x]["user"]]["user"])  
                }
                if allElements[x]["role"] != ""{
                    groupRoles = append(groupRoles, allRoles[allElements[x]["role"]]["role"])  
                }
            }
        }
        allGroups[y]["roles"] = strings.Join(groupRoles, ",") 
        allGroups[y]["users"] = strings.Join(groupUsers, ",")
    }

    return allGroups, err
}

func GetRolesForGroups(roleID string)(data map[string]map[string]string, err error) {
    roles,err := ndb.GetUserRoles()
    if err != nil{logs.Error("master/GetRolesForGroups Error getting roles: "+err.Error()); return nil, err}    
    all,err := ndb.GetUserGroupRoles()
    if err != nil{logs.Error("master/GetRolesForGroups Error getting all: "+err.Error()); return nil, err}    
    for x := range all{
        for y := range roles{
            if all[x]["role"] == y && all[x]["group"] == roleID{ delete(roles,y) }
        }
    }

    if err != nil{logs.Error("master/GetRolesForGroups Error getting roles: "+err.Error()); return nil, err}    
    return roles, err
}

func AddRoleToGroup(anode map[string]string) (err error) {
    values := strings.Split(anode["values"], ",")
    for _,x := range values{
        uuid := utils.Generate()
        err = ndb.InsertUserGroupRole(uuid,"group", anode["group"])
        if err != nil{logs.Error("master/AddRoleToGroup Error inserting user into InsertUserGroupRole: "+err.Error()); return err} 
        err = ndb.InsertUserGroupRole(uuid,"role", x)
        if err != nil{logs.Error("master/AddRoleToGroup Error inserting role into InsertUserGroupRole: "+err.Error()); return err}    
    }
    
    return nil
}
