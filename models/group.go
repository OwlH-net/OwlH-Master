package models 

import (
    "owlhmaster/group"
)

func CreateGroup(data map[string]string) (err error) {
    err = group.CreateGroup(data)
    return err
}

func EditGroup(data map[string]string) (err error) {
    err = group.EditGroup(data)
    return err
}

func DeleteGroup(groupId string) (err error) {
    err = group.DeleteGroup(groupId)
    return err
}

func GetAllGroups() (data map[string]map[string]string, err error) {
    data, err = group.GetAllGroups()
    return data, err
}