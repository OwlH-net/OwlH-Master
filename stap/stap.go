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
   _ "github.com/mattn/go-sqlite3"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
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

func GetServer(uuid string, serveruuid string)(data map[string]map[string]string, err error){
    rData := make(map[string]map[string]string)
    logs.Info("stap/stap.go stap GetServer()")
    ipuuid,portuuid,err := utils.ObtainPortIp(uuid)
    url := "https://"+ipuuid+":"+portuuid+"/node/stap/server/"+serveruuid

    req, err := http.NewRequest("GET", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    logs.Info("GetServer Request -------> ",req.Body)
    if err != nil {
        return nil,err
	}
    defer resp.Body.Close()
    responseData, _ := ioutil.ReadAll(resp.Body)

    json.Unmarshal(responseData, &rData)
    logs.Info(rData)

    return rData,nil
}
//ping to Stap function at node.html. Create or update this function if is needed
func Stap(n string) (data map[string]bool, err error) {
    logs.Info("Node Stap -> IN")

    ip,port,err := utils.ObtainPortIp(n)
    if err != nil {
        logs.Info("Stap - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Stap IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Stap(ip,port,n)
    if err != nil {
        return nil,err
    }
    return data,nil
}

//Launch stap main server
func RunStap(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunStap -- Can't acces to database")
        return "", errors.New("RunStap -- Can't acces to database")
    }
    
    ipnid,portnid,err := utils.ObtainPortIp(uuid)
    
    url := "https://"+ipnid+":"+portnid+"/node/stap/RunStap/"+uuid
    req, err := http.NewRequest("PUT", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("RunStap function "+string(body))
    return string(body),nil
}

//Stop stap main server
func StopStap(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStap -- Can't acces to database")
        return "", errors.New("StopStap -- Can't acces to database")
    }

    ipnid,portnid,err := utils.ObtainPortIp(uuid)

    url := "https://"+ipnid+":"+portnid+"/node/stap/StopStap/"+uuid
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

//Launch stap specific server
func RunStapServer(uuid string, server string)(data string, err error){
    logs.Info("RunStapServer uuid "+uuid)
    logs.Info("RunStapServer server "+server)
    if ndb.Db == nil {
        logs.Error("RunStapServer -- Can't acces to database")
        return "", errors.New("RunStapServer -- Can't acces to database")
    }
    
    ipnid,portnid,err := utils.ObtainPortIp(uuid)
    
    url := "https://"+ipnid+":"+portnid+"/node/stap/RunStapServer/"+server
    logs.Info("URL --> "+url)
    req, err := http.NewRequest("PUT", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    logs.Warn("RunStapServer function "+string(body))
    return string(body),nil
}

//Stop stap specific server
func StopStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStapServer -- Can't acces to database")
        return "", errors.New("StopStapServer -- Can't acces to database")
    }

    ipnid,portnid,err := utils.ObtainPortIp(uuid)

    url := "https://"+ipnid+":"+portnid+"/node/stap/StopStapServer/"+server
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

func PingServerStap(uuid string, server string) (data map[string]bool, err error) {
    ip,port,err := utils.ObtainPortIp(uuid)
    if err != nil {
        logs.Info("PingServerStap - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    

    tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    url := "https://"+ip+":"+port+"/node/stap/PingServerStap/"+server
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
    err = json.Unmarshal(body, &data)
    if err != nil {
        return nil,err
    }
    return data,nil
}






