package nodeclient

import (
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"encoding/json"
	"errors"
	"owlhmaster/utils"
	"bytes"
)

func init() {

}
 
func Echo() {
    logs.Info("NODE CLIENT -> ECHO")
}

func PingNode(ip string, port string) (err error) {
    logs.Info("NodeClient PingNode -> %s, %s", ip, port)
	// url := "https://"+ip+":"+port+"/node/node"
	url := "https://"+ip+":"+port+"/node/ping"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/PingNode ERROR connection through http new Request: "+err.Error())
        return err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("response Body:", string(body))
    defer resp.Body.Close()
    return nil
}

func Suricata(ip string, port string) (data map[string]bool, err error ) {
    logs.Info("NodeClient suricata status -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/suricata"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/Suricata ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
      defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    
    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("nodeClient/Suricata -- ERROR JSON unmarshal: "+err.Error())
        return nil,err
	}
    return data,nil
}

func Zeek(ip string, port string) (data map[string]bool, err error ) {
	logs.Info("NodeClient zeek status -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/zeek"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/Zeek ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	
    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("nodeClient/Zeek -- ERROR JSON unmarshal: "+err.Error())
        return nil,err
    }
	defer resp.Body.Close()
    return data,nil
}

func Wazuh(ip string, port string) (data map[string]bool, err error ) {
    logs.Info("NodeClient wazuh status -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/wazuh"
    resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/Wazuh ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	
    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
		return nil,err
    }
	defer resp.Body.Close()
    return data,nil
}

func Stap(ip string, port string, uuid string) (data map[string]bool, err error ) {
    logs.Info("NodeClient Stap status -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/stap/ping/"+uuid
    resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/Stap ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
	
	// data[status ] = resp.status
	// data[data] = restp.body 
	 
    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
		return nil,err
    }
	defer resp.Body.Close()
    return data,nil
}

func GetAllFiles(ipData string, portData string, uuid string)(rData map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/file"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
	
    logs.Info("GetAllFiles response Status:", resp.Status)
    logs.Info("GetAllFiles response Headers:", resp.Header)
    responseData, err := ioutil.ReadAll(resp.Body)
    logs.Info("GetAllFiles response Body:", responseData)
	
    json.Unmarshal(responseData, &rData)
    logs.Info("rData Response: ")
    logs.Info(rData)
    rData["nodeUUID"] = uuid
	
	defer resp.Body.Close()
	return rData,nil;
}

