package models 

import (
    "owlhmaster/master"
    "owlhmaster/database"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/getMasterTitle \
// }
func GetMasterTitle() (data string, err error) {
    data, err = master.GetMasterTitle()
    return data, err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/editFile/:uuid \
// }
func GetFileContent(file string) (data map[string]string, err error) {
    data, err = master.GetFileContent(file)
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
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingPlugins \
// }
func PingPlugins() (data map[string]map[string]string, err error) {
    data,err = ndb.PingPlugins()
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingFlow \
// }
func PingFlow() (data map[string]map[string]string, err error) {
    data,err = ndb.PingFlow()
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
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/interface \
// }
func GetNetworkInterface()(data map[string]string, err error) {
    data,err = master.GetNetworkInterface()
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
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/loadMasterNetworkValuesSelected \
// }
func LoadMasterNetworkValuesSelected()(data map[string]map[string]string ,err error) {
    data,err = master.LoadMasterNetworkValuesSelected()
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/pingservice \
// }
func PingServiceMaster()(err error) {
    err = master.PingServiceMaster()
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/master/deployservice \
// }
func DeployServiceMaster()(err error) {
    err = master.DeployServiceMaster()
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
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/master/incidents \
// }
func GetIncidents()(data map[string]map[string]string, err error){
    data,err = master.GetIncidents()
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
    return err
}