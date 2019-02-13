package nodeclient

import (
    "github.com/astaxie/beego/logs"
    "net/http"
    //"bytes"
    "io/ioutil"
    "crypto/tls"
    //"encoding/json"
)

func init() {

}
 
func Echo() {
    logs.Info("NODE CLIENT -> ECHO")
}

func PingNode(ip string, port string) (err error) {
    logs.Info("NodeClient PingNode -> %s, %s", ip, port)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    url := "https://"+ip+":"+port+"/node/node"
    req, err := http.NewRequest("GET", url, nil)
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("response Body:", string(body))
    return nil
}

func Suricata(ip string, port string) (data []byte, err error ) {
    logs.Info("NodeClient suricata status -> %s, %s", ip, port)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    url := "https://"+ip+":"+port+"/node/suricata"
    req, err := http.NewRequest("GET", url, nil)
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    //data, _ := json.Marshal(body)
    //logs.Info("response Body:", string(data))
    return body,nil//data,nil
}

func Zeek(ip string, port string) (data []byte, err error ) {
    logs.Info("NodeClient zeek status -> %s, %s", ip, port)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    url := "https://"+ip+":"+port+"/node/zeek"
    req, err := http.NewRequest("GET", url, nil)
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    //data, _ := json.Marshal(body)
    //logs.Info("response Body:", string(data))
    return body,nil//data,nil
}

func Wazuh(ip string, port string) (data []byte, err error ) {
    logs.Info("NodeClient wazuh status -> %s, %s", ip, port)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    url := "https://"+ip+":"+port+"/node/wazuh"
    req, err := http.NewRequest("GET", url, nil)
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)
    if err != nil {
        return nil,err
    }
    defer resp.Body.Close()
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    //data, _ := json.Marshal(body)
    //logs.Info("response Body:", string(data))
    return body,nil//data,nil
}