func SyncRulesetToNode(ipData string, portData string, data []byte)(err error){
	values := make(map[string][]byte)
	values["data"] = data
	url := "https://"+ipData+":"+portData+"/node/suricata/sync"
	valuesJSON,err := json.Marshal(values)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/SetRuleset ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func SetNodeFile(ipData string, portData string, loadFile map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/file"
	valuesJSON,err := json.Marshal(loadFile)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/SetNodeFile ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func GetNodeFile(ipData string, portData string, loadFile map[string]string)(rData map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/file/"+loadFile["file"]
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/GetNodeFile ERROR connection through http new Request: "+err.Error())
		return nil, err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/GetNodeFile ERROR reading file: "+err.Error())
		return nil, err
	}
    json.Unmarshal(responseData, &rData)
    rData["nodeUUID"] = loadFile["uuid"]
	
	defer resp.Body.Close()
	return rData, nil
}

func PutSuricataBPF(ipnid string, portnid string, anode map[string]string)(err error){
	valuesJSON,err := json.Marshal(anode)
	url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/PutSuricataBPF ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return  nil
}

// func GetSuricataBPF(ipnid string, portnid string)(bpf string, err error){
// 	url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
// 	resp,err := utils.NewRequestHTTP("GET", url, nil)
// 	if err != nil {
// 		logs.Error("nodeclient/GetNodeFile ERROR connection through http new Request: "+err.Error())
// 		return "", err
// 	}
// 	responseData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logs.Error("nodeclient/GetNodeFile ERROR reading file: "+err.Error())
// 		return "", err
// 	}
//     json.Unmarshal(responseData, &bpf)
// 	defer resp.Body.Close()
// 	return bpf, nil
// }

func RunSuricata(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/suricata/RunSuricata"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunSuricata ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunSuricata ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func StopSuricata(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/suricata/StopSuricata"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopSuricata ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopSuricata ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func RunWazuh(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/wazuh/RunWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func StopWazuh(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/wazuh/StopWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func RunZeek(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/zeek/RunZeek"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunZeek ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunZeek ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func StopZeek(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/zeek/StopZeek"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopZeek ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopZeek ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func AddServer(ipuuid string,portuuid string, data map[string]string )(err error){
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"
	valuesJSON,err := json.Marshal(data)
	if err != nil {
		logs.Error("nodeclient/AddServer Error Marshal new JSON data: "+err.Error())
        return err
	}
	resp,err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
		logs.Error("nodeclient/AddServer ERROR on the new HTTP request response: "+err.Error())
        return err
	}
	defer resp.Body.Close()
	return nil
}

func GetAllServers(ipuuid string,portuuid string)(data map[string]map[string]string, err error){
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("GetAllServers ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("GetAllServers ERROR reading requested data: "+err.Error())
        return nil,err
	}
    json.Unmarshal(responseData, &data)
	defer resp.Body.Close()
    return data,nil
}

func GetServer(ipuuid string,portuuid string, serveruuid string)(data map[string]map[string]string, err error){
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/server/"+serveruuid
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("GetServer ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("GetServer ERROR reading requested data: "+err.Error())
        return nil,err
	}
    json.Unmarshal(responseData, &data)
	defer resp.Body.Close()
    return data,nil
}

func RunStap(ipnid string, portnid string, uuid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/RunStap/"+uuid
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunStap ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunStap ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func StopStap(ipnid string, portnid string, uuid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/StopStap/"+uuid
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopStap ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopStap ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func RunStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/RunStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func StopStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/StopStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func DeleteStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/DeleteStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/DeleteStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/DeleteStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return string(body),nil
}

func PingServerStap(ipnid string, portnid string, server string)(data map[string]string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/PingServerStap/"+server
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/PingServerStap ERROR connection through http new Request: "+err.Error())
        return nil, err
    }
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/PingServerStap ERROR reading request data: "+err.Error())
        return nil,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("PingServerStap ERROR doing unmarshal JSON: "+err.Error())
        return nil,err
	}
	defer resp.Body.Close()
	return data,nil
}

func EditStapServer(ip string, port string, data map[string]string)(err error){
	url := "https://"+ip+":"+port+"/node/stap/EditStapServer"
	valuesJSON,err := json.Marshal(data)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
		logs.Error("EditStapServer ERROR on the new HTTP request response: "+err.Error())
        return err
	}
    defer resp.Body.Close()
    return nil
}

func PlayCollector(ipnid string, portnid string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/collector/play"
	_,err = utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/playCollector ERROR connection through http new Request: "+err.Error())
        return err
    }
	return nil
}
func StopCollector(ipnid string, portnid string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/collector/stop"
	_,err = utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopCollector ERROR connection through http new Request: "+err.Error())
        return err
    }
	return nil
}
func ShowCollector(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/collector/show"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/ShowCollector ERROR connection through http new Request: "+err.Error())
        return "",err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/ShowCollector ERROR reading request data: "+err.Error())
        return "",err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("ShowCollector ERROR doing unmarshal JSON: "+err.Error())
        return "",err
	}
	defer resp.Body.Close()
	return data,nil
}

func DeployZeek(ipnid string, portnid string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/zeek/DeployZeek"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/DeployZeek ERROR connection through http new Request: "+err.Error())
        return err
    }

	defer resp.Body.Close()
	return nil
}

func ShowPorts(ipnid string, portnid string)(data map[string]map[string]string ,err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/ShowPorts ERROR connection through http new Request: "+err.Error())
        return data,err
    }

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/ShowPorts ERROR reading request data: "+err.Error())
        return data,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("ShowPorts ERROR doing unmarshal JSON: "+err.Error())
        return data,err
	}
	defer resp.Body.Close()
	return data,nil
}

