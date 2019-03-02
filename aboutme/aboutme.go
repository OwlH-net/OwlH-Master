package aboutme


import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
//    "bytes"
//    "fmt"
//    "log"
//    "os/exec"
//    "errors"
//    "strings"
//    "database/sql"
//    _ "github.com/mattn/go-sqlite3"
)

var (
    paramList map[string]*string
)

func AboutMe() (plist map[string]*string, err error) {
    paramList = make(map[string]*string)
    if ndb.Db == nil {
        logs.Error("Can't Access Database")
        return nil, errors.New("Can't Access Database")
    }
    rows, err := ndb.Db.Query("SELECT * FROM aboutme;")
    if err != nil {
        logs.Error(err.Error())
        return nil, err
    }
    for rows.Next() {
        var id int
        var name string
        var value string
        err = rows.Scan(&id, &name, &value)
        if err != nil {
            logs.Error("No hemos podido recoger los campos del registro leido (name o value)")
            continue
        }
        logs.Info("name: %s, Value: %s",name, value)
        paramList[name]=&value
    }
    return paramList, nil
}

func UpdateMe(param string, value string) (err error) {

    // consulto si el param existe 

    //insert or replace into aboutme (am_param, am_value) values ("floor","third");
    if ndb.Db == nil {
        logs.Error("no hemos podido acceder a la base de datos")
        return errors.New("no hemos podido acceder a la bbdd")
    }
    stmt, err := ndb.Db.Prepare("insert or replace into aboutme (am_param, am_value) values(?,?)")
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    res, err := stmt.Exec(&param, &value)
    logs.Info("DB -> info", res)
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    return nil
}
