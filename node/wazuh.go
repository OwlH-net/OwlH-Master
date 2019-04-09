package node

import (
    "github.com/astaxie/beego/logs"
// 	  "strings"
//    "database/sql"
//    "fmt"
//    "time"
   _ "github.com/mattn/go-sqlite3"
    // "crypto/tls"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
    // "encoding/json"
//    "regexp"
    "io/ioutil"
    // "net/http"
    // "bytes"
)

func Wazuh(n string) (data map[string]bool, err error) {
    ip,port,err := ndb.ObtainPortIp(n)
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
    
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    url := "https://"+ipnid+":"+portnid+"/node/wazuh/RunWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("node/RunWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

func StopWazuh(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopWazuh -- Can't acces to database")
        return "", errors.New("StopWazuh -- Can't acces to database")
	}
	
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    url := "https://"+ipnid+":"+portnid+"/node/wazuh/StopWazuh"
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
	if err != nil {
		logs.Error("node/StopWazuh ERROR connection through http new Request: "+err.Error())
        return "", err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}