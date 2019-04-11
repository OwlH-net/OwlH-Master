package nodeclient

import (
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"encoding/json"
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
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("response Body:", string(body))
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
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	
    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("nodeClient/Zeek -- ERROR JSON unmarshal: "+err.Error())
        return nil,err
    }
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
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil,err
    }
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
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func GetAllFiles(ipData string, portData string, uuid string)(rData map[string]string, err error){
	url := "https://"+ipData+":"+portData+"/node/file"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR connection through http new Request: "+err.Error())
        return nil,err
    }
	defer resp.Body.Close()
	
    logs.Info("GetAllFiles response Status:", resp.Status)
    logs.Info("GetAllFiles response Headers:", resp.Header)
    responseData, err := ioutil.ReadAll(resp.Body)
    logs.Info("GetAllFiles response Body:", responseData)

    json.Unmarshal(responseData, &rData)
    logs.Info("rData Response: ")
    logs.Info(rData)
    rData["nodeUUID"] = uuid

	return rData,nil;
}

func SetRuleset(ipData string, portData string, data []byte)(err error){
	values := make(map[string][]byte)
	values["data"] = data
	url := "https://"+ipData+":"+portData+"/node/suricata/retrieve"
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
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/GetNodeFile ERROR reading file: "+err.Error())
		return nil, err
	}
    json.Unmarshal(responseData, &rData)
    rData["nodeUUID"] = loadFile["uuid"]
		
	return rData, nil
}

func PutSuricataBPF(ipnid string, portnid string, jsonnid string, jsonbpf string)(err error){
	// //create json with nid y bpf 
    values := make(map[string]string)
    values["nid"] = jsonnid
	values["bpf"] = jsonbpf
	valuesJSON,err := json.Marshal(values)
	url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("nodeclient/PutSuricataBPF ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return  nil
}

func GetSuricataBPF(ipnid string, portnid string)(bpf string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/GetNodeFile ERROR connection through http new Request: "+err.Error())
		return "", err
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/GetNodeFile ERROR reading file: "+err.Error())
		return "", err
	}
    json.Unmarshal(responseData, &bpf)
	return bpf, nil
}

func RunSuricata(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/suricata/RunSuricata"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunSuricata ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunSuricata ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func StopSuricata(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/suricata/StopSuricata"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopSuricata ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopSuricata ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func RunWazuh(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/wazuh/RunWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func StopWazuh(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/wazuh/StopWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunWazuh ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func RunZeek(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/zeek/RunZeek"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunZeek ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunZeek ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func StopZeek(ipnid string, portnid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/zeek/StopZeek"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopZeek ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopZeek ERROR reading request data: "+err.Error())
        return "",err
	}
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
    defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("GetAllServers ERROR reading requested data: "+err.Error())
        return nil,err
	}
    json.Unmarshal(responseData, &data)
    return data,nil
}

func GetServer(ipuuid string,portuuid string, serveruuid string)(data map[string]map[string]string, err error){
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/server/"+serveruuid
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("GetServer ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
    defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("GetServer ERROR reading requested data: "+err.Error())
        return nil,err
	}
    json.Unmarshal(responseData, &data)
    return data,nil
}

func RunStap(ipnid string, portnid string, uuid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/RunStap/"+uuid
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunStap ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunStap ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func StopStap(ipnid string, portnid string, uuid string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/StopStap/"+uuid
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopStap ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopStap ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func RunStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/RunStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/RunStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/RunStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func StopStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/StopStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/StopStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/StopStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func DeleteStapServer(ipnid string, portnid string, server string)(data string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/DeleteStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("nodeclient/DeleteStapServer ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("nodeclient/DeleteStapServer ERROR reading request data: "+err.Error())
        return "",err
	}
	return string(body),nil
}

func PingServerStap(ipnid string, portnid string, server string)(data map[string]string, err error){
	url := "https://"+ipnid+":"+portnid+"/node/stap/PingServerStap/"+server
	resp,err := utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("nodeclient/PingServerStap ERROR connection through http new Request: "+err.Error())
        return nil, err
    }
    defer resp.Body.Close()
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