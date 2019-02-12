package node

import (
    "github.com/astaxie/beego/logs"
    "strings"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
    "regexp"
)

func findNode (s string) (id string, err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return "", errors.New("no hemos podido acceder a la bbdd")
    }
    ip, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
    uuid, _ := regexp.Compile(`\w{8}-\w{4}-\w{4}-\w{4}-\w{12}`)
    sql := "SELECT node_uniqueid FROM nodes where node_param = 'name' and node_value='"+s+"';"
    if ip.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'ip' and node_value='"+s+"';"
    } else if uuid.MatchString(s) {
        sql = "SELECT node_uniqueid FROM nodes where node_param = 'UUID' and node_value='"+s+"';"
    }
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&id); err == nil {
            return id, nil
        }
    }
    return "", err
}

func DeleteNode(nodeid string)(err error) {
    logs.Info("NODE Delete -> IN")
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return errors.New("no hemos podido acceder a la bbdd")
    }
    stmt, err := ndb.Db.Prepare("delete from nodes where node_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nodeid)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}

func getNodeIPbyUID (nk string) (ip string, err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return "", errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "SELECT node_value FROM nodes where node_param = 'ip' and node_uniqueid='"+nk+"';"
    logs.Info("GetNodeIP -> SQL -> %s", sql)
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&ip); err == nil {
            return ip, nil
        }
    }
    return "", err
}

func getNodeConf (nodeKey string)(conf map[string]string, err error) {
    var param string
    var value string

    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return nil, errors.New("no hemos podido acceder a la bbdd")
    }
    
    sql := "SELECT node_param, node_value FROM nodes where node_uniqueid='"+nodeKey+"';"
    logs.Info("GetNodeConf -> SQL -> %s", sql)
    
    rows, err := ndb.Db.Query(sql)
    
    if err != nil {
        logs.Error(err.Error())
        return nil, err
    }
    
    defer rows.Close()
    for rows.Next() {
        if err = rows.Scan(&param, &value); err != nil {
            logs.Info (" Error en el scan -> %s",err.Error())
            continue
        }
        conf[param]=value
    }
    return conf, nil
}

func getNodePortbyUID (nk string) (port string, err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return "", errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "SELECT node_value FROM nodes where node_param = 'port' and node_uniqueid='"+nk+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&port); err == nil {
            return port, nil
        }
    }
    return "", err
}

func getAllNodesIp () (ips map[string]string, err error) {
    var uid string
    var ip string
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return ips, errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "SELECT node_uniqueid, node_value FROM nodes where node_param = 'ip';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error("Error al ejecutar la query %s -> %s", sql, err.Error())
        return ips, err
    }
    defer rows.Close()
    for rows.Next() {
        if err = rows.Scan(&uid, &ip); err != nil {
            logs.Info (" Error en el scan -> %s",err.Error())
        }
        ips[uid]=ip
    }
    return ips, nil
}

func nodeKeyExists (nodekey string, key string) (id int, err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return 0, errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "SELECT node_id FROM nodes where node_uniqueid = '"+nodekey+"' and node_param = '"+key+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return 0, err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&id); err == nil {
            return id, err
        }
    }
    return 0, nil
}

func nodeExists (nodeid string) (err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "SELECT * FROM nodes where node_uniqueid = '"+nodeid+"';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return errors.New("ya existe el node buscado")
    } else {
        return nil
    }
}

func nodeKeyUpdate(id int, nkey string, key string, value string) (err error) {
    logs.Info("NODE Key Insert -> IN")
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return errors.New("no hemos podido acceder a la bbdd")
    }
    logs.Info("nkey: %s, key: %s, value: %s", nkey, key, value)
    stmt, err := ndb.Db.Prepare("update nodes set node_param = ?, node_value = ? where node_id = ? and node_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&key, &value, &id, &nkey)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}

func nodeKeyInsert(nkey string, key string, value string) (err error) {
    logs.Info("NODE Insert -> IN")
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return errors.New("no hemos podido acceder a la bbdd")
    }
    logs.Info("nkey: %s, key: %s, value: %s", nkey, key, value)
    stmt, err := ndb.Db.Prepare("insert into nodes (node_uniqueid, node_param, node_value) values(?,?,?)")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nkey, &key, &value)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}

