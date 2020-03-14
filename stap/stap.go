package stap

import(
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/nodeclient"
)

//Add server to software TAP
func AddServer(data map[string]string)(err error) {
    uuid := data["uuid"]
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ipuuid,portuuid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("AddServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return err
    }
    err = nodeclient.AddServer(ipuuid,portuuid, data)
    if err != nil {
        logs.Error("node/AddServer ERROR http data request: "+err.Error())
        return err
    }
    return nil
}

//Get all STAP servers
func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error){
    err = ndb.GetTokenByUuid(nodeuuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipuuid,portuuid,err := ndb.ObtainPortIp(nodeuuid)
    if err != nil {
        logs.Error("GetAllServers ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return nil,err
    }
    data, err = nodeclient.GetAllServers(ipuuid,portuuid)
    if err != nil {
        logs.Error("node/GetAllServers ERROR http data request: "+err.Error())
        return nil, err
    }
    return data,nil
}

//Get a specific server
func GetServer(uuid string, serveruuid string)(data map[string]map[string]string, err error){
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ipuuid,portuuid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("GetServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return nil,err
    }
    data, err = nodeclient.GetServer(ipuuid,portuuid, serveruuid)
    if err != nil {
        logs.Error("node/GetServer ERROR http data request: "+err.Error())
        return nil, err
    }
    return data,nil
}

//ping to Stap function at node.html. Create or update this function if is needed
func Stap(n string) (data map[string]bool, err error) {
    err = ndb.GetTokenByUuid(n); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
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
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("RunStap ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    data, err = nodeclient.RunStap(ipnid,portnid,uuid)
    if err != nil {
        logs.Error("Stap run ERROR: "+err.Error())
        return "",err
    }
    return data,nil
}

//Stop stap main server
func StopStap(uuid string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStap -- Can't acces to database")
        return "", errors.New("StopStap -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("StopStap ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    data, err = nodeclient.StopStap(ipnid,portnid,uuid)
    if err != nil {
        logs.Error("Stap stop ERROR: "+err.Error())
        return "",err
    }
    return data,nil
}

//Launch stap specific server
func RunStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("RunStapServer -- Can't acces to database")
        return "", errors.New("RunStapServer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("RunStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    data, err = nodeclient.RunStapServer(ipnid,portnid,server)
    if err != nil {
        logs.Error("RunStapServer ERROR: "+err.Error())
        return "",err
    }
    return data,nil
}

//Stop stap specific server
func StopStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("StopStapServer -- Can't acces to database")
        return "", errors.New("StopStapServer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("StopStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    data, err = nodeclient.StopStapServer(ipnid,portnid,server)
    if err != nil {
        logs.Error("StopStapServer ERROR: "+err.Error())
        return "",err
    }
    return data,nil
}

//Delete specific stap server
func DeleteStapServer(uuid string, server string)(data string, err error){
    if ndb.Db == nil {
        logs.Error("DeleteStapServer -- Can't acces to database")
        return "", errors.New("DeleteStapServer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return "",err}
    ipnid,portnid,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("DeleteStapServer ERROR Obtaining Port and IP for Add a new server into STAP: "+err.Error())
        return "",err
    }
    data, err = nodeclient.DeleteStapServer(ipnid,portnid,server)
    if err != nil {
        logs.Error("StopStapServer ERROR: "+err.Error())
        return "",err
    }
    return data,nil
}

func PingServerStap(uuid string, server string) (data map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("DeleteStapServer -- Can't acces to database")
        return nil, errors.New("DeleteStapServer -- Can't acces to database")
    }
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return nil,err}
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Info("PingServerStap - get IP and PORT Error -> %s", err.Error())
        return nil,err
    }    
    data, err = nodeclient.PingServerStap(ip,port,server)
    if err != nil {
        logs.Error("PingServerStap ERROR: "+err.Error())
        return nil,err
    }
    return data,nil
}

func EditStapServer(data map[string]string) (err error) {
    uuid := data["uuid"]
    err = ndb.GetTokenByUuid(uuid); if err!=nil{logs.Error("Error loading node token: %s",err); return err}
    ip,port,err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Info("EditStapServer - get IP and PORT Error -> %s", err.Error())
        return err
    }    
    err = nodeclient.EditStapServer(ip,port,data)
    if err != nil {
        logs.Error("EditStapServer ERROR: "+err.Error())
        return err
    }
    return nil
}