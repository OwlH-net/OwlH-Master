package collector

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/nodeclient"
    "owlhmaster/database"
)

func PlayCollector(uuid string) (err error) {
	ip,port,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return err
	}

	err = nodeclient.PlayCollector(ip,port)
	if err != nil {
		logs.Error("nodeclient.playCollector ERROR connection through http new Request: "+err.Error())
		return err
	}
	return nil
}

func StopCollector(uuid string) (err error) {
	ip,port,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return err
	}   
	err = nodeclient.StopCollector(ip,port)
	if err != nil {
		logs.Error("nodeclient.StopCollector ERROR connection through http new Request: "+err.Error())
		return err
	}
	return nil
}

func ShowCollector(uuid string) (data string, err error) {
	ip,port,err := ndb.ObtainPortIp(uuid)
	if err != nil {
		logs.Error("node/GetAllFiles ERROR getting node port/ip : "+err.Error())
        return data, err
	}   
	data, err = nodeclient.ShowCollector(ip,port)
	if err != nil {
		logs.Error("nodeclient.ShowCollector ERROR connection through http new Request: "+err.Error())
		return "",err
	}
	
	return data,nil
}