func AddNode (n map[string]string) (err error) {
    logs.Info("ADD NODE -> IN")
    var nodeKey string
    if _, ok := n["name"]; !ok {
        return errors.New("name está vacio")
    }
    if _, ok := n["ip"]; !ok {
        return errors.New("ip está vacio")
    }
    if _, ok := n["id"]; !ok {
        nodeKey = strings.Replace(n["name"], " ", "-",0)+"-"+strings.Replace(n["ip"], ".", "-",0)
    } else {
        nodeKey = n["id"]
    }
    if err := nodeExists(nodeKey); err != nil {
        return err
    }
    
    for key, value := range n {
        err = nodeKeyInsert(nodeKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func UpdateNode (n map[string]string) (err error) {
    logs.Info("UPDATE NODE -> IN name es -  %s", n["name"])
    var nodeKey string

    if _, ok := n["name"]; !ok {
        return errors.New("name está vacio")
    }
    if _, ok := n["ip"]; !ok {
        return errors.New("ip está vacio")
    }
    if _, ok := n["id"]; !ok {
        nodeKey = strings.Replace(n["name"], " ", "-",0)+"-"+strings.Replace(n["ip"], ".", "-",0)
    } else {
        nodeKey = n["id"]
    }
    if err := nodeExists(nodeKey); err == nil {
        return errors.New("El nodo no existe. Hay que crearlo antes")
    }
    for key, value := range n {
        if id, _ := nodeKeyExists(nodeKey, key); id != 0 {
            err = nodeKeyUpdate(id, nodeKey, key, value)
        } else {
            err = nodeKeyInsert(nodeKey, key, value)
        }
    }
    if err != nil {
        return err
    }
    return nil
}

func getNodeIpbyName(n string)(ip string, err error) {
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return "", errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "select node_value from nodes where node_uniqueid like '%"+n+"%' and node_param = 'ip';"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        if err = rows.Scan(&ip); err == nil {
            return ip, err
        }
    }
    return "", errors.New("no hay IP asociada al nombre")
    //select node_value from nodes where node_uniqueid like '%que-rico%' and node_param = "ip";
}

func GetAllNodes() (nodes *map[string]map[string]string, err error) {
    logs.Info("Node - getAllNodes - IN")
    var allnodes = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return nil, errors.New("no hemos podido acceder a la bbdd")
    }
    sql := "select node_uniqueid, node_param, node_value from nodes;"
    rows, err := ndb.Db.Query(sql)
    if err != nil {
        logs.Error("Error al ejecutar la query: %s", err.Error())
        return nil, err
    }
    for rows.Next() {
        if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("no hemos podido leer del resultado de la query: %s", err.Error())
            return nil, err
        }
        logs.Info ("uniqid: %s, param: %s, value: %s", uniqid,param,value)
        if allnodes[uniqid] == nil { allnodes[uniqid] = map[string]string{}}
        allnodes[uniqid][param]=value
        logs.Info ("GET all nodes -> vamos a por otro")
    } 
    return &allnodes, nil
}

func NodePing (n string) (err error) {
    logs.Info("Node PING -> IN")
    logs.Info("Ping - UID -> %s", n)
//    uid, err := findNode(n)
//    if err != nil {
//        logs.Info("Ping - UID Error -> %s", err.Error())
//        return err
//    }
    ip, err := getNodeIPbyUID(n)
    if err != nil {
        logs.Info("Ping - IP Error -> %s", err.Error())
        return err
    }
    port, err := getNodePortbyUID(n)
    if err != nil {
        logs.Info("Ping - PORT Error -> %s", err.Error())
        return err
    }    
    logs.Info("Ping - vamos a por el ping -> %s, %s", ip, port)
    err = nodeclient.PingNode(ip,port)
    if err != nil {
        return err
    }
    return nil
}

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

func Zeek (n string) (data []byte, err error) {
    logs.Info("Node Zeek -> IN")
    logs.Info("Zeek - UID -> %s", n)

    ip, err := getNodeIPbyUID(n)
    if err != nil {
        logs.Info("Zeek - IP Error -> %s", err.Error())
        return nil,err
    }
    port, err := getNodePortbyUID(n)
    if err != nil {
        logs.Info("Zeek - PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Zeek - vamos a por el zeek -> %s, %s", ip, port)
    data, err = nodeclient.Zeek(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func Wazuh (n string) (data []byte, err error) {
    logs.Info("Node wazuh -> IN")
    logs.Info("Wazuh - UID -> %s", n)

    ip, err := getNodeIPbyUID(n)
    if err != nil {
        logs.Info("Wazuh - IP Error -> %s", err.Error())
        return nil,err
    }
    port, err := getNodePortbyUID(n)
    if err != nil {
        logs.Info("Wazuh - PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Wazuh - vamos a por el wazuh -> %s, %s", ip, port)
    data, err = nodeclient.Wazuh(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func SuricataBPF(n string)(bpf string, err error) {
    if ndb.Db == nil {
        logs.Error("suricataBPF -- Can't acces to database")
        return "", errors.New("suricataBPF -- Can't acces to database")
    }
    var res string
    sql := "select node_value from nodes where node_uniqueid = \""+n+"\" and node_param = \"bpf\";"
    logs.Info("BPF Suricata query sql %s",sql)
    rows, err := ndb.Db.Query(sql)
    
    if err != nil {
        logs.Info("BPF Suricata query error %s",err.Error())
        return "", err
    }
    defer rows.Close()
    if rows.Next() {
        err = rows.Scan(&res)
        logs.Info("BPF Suricata there are rows")
        if  err != nil {
            logs.Info("BPF Suricata scan error %s",err.Error())
            return "", err
        }
        logs.Info("BPF Suricata res: "+res)
        return res, err
    }
    return "", errors.New("suricataBPF -- There is no defined BPF")
    //select node_value from nodes where node_uniqueid like '%que-rico%' and node_param = "ip";
}