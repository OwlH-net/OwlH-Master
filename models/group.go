package models 

import (
    "owlhmaster/group"
)

func CreateGroup(data map[string]string) (err error) {
    err = group.CreateGroup(data)
    return err
}