package nodeclient

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
    "owlhmaster/database"
    "owlhmaster/utils"
    "time"
)

//Zeek data struct
type ZeekData struct {
    Path        bool              `json:"path"`
    Rol         string            `json:"role"`
    Bin         bool              `json:"bin"`
    Action      string            `json:"action"`
    Running     []ZeekNodeStatus  `json:"running"`
    Mode        string            `json:"mode"`
    Managed     bool              `json:"managed"`
    Manager     bool              `json:"manager"`
    ManagerIP   string            `json:"managerip"`
    ManagerUUID string            `json:"manageruuid"`
    ManagerName string            `json:"managername"`
    Nodes       []ZeekNode        `json:"nodes"`
    Extra       map[string]string `json:"extra"`
}
type ZeekKeys struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
type ZeekNode struct {
    Name       string     `json:"name"`
    Host       string     `json:"host"`
    Status     string     `json:"status"`
    Type       string     `json:"type"`
    NInterface string     `json:"interface"`
    Pid        string     `json:"pid"`
    Started    string     `json:"started"`
    Extra      []ZeekKeys `json:"extra"`
}
type ZeekNodeStatus struct {
    Status string `json:"status"`
    Nodes  int    `json:"nodes"`
}

func PingNode(ip string, port string) (nodeResp map[string]string, err error) {
    url := "https://" + ip + ":" + port + "/node/ping"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        //add incident
        uuid := utils.Generate()
        var controlError error
        currentTime := time.Now()
        timeFormated := currentTime.Format("2006-01-02T15:04:05")
        controlError = ndb.PutIncident(uuid, "date", timeFormated)
        controlError = ndb.PutIncident(uuid, "desc", "Master description")
        controlError = ndb.PutIncident(uuid, "status", "new")  // new, open, closed, delayed
        controlError = ndb.PutIncident(uuid, "level", "info")  // warning, info or danger
        controlError = ndb.PutIncident(uuid, "NodeIp", ip)     // warning, info or danger
        controlError = ndb.PutIncident(uuid, "NodePort", port) // warning, info or danger
        if controlError != nil {
            logs.Error("PingNode error creating incident: " + controlError.Error())
        }

        logs.Error("nodeClient/PingNode ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/PingNode ERROR reading request data: " + err.Error())
        return nil, err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/PingNode ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    // if returnValues["nodeToken"] == "none"{
    //     logs.Error("nodeclient/PingNode ERROR from node: "+returnValues["error"])
    //     values := make(map[string]string)
    //     values["nodeToken"] = "none"
    //     values["error"] = "nodeclient/PingNode Node token ERROR: "+values["error"]
    //     values["ack"] = "false"
    //     return values,nil
    // }

    return returnValues, nil
}

