package node

import (
    "github.com/astaxie/beego/logs"
    // "strings"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    // "regexp"
)

func Suricata (n string) (data []byte, err error) {
    logs.Info("Node suricata -> IN")
    logs.Info("Suricata - UID -> %s", n)

    ip, err := getNodeIPbyUID(n)
    if err != nil {
        logs.Info("Suricata - IP Error -> %s", err.Error())
        return nil,err
    }
    port, err := getNodePortbyUID(n)
    if err != nil {
        logs.Info("Suricata - PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Suricata - vamos a por el suricata -> %s, %s", ip, port)
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
        //sql = "update nodes set node_value = '"+jsonbpf+"' where node_uniqueid = '"+jsonnid+"' and node_param = 'bpf';"
        //_, err := ndb.Db.Query(sql)
        updtbpf, err := ndb.Db.Prepare("update nodes set node_value = ? where node_uniqueid = ? and node_param = ?;")
        //ndb.Db.Exec("PRAGMA journal_mode=WAL;")
        //updtbpf, err := ndb.Db.Prepare("delete from nodes where node_uniqueid = ? and node_param = ?;")
        if (err != nil){
            logs.Info("Put BPF Suricata prepare UPDATE -- "+err.Error())
            return "", err
        }
        _, err = updtbpf.Exec(&jsonbpf, &jsonnid, bpftext)  
        defer updtbpf.Close()      
        //logs.Info("PUT BPF UPDATE -- "+updtbpf)
        // if (err != nil){
        //     logs.Info("Put BPF Suricata res DELETE -- "+err.Error())
        //     return "", err
        // }
        // indtbpf, err := ndb.Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values (?,?,?);")
        // logs.Info("PUT BPF INSERT -- "+updtbpf)
        // _, err = indtbpf.Exec(&jsonnid, bpftext, &jsonbpf)  
        // defer updtbpf.Close()
        // if (err != nil){
        //     return "", err
        // }
        // return "SuccessInsert", err
        return "SuccessUpdate", err
    }else{
        logs.Info("Put BPF Suricata res INSERT")
        //sql = "insert into nodes (node_uniqueid, node_param, node_value) values ('"+jsonnid+"', 'bpf', '"+jsonbpf+"');"
        //_, err := ndb.Db.Query(sql)
        indtbpf, err := ndb.Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values (?,?,?);")
        //logs.Info("PUT BPF INSERT -- "+updtbpf)
        _, err = indtbpf.Exec(&jsonnid, bpftext, &jsonbpf)  
        defer indtbpf.Close()
        if (err != nil){
            return "", err
        }
        return "SuccessInsert", err
    }
    return "Error", errors.New("Put SuricataBPF -- There is no defined BPF")
    //select node_value from nodes where node_uniqueid like '%que-rico%' and node_param = "ip";
}