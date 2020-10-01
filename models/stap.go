package models 

import (
    "owlhmaster/stap"
    "owlhmaster/changeControl"
)

// curl -X POST \
//   https://52.47.197.22:50002/v1/stap/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "nodeName": "d",
//     "nodeIp": "d",
// }
func AddServer(data map[string]string, username string)(err error) {
    err = stap.AddServer(data)
    changecontrol.ChangeControlInsertData(err, "AddServer", username)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/:uuid \
func GetAllServers(nodeuuid string, username string)(data map[string]map[string]string, err error) {
    data,err = stap.GetAllServers(nodeuuid)
    changecontrol.ChangeControlInsertData(err, "GetAllServers", username)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/server/:uuid/:serveruuid \
func GetServer(uuid string, serveruuid string, username string)(data map[string]map[string]string, err error) {
    data,err = stap.GetServer(uuid, serveruuid)
    changecontrol.ChangeControlInsertData(err, "GetServer", username)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/stap/:uuid  \
func Stap(n string, username string) (data map[string]bool, err error) {
    data,err = stap.Stap(n)
    changecontrol.ChangeControlInsertData(err, "Stap", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/RunStap/:uuid \
func RunStap(n string, username string) (data string, err error) {
    data,err = stap.RunStap(n)
    changecontrol.ChangeControlInsertData(err, "RunStap", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStap/:uuid \
func StopStap(n string, username string) (data string, err error) {
    data,err = stap.StopStap(n)
    changecontrol.ChangeControlInsertData(err, "StopStap", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/RunStapServer/:uuid/:server \
func RunStapServer(uuid string, server string, username string) (data string, err error) {
    data,err = stap.RunStapServer(uuid, server)
    changecontrol.ChangeControlInsertData(err, "RunStapServer", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStapServer/:uuid/:server \
func StopStapServer(uuid string, server string, username string) (data string, err error) {
    data,err = stap.StopStapServer(uuid, server)
    changecontrol.ChangeControlInsertData(err, "StopStapServer", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStapServer/:uuid/:server \
func PingServerStap(uuid string, server string, username string) (data map[string]string, err error) {
    data,err = stap.PingServerStap(uuid, server)
    changecontrol.ChangeControlInsertData(err, "PingServerStap", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/EditStapServer/:uuid/:server \
func DeleteStapServer(uuid string, server string, username string) (data string, err error) {
    data,err = stap.DeleteStapServer(uuid, server)
    changecontrol.ChangeControlInsertData(err, "DeleteStapServer", username)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/EditStapServer \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "server": "v",
//     "param": "d",
//     "value": "d",
//     "uuid": "d",
// }
func EditStapServer(data map[string]string, username string) (err error) {
    err = stap.EditStapServer(data)
    changecontrol.ChangeControlInsertData(err, "EditStapServer", username)
    return err
}