func PingPluginsNode(ipnid string, portnid string)(data map[string]map[string]string ,err error){
	url := "https://"+ipnid+":"+portnid+"/node/ping/PingPluginsNode/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/PingPluginsNode ERROR connection through http new Request: "+err.Error())
        return data,err
    }

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/PingPluginsNode ERROR reading request data: "+err.Error())
        return data,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("PingPluginsNode ERROR doing unmarshal JSON: "+err.Error())
        return data,err
	}
	defer resp.Body.Close()
	return data,nil
}

func ChangeMode(ipnid string, portnid string, mode string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/mode"
	values := make(map[string]string)
    values["mode"] = mode
	valuesJSON,err := json.Marshal(values)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/ChangeMode ERROR connection through http new Request: "+err.Error())
        return err
    }
	defer resp.Body.Close()
	return nil
}

func PingPorts(ipnid string, portnid string)(data map[string]map[string]string ,err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/PingPorts/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {logs.Error("nodeclient/PingPorts ERROR connection through http new Request: "+err.Error());return data,err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {logs.Error("nodeclient/PingPorts ERROR reading request data: "+err.Error()); return data,err}
	err = json.Unmarshal(body, &data)
	if err != nil {logs.Error("PingPorts ERROR doing unmarshal JSON: "+err.Error()); return data,err}
	
	defer resp.Body.Close()
	return data,nil
}

func ChangeStatus(ipnid string, portnid string, status string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/status"
	values := make(map[string]string)
    values["status"] = status
	valuesJSON,err := json.Marshal(values)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/ChangeStatus ERROR connection through http new Request: "+err.Error())
        return err
    }
	defer resp.Body.Close()
	return nil
}

func DeletePorts(ipnid string, portnid string, ports map[string]string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/delete"
 
	valuesJSON,err := json.Marshal(ports)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/DeletePorts ERROR connection through http new Request: "+err.Error())
        return err
    }
	defer resp.Body.Close()
	return nil
}

func DeleteAllPorts(ipnid string, portnid string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/ports/deleteAll"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/DeleteAllPorts ERROR connection through http new Request: "+err.Error())
        return err
    }
	defer resp.Body.Close()
	return nil
}

func PingAnalyzer(ipnid string, portnid string)(data map[string]string ,err error){
	url := "https://"+ipnid+":"+portnid+"/node/analyzer/pingAnalyzer/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/PingAnalyzer ERROR connection through http new Request: "+err.Error())
        return data,err
    }

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/PingAnalyzer ERROR reading request data: "+err.Error())
        return data,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("PingAnalyzer ERROR doing unmarshal JSON: "+err.Error())
        return data,err
	}
	
	logs.Warn("Analyzer data from nodeclient")
	logs.Warn(data)

	defer resp.Body.Close()
	return data,nil
}

func ChangeAnalyzerStatus(ipnid string, portnid string, anode map[string]string)(err error){
	url := "https://"+ipnid+":"+portnid+"/node/analyzer/changeAnalyzerStatus/"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/ChangeAnalyzerStatus ERROR connection through http new Request: "+err.Error())
        return err
    }
	defer resp.Body.Close()
	return nil
}

func DeployNode(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/deploy"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/Deploy ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func CheckDeploy(ipData string, portData string)(data map[string]string){
	url := "https://"+ipData+":"+portData+"/node/deploy"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/CheckDeploy ERROR connection through http new Request: "+err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/PingAnalyzer ERROR reading request data: "+err.Error())
        return nil
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("PingAnalyzer ERROR doing unmarshal JSON: "+err.Error())
        return nil
	}
	
	defer resp.Body.Close()
	return data
}

func ChangeDataflowValues(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/changeDataflowValues"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/ChangeDataflowValues ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func UpdateNetworkInterface(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/net"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/UpdateNetworkInterface ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func LoadDataflowValues(ipData string, portData string)(data map[string]map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/loadDataflowValues"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/LoadDataflowValues ERROR connection through http new Request: "+err.Error())
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/LoadDataflowValues ERROR reading request data: "+err.Error())
        return nil,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("LoadDataflowValues ERROR doing unmarshal JSON: "+err.Error())
        return nil,err
	}
	
	defer resp.Body.Close()
	return data,nil
}

