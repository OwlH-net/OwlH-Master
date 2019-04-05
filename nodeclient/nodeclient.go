package nodeclient

import (
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"encoding/json"
	"owlhmaster/utils"
    // "net/http"
    //"bytes"
    // "crypto/tls"
)

func init() {

}
 
func Echo() {
    logs.Info("NODE CLIENT -> ECHO")
}

func PingNode(ip string, port string) (err error) {
    logs.Info("NodeClient PingNode -> %s, %s", ip, port)
	url := "https://"+ip+":"+port+"/node/node"
    // tr := &http.Transport{
		//     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// }
		// req, err := http.NewRequest("GET", url, nil)
		// client := &http.Client{Transport: tr}
		// resp, err := client.Do(req)
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
    // tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    // }
    // req, err := http.NewRequest("GET", url, nil)
    // client := &http.Client{Transport: tr}
	// resp, err := client.Do(req)
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
    // tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    // }
    // req, err := http.NewRequest("GET", url, nil)
    // client := &http.Client{Transport: tr}
    // resp, err := client.Do(req)
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
    // tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    // }
    // req, err := http.NewRequest("GET", url, nil)
    // client := &http.Client{Transport: tr}
    // resp, err := client.Do(req)
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
    // tr := &http.Transport{
    //     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    // }
    // req, err := http.NewRequest("GET", url, nil)
    // client := &http.Client{Transport: tr}
    // resp, err := client.Do(req)
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