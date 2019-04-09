package stap

import(
    "io/ioutil"
    // "fmt"
	"github.com/astaxie/beego/logs"
	"bytes"
	"encoding/json"
	// "net/http"
	// "owlhmaster/node"
    // "bufio" //read line by line the doc
    // "regexp"
    // "os"
    // "os/exec"
    //"strconv"
	// "crypto/tls"
    // "owlhmaster/database"
    "owlhmaster/utils"
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

//Add server to software TAP
func AddServer(data map[string]string)(err error) {
	uuid := data["uuid"]
	ipuuid,portuuid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("AddServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return err
	}
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"
	valuesJSON,err := json.Marshal(data)
	if err != nil {
		logs.Error("Error Marshal new JSON data: "+err.Error())
        return err
	}
	resp,err := utils.NewRequestHTTP("POST", url, bytes.NewBuffer(valuesJSON))
    if err != nil {
		logs.Error("AddServer ERROR on the new HTTP request response: "+err.Error())
        return err
	}
    defer resp.Body.Close()
    return nil
}

//Get all STAP servers
func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error){
    rData := make(map[string]map[string]string)
    ipuuid,portuuid,err := ndb.ObtainPortIp(nodeuuid)
	if err != nil {
		logs.Error("GetAllServers ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return nil,err
	}
	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("GetAllServers ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
    defer resp.Body.Close()
    responseData, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(responseData, &rData)
    return rData,nil
}

//Get a specific server
func GetServer(uuid string, serveruuid string)(data map[string]map[string]string, err error){
    rData := make(map[string]map[string]string)
	ipuuid,portuuid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("GetServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return nil,err
	}
    url := "https://"+ipuuid+":"+portuuid+"/node/stap/server/"+serveruuid
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("GetServer ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
    defer resp.Body.Close()
    responseData, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(responseData, &rData)
    return rData,nil
}

//ping to Stap function at node.html. Create or update this function if is needed
func Stap(n string) (data map[string]bool, err error) {
    ip,port,err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Error("Stap ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return nil,err
    }    
	data, err = nodeclient.Stap(ip,port,n)
    if err != nil {
		logs.Error("Stap ERROR: "+err.Error())
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
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("RunStap ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    url := "https://"+ipnid+":"+portnid+"/node/stap/RunStap/"+uuid
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("RunStap ERROR on the new HTTP request response: "+err.Error())
        return "",err
	}
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

//Stop stap main server
func StopStap(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStap -- Can't acces to database")
        return "", errors.New("StopStap -- Can't acces to database")
    }

	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("StopStap ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    url := "https://"+ipnid+":"+portnid+"/node/stap/StopStap/"+uuid
    resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("StopStap ERROR on the new HTTP request response: "+err.Error())
        return "",err
	}
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

//Launch stap specific server
func RunStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunStapServer -- Can't acces to database")
        return "", errors.New("RunStapServer -- Can't acces to database")
    }
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("RunStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    url := "https://"+ipnid+":"+portnid+"/node/stap/RunStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("RunStapServer ERROR on the new HTTP request response: "+err.Error())
        return "",err
	}
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

//Stop stap specific server
func StopStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStapServer -- Can't acces to database")
        return "", errors.New("StopStapServer -- Can't acces to database")
    }
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("StopStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
	}
	url := "https://"+ipnid+":"+portnid+"/node/stap/StopStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("StopStapServer ERROR on the new HTTP request response: "+err.Error())
        return "",err
	}
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

//Delete specific stap server
func DeleteStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("DeleteStapServer -- Can't acces to database")
        return "", errors.New("DeleteStapServer -- Can't acces to database")
    }
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("DeleteStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    url := "https://"+ipnid+":"+portnid+"/node/stap/DeleteStapServer/"+server
	resp,err := utils.NewRequestHTTP("PUT", url, nil)
    if err != nil {
		logs.Error("DeleteStapServer ERROR on the new HTTP request response: "+err.Error())
        return "",err
	}
    defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
    return string(body),nil
}

func PingServerStap(uuid string, server string) (data map[string]string, err error) {
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Info("PingServerStap - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
	url := "https://"+ip+":"+port+"/node/stap/PingServerStap/"+server
	resp,err := utils.NewRequestHTTP("GET", url, nil)
    if err != nil {
		logs.Error("PingServerStap ERROR on the new HTTP request response: "+err.Error())
        return nil,err
	}
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, &data)
    if err != nil {
		logs.Error("PingServerStap ERROR doing unmarshal JSON: "+err.Error())
        return nil,err
	}
    return data,nil
}

func EditStapServer(data map[string]string) (err error) {
	uuid := data["uuid"]
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Info("EditStapServer - get IP and PORT Error -> %s", err.Error())
        return err
    }    
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