package node

import (
    "github.com/astaxie/beego/logs"
// 	  "strings"
//    "database/sql"
//    "fmt"
//    "time"
   _ "github.com/mattn/go-sqlite3"
    "crypto/tls"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
    // "encoding/json"
//    "regexp"
    "io/ioutil"
    "net/http"
    // "bytes"
)

func Wazuh(n string) (data map[string]bool, err error) {
    logs.Info("Node Wazuh -> IN")

    ip,port,err := utils.ObtainPortIp(n)
    if err != nil {
        logs.Info("Wazuh - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Wazuh IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Wazuh(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func RunWazuh(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunWazuh -- Can't acces to database")
        return "", errors.New("RunWazuh -- Can't acces to database")
    }
    
    // ipnid,portnid,err := GetSuricataIpPort(uuid)
    ipnid,portnid,err := utils.ObtainPortIp(uuid)
    
    url := "https://"+ipnid+":"+portnid+"/node/wazuh/RunWazuh"
    req, err := http.NewRequest("PUT", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("RunWazuh function "+string(body))
    return string(body),nil
}

func StopWazuh(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopWazuh -- Can't acces to database")
        return "", errors.New("StopWazuh -- Can't acces to database")
    }

    // ipnid,portnid,err := GetSuricataIpPort(uuid)
    ipnid,portnid,err := utils.ObtainPortIp(uuid)

    url := "https://"+ipnid+":"+portnid+"/node/wazuh/StopWazuh"
    req, err := http.NewRequest("PUT", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}