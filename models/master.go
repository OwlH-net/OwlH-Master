package models 

import (
    "owlhmaster/master"
    "owlhmaster/changeControl"
    "owlhmaster/database"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/getMasterTitle \
// }
func GetMasterTitle() (data string, err error) {
    data, err = master.GetMasterTitle()
    changecontrol.ChangeControlInsertData(err, "GetMasterTitle")
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/editFile/:uuid \
// }
func GetFileContent(file string) (data map[string]string, err error) {
    data, err = master.GetFileContent(file)
    changecontrol.ChangeControlInsertData(err, "GetFileContent")
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/savefile \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "file": "v",
//     "content": "v"
//  }
// }
func SaveFileContent(anode map[string]string) (err error) {
    err = master.SaveFileContent(anode)
    changecontrol.ChangeControlInsertData(err, "SaveFileContent")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingPlugins \
// }
func PingPlugins() (data map[string]map[string]string, err error) {
    data,err = master.PingPlugins()
    changecontrol.ChangeControlInsertData(err, "PingPlugins")
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingFlow \
// }
func PingFlow() (data map[string]map[string]string, err error) {
    data,err = ndb.PingFlow()
    changecontrol.ChangeControlInsertData(err, "PingFlow")
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/changePluginStatus \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v"
//  }
// }
func ChangePluginStatus(anode map[string]string) (err error) {
    err = ndb.UpdatePluginValueMaster(anode["uuid"], anode["param"], anode["value"])
    changecontrol.ChangeControlInsertData(err, "ChangePluginStatus")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/saveStapInterface \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v"
//     "value": "v"
//  }
// }
func SaveStapInterface(anode map[string]string) (err error) {
    err = master.SaveStapInterface(anode)
    changecontrol.ChangeControlInsertData(err, "SaveStapInterface")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/interface \
// }
func GetNetworkInterface()(data map[string]string, err error) {
    data,err = master.GetNetworkInterface()
    changecontrol.ChangeControlInsertData(err, "GetNetworkInterface")
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/changeDataflowStatus \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "value": "v"
//  }
// }
func ChangeDataflowStatus(anode map[string]string) (err error) {
    err = ndb.ChangeDataflowStatus(anode)
    changecontrol.ChangeControlInsertData(err, "ChangeDataflowStatus")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/deployMaster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "value": "v"
//  }
// }
func DeployMaster(anode map[string]string)(err error) {
    err = master.DeployMaster(anode)
    changecontrol.ChangeControlInsertData(err, "DeployMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/updateMasterNetworkInterface \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "value": "v"
//  }
// }
func UpdateMasterNetworkInterface(anode map[string]string)(err error) {
    err = master.UpdateMasterNetworkInterface(anode)
    changecontrol.ChangeControlInsertData(err, "UpdateMasterNetworkInterface")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/loadMasterNetworkValuesSelected \
// }
func LoadMasterNetworkValuesSelected()(data map[string]map[string]string ,err error) {
    data,err = master.LoadMasterNetworkValuesSelected()
    changecontrol.ChangeControlInsertData(err, "LoadMasterNetworkValuesSelected")
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingservice \
// }
func PingServiceMaster()(err error) {
    err = master.PingServiceMaster()
    changecontrol.ChangeControlInsertData(err, "PingServiceMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/deployservice \
// }
func DeployServiceMaster()(err error) {
    err = master.DeployServiceMaster()
    changecontrol.ChangeControlInsertData(err, "DeployServiceMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/add \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "type": "v",
//     "cert": "v",
//     "port": "v",
//     "interface": "v"
//  }
// }
func AddPluginServiceMaster(anode map[string]string)(err error){
    err = master.AddPluginServiceMaster(anode)
    changecontrol.ChangeControlInsertData(err, "AddPluginServiceMaster")
    return err
}

// curl -X DELETE \
//   https://52.47.197.22:50002/v1/master/deleteService \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
//  }
// }
func DeleteServiceMaster(anode map[string]string)(err error){
    err = master.DeleteServiceMaster(anode)
    changecontrol.ChangeControlInsertData(err, "DeleteServiceMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/modifyStapValues \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "name": "v",
//     "type": "v",
//     "cert": "v",
//     "port": "v",
//     "interface": "v"
//  }
// }
func ModifyStapValuesMaster(anode map[string]string)(err error){
    err = master.ModifyStapValuesMaster(anode)
    changecontrol.ChangeControlInsertData(err, "ModifyStapValuesMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/updateMasterStapInterface \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "value": "v"
//  }
// }
func UpdateMasterStapInterface(anode map[string]string)(err error){
    err = master.UpdateMasterStapInterface(anode)
    changecontrol.ChangeControlInsertData(err, "UpdateMasterStapInterface")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/setbpf \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "value": "v"
//  }
// }
func SetBPF(anode map[string]string)(err error){
    err = master.SetBPF(anode)
    changecontrol.ChangeControlInsertData(err, "SetBPF")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/deployStapServiceMaster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v",
//     "port": "v",
//     "interface": "v",
//     "collector": "v"
//  }
// }
func DeployStapServiceMaster(anode map[string]string)(err error){
    err = master.DeployStapServiceMaster(anode)
    changecontrol.ChangeControlInsertData(err, "DeployStapServiceMaster")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/stopStapServiceMaster \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "type": "v"
//  }
// }
func StopStapServiceMaster(anode map[string]string)(err error){
    err = master.StopStapServiceMaster(anode)
    changecontrol.ChangeControlInsertData(err, "StopStapServiceMaster")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/incidents \
// }
func GetIncidents()(data map[string]map[string]string, err error){
    data,err = master.GetIncidents()
    changecontrol.ChangeControlInsertData(err, "GetIncidents")
    return data,err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/master/incidents \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "value": "v",
// }
func PutIncident(anode map[string]string)(err error){
    err = master.PutIncident(anode)
    changecontrol.ChangeControlInsertData(err, "PutIncident")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/zeek/saveZeekValues \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "param": "v",
//     "node": "v"
// }
func SaveZeekValues(anode map[string]string)(err error){
    err = master.SaveZeekValues(anode)
    changecontrol.ChangeControlInsertData(err, "SaveZeekValues")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/incidents \
// }
func PingPluginsMaster()(data map[string]map[string]string, err error){
    data,err = master.PingPluginsMaster()
    changecontrol.ChangeControlInsertData(err, "PingPluginsMaster")
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/editPathFile/:path \
// }
func GetPathFileContent(param string) (data map[string]string, err error) {
    data, err = master.GetPathFileContent(param)
    changecontrol.ChangeControlInsertData(err, "GetPathFileContent")
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/savefilePath \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "file": "v",
//     "content": "v"
//  }
// }
func SaveFilePathContent(anode map[string]string) (err error) {
    err = master.SaveFilePathContent(anode)
    changecontrol.ChangeControlInsertData(err, "SaveFilePathContent")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/login \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "user": "v",
//     "password": "v"
//  }
// }
func Login(anode map[string]string) (err error) {
    err = master.Login(anode)
    changecontrol.ChangeControlInsertData(err, "ChangePluginStatus")
    return err
}