func UpdateNodeData(ipData string, portData string, loadFile map[string]map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/ping/updateNode"
    valuesJSON, err := json.Marshal(loadFile)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/UpdateNodeData ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func Suricata(ip string, port string) (data map[string]bool, err error) {
    logs.Info("NodeClient suricata status -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/suricata"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeClient/Suricata ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeClient/Suricata -- ERROR JSON unmarshal: " + err.Error())
        return nil, err
    }
    return data, nil
}

func Zeek(ip string, port string) (data ZeekData, err error) {
    logs.Info("NodeClient zeek status -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/zeek"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeClient/Zeek ERROR connection through http new Request: " + err.Error())
        return data, err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeClient/Zeek -- ERROR JSON unmarshal: " + err.Error())
        return data, err
    }
    defer resp.Body.Close()
    return data, nil
}

func Wazuh(ip string, port string) (data map[string]bool, err error) {
    logs.Info("NodeClient wazuh status -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/wazuh"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeClient/Wazuh ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return data, nil
}

func ZeekCommand(ip, port, cmd string) (data map[string]string, err error) {
    logs.Info("NodeClient ZeekCommand -> %s, %s, %s", ip, port, cmd)
    url := "https://" + ip + ":" + port + "/node/zeek/" + cmd
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeClient/ZeekCommand ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, _ := ioutil.ReadAll(resp.Body)

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return data, nil
}

func Stap(ip string, port string, uuid string) (data map[string]bool, err error) {
    logs.Info("NodeClient Stap status -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/stap/ping/" + uuid
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeClient/Stap ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)

    // data[status ] = resp.status
    // data[data] = restp.body

    //Convert []byte to map[string]bool
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return data, nil
}

func GetAllFiles(ipData string, portData string, uuid string) (rData map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/file"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("node/GetAllFiles ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    responseData, err := ioutil.ReadAll(resp.Body)
    json.Unmarshal(responseData, &rData)
    rData["nodeUUID"] = uuid

    defer resp.Body.Close()
    return rData, nil
}

// func SyncRulesetToNode2(ipData string, portData string, token string, data []byte, service string)(err error){
//     utils.TokenMasterValidated = token
//     if data == nil || len(data) <= 0 { return errors.New("SyncRulesetToNode error - Can't synchronize an empty ruleset")}

//     values := make(map[string][]byte)
//     values[service] = data
//     url := "https://"+ipData+":"+portData+"/node/suricata/sync"
//     valuesJSON,err := json.Marshal(values)
//     resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
//     if err != nil {logs.Error("nodeclient/SetRuleset ERROR connection through http new Request: "+err.Error()); return err}
//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil { logs.Error("nodeclient/SyncRulesetToNode ERROR reading request data: "+err.Error()); return err}
//     nodeResponse := make(map[string]string)
//     err = json.Unmarshal(body, &nodeResponse)
//     if err != nil { logs.Error("nodeclient/GetNodeAutentication ERROR doing unmarshal JSON: "+err.Error()); return err}
//     if nodeResponse["ack"] == "false" {
//         return errors.New(nodeResponse["error"])
//     }

//     return nil
// }

func SyncGroupRulesetToNode(ipData string, portData string, data map[string][]byte) (err error) {
    if data == nil || len(data) <= 0 {
        return errors.New("SyncGroupRulesetToNode error - Can't synchronize an empty ruleset")
    }

    url := "https://" + ipData + ":" + portData + "/node/group/groupSync"

    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncGroupRulesetToNode ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncGroupRulesetToNode ERROR reading request data: " + err.Error())
        return err
    }
    nodeResponse := make(map[string]string)
    err = json.Unmarshal(body, &nodeResponse)
    if err != nil {
        logs.Error("nodeclient/SyncGroupRulesetToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if nodeResponse["ack"] == "false" {
        return errors.New(nodeResponse["error"])
    }

    return nil
}

func SyncRulesetToNode(ipData string, portData string, data []byte, service string, rulesetName string) (err error) {
    if data == nil || len(data) <= 0 {
        return errors.New("SyncRulesetToNode error - Can't synchronize an empty ruleset")
    }

    values := make(map[string][]byte)
    values["data"] = data
    values["service"] = []byte(service)
    values["rulesetName"] = []byte(rulesetName)
    url := "https://" + ipData + ":" + portData + "/node/suricata/sync"
    valuesJSON, err := json.Marshal(values)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncRulesetToNode ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncRulesetToNode ERROR reading request data: " + err.Error())
        return err
    }
    nodeResponse := make(map[string]string)
    err = json.Unmarshal(body, &nodeResponse)
    if err != nil {
        logs.Error("nodeclient/SyncRulesetToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if nodeResponse["ack"] == "false" {
        return errors.New(nodeResponse["error"])
    }

    return nil
}

func SetNodeFile(ipData string, portData string, loadFile map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/file"
    valuesJSON, err := json.Marshal(loadFile)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SetNodeFile ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func GetNodeFile(ipData string, portData string, loadFile map[string]string) (rData map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/file/" + loadFile["file"]
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetNodeFile ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetNodeFile ERROR reading file: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    json.Unmarshal(responseData, &rData)
    if rData["ack"] == "false" {
        return nil, errors.New("GetNodeFile Error getting node file: " + rData["error"])
    }
    rData["nodeUUID"] = loadFile["uuid"]

    return rData, nil
}

func PutSuricataBPF(ipnid string, portnid string, anode map[string]string) (err error) {
    valuesJSON, err := json.Marshal(anode)
    url := "https://" + ipnid + ":" + portnid + "/node/suricata/bpf"
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/PutSuricataBPF ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

// func GetSuricataBPF(ipnid string, portnid string)(bpf string, err error){
//     url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
//     resp,err := utils.NewRequestHTTP("GET", url, nil)
//     if err != nil {
//         logs.Error("nodeclient/GetNodeFile ERROR connection through http new Request: "+err.Error())
//         return "", err
//     }
//     responseData, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         logs.Error("nodeclient/GetNodeFile ERROR reading file: "+err.Error())
//         return "", err
//     }
//     json.Unmarshal(responseData, &bpf)
//     defer resp.Body.Close()
//     return bpf, nil
// }

func RunSuricata(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/suricata/RunSuricata"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunSuricata ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunSuricata ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func StopSuricata(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/suricata/StopSuricata"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/StopSuricata ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/StopSuricata ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func RunWazuh(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/wazuh/RunWazuh"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunWazuh ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func StopWazuh(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/wazuh/StopWazuh"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunWazuh ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunWazuh ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func RunZeek(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/RunZeek"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunZeek ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunZeek ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func StopZeek(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/StopZeek"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/StopZeek ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/StopZeek ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func AddServer(ipuuid string, portuuid string, data map[string]string) (err error) {
    url := "https://" + ipuuid + ":" + portuuid + "/node/stap/"
    valuesJSON, err := json.Marshal(data)
    if err != nil {
        logs.Error("nodeclient/AddServer Error Marshal new JSON data: " + err.Error())
        return err
    }
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddServer ERROR on the new HTTP request response: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func GetAllServers(ipuuid string, portuuid string) (data map[string]map[string]string, err error) {
    url := "https://" + ipuuid + ":" + portuuid + "/node/stap/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("GetAllServers ERROR on the new HTTP request response: " + err.Error())
        return nil, err
    }
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("GetAllServers ERROR reading requested data: " + err.Error())
        return nil, err
    }
    json.Unmarshal(responseData, &data)
    defer resp.Body.Close()
    return data, nil
}

func GetServer(ipuuid string, portuuid string, serveruuid string) (data map[string]map[string]string, err error) {
    url := "https://" + ipuuid + ":" + portuuid + "/node/stap/server/" + serveruuid
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("GetServer ERROR on the new HTTP request response: " + err.Error())
        return nil, err
    }
    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("GetServer ERROR reading requested data: " + err.Error())
        return nil, err
    }
    json.Unmarshal(responseData, &data)
    defer resp.Body.Close()
    return data, nil
}

func RunStap(ipnid string, portnid string, uuid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/RunStap/" + uuid
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunStap ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunStap ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func StopStap(ipnid string, portnid string, uuid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/StopStap/" + uuid
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/StopStap ERROR connection through http new Request: " + err.Error())
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/StopStap ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func RunStapServer(ipnid string, portnid string, server string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/RunStapServer/" + server
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/RunStapServer ERROR connection through http new Request: " + err.Error())
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/RunStapServer ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func StopStapServer(ipnid string, portnid string, server string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/StopStapServer/" + server
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/StopStapServer ERROR connection through http new Request: " + err.Error())
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/StopStapServer ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func DeleteStapServer(ipnid string, portnid string, server string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/DeleteStapServer/" + server
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/DeleteStapServer ERROR connection through http new Request: " + err.Error())
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/DeleteStapServer ERROR reading request data: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return string(body), nil
}

func PingServerStap(ipnid string, portnid string, server string) (data map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/stap/PingServerStap/" + server
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingServerStap ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/PingServerStap ERROR reading request data: " + err.Error())
        return nil, err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("PingServerStap ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()
    return data, nil
}

func EditStapServer(ip string, port string, data map[string]string) (err error) {
    url := "https://" + ip + ":" + port + "/node/stap/EditStapServer"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("EditStapServer ERROR on the new HTTP request response: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func PlayCollector(ipnid string, portnid string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/collector/play"
    _, err = utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/playCollector ERROR connection through http new Request: " + err.Error())
        return err
    }
    return nil
}
func StopCollector(ipnid string, portnid string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/collector/stop"
    _, err = utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/StopCollector ERROR connection through http new Request: " + err.Error())
        return err
    }
    return nil
}
func ShowCollector(ipnid string, portnid string) (data string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/collector/show"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/ShowCollector ERROR connection through http new Request: " + err.Error())
        return "", err
    }
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        logs.Error("nodeclient/ShowCollector ERROR reading request data: " + err.Error())
        return "", err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("ShowCollector ERROR doing unmarshal JSON: " + err.Error())
        return "", err
    }
    defer resp.Body.Close()
    return data, nil
}

func DeployZeek(ipnid string, portnid string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/DeployZeek"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/DeployZeek ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func ShowPorts(ipnid string, portnid string) (data map[string]map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/ShowPorts ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        logs.Error("nodeclient/ShowPorts ERROR reading request data: " + err.Error())
        return data, err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("ShowPorts ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }
    defer resp.Body.Close()
    return data, nil
}

func PingPluginsNode(ipnid string, portnid string) (data map[string]map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ping/PingPluginsNode/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingPluginsNode ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/PingPluginsNode ERROR reading request data: " + err.Error())
        return data, err
    }

    // errorMap := make(map[string]string)
    // err = json.Unmarshal(body, &errorMap)
    // if errorMap["ack"] == "false"{ logs.Error("nodeclient/PingPluginsNode ERROR: "+err.Error()); return nil,errors.New(errorMap["error"])}

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/PingPluginsNode ERROR doing unmarshal JSON: " + err.Error())
    }

    if data["hasError"]["ack"] == "false" {
        logs.Error("nodeclient/PingPluginsNode ERROR: " + data["hasError"]["error"])
        return nil, errors.New(data["hasError"]["error"])
    }

    return data, nil
}

func ChangeMode(ipnid string, portnid string, mode string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/mode"
    values := make(map[string]string)
    values["mode"] = mode
    valuesJSON, err := json.Marshal(values)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeMode ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func PingPorts(ipnid string, portnid string) (data map[string]map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/PingPorts/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingPorts ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/PingPorts ERROR reading request data: " + err.Error())
        return data, err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("PingPorts ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }

    defer resp.Body.Close()
    return data, nil
}

func ChangeStatus(ipnid string, portnid string, status string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/status"
    values := make(map[string]string)
    values["status"] = status
    valuesJSON, err := json.Marshal(values)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeStatus ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func DeletePorts(ipnid string, portnid string, ports map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/delete"

    valuesJSON, err := json.Marshal(ports)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeletePorts ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func DeleteAllPorts(ipnid string, portnid string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ports/deleteAll"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeclient/DeleteAllPorts ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func PingAnalyzer(ipnid string, portnid string) (data map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/analyzer/pingAnalyzer/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingAnalyzer ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/PingAnalyzer ERROR reading request data: " + err.Error())
        return data, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("PingAnalyzer ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }

    defer resp.Body.Close()
    return data, nil
}

func ChangeAnalyzerStatus(ipnid string, portnid string, anode map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/analyzer/changeAnalyzerStatus/"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeAnalyzerStatus ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func DeployNode(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/deploy"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/Deploy ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func CheckDeploy(ipData string, portData string) (data map[string]string) {
    url := "https://" + ipData + ":" + portData + "/node/deploy"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/CheckDeploy ERROR connection through http new Request: " + err.Error())
        return nil
    }
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        logs.Error("nodeclient/PingAnalyzer ERROR reading request data: " + err.Error())
        return nil
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("PingAnalyzer ERROR doing unmarshal JSON: " + err.Error())
        return nil
    }

    defer resp.Body.Close()
    return data
}

func ChangeDataflowValues(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/changeDataflowValues"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeDataflowValues ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func UpdateNetworkInterface(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/net"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/UpdateNetworkInterface ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func LoadDataflowValues(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/loadDataflowValues"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/LoadDataflowValues ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        logs.Error("nodeclient/LoadDataflowValues ERROR reading request data: " + err.Error())
        return nil, err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("LoadDataflowValues ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func LoadNetworkValues(ipData string, portData string) (data map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/net/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/LoadNetworkValues ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/LoadNetworkValues ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("LoadNetworkValues ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func LoadNetworkValuesSelected(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/net/values"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/LoadNetworkValuesSelected ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/LoadNetworkValuesSelected ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("LoadNetworkValuesSelected ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func GetServiceStatus(ip string, port string) (err error) {
    logs.Info("NodeClient GetServiceStatus -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/ping/services"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeClient/GetServiceStatus ERROR connection through http new Request: " + err.Error())
        return err
    }
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    var raw map[string]interface{}
    json.Unmarshal(body, &raw)

    if raw["ack"] == "false" {
        return errors.New("Service don't exists")
    } else if raw["ack"] == "true" {
        return nil
    }

    return nil
}

func DeployService(ip string, port string) (err error) {
    logs.Info("NodeClient DeployService -> %s, %s", ip, port)
    url := "https://" + ip + ":" + port + "/node/ping/deployservice"
    resp, err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
        logs.Error("nodeClient/DeployService ERROR connection through http new Request: " + err.Error())
        return err
    }
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    var raw map[string]interface{}
    json.Unmarshal(body, &raw)

    if raw["ack"] == "false" {
        return errors.New("Service don't exists")
    } else if raw["ack"] == "true" {
        return nil
    }

    return nil
}

func SaveSocketToNetwork(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/saveSocketToNetwork"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveSocketToNetwork ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    var raw map[string]interface{}
    json.Unmarshal(body, &raw)

    if raw["ack"] == "false" {
        return errors.New("Name used. Use other name.")
    } else if raw["ack"] == "true" {
        return nil
    }

    return nil
}

func SaveNewLocal(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/saveNewLocal"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveNewLocal ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    var raw map[string]interface{}
    json.Unmarshal(body, &raw)

    if raw["ack"] == "false" {
        return errors.New("Name used. Use other name.")
    } else if raw["ack"] == "true" {
        return nil
    }

    return nil
}

func SaveVxLAN(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/saveVxLAN"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveVxLAN ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    var raw map[string]interface{}
    json.Unmarshal(body, &raw)

    if raw["ack"] == "false" {
        return errors.New("Name used. Use other name.")
    } else if raw["ack"] == "true" {
        return nil
    }

    return nil
}

func SocketToNetworkList(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/socketToNetworkList"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/SocketToNetworkList ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        logs.Error("nodeclient/SocketToNetworkList ERROR reading request data: " + err.Error())
        return nil, err
    }
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("SocketToNetworkList ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func SaveSocketToNetworkSelected(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/saveSocketToNetworkSelected"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveSocketToNetworkSelected ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func DeleteDataFlowValueSelected(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/dataflow/deleteDataFlowValueSelected"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeleteDataFlowValueSelected ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()
    return nil
}

func GetNodeMonitor(ipData string, portData string) (data map[string]interface{}, err error) {
    url := "https://" + ipData + ":" + portData + "/node/monitor/"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetNodeMonitor ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetNodeMonitor ERROR reading request data: " + err.Error())
        return data, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/GetNodeMonitor ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }

    defer resp.Body.Close()
    return data, nil
}

func AddPluginService(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/addService"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddPluginService ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/AddPluginService ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/AddPluginService ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func GetSuricataServices(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/suricata/get"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetSuricataServices ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetSuricataServices ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("GetSuricataServices ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func GetMainconfData(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/ping/mainconf"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetMainconfData ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetMainconfData ERROR reading request data: " + err.Error())
        return data, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/GetMainconfData ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }
    /*    if data["ack"] == "false"{
          defer resp.Body.Close()
          return errors.New(data["error"])
      } */

    defer resp.Body.Close()
    return data, nil
}

func ChangeServiceStatus(ipData string, portData string, anode map[string]string) (err error) {
    data := make(map[string]string)
    url := "https://" + ipData + ":" + portData + "/node/plugin/ChangeServiceStatus"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: " + err.Error())
        return err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        defer resp.Body.Close()
        return errors.New(data["error"])
    }

    return nil
}

func ChangeMainServiceStatus(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/ChangeMainServiceStatus"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeMainServiceStatus ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ChangeMainServiceStatus ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ChangeMainServiceStatus ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    /*    if data["ack"] == "false"{
          defer resp.Body.Close()
          return errors.New(data["error"])
      } */
    defer resp.Body.Close()
    return nil
}

func DeleteService(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/deleteService"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeleteService ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/DeleteService ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/DeleteService ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    /*    if data["ack"] == "false"{
          defer resp.Body.Close()
          return errors.New(data["error"])
      } */
    defer resp.Body.Close()
    return nil
}

func UpdateSuricataValue(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/updateSuricataValue"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/UpdateSuricataValue ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/UpdateSuricataValue ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/UpdateSuricataValue ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    /*    if data["ack"] == "false"{
          defer resp.Body.Close()
          return errors.New(data["error"])
      } */

    defer resp.Body.Close()
    return nil
}

func DeployStapService(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/deployStapService"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeployStapService ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func StopStapService(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/stopStapService"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/StopStapService ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ChangeServiceStatus ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func ModifyNodeOptionValues(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/modifyNodeOptionValues"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ModifyNodeOptionValues ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ModifyNodeOptionValues ERROR reading request data: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ModifyNodeOptionValues ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func PingWazuhFiles(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/wazuh/pingWazuhFiles"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingWazuhFiles ERROR connection through http new Request: " + err.Error())
        return data, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/PingWazuhFiles ERROR reading request data: " + err.Error())
        return data, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/PingWazuhFiles ERROR doing unmarshal JSON: " + err.Error())
        return data, err
    }

    defer resp.Body.Close()
    return data, nil
}

func DeleteWazuhFile(ipData string, portData string, anode map[string]interface{}) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/wazuh/deleteWazuhFile"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeleteWazuhFile ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func AddWazuhFile(ipData string, portData string, anode map[string]interface{}) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/wazuh/addWazuhFile"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddWazuhFile ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func LoadFileLastLines(ipData string, portData string, anode map[string]string) (data map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/wazuh/loadFileLastLines"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/LoadFileLastLines ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/LoadFileLastLines ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/LoadFileLastLines ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func SaveFileContentWazuh(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/wazuh/saveFileContentWazuh"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveFileContentWazuh ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    data := make(map[string]string)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/LoadFileLastLines ERROR reading request data: " + err.Error())
        return err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/LoadFileLastLines ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func ReloadFilesData(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/file/reloadFilesData"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/ReloadFilesData ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ReloadFilesData ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ReloadFilesData ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func AddMonitorFile(ipuuid string, portuuid string, data map[string]string) (err error) {
    url := "https://" + ipuuid + ":" + portuuid + "/node/monitor/addFile"
    valuesJSON, err := json.Marshal(data)
    if err != nil {
        logs.Error("nodeclient/AddMonitorFile Error Marshal new JSON data: " + err.Error())
        return err
    }
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddMonitorFile ERROR on the new HTTP request response: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func PingMonitorFiles(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/monitor/pingMonitorFiles"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingMonitorFiles ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/PingMonitorFiles ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/PingMonitorFiles ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return data, nil
}

func DeleteMonitorFile(ipuuid string, portuuid string, data map[string]string) (err error) {
    url := "https://" + ipuuid + ":" + portuuid + "/node/monitor/deleteFile"
    valuesJSON, err := json.Marshal(data)
    if err != nil {
        logs.Error("nodeclient/DeleteMonitorFile Error Marshal new JSON data: " + err.Error())
        return err
    }
    resp, err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeleteMonitorFile ERROR on the new HTTP request response: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func ChangeZeekMode(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/changeZeekMode"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeZeekMode ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func AddClusterValue(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/addClusterValue"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddClusterValue ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func PingCluster(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/zeek/pingCluster"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/PingCluster ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/PingCluster ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/PingCluster ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    if data["hasError"]["ack"] == "false" {
        return nil, errors.New(data["hasError"]["error"])
    }
    return data, nil
}

func EditClusterValue(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/editClusterValue"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/EditClusterValue ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func DeleteClusterValue(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/deleteClusterValue"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("DELETE", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/DeleteClusterValue ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func SyncCluster(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek/syncCluster"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncCluster ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func GetChangeControlNode(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/changecontrol"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetChangeControlNode ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetChangeControlNode ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/GetChangeControlNode ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    if data["error"] != nil {
        return nil, errors.New(data["error"]["error"])
    }
    return data, nil
}

func GetIncidentsNode(ipData string, portData string) (data map[string]map[string]string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/incidents"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/GetIncidentsNode ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetIncidentsNode ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/GetIncidentsNode ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    if data["error"] != nil {
        return nil, errors.New(data["error"]["error"])
    }
    return data, nil
}

func PutIncidentNode(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/incidents"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/PutIncidentNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    return nil
}

func ChangeSuricataTable(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/plugin/changeSuricataTable"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeSuricataTable ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/ChangeSuricataTable ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    _ = json.Unmarshal(body, &returnValues)

    if returnValues["ack"] == "false" {
        return errors.New("Error getting node group http response: " + returnValues["error"])
    }
    return nil
}

func ChangeSuricataPathsGroups(ipnid string, portnid string, data map[string][]byte) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/suricata"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeSuricataPathsGroups ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/ChangeSuricataPathsGroups ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    _ = json.Unmarshal(body, &returnValues)

    if returnValues["ack"] == "false" {
        return errors.New("Error getting node group http response: " + returnValues["error"])
    }
    return nil
}

func ChangeZeekPathsGroups(ipnid string, portnid string, data map[string]map[string][]byte) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/zeek"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeZeekPathsGroups ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/ChangeZeekPathsGroups ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    _ = json.Unmarshal(body, &returnValues)

    if returnValues["ack"] == "false" {
        return errors.New("Error getting node group http response: " + returnValues["error"])
    }
    return nil
}

func PutSuricataServicesFromGroup(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/suricta/putServices"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/PutSuricataServicesFromGroup ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SuricataGroupService ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    _ = json.Unmarshal(body, &returnValues)

    if returnValues["ack"] == "false" {
        return errors.New("Error getting node group http response: " + returnValues["error"])
    }

    return nil
}

func SyncAnalyzerToAllGroupNodes(ipData string, portData string, data []byte) (err error) {
    values := make(map[string][]byte)
    values["data"] = data
    url := "https://" + ipData + ":" + portData + "/node/analyzer/sync"
    valuesJSON, err := json.Marshal(values)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncAnalyzerToAllGroupNodes ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SyncAnalyzerToAllGroupNodes ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/SyncAnalyzerToAllGroupNodes ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func SyncClusterFileNode(ipData string, portData string, data []byte) (err error) {
    values := make(map[string][]byte)
    values["data"] = data
    url := "https://" + ipData + ":" + portData + "/node/zeek/syncClusterFile"
    valuesJSON, err := json.Marshal(values)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncClusterFileNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SyncClusterFileNode ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/SyncClusterFileNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func StartSuricataMainConf(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/suricata/StartSuricataMainConf"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/StartSuricataMainConf ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/StartSuricataMainConf ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/StartSuricataMainConf ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func StopSuricataMainConf(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/suricata/StopSuricataMainConf"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/StopSuricataMainConf ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/StopSuricataMainConf ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/StopSuricataMainConf ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func KillSuricataMainConf(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/suricata/KillSuricataMainConf"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/KillSuricataMainConf ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/KillSuricataMainConf ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/KillSuricataMainConf ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func ReloadSuricataMainConf(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/suricata/ReloadSuricataMainConf"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ReloadSuricataMainConf ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/ReloadSuricataMainConf ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/ReloadSuricataMainConf ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func LaunchZeekMainConf(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/zeek/LaunchZeekMainConf"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/LaunchZeekMainConf ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/LaunchZeekMainConf ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/LaunchZeekMainConf ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

// func SaveZeekValues(ipData string, portData string, anode map[string]string)(err error){
//     url := "https://"+ipData+":"+portData+"/node/zeek/saveZeekValues"
//     valuesJSON,err := json.Marshal(anode)
//     resp,err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
//     if err != nil {
//         logs.Error("nodeclient/SaveZeekValues ERROR connection through http new Request: "+err.Error())
//         return err
//     }

//     body, err := ioutil.ReadAll(resp.Body)
//     defer resp.Body.Close()
//     if err != nil { logs.Error("nodeclient/SaveZeekValues ERROR reading request data: "+err.Error()); return err}

//     returnValues := make(map[string]string)
//     err = json.Unmarshal(body, &returnValues)
//     if err != nil { logs.Error("nodeclient/SaveZeekValues ERROR doing unmarshal JSON: "+err.Error()); return err}

//     if returnValues["ack"] == "false"{
//         return errors.New(returnValues["error"])
//     }

//     return nil
// }

func SyncZeekValues(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/zeek/syncZeekValues"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncZeekValues ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SyncZeekValues ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/SyncZeekValues ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func SyncAllSuricataGroup(ipData string, portData string, data map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/group/sync"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncAllSuricataGroup ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SyncAllSuricataGroup ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    err = json.Unmarshal(body, &returnValues)
    if err != nil {
        logs.Error("nodeclient/SyncAllSuricataGroup ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if returnValues["ack"] == "false" {
        return errors.New(returnValues["error"])
    }

    return nil
}

func SuricataGroupService(ipData string, portData string, data map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/group/suricata"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SuricataGroupService ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/SuricataGroupService ERROR reading request data: " + err.Error())
        return err
    }

    returnValues := make(map[string]string)
    _ = json.Unmarshal(body, &returnValues)

    if returnValues["ack"] == "false" {
        return errors.New("Error getting node group http response: " + returnValues["error"])
    }

    return nil
}

func GetNodeToken(ipData string, portData string, login map[string]string) (token string, err error) {
    url := "https://" + ipData + ":" + portData + "/node/autentication"
    valuesJSON, err := json.Marshal(login)

    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/GetNodeAutentication ERROR connection through http new Request: " + err.Error())
        //add to db with wait status
        return "", errors.New("connection refused")
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetNodeAutentication ERROR reading request data: " + err.Error())
        return "", err
    }
    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/GetNodeAutentication ERROR doing unmarshal JSON: " + err.Error())
        return "", err
    }
    if data["ack"] == "false" {
        return "", errors.New(data["error"])
    }
    return data["token"], nil
}

func SaveNodeInformation(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/ping/saveNodeInformation"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveNodeInformation ERROR connection through http new Request: " + err.Error())
        return err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SaveNodeInformation ERROR reading request data: " + err.Error())
        return err
    }
    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SaveNodeInformation ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func DeleteNode(ipData string, portData string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/ping"
    resp, err := utils.NewRequestHTTP("DELETE", url, nil)
    if err != nil {
        logs.Error("nodeclient/DeleteNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/DeleteNode ERROR reading request data: " + err.Error())
        return err
    }
    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/DeleteNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if data["ack"] == "false" {
        return errors.New(data["error"])
    }
    return nil
}

func ChangeRotationStatus(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/monitor/changeRotationStatus"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/ChangeRotationStatus ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ChangeRotationStatus ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/ChangeRotationStatus ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }

    defer resp.Body.Close()
    return nil
}

func EditRotation(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/monitor/editRotation"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/EditRotation ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/EditRotation ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/EditRotation ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }

    defer resp.Body.Close()
    return nil
}

func SyncUsersToNode(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addUser"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncUsersToNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncUsersToNode ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncUsersToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func SyncRolesToNode(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addRole"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncRolesToNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncRolesToNode ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncRolesToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func SyncGroupsToNode(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addGroup"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncGroupsToNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncGroupsToNode ERROR reading request data: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncGroupsToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    return nil
}

func SyncUserGroupRolesToNode(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addUgr"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncUserGroupRolesToNode ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncUserGroupRolesToNode ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncUserGroupRolesToNode ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func GetServiceCommands(ipnid string, portnid string, data map[string]string) (values map[string]map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/plugin/getCommands"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/GetServiceCommands ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetServiceCommands ERROR reading request data: " + err.Error())
        return nil, err
    }

    err = json.Unmarshal(body, &values)
    if err != nil {
        logs.Error("nodeclient/GetServiceCommands ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    defer resp.Body.Close()
    return values, nil
}

func SyncRolePermissions(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addRolePerm"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncRolePermissions ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncRolePermissions ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncRolePermissions ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func SyncPermissions(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addPerm"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncPermissions ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncPermissions ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncPermissions ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func SyncRoleGroups(ipnid string, portnid string, data map[string]map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/autentication/addRoleGroups"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SyncRoleGroups ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SyncRoleGroups ERROR reading request data: " + err.Error())
        return err
    }

    mapData := make(map[string]string)
    err = json.Unmarshal(body, &mapData)
    if err != nil {
        logs.Error("nodeclient/SyncRoleGroups ERROR doing unmarshal JSON: " + err.Error())
        return err
    }
    if mapData["ack"] == "false" {
        return errors.New(mapData["error"])
    }
    defer resp.Body.Close()
    return nil
}

func GetMD5files(ipnid string, portnid string, data map[string]map[string]string) (allData map[string]map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/suricata/getMD5files"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/GetMD5files ERROR connection through http new Request: " + err.Error())
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/GetMD5files ERROR reading request data: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    err = json.Unmarshal(body, &allData)
    if err != nil {
        logs.Error("nodeclient/GetMD5files ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    if allData["hasError"]["ack"] == "false" {
        return nil, errors.New(allData["hasError"]["error"])
    }

    return allData, nil
}

func SaveSurictaRulesetSelected(ipData string, portData string, anode map[string]string) (err error) {
    url := "https://" + ipData + ":" + portData + "/node/plugin/setRuleset"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/SaveSurictaRulesetSelected ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/SaveSurictaRulesetSelected ERROR reading request data: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/SaveSurictaRulesetSelected ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

//
func AddMacIp(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/hwaddmng/"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddMacIp ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    return nil
}

func LoadConfig(ipnid string, portnid string, data map[string]string) (values map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/hwaddmng/config"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/LoadConfig ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/LoadConfig ERROR reading request data: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    // data := make(map[string]string)
    err = json.Unmarshal(body, &values)
    if err != nil {
        logs.Error("nodeclient/LoadConfig ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    return values, nil
}

func Config(ipnid string, portnid string, data map[string]interface{}) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/hwaddmng/config"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/Config ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    return nil
}

func Db(ipnid string, portnid string, data map[string]string) (err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/hwaddmng/db"
    valuesJSON, err := json.Marshal(data)
    resp, err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/Db ERROR connection through http new Request: " + err.Error())
        return err
    }
    defer resp.Body.Close()

    return nil
}

func ConfigGet(ipnid string, portnid string) (data map[string]string, err error) {
    url := "https://" + ipnid + ":" + portnid + "/node/hwaddmng/config"
    resp, err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
        logs.Error("nodeclient/ConfigGet ERROR connection through http new Request: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logs.Error("nodeclient/ConfigGet ERROR reading request data: " + err.Error())
        return nil, err
    }
    defer resp.Body.Close()

    // data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/ConfigGet ERROR doing unmarshal JSON: " + err.Error())
        return nil, err
    }

    return data, nil
}

func AddSuricataService(uuid string, anode map[string]string) (err error) {
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    url := "https://" + ipData + ":" + portData + "/node/suricata/service"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddSuricataService ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/AddSuricataService ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/AddSuricataService ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}

func AddStapService(uuid string, anode map[string]string) (err error) {
    ipData, portData, err := ndb.ObtainPortIp(uuid)
    url := "https://" + ipData + ":" + portData + "/node/plugin/addService"
    valuesJSON, err := json.Marshal(anode)
    resp, err := utils.NewRequestHTTP("PUT", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
        logs.Error("nodeclient/AddStapService ERROR connection through http new Request: " + err.Error())
        return err
    }

    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        logs.Error("nodeclient/AddStapService ERROR reading request data: " + err.Error())
        return err
    }

    data := make(map[string]string)
    err = json.Unmarshal(body, &data)
    if err != nil {
        logs.Error("nodeclient/AddStapService ERROR doing unmarshal JSON: " + err.Error())
        return err
    }

    if data["ack"] == "false" {
        return errors.New(data["error"])
    }

    return nil
}
