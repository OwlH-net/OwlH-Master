package models 

import (
    "owlhmaster/master"
    "owlhmaster/database"
)

func GetMasterTitle() (data string, err error) {
    data, err = master.GetMasterTitle()
    return data, err
}

func GetFileContent(file string) (data map[string]string, err error) {
    data, err = master.GetFileContent(file)
    return data, err
}

func SaveFileContent(anode map[string]string) (err error) {
    err = master.SaveFileContent(anode)
    return err
}

func PingPlugins() (data map[string]map[string]string, err error) {
    data,err = ndb.PingPlugins()
    return data,err
}

func PingFlow() (data map[string]map[string]string, err error) {
    data,err = ndb.PingFlow()
    return data,err
}

func ChangePluginStatus(anode map[string]string) (err error) {
    err = ndb.UpdatePluginValueMaster(anode["uuid"], anode["param"], anode["value"])
    return err
}

func SaveStapInterface(anode map[string]string) (err error) {
    err = master.SaveStapInterface(anode)
    return err
}

func GetNetworkInterface()(data map[string]string, err error) {
	data,err = master.GetNetworkInterface()
	return data,err
}

func ChangeDataflowStatus(anode map[string]string) (err error) {
    err = ndb.ChangeDataflowStatus(anode)
    return err
}

func DeployMaster(anode map[string]string)(err error) {
    err = master.DeployMaster(anode)
    return err
}

func UpdateMasterNetworkInterface(anode map[string]string)(err error) {
    err = master.UpdateMasterNetworkInterface(anode)
    return err
}

func LoadMasterNetworkValuesSelected()(data map[string]map[string]string ,err error) {
    data,err = master.LoadMasterNetworkValuesSelected()
    return data,err
}

func PingServiceMaster()(err error) {
    err = master.PingServiceMaster()
    return err
}

func DeployServiceMaster()(err error) {
    err = master.DeployServiceMaster()
    return err
}

func AddPluginServiceMaster(anode map[string]string)(err error){
    err = master.AddPluginServiceMaster(anode)
    return err
}

func DeleteServiceMaster(anode map[string]string)(err error){
    err = master.DeleteServiceMaster(anode)
    return err
}

func ModifyStapValuesMaster(anode map[string]string)(err error){
    err = master.ModifyStapValuesMaster(anode)
    return err
}

func UpdateMasterStapInterface(anode map[string]string)(err error){
    err = master.UpdateMasterStapInterface(anode)
    return err
}

func SetBPF(anode map[string]string)(err error){
    err = master.SetBPF(anode)
    return err
}

func DeployStapServiceMaster(anode map[string]string)(err error){
    err = master.DeployStapServiceMaster(anode)
    return err
}

func StopStapServiceMaster(anode map[string]string)(err error){
    err = master.StopStapServiceMaster(anode)
    return err
}