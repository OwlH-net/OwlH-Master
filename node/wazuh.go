package node

import (
    "errors"
    "fmt"
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "owlhmaster/nodeclient"
    "owlhmaster/utils"
)

func Wazuh(n string) (data utils.WazuhDetails, err error) {
    err = ndb.GetTokenByUuid(n)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return data, err
    }
    ip, port, err := ndb.ObtainPortIp(n)
    if err != nil {
        logs.Info("Wazuh - get IP and PORT Error -> %s", err.Error())
        return data, err
    }
    logs.Info("Wazuh IP and PORT -> %s, %s", ip, port)
    data, err = nodeclient.Wazuh(ip, port)
    if err != nil {
        return data, err
    }
    return data, nil
}

func RunWazuh(uuid string) (data string, err error) {
    if ndb.Db == nil {
        logs.Error("RunWazuh -- Can't acces to database")
        return "", errors.New("RunWazuh -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return "", err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/RunWazuh ERROR Obtaining Port and Ip: " + err.Error())
        return "", err
    }
    data, err = nodeclient.RunWazuh(ipnid, portnid)
    if err != nil {
        logs.Error("node/RunWazuh ERROR http data request: " + err.Error())
        return "", err
    }
    return data, nil
}

func StopWazuh(uuid string) (data string, err error) {
    if ndb.Db == nil {
        logs.Error("StopWazuh -- Can't acces to database")
        return "", errors.New("StopWazuh -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return "", err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/StopWazuh ERROR Obtaining Port and Ip: " + err.Error())
        return "", err
    }

    data, err = nodeclient.StopWazuh(ipnid, portnid)
    if err != nil {
        logs.Error("node/StopWazuh ERROR http data request: " + err.Error())
        return "", err
    }
    return data, nil
}

func PingWazuhFiles(uuid string) (anode map[string]map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("PingWazuhFiles -- Can't acces to database")
        return nil, errors.New("PingWazuhFiles -- Can't acces to database")
    }

    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("node/PingWazuhFiles ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    anode, err = nodeclient.PingWazuhFiles(ipnid, portnid)
    if err != nil {
        logs.Error("node/PingWazuhFiles ERROR http data request: " + err.Error())
        return nil, err
    }
    return anode, nil
}

func DeleteWazuhFile(anode map[string]interface{}) (err error) {
    if ndb.Db == nil {
        logs.Error("DeleteWazuhFile Error -- Can't acces to database: ")
        return errors.New("DeleteWazuhFile -- Can't acces to database")
    }

    var uuid = fmt.Sprintf("%v", anode["uuid"]) //get strinmg from interface map
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("DeleteWazuhFile ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.DeleteWazuhFile(ipnid, portnid, anode)
    if err != nil {
        logs.Error("DeleteWazuhFile error HTTP data request: " + err.Error())
        return err
    }

    return nil
}

func AddWazuhFile(anode map[string]interface{}) (err error) {
    if ndb.Db == nil {
        logs.Error("AddWazuhFile Error -- Can't acces to database: ")
        return errors.New("AddWazuhFile -- Can't acces to database")
    }

    var uuid = fmt.Sprintf("%v", anode["uuid"])
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("AddWazuhFile ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.AddWazuhFile(ipnid, portnid, anode)
    if err != nil {
        logs.Error("AddWazuhFile error HTTP data request: " + err.Error())
        return err
    }

    return nil
}

func LoadFileLastLines(anode map[string]string) (data map[string]string, err error) {
    if ndb.Db == nil {
        logs.Error("LoadFileLastLines Error -- Can't acces to database: ")
        return nil, errors.New("LoadFileLastLines -- Can't acces to database")
    }

    var uuid = fmt.Sprintf("%v", anode["uuid"])
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return nil, err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("LoadFileLastLines ERROR Obtaining Port and Ip: " + err.Error())
        return nil, err
    }

    data, err = nodeclient.LoadFileLastLines(ipnid, portnid, anode)
    if err != nil {
        logs.Error("LoadFileLastLines error HTTP data request: " + err.Error())
        return nil, err
    }

    return data, nil
}

func SaveFileContentWazuh(anode map[string]string) (err error) {
    if ndb.Db == nil {
        logs.Error("SaveFileContentWazuh Error -- Can't acces to database: ")
        return errors.New("SaveFileContentWazuh -- Can't acces to database")
    }

    var uuid = fmt.Sprintf("%v", anode["uuid"])
    err = ndb.GetTokenByUuid(uuid)
    if err != nil {
        logs.Error("Error loading node token: %s", err)
        return err
    }
    ipnid, portnid, err := ndb.ObtainPortIp(uuid)
    if err != nil {
        logs.Error("SaveFileContentWazuh ERROR Obtaining Port and Ip: " + err.Error())
        return err
    }

    err = nodeclient.SaveFileContentWazuh(ipnid, portnid, anode)
    if err != nil {
        logs.Error("SaveFileContentWazuh error HTTP data request: " + err.Error())
        return err
    }

    return nil
}
