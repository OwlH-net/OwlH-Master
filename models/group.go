package models 

import (
    "owlhmaster/group"
)

func CreateGroup(data map[string]string)(err error) {
    err = group.CreateGroup(data)
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