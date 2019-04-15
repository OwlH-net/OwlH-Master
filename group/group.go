package group

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/utils"
)


func CreateGroup(n map[string]string) (err error) {
	logs.Warn(n["name"])
	logs.Warn(n["desc"])
	groupKey := utils.Generate()
    if _, ok := n["name"]; !ok {
		logs.Error("name empty: "+err.Error())
        return errors.New("name empty")
    }
    if _, ok := n["desc"]; !ok {
		logs.Error("desc empty: "+err.Error())
        return errors.New("desc empty")
    }

    if err := groupExists(groupKey); err != nil {
		logs.Error("Group exist: "+err.Error())
        return errors.New("name empty")
    }
    
    for key, value := range n {
        err = groupKeyInsert(groupKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func groupExists(nodeid string) (err error) {
    if ndb.Gdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    sql := "SELECT * FROM groups where node_uniqueid = '"+nodeid+"';"
    rows, err := ndb.Gdb.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return errors.New("Node Exists " + nodeid)
    } else {
        return nil
    }
}

func groupKeyInsert(nkey string, key string, value string) (err error) {
    if ndb.Gdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := ndb.Gdb.Prepare("insert into groups (node_uniqueid, node_param, node_value) values(?,?,?)")
    if err != nil {
        logs.Error("Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nkey, &key, &value)
    if err != nil {
        logs.Error("Execute -> %s", err.Error())
        return err
    }
    return nil
}