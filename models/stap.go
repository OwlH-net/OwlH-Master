package models 

import (
    "owlhmaster/stap"
)

// curl -X POST \
//   https://52.47.197.22:50002/v1/stap/ \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "nodeName": "d",
//     "nodeIp": "d",
// }
func AddServer(data map[string]string)(err error) {
    err = stap.AddServer(data)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/:uuid \
func GetAllServers(nodeuuid string)(data map[string]map[string]string, err error) {
    data,err = stap.GetAllServers(nodeuuid)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/server/:uuid/:serveruuid \
func GetServer(uuid string, serveruuid string)(data map[string]map[string]string, err error) {
    data,err = stap.GetServer(uuid, serveruuid)
    return data,err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/stap/stap/:uuid  \
func Stap(n string) (data map[string]bool, err error) {
    data,err = stap.Stap(n)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/RunStap/:uuid \
func RunStap(n string) (data string, err error) {
    data,err = stap.RunStap(n)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStap/:uuid \
func StopStap(n string) (data string, err error) {
    data,err = stap.StopStap(n)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/RunStapServer/:uuid/:server \
func RunStapServer(uuid string, server string) (data string, err error) {
    data,err = stap.RunStapServer(uuid, server)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStapServer/:uuid/:server \
func StopStapServer(uuid string, server string) (data string, err error) {
    data,err = stap.StopStapServer(uuid, server)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/StopStapServer/:uuid/:server \
func PingServerStap(uuid string, server string) (data map[string]string, err error) {
    data,err = stap.PingServerStap(uuid, server)
    return data,err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/stap/EditStapServer/:uuid/:server \
func DeleteStapServer(uuid string, server string) (data string, err error) {
    data,err = stap.DeleteStapServer(uuid, server)
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
func EditStapServer(data map[string]string) (err error) {
    err = stap.EditStapServer(data)
    return err
}