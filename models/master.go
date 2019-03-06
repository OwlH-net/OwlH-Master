package models 

import (
    // "github.com/astaxie/beego/logs"
//    "database/sql"
//    "fmt"
//   "time"
//    _ "github.com/mattn/go-sqlite3"
    "owlhmaster/database"
    "owlhmaster/aboutme"
)

var (
        Pmaster Master
)

func init() {
}


type Master struct {
    Id       string
    Name     string
    Ip       string
    Port     int
}

func InitMaster() string {
    return "go"
}

func GetMaster() (master map[string]*string, err error) {
    ndb.Get_master()
    lmaster, err := aboutme.AboutMe()
	return lmaster, err
}

func UpdateMaster(m map[string]string) (err error) {
    for key, value := range m {
        err = aboutme.UpdateMe(key, value)
    }
    return err
}