func LoadNetworkValues(ipData string, portData string)(data map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/net/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/LoadNetworkValues ERROR connection through http new Request: "+err.Error())
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {logs.Error("nodeclient/LoadNetworkValues ERROR reading request data: "+err.Error());return nil,err}

	err = json.Unmarshal(body, &data)
    if err != nil {logs.Error("LoadNetworkValues ERROR doing unmarshal JSON: "+err.Error());return nil,err}

	defer resp.Body.Close()
	return data,nil
}

func LoadNetworkValuesSelected(ipData string, portData string)(data map[string]map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/net/values"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/LoadNetworkValuesSelected ERROR connection through http new Request: "+err.Error())
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {logs.Error("nodeclient/LoadNetworkValuesSelected ERROR reading request data: "+err.Error());return nil,err}

	err = json.Unmarshal(body, &data)
    if err != nil {logs.Error("LoadNetworkValuesSelected ERROR doing unmarshal JSON: "+err.Error());return nil,err}

	defer resp.Body.Close()
	return data,nil
}

func GetServiceStatus(ip string, port string) (err error) {
    logs.Info("NodeClient GetServiceStatus -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/ping/services"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("nodeClient/GetServiceStatus ERROR connection through http new Request: "+err.Error())
        return err
    }
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	
	if raw["ack"] == "false"{
		return errors.New("Service don't exists")
	}else if raw["ack"] == "true"{
		return nil
	}

	return nil
}

func DeployService(ip string, port string) (err error) {
    logs.Info("NodeClient DeployService -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/ping/deployservice"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("nodeClient/DeployService ERROR connection through http new Request: "+err.Error())
        return err
    }
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	
	if raw["ack"] == "false"{
		return errors.New("Service don't exists")
	}else if raw["ack"] == "true"{
		return nil
	}

	return nil
}

func SaveSocketToNetwork(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/saveSocketToNetwork"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil { logs.Error("nodeclient/SaveSocketToNetwork ERROR connection through http new Request: "+err.Error()); return err}
	
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	
	if raw["ack"] == "false"{
		return errors.New("Name used. Use other name.")
	}else if raw["ack"] == "true"{
		return nil
	}
	

	return nil
}

func SaveNewLocal(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/saveNewLocal"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/SaveNewLocal ERROR connection through http new Request: "+err.Error()); return err}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	
	if raw["ack"] == "false"{
		return errors.New("Name used. Use other name.")
	}else if raw["ack"] == "true"{
		return nil
	}

	return nil
}

func SaveVxLAN(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/saveVxLAN"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil { logs.Error("nodeclient/SaveVxLAN ERROR connection through http new Request: "+err.Error()); return err}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	
	if raw["ack"] == "false"{
		return errors.New("Name used. Use other name.")
	}else if raw["ack"] == "true"{
		return nil
	}

	return nil
}

func SocketToNetworkList(ipData string, portData string)(data map[string]map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/socketToNetworkList"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/SocketToNetworkList ERROR connection through http new Request: "+err.Error())
		return nil,err
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Error("nodeclient/SocketToNetworkList ERROR reading request data: "+err.Error())
        return nil,err
	}
	err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("SocketToNetworkList ERROR doing unmarshal JSON: "+err.Error())
        return nil,err
	}
	
	defer resp.Body.Close()
	return data,nil
}

func SaveSocketToNetworkSelected(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/saveSocketToNetworkSelected"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/SaveSocketToNetworkSelected ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func DeleteDataFlowValueSelected(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/dataflow/deleteDataFlowValueSelected"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/DeleteDataFlowValueSelected ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func GetNodeMonitor(ipData string, portData string)(data map[string]interface{}, err error){
	url := "https://"+ipData+":"+portData+"/node/monitor/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil { logs.Error("nodeclient/GetNodeMonitor ERROR connection through http new Request: "+err.Error()); return data,err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/GetNodeMonitor ERROR reading request data: "+err.Error()); return data,err}
	
	err = json.Unmarshal(body, &data)
    if err != nil { logs.Error("nodeclient/GetNodeMonitor ERROR doing unmarshal JSON: "+err.Error()); return data,err}

	defer resp.Body.Close()
	return data,nil
}

