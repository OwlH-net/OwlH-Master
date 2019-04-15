package models 

import (
    "owlhmaster/master"
)

func GetMasterTitle() (data string, err error) {
    data, err = master.GetMasterTitle()
    return data, err
}