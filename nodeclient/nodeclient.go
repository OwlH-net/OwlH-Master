package nodeclient

import (
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"encoding/json"
	"owlhmaster/utils"
	"bytes"
	"net/http"
)

func init() {

}
 
func Echo() {
    logs.Info("NODE CLIENT -> ECHO")
}

func SetRuleset(ipData string, portData string, data []byte)(err error){
	values := make(map[string][]byte)
	values["data"] = data
	url := "https://"+ipData+":"+portData+"/node/suricata/retrieve"
	valuesJSON,err := json.Marshal(values)
	resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
	if err != nil {
		logs.Error("node/SetRuleset ERROR connection through http new Request: "+err.Error())
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
		logs.Error("node/SetRuleset ERROR connection through http new Request: "+err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}
func GetNodeFile(ipData string, portData string, loadFile string)(resp *http.Response, err error){
	url := "https://"+ipData+":"+portData+"/node/file/"+loadFile
	resp,err = utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("node/GetNodeFile ERROR connection through http new Request: "+err.Error())
		return resp,err;
	}
    defer resp.Body.Close()
	return resp,nil;
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

func GetAllFiles(ipData string, portData string)(resp *http.Response, err error){
	url := "https://"+ipData+":"+portData+"/node/file"
	resp,err = utils.NewRequestHTTP("GET", url, nil)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR connection through http new Request: "+err.Error())
        return resp,err
    }
    defer resp.Body.Close()
	return resp,nil;
}