package node

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
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
	if err != nil {
		logs.Error("node/RunWazuh ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
	data,err = nodeclient.RunWazuh(ipnid,portnid)
	if err != nil {
		logs.Error("node/RunWazuh ERROR http data request: "+err.Error())
        return "",err
    }
	return data,nil
}

func StopWazuh(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopWazuh -- Can't acces to database")
        return "", errors.New("StopWazuh -- Can't acces to database")
	}
	
	ipnid,portnid,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/StopWazuh ERROR Obtaining Port and Ip: "+err.Error())
        return "",err
    }
	data,err = nodeclient.StopWazuh(ipnid,portnid)
	if err != nil {
		logs.Error("node/StopWazuh ERROR http data request: "+err.Error())
        return "",err
    }
	return data,nil
}