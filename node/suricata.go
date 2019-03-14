package node

import (
    "github.com/astaxie/beego/logs"
// 	  "strings"
//    "database/sql"
//    "fmt"
//    "time"
//    _ "github.com/mattn/go-sqlite3"
    "crypto/tls"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
    "encoding/json"
//    "regexp"
    "io/ioutil"
    "net/http"
    "bytes"
)

func Suricata(n string) (data map[string]bool, err error) {
    logs.Info("Node suricata -> IN")

    // ip, err := getNodeIPbyUID(n)
    // if err != nil {
    //     logs.Info("Suricata - IP Error -> %s", err.Error())
    //     return nil,err
    // }
    // port, err := getNodePortbyUID(n)
    // if err != nil {
    //     logs.Info("Suricata - PORT Error -> %s", err.Error())
    //     return nil,err
    // }    
    ip,port,err := utils.ObtainPortIp(n)
    if err != nil {
        logs.Info("Suricata - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Suricata IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Suricata(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func GetSuricataBPF(n string)(bpf string, err error) {
    if ndb.Db == nil {
        logs.Error("GetSuricataBPF -- Can't acces to database")
        return "", errors.New("GetSuricataBPF -- Can't acces to database")
    }
    var res string
    sql := "select node_value from nodes where node_uniqueid = \""+n+"\" and node_param = \"bpf\";"
    logs.Info("Get BPF Suricata query sql %s",sql)
    rows, err := ndb.Db.Query(sql)
    
    if err != nil {
        logs.Info("Get BPF Suricata query error %s",err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        err = rows.Scan(&res)
        logs.Info("Get BPF Suricata there are rows")
        if  err != nil {
            logs.Info("Get BPF Suricata scan error %s",err.Error())
            return "", err
        }
        logs.Info("Get BPF Suricata res: "+res)
        return res, err
    }
    return "", errors.New("Get SuricataBPF -- There is no defined BPF")
    //select node_value from nodes where node_uniqueid like '%que-rico%' and node_param = "ip";
}

func PutSuricataBPF(n map[string]string)(bpf string, err error) {
    if ndb.Db == nil {
        logs.Error("PutSuricataBPF -- Can't acces to database")
        return "", errors.New("PutSuricataBPF -- Can't acces to database")
    }
    jsonnid := n["nid"]
    jsonbpf := n["bpf"]
    bpftext := "bpf"

    // ipnid,portnid,err := GetSuricataIpPort(jsonnid)
    ipnid,portnid,err := utils.ObtainPortIp(jsonnid)
    
    //crear map con nid y bpf
    values := make(map[string]string)
    values["nid"] = jsonnid
    values["bpf"] = jsonbpf

    //pasar json al cliente (MIRAR COMO SE HACE)
    valuesJSON,err := json.Marshal(values)

    url := "https://"+ipnid+":"+portnid+"/node/suricata/bpf"
    logs.Info("\n"+url+"\n")
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(valuesJSON))
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    logs.Info("Request newBuffer(JSON) -------> ",req.Body)
    logs.Info("Resp cliente Do (request) -------->",resp.Body)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()
    
    logs.Info("response Status:", resp.Status)
    logs.Info("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("response Body:", string(body))

    sql := "select node_value from nodes where node_uniqueid = \""+jsonnid+"\" and node_param = \"bpf\";"
    logs.Info("Put BPF Suricata query sql %s",sql)
    rows, err := ndb.Db.Query(sql)
    
    if err != nil {
        logs.Info("Put BPF Suricata query error %s",err.Error())
        return "", err
    }

    defer rows.Close()
    if rows.Next() {
        rows.Close()
        logs.Info("Put BPF Suricata res UPDATE")
        updtbpf, err := ndb.Db.Prepare("update nodes set node_value = ? where node_uniqueid = ? and node_param = ?;")

        if (err != nil){
            logs.Info("Put BPF Suricata prepare UPDATE -- "+err.Error())
            return "", err
        }
        _, err = updtbpf.Exec(&jsonbpf, &jsonnid, bpftext)  
        defer updtbpf.Close()      

        return "SuccessUpdate", err
    }else{
        logs.Info("Put BPF Suricata res INSERT")
        indtbpf, err := ndb.Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values (?,?,?);")
        _, err = indtbpf.Exec(&jsonnid, bpftext, &jsonbpf)  
        defer indtbpf.Close()
        if (err != nil){
            return "", err
        }
        return "SuccessInsert", err
    }
    return "Error", errors.New("Put SuricataBPF -- There is no defined BPF")
}

// func GetSuricataIpPort(jsonnid string)(ip string, port string, err error){ //(ipReturn string, portReturn string,  err error ) {
//     if ndb.Db == nil {
//         logs.Error("GetSuricataIpPort -- Can't acces to database")
//         return "","", errors.New("GetSuricataIpPort -- Can't acces to database")
//     }

//     var ipObtained string
//     var portObtained string

//     sqlIP := "select node_value from nodes where node_uniqueid = \""+jsonnid+"\" and node_param = \"ip\";"
//     rowIP, err := ndb.Db.Query(sqlIP)
//     sqlPORT := "select node_value from nodes where node_uniqueid = \""+jsonnid+"\" and node_param = \"port\";"
//     rowPORT, err := ndb.Db.Query(sqlPORT)
    
//     defer rowIP.Close()
//     defer rowPORT.Close()
    
//     if rowIP.Next() {
//         err = rowIP.Scan(&ipObtained)
//         if  err != nil {
//             logs.Info("--- NO IP FOR THIS NID ---"+err.Error())
//             return "","",err
//         }
//     }
//     if rowPORT.Next() {
//         err = rowPORT.Scan(&portObtained)
//         if  err != nil {
//             logs.Info("--- NO PORT FOR THIS NID ---"+err.Error())
//             return "","",err
//         }
//     }
//     return ipObtained,portObtained,err
// }

func RunSuricata(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunSuricata -- Can't acces to database")
        return "", errors.New("RunSuricata -- Can't acces to database")
    }
    
    // ipnid,portnid,err := GetSuricataIpPort(uuid)
    ipnid,portnid,err := utils.ObtainPortIp(uuid)
    
    url := "https://"+ipnid+":"+portnid+"/node/suricata/RunSuricata"
    req, err := http.NewRequest("PUT", url, nil)
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    if err != nil {
        return "",err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    logs.Info("RunSuricata function "+string(body))
    return string(body),nil
}

func StopSuricata(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopSuricata -- Can't acces to database")
        return "", errors.New("StopSuricata -- Can't acces to database")
    }

    // ipnid,portnid,err := GetSuricataIpPort(uuid)
    ipnid,portnid,err := utils.ObtainPortIp(uuid)

    url := "https://"+ipnid+":"+portnid+"/node/suricata/StopSuricata"
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