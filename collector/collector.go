package collector

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/nodeclient"
    "owlhmaster/node"
)

func PlayCollector(uuid string) (err error) {
	ip, err := node.GetNodeIPbyUUID(uuid)
    if err != nil {
        logs.Info("playCollector - IP Error -> %s", err.Error())
        return err
	}
    port, err := node.GetNodePortbyUUID(uuid)
    if err != nil {
        logs.Info("playCollector - PORT Error -> %s", err.Error())
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
	ip, err := node.GetNodeIPbyUUID(uuid)
    if err != nil {
        logs.Info("StopCollector - IP Error -> %s", err.Error())
        return err
	}
    port, err := node.GetNodePortbyUUID(uuid)
    if err != nil {
        logs.Info("StopCollector - PORT Error -> %s", err.Error())
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
	ip, err := node.GetNodeIPbyUUID(uuid)
    if err != nil {
        logs.Info("ShowCollector - IP Error -> %s", err.Error())
        return "",err
	}
    port, err := node.GetNodePortbyUUID(uuid)
    if err != nil {
        logs.Info("ShowCollector - PORT Error -> %s", err.Error())
        return "",err
	}    
	data, err = nodeclient.ShowCollector(ip,port)
	if err != nil {
		logs.Error("nodeclient.ShowCollector ERROR connection through http new Request: "+err.Error())
		return "",err
	}
	
	return data,nil
}