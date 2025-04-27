package collector

import (
  "os/exec"

  ndb "github.com/OwlH-net/OwlH-Master/database"
  "github.com/OwlH-net/OwlH-Master/nodeclient"
  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
)

func PlayCollector(uuid string) (err error) {
  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ip, port, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
    return err
  }

  err = nodeclient.PlayCollector(ip, port)
  if err != nil {
    logs.Error("nodeclient.playCollector ERROR connection through http new Request: " + err.Error())
    return err
  }
  return nil
}

func StopCollector(uuid string) (err error) {
  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return err
  }
  ip, port, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
    return err
  }
  err = nodeclient.StopCollector(ip, port)
  if err != nil {
    logs.Error("nodeclient.StopCollector ERROR connection through http new Request: " + err.Error())
    return err
  }
  return nil
}

func ShowCollector(uuid string) (data string, err error) {
  err = ndb.GetTokenByUuid(uuid)
  if err != nil {
    logs.Error("Error loading node token: %s", err)
    return "", err
  }
  ip, port, err := ndb.ObtainPortIp(uuid)
  if err != nil {
    logs.Error("node/GetAllFiles ERROR getting node port/ip : " + err.Error())
    return data, err
  }
  data, err = nodeclient.ShowCollector(ip, port)
  if err != nil {
    logs.Error("nodeclient.ShowCollector ERROR connection through http new Request: " + err.Error())
    return "", err
  }

  return data, nil
}

func PlayMasterCollector() (err error) {
  cmd, err := utils.GetKeyValueString("execute", "command")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }
  param, err := utils.GetKeyValueString("execute", "param")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }
  list, err := utils.GetKeyValueString("execute", "list")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }

  _, err = exec.Command(cmd, param, list).Output()
  if err != nil {
    logs.Error("PlayMasterCollector Error executing command in StopCollector function: " + err.Error())
    return err
  }
  return nil
}

func StopMasterCollector() (err error) {
  cmd, err := utils.GetKeyValueString("execute", "command")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }
  param, err := utils.GetKeyValueString("execute", "param")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }
  list, err := utils.GetKeyValueString("execute", "list")
  if err != nil {
    logs.Error("PlayMasterCollector Error getting data from main.conf: " + err.Error())
    return err
  }

  _, err = exec.Command(cmd, param, list).Output()
  if err != nil {
    logs.Error("StopMasterCollector Error executing command in StopCollector function: " + err.Error())
    return err
  }
  return nil
}

func ShowMasterCollector() (data string, err error) {
  status, err := utils.GetKeyValueString("stapCollector", "status")
  if err != nil {
    logs.Error("ShowMasterCollector Error getting data from main.conf: " + err.Error())
    return "", err
  }
  param, err := utils.GetKeyValueString("execute", "param")
  if err != nil {
    logs.Error("ShowMasterCollector Error getting data from main.conf: " + err.Error())
    return "", err
  }
  command, err := utils.GetKeyValueString("execute", "command")
  if err != nil {
    logs.Error("ShowMasterCollector Error getting data from main.conf: " + err.Error())
    return "", err
  }

  output, err := exec.Command(command, param, status).Output()
  if err != nil {
    logs.Error("ShowMasterCollector Error executing command in ShowCollector function: " + err.Error())
    return "", err
  }

  return string(output), nil
}