func AddPluginService(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/addService"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/AddPluginService ERROR connection through http new Request: "+err.Error());return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {logs.Error("nodeclient/AddPluginService ERROR reading request data: "+err.Error()); return err}

	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/AddPluginService ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */



	defer resp.Body.Close()

	return nil
}

func GetSuricataServices(ipData string, portData string)(data map[string]map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/suricata/get"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {logs.Error("nodeclient/GetSuricataServices ERROR connection through http new Request: "+err.Error()); return nil,err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {logs.Error("nodeclient/GetSuricataServices ERROR reading request data: "+err.Error()); return nil,err}

	err = json.Unmarshal(body, &data)
    if err != nil {logs.Error("GetSuricataServices ERROR doing unmarshal JSON: "+err.Error()); return nil,err}
	
	defer resp.Body.Close()
	return data,nil
}

func GetMainconfData(ipData string, portData string)(data map[string]map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/ping/mainconf"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil { logs.Error("nodeclient/GetMainconfData ERROR connection through http new Request: "+err.Error()); return data,err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/GetMainconfData ERROR reading request data: "+err.Error()); return data,err}
	
	err = json.Unmarshal(body, &data)
    if err != nil { logs.Error("nodeclient/GetMainconfData ERROR doing unmarshal JSON: "+err.Error()); return data,err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */

	defer resp.Body.Close()
	return data,nil
}

func ChangeServiceStatus(ipData string, portData string, anode map[string]string)(err error){
	var data interface{}
	url := "https://"+ipData+":"+portData+"/node/plugin/ChangeServiceStatus"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: "+err.Error()); return err}
	
	err = json.Unmarshal(body, &data)
    if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */

	defer resp.Body.Close()
	return nil
}

func ChangeMainServiceStatus(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/ChangeMainServiceStatus"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/ChangeMainServiceStatus ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/ChangeMainServiceStatus ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/ChangeMainServiceStatus ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */
	defer resp.Body.Close()
	return nil
}

func DeleteService(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/deleteService"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/DeleteService ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/DeleteService ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/DeleteService ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */
	defer resp.Body.Close()
	return nil
}

func SaveSuricataInterface(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/SaveSuricataInterface"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/SaveSuricataInterface ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/SaveSuricataInterface ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/SaveSuricataInterface ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */

	defer resp.Body.Close()
	return nil
}

func DeployStapService(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/deployStapService"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/DeployStapService ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */

	defer resp.Body.Close()
	return nil
}

func StopStapService(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/stopStapService"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/StopStapService ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */

	defer resp.Body.Close()
	return nil
}

func ModifyStapValues(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/plugin/modifyStapValues"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/ModifyStapValues ERROR connection through http new Request: "+err.Error()); return err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: "+err.Error()); return err}
	
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil { logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: "+err.Error()); return err}
/*	if data["ack"] == "false"{
		defer resp.Body.Close()
		return errors.New(data["error"])
	} */


	defer resp.Body.Close()
	return nil
}

func PingWazuhFiles(ipData string, portData string)(data map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/wazuh/pingWazuhFiles"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil { logs.Error("nodeclient/PingWazuhFiles ERROR connection through http new Request: "+err.Error()); return data,err}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { logs.Error("nodeclient/PingWazuhFiles ERROR reading request data: "+err.Error()); return data,err}
	
	err = json.Unmarshal(body, &data)
    if err != nil { logs.Error("nodeclient/PingWazuhFiles ERROR doing unmarshal JSON: "+err.Error()); return data,err}
	
	// if data["ack"] == "false"{
	// 	defer resp.Body.Close()
	// 	return errors.New(data["error"])
	// } 

	defer resp.Body.Close()
	return data,nil
}

func DeleteWazuhFile(ipData string, portData string, anode map[string]string)(err error){
	url := "https://"+ipData+":"+portData+"/node/wazuh/deleteWazuhFile"
	valuesJSON,err := json.Marshal(anode)
	resp,err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
	if err != nil {logs.Error("nodeclient/DeleteWazuhFile ERROR connection through http new Request: "+err.Error()); return err}
	
	defer resp.Body.Close()
	return nil
}