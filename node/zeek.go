package node

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
)

func Zeek(n string) (data map[string]bool, err error) {
    ip,port,err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Info("Zeek - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    logs.Info("Zeek IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Zeek(ip,port)
    if err != nil {
        return nil,err
    }
    return data,nil
}

func RunZeek(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunZeek -- Can't acces to database")
        return "", errors.New("RunZeek -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/RunZeek ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
	data,err = nodeclient.RunZeek(ipnid,portnid)
	if err != nil {
		logs.Error("node/RunZeek ERROR http data request: "+err.Error())
        return "",err
    }
	return data,nil
}

func StopZeek(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopZeek -- Can't acces to database")
        return "", errors.New("StopZeek -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/StopZeek ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
	data,err = nodeclient.StopZeek(ipnid,portnid)
	if err != nil {
		logs.Error("node/StopZeek ERROR http data request: "+err.Error())
        return "",err
    }
	return data,nil
}

func DeployZeek(uuid string)(err error){
    if ndb.Db == nil {
        logs.Error("DeployZeek -- Can't acces to database")
        return errors.New("DeployZeek -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/DeployZeek ERROR Obtaining Port and Ip: "+err.Error())
        return err
    }
	err = nodeclient.DeployZeek(ipnid,portnid)
	if err != nil {
		logs.Error("node/DeployZeek ERROR http data request: "+err.Error())
        return err
    }
	return nil
}

func ChangeZeekMode(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("ChangeZeekMode -- Can't acces to database"); return err}

	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil { logs.Error("node/ChangeZeekMode ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.ChangeZeekMode(ipnid,portnid,anode)
	if err != nil { logs.Error("node/ChangeZeekMode ERROR http data request: "+err.Error()); return err}

	return nil
}

func AddClusterValue(anode map[string]string)(err error){
    if ndb.Db == nil { logs.Error("AddClusterValue -- Can't acces to database"); return err}

	ipnid,portnid,err := ndb.ObtainPortIp(anode["uuid"])
	if err != nil { logs.Error("node/AddClusterValue ERROR Obtaining Port and Ip: "+err.Error()); return err}
    
    err = nodeclient.AddClusterValue(ipnid,portnid,anode)
	if err != nil { logs.Error("node/AddClusterValue ERROR http data request: "+err.Error()); return err}

	return nil
}