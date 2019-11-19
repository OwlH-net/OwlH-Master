package models 

import (
    "owlhmaster/group"
)

func AddGroupElement(data map[string]string)(err error) {
    err = group.AddGroupElement(data)
    return err
}

func EditGroup(data map[string]string)(err error) {
    err = group.EditGroup(data)
    return err
}

func DeleteGroup(groupId string)(err error) {
    err = group.DeleteGroup(groupId)
    return err
}

func GetAllGroups()(Groups []group.Group, err error) {
    Groups, err = group.GetAllGroups()
    return Groups, err
}

func GetAllNodesGroup(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetAllNodesGroup(uuid)
    return data, err
}

func AddGroupNodes(data map[string]interface{})(err error) {
    err = group.AddGroupNodes(data)
    return err
}

func PingGroupNodes()(data map[string]map[string]string, err error) {
    data, err = group.PingGroupNodes()
    return data, err
}

func GetNodeValues(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetNodeValues(uuid)
    return data, err
}

func DeleteNodeGroup(uuid string)(err error) {
    err = group.DeleteNodeGroup(uuid)
    return err
}

func ChangeGroupRuleset(data map[string]string)(err error) {
    err = group.ChangeGroupRuleset(data)
    return err
}

func ChangePathsGroups(data map[string]string)(err error) {
    err = group.ChangePathsGroups(data)
    return err
}

func UpdateGroupService(data map[string]string)(err error) {
    err = group.UpdateGroupService(data)
    return err
}

func SyncPathGroup(data map[string]string)(err error) {
    err = group.SyncPathGroup(data)
    return err
}

func SyncAll(uuid string, data map[string]map[string]string)(err error) {
    err = group.SyncAll(uuid, data)
    return err
}

func AddCluster(data map[string]string)(err error) {
    err = group.AddCluster(data)
    return err
}

func GetClusterFiles(uuid string)(data map[string]map[string]string, err error) {
    data, err = group.GetClusterFiles(uuid)
    return data, err
}

func GetClusterFileContent(uuid string)(data map[string]string, err error) {
    data, err = group.GetClusterFileContent(uuid)
    return data, err
}

func DeleteCluster(data map[string]string)(err error) {
    err = group.DeleteCluster(data)
    return err
}

func ChangeClusterValue(data map[string]string)(err error) {
    err = group.ChangeClusterValue(data)
    return err
}

func SaveClusterFileContent(data map[string]string)(err error) {
    err = group.SaveClusterFileContent(data)
    return err
}

func SyncClusterFile(data map[string]string)(err error) {
    err = group.SyncClusterFile(data)
    return err
}

func SyncAllGroupCluster(data map[string]string)(err error) {
    err = group.SyncAllGroupCluster(data)
    return err
}