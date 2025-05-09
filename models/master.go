package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  ndb "github.com/OwlH-net/OwlH-Master/database"
  "github.com/OwlH-net/OwlH-Master/master"
)

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getMasterTitle \
//  }
func GetMasterTitle(username string) (data string, err error) {
  data, err = master.GetMasterTitle()
  changecontrol.ChangeControlInsertData(err, "GetMasterTitle", username)
  return data, err
}

func CheckDefaultAdmin() (isDefault bool, err error) {
  isDefault, err = master.CheckDefaultAdmin()
  return isDefault, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/editFile/:uuid \
//  }
func GetFileContent(file string, username string) (data map[string]string, err error) {
  data, err = master.GetFileContent(file)
  changecontrol.ChangeControlInsertData(err, "GetFileContent", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/GetAllRulesetsForAllNodes \
//  }
func GetAllRulesetsForAllNodes(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetAllRulesetsForAllNodes()
  changecontrol.ChangeControlInsertData(err, "GetAllRulesetsForAllNodes", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/savefile \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "file": "v",
//      "content": "v"
//   }
//  }
func SaveFileContent(anode map[string]string, username string) (err error) {
  err = master.SaveFileContent(anode)
  changecontrol.ChangeControlInsertData(err, "SaveFileContent", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/pingPlugins \
//  }
func PingPlugins(username string) (data map[string]map[string]string, err error) {
  data, err = master.PingPlugins()
  changecontrol.ChangeControlInsertData(err, "PingPlugins", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/pingFlow \
//  }
func PingFlow(username string) (data map[string]map[string]string, err error) {
  data, err = ndb.PingFlow()
  changecontrol.ChangeControlInsertData(err, "PingFlow", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/changePluginStatus \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v"
//   }
//  }
func ChangePluginStatus(anode map[string]string, username string) (err error) {
  err = master.ChangePluginStatus(anode)
  changecontrol.ChangeControlInsertData(err, "ChangePluginStatus", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/saveStapInterface \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v"
//      "value": "v"
//   }
//  }
func SaveStapInterface(anode map[string]string, username string) (err error) {
  err = master.SaveStapInterface(anode)
  changecontrol.ChangeControlInsertData(err, "SaveStapInterface", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/interface \
//  }
func GetNetworkInterface(username string) (data map[string]string, err error) {
  data, err = master.GetNetworkInterface()
  changecontrol.ChangeControlInsertData(err, "GetNetworkInterface", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/changeDataflowStatus \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "value": "v"
//   }
//  }
func ChangeDataflowStatus(anode map[string]string, username string) (err error) {
  err = ndb.ChangeDataflowStatus(anode)
  changecontrol.ChangeControlInsertData(err, "ChangeDataflowStatus", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/deployMaster \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "value": "v"
//   }
//  }
func DeployMaster(anode map[string]string, username string) (err error) {
  err = master.DeployMaster(anode)
  changecontrol.ChangeControlInsertData(err, "DeployMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/updateMasterNetworkInterface \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "value": "v"
//   }
//  }
func UpdateMasterNetworkInterface(anode map[string]string, username string) (err error) {
  err = master.UpdateMasterNetworkInterface(anode)
  changecontrol.ChangeControlInsertData(err, "UpdateMasterNetworkInterface", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/loadMasterNetworkValuesSelected \
//  }
func LoadMasterNetworkValuesSelected(username string) (data map[string]map[string]string, err error) {
  data, err = master.LoadMasterNetworkValuesSelected()
  changecontrol.ChangeControlInsertData(err, "LoadMasterNetworkValuesSelected", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/pingservice \
//  }
func PingServiceMaster(username string) (err error) {
  err = master.PingServiceMaster()
  changecontrol.ChangeControlInsertData(err, "PingServiceMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/deployservice \
//  }
func DeployServiceMaster(username string) (err error) {
  err = master.DeployServiceMaster()
  changecontrol.ChangeControlInsertData(err, "DeployServiceMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/add \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v",
//      "type": "v",
//      "cert": "v",
//      "port": "v",
//      "interface": "v"
//   }
//  }
func AddPluginServiceMaster(anode map[string]string, username string) (err error) {
  err = master.AddPluginServiceMaster(anode)
  changecontrol.ChangeControlInsertData(err, "AddPluginServiceMaster", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteService \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v"
//   }
//  }
func DeleteServiceMaster(anode map[string]string, username string) (err error) {
  err = master.DeleteServiceMaster(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteServiceMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/modifyStapValues \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v",
//      "type": "v",
//      "cert": "v",
//      "port": "v",
//      "interface": "v"
//   }
//  }
func ModifyStapValuesMaster(anode map[string]string, username string) (err error) {
  err = master.ModifyStapValuesMaster(anode)
  changecontrol.ChangeControlInsertData(err, "ModifyStapValuesMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/updateMasterStapInterface \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "value": "v"
//   }
//  }
func UpdateMasterStapInterface(anode map[string]string, username string) (err error) {
  err = master.UpdateMasterStapInterface(anode)
  changecontrol.ChangeControlInsertData(err, "UpdateMasterStapInterface", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/setbpf \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "value": "v"
//   }
//  }
func SetBPF(anode map[string]string, username string) (err error) {
  err = master.SetBPF(anode)
  changecontrol.ChangeControlInsertData(err, "SetBPF", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/deployStapServiceMaster \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "type": "v",
//      "port": "v",
//      "interface": "v",
//      "collector": "v"
//   }
//  }
func DeployStapServiceMaster(anode map[string]string, username string) (err error) {
  err = master.DeployStapServiceMaster(anode)
  changecontrol.ChangeControlInsertData(err, "DeployStapServiceMaster", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/stopStapServiceMaster \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "type": "v"
//   }
//  }
func StopStapServiceMaster(anode map[string]string, username string) (err error) {
  err = master.StopStapServiceMaster(anode)
  changecontrol.ChangeControlInsertData(err, "StopStapServiceMaster", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/incidents \
//  }
func GetIncidents(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetIncidents()
  changecontrol.ChangeControlInsertData(err, "GetIncidents", username)
  return data, err
}

//  curl -X POST \
//    https://52.47.197.22:50002/v1/master/incidents \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "value": "v",
//  }
func PutIncident(anode map[string]string, username string) (err error) {
  err = master.PutIncident(anode)
  changecontrol.ChangeControlInsertData(err, "PutIncident", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/zeek/saveZeekValues \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "param": "v",
//      "node": "v"
//  }
func SaveZeekValues(anode map[string]string, username string) (err error) {
  err = master.SaveZeekValues(anode)
  changecontrol.ChangeControlInsertData(err, "SaveZeekValues", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/incidents \
//  }
func PingPluginsMaster(username string) (data map[string]map[string]string, err error) {
  data, err = master.PingPluginsMaster()
  changecontrol.ChangeControlInsertData(err, "PingPluginsMaster", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/editPathFile/:path \
//  }
func GetPathFileContent(param string, username string) (data map[string]string, err error) {
  data, err = master.GetPathFileContent(param)
  changecontrol.ChangeControlInsertData(err, "GetPathFileContent", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/savefilePath \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "file": "v",
//      "content": "v"
//   }
//  }
func SaveFilePathContent(anode map[string]string, username string) (err error) {
  err = master.SaveFilePathContent(anode)
  changecontrol.ChangeControlInsertData(err, "SaveFilePathContent", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/login \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "user": "v",
//      "password": "v"
//   }
//  }
func Login(anode map[string]string, username string) (token string, err error) {
  token, err = master.Login(anode)
  changecontrol.ChangeControlInsertData(err, "Login", username)
  return token, err
}

//  curl -X POST \
//    https://52.47.197.22:50002/v1/master/addUser \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "user": "v",
//      "pass": "v"
//  }
func AddUser(anode map[string]string, username string) (err error) {
  err = master.AddUser(anode)
  changecontrol.ChangeControlInsertData(err, "AddUser", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getAllUsers \
//  }
func GetAllUsers(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetAllUsers()
  changecontrol.ChangeControlInsertData(err, "GetAllUsers", username)
  return data, err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteUser \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v"
//   }
//  }
func DeleteUser(anode map[string]string, username string) (err error) {
  err = master.DeleteUser(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteUser", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addGroupUsers \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "group": "v"
//   }
//  }
func AddGroupUsers(anode map[string]string, username string) (err error) {
  err = master.AddGroupUsers(anode)
  changecontrol.ChangeControlInsertData(err, "AddGroupUsers", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "role": "v"
//   }
//  }
func AddRole(anode map[string]string, username string) (err error) {
  err = master.AddRole(anode)
  changecontrol.ChangeControlInsertData(err, "AddRole", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getRolesForUser/:roleID \
//  }
func GetRolesForUser(id string, username string) (data map[string]map[string]string, err error) {
  data, err = master.GetRolesForUser(id)
  changecontrol.ChangeControlInsertData(err, "GetRolesForUser", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getGroupsForUser/:groupID \
//  }
func GetGroupsForUser(id string, username string) (data map[string]map[string]string, err error) {
  data, err = master.GetGroupsForUser(id)
  changecontrol.ChangeControlInsertData(err, "GetGroupsForUser", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addUsersTo \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "user": "v",
//      "type": "v",
//      "values": [x,y,z]
//   }
//  }
func AddUsersTo(anode map[string]string, username string) (err error) {
  err = master.AddUsersTo(anode)
  changecontrol.ChangeControlInsertData(err, "AddUsersTo", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/changePassword \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "user": "v",
//      "pass": "v",
//   }
//  }
func ChangePassword(anode map[string]string, username string) (err error) {
  err = master.ChangePassword(anode)
  changecontrol.ChangeControlInsertData(err, "ChangePassword", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/getFileContentByType \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "file": "v",
//      "type": "v",
//   }
//  }
func GetFileContentByType(anode map[string]string, username string) (content map[string]string, err error) {
  content, err = master.GetFileContentByType(anode)
  changecontrol.ChangeControlInsertData(err, "GetFileContentByType", username)
  return content, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/saveNewFileContent \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "file": "v",
//      "type": "v",
//   }
//  }
func SaveNewFileContent(anode map[string]string, username string) (err error) {
  err = master.SaveNewFileContent(anode)
  changecontrol.ChangeControlInsertData(err, "SaveNewFileContent", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteUserRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "role": "v"
func DeleteUserRole(anode map[string]string, username string) (err error) {
  err = master.DeleteUserRole(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteUserRole", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteUserGroup \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "group": "v"
func DeleteUserGroup(anode map[string]string, username string) (err error) {
  err = master.DeleteUserGroup(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteUserGroup", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getAllRoles \
//  }
func GetAllRoles(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetAllRoles()
  changecontrol.ChangeControlInsertData(err, "GetAllRoles", username)
  return data, err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v"
func DeleteRole(anode map[string]string, username string) (err error) {
  err = master.DeleteRole(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteRole", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/editRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "name": "v",
//      "permissions": "v",
func EditRole(anode map[string]string, username string) (err error) {
  err = master.EditRole(anode)
  changecontrol.ChangeControlInsertData(err, "EditRole", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/editUserGroup \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "name": "v"
func EditUserGroup(anode map[string]string, username string) (err error) {
  err = master.EditUserGroup(anode)
  changecontrol.ChangeControlInsertData(err, "EditUserGroup", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getAllUserGroups \
//  }
func GetAllUserGroups(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetAllUserGroups()
  changecontrol.ChangeControlInsertData(err, "GetAllUserGroups", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getRolesForGroups/:roleID \
//  }
func GetRolesForGroups(id string, username string) (data map[string]map[string]string, err error) {
  data, err = master.GetRolesForGroups(id)
  changecontrol.ChangeControlInsertData(err, "GetRolesForGroups", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addRoleToGroup \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "id",
//      "name": "get,put"
func AddRoleToGroup(anode map[string]string, username string) (err error) {
  err = master.AddRoleToGroup(anode)
  changecontrol.ChangeControlInsertData(err, "AddRoleToGroup", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteRoleUser \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "user": "v"
func DeleteRoleUser(anode map[string]string, username string) (err error) {
  err = master.DeleteRoleUser(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteRoleUser", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteRoleGroup \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "group": "v"
func DeleteRoleGroup(anode map[string]string, username string) (err error) {
  err = master.DeleteRoleGroup(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteRoleGroup", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteGroupUser \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "user": "v"
func DeleteGroupUser(anode map[string]string, username string) (err error) {
  err = master.DeleteGroupUser(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteGroupUser", username)
  return err
}

//  curl -X DELETE \
//    https://52.47.197.22:50002/v1/master/deleteGroupRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "id": "v",
//      "role": "v"
func DeleteGroupRole(anode map[string]string, username string) (err error) {
  err = master.DeleteGroupRole(anode)
  changecontrol.ChangeControlInsertData(err, "DeleteGroupRole", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getPermissions \
//  }
func GetPermissions(username string) (data map[string]map[string]string, err error) {
  data, err = master.GetPermissions()
  changecontrol.ChangeControlInsertData(err, "GetPermissions", username)
  return data, err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/GetPermissionsByRole/:uuid \
//  }
func GetPermissionsByRole(uuid string, username string) (data map[string]map[string]string, err error) {
  data, err = master.GetPermissionsByRole(uuid)
  changecontrol.ChangeControlInsertData(err, "GetPermissionsByRole", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addNewRole \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "role": "v"
//   }
//  }
func AddNewRole(anode map[string]string, username string) (err error) {
  err = master.AddNewRole(anode)
  changecontrol.ChangeControlInsertData(err, "AddNewRole", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/master/addOrganization \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "name": "v"
//      "desc": "v"
//      "default": "v"
//   }
//  }
func AddOrganization(anode map[string]string, username string) (err error) {
  err = master.AddOrganization(anode)
  changecontrol.ChangeControlInsertData(err, "AddOrganization", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/master/getAllOrganizationNodes/:uuid \
//  }
func GetAllOrganizationNodes(file string, username string) (data map[string]string, err error) {
  data, err = master.GetAllOrganizationNodes(file)
  changecontrol.ChangeControlInsertData(err, "GetAllOrganizationNodes", username)
  return data, err
}
