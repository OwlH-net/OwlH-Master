package stap

import(
    "io/ioutil"
    // "fmt"
	"github.com/astaxie/beego/logs"
	"bytes"
	"encoding/json"
	"net/http"
	// "owlhmaster/node"
    // "bufio" //read line by line the doc
    // "regexp"
    // "os"
    // "os/exec"
    //"strconv"
	"owlhmaster/utils"
	"crypto/tls"
    // "owlhmaster/database"
    // "errors"
    // "database/sql"
    // "strings"
    // "time"
    // "strconv"
)

func AddServer(data map[string]string)(err error) {
	
	logs.Info("stap/stap.go stap AddServer()")
	uuid := data["uuid"]
	ipuuid,portuuid,err := utils.ObtainPortIp(uuid)
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"
	valuesJSON,err := json.Marshal(data)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(valuesJSON))
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    logs.Info("Request newBuffer(JSON) -------> ",req.Body)
    if err != nil {
        return err
	}
    defer resp.Body.Close()
    return nil
}


func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error){
    rData := make(map[string]map[string]string)
    logs.Info("stap/stap.go stap GetAllServers()")
    ipuuid,portuuid,err := utils.ObtainPortIp(nodeuuid)
    url := "https://"+ipuuid+":"+portuuid+"/node/stap/"

    req, err := http.NewRequest("GET", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    logs.Info("GetAllServers Request -------> ",req.Body)
    if err != nil {
        return nil,err
	}
    defer resp.Body.Close()
    responseData, _ := ioutil.ReadAll(resp.Body)

    json.Unmarshal(responseData, &rData)
    logs.Info(rData)


    return rData,nil
}