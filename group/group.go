package group

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "errors"
    "owlhmaster/utils"
)


func CreateGroup(n map[string]string) (err error) {
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
        return errors.New("Group exist")
    }
    
    for key, value := range n {
        err = groupKeyInsert(groupKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func DeleteGroup(groupId string) (err error) {
	if ndb.Gdb == nil {
        logs.Error("DeleteGroup -- Can't acces to database")
        return errors.New("DeleteGroup -- Can't acces to database")
    }
	stmt, err := ndb.Gdb.Prepare("delete from groups where group_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare DeleteGroup -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&groupId)
    if err != nil {
        logs.Error("Execute DeleteGroup -> %s", err.Error())
        return err
    }

	return nil
}

func groupExists(nodeid string) (err error) {
    if ndb.Gdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    sql := "SELECT * FROM groups where group_uniqueid = '"+nodeid+"';"
    rows, err := ndb.Gdb.Query(sql)
    if err != nil {
        logs.Error("Error on query groupExist at group.go "+err.Error())
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
    stmt, err := ndb.Gdb.Prepare("insert into groups (group_uniqueid, group_param, group_value) values(?,?,?)")
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

func GetAllGroups()(groups map[string]map[string]string, err error){
	var allgroups = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    if ndb.Gdb == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
    }
    sql := "select group_uniqueid, group_param, group_value from groups;"
    rows, err := ndb.Gdb.Query(sql)
    if err != nil {
        logs.Error("ndb.Gdb.Query Error : %s", err.Error())
        return nil, err
    }
    for rows.Next() {
        if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("GetAllGroups rows.Scan: %s", err.Error())
            return nil, err
        }
        if allgroups[uniqid] == nil { allgroups[uniqid] = map[string]string{}}
        allgroups[uniqid][param]=value
	} 
    return allgroups, nil
}

func EditGroup(data map[string]string) (err error) { 
	var name = data["name"]
	var desc = data["desc"]
	var groupid = data["groupid"]
    if ndb.Gdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
	}
	
	//insert name
    insertName, err := ndb.Gdb.Prepare("update groups set group_value = ? where group_param = ? and group_uniqueid = ?")
    if err != nil {
		logs.Error("Prepare EditGroup-> %s", err.Error())
        return err
    }
    _, err = insertName.Exec(&name, "name", &groupid)
    if err != nil {
		logs.Error("Execute EditGroup-> %s", err.Error())
        return err
	}
	
	//insert desc
	insertDesc, err := ndb.Gdb.Prepare("update groups set group_value = ? where group_param = ? and group_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare EditGroup-> %s", err.Error())
        return err
    }
    _, err = insertDesc.Exec(&desc, "desc", &groupid)
    if err != nil {
        logs.Error("Execute EditGroup-> %s", err.Error())
        return err
	}
	
    return nil
}