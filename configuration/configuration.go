package configuration

import (
    // "encoding/json"
    // "strconv"
    "github.com/astaxie/beego/logs"
    "database/sql"
    // "io/ioutil"
    // "io"
    // "errors"
    "owlhmaster/utils"
    "owlhmaster/validation"
    "os"
    // "time"
    // "os/exec"
    // "fmt"
    // "crypto/rand"
    _ "github.com/mattn/go-sqlite3"
)

type Dbconfig struct {
    Dbname          string
    Dbconn          string
    Dbpath          string
    Dbtables        []Table
}

type Table struct {
    Tconn           string
    Tname           string
    Tcreate         string
    Tfields         []Field
}

type Field struct {
    Fconn           string
    Ftable          string
    Fname           string
    Fquery          string
    Finsert         string
}

var DBCONFIG        []Dbconfig


func MainCheck()(cancontinue bool){

    ok := checkDatabases()
    if !ok {
        return false
    }
    ok = checkTables()
    if !ok {
        return false
    }
    ok = checkFields()
    if !ok {
        return false
    }
    return true
}

func checkDatabases()(ok bool){
    dbs := []string{"masterConn","dbsConn","rulesetConn"}
    for db := range dbs {
        logs.Warn("lets check db -> "+dbs[db])
        ok := CheckDB(dbs[db])
        if !ok {
            return false
        }
    }
    return true
}


func checkTables()(ok bool){
    var table Table

    table.Tname = "usergrouproles"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE usergrouproles (ugr_id integer PRIMARY KEY AUTOINCREMENT,ugr_uniqueid text NOT NULL,ugr_param text NOT NULL,ugr_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "userRoles"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE userRoles (ur_id integer PRIMARY KEY AUTOINCREMENT,ur_uniqueid text NOT NULL,ur_param text NOT NULL,ur_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    // table.Tname = "userPrivileges"
    // table.Tconn = "masterConn"
    // table.Tcreate = "CREATE TABLE userPrivileges (priv_id integer PRIMARY KEY AUTOINCREMENT,priv_uniqueid text NOT NULL,priv_param text NOT NULL,priv_value text NOT NULL)"
    // ok = CheckTable(table)
    // if !ok {
    //     return false
    // }

    table.Tname = "userGroups"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE userGroups (ug_id integer PRIMARY KEY AUTOINCREMENT,ug_uniqueid text NOT NULL,ug_param text NOT NULL,ug_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "users"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE users (user_id integer PRIMARY KEY AUTOINCREMENT,user_uniqueid text NOT NULL,user_param text NOT NULL,user_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "plugins"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE plugins (plugin_id integer PRIMARY KEY AUTOINCREMENT,plugin_uniqueid text NOT NULL,plugin_param text NOT NULL,plugin_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "groupcluster"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE groupcluster (gc_id integer PRIMARY KEY AUTOINCREMENT,gc_uniqueid text NOT NULL,gc_param text NOT NULL,gc_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "groups"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE groups (group_id integer PRIMARY KEY AUTOINCREMENT,group_uniqueid text NOT NULL,group_param text NOT NULL,group_value text NOT NULL)"

    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "masterconfig"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE masterconfig (config_id integer PRIMARY KEY AUTOINCREMENT,config_uniqueid text NOT NULL,config_param text NOT NULL,config_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "dataflow"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE dataflow (flow_id integer PRIMARY KEY AUTOINCREMENT,flow_uniqueid text NOT NULL,flow_param text NOT NULL,flow_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "changerecord"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE changerecord (control_id integer PRIMARY KEY AUTOINCREMENT,control_uniqueid text NOT NULL,control_param text NOT NULL,control_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "incidents"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE incidents (incidents_id integer PRIMARY KEY AUTOINCREMENT,incidents_uniqueid text NOT NULL,incidents_param text NOT NULL,incidents_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "groupnodes"
    table.Tconn = "masterConn"
    table.Tcreate = "CREATE TABLE groupnodes (gn_id integer PRIMARY KEY AUTOINCREMENT,gn_uniqueid text NOT NULL,gn_param text NOT NULL,gn_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "nodes"
    table.Tconn = "dbsConn"
    table.Tcreate = "CREATE TABLE nodes (node_id integer PRIMARY KEY AUTOINCREMENT,node_uniqueid text NOT NULL,node_param text NOT NULL,node_value text NOT NULL)"

    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "ruleset"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE ruleset (ruleset_id integer PRIMARY KEY AUTOINCREMENT,ruleset_uniqueid text NOT NULL,ruleset_param text NOT NULL,ruleset_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "ruleset_node"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE ruleset_node (ruleset_id integer PRIMARY KEY AUTOINCREMENT,ruleset_uniqueid text NOT NULL,node_uniqueid text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "rule_note"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE rule_note (ruleset_id integer PRIMARY KEY AUTOINCREMENT,ruleset_uniqueid text NOT NULL,rule_sid text NOT NULL,note_date text NOT NULL,ruleNote text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "rule_files"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE rule_files (rule_id integer PRIMARY KEY AUTOINCREMENT,rule_uniqueid text NOT NULL,rule_param text NOT NULL,rule_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "scheduler"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE scheduler (scheduler_id integer PRIMARY KEY AUTOINCREMENT,scheduler_uniqueid text NOT NULL,scheduler_param text NOT NULL,scheduler_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }

    table.Tname = "scheduler_log"
    table.Tconn = "rulesetConn"
    table.Tcreate = "CREATE TABLE scheduler_log (log_id integer PRIMARY KEY AUTOINCREMENT,log_uniqueid text NOT NULL,log_param text NOT NULL,log_value text NOT NULL)"
    ok = CheckTable(table)
    if !ok {
        return false
    }
    return true
}

func checkFields()(ok bool){

    var field Field
    
    //add admin user by default
    secret := utils.Generate()
    pass,err := validation.HashPassword("admin")
    if err!=nil {logs.Error("Error hashing password at configuration.")}
    field.Fconn      = "masterConn"
    field.Ftable     = "users"
    field.Fquery     = "select user_param from users where user_param='user' and user_value='admin'"
    field.Finsert    = "insert into users (user_uniqueid,user_param,user_value) values ('00000000-0000-0000-0000-000000000000','user','admin')"
    field.Fname      = "users - user"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "users"
    field.Fquery     = "select user_param from users where user_param='secret' and user_uniqueid='00000000-0000-0000-0000-000000000000'"
    field.Finsert    = "insert into users (user_uniqueid,user_param,user_value) values ('00000000-0000-0000-0000-000000000000','secret','"+secret+"')"
    field.Fname      = "users - secret"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "users"
    field.Fquery     = "select user_param from users where user_param='pass' and user_uniqueid='00000000-0000-0000-0000-000000000000'"
    field.Finsert    = "insert into users (user_uniqueid,user_param,user_value) values ('00000000-0000-0000-0000-000000000000','pass','"+pass+"')"
    field.Fname      = "users - pass"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "users"
    field.Fquery     = "select user_param from users where user_param='deleteable' and user_uniqueid='00000000-0000-0000-0000-000000000000'"
    field.Finsert    = "insert into users (user_uniqueid,user_param,user_value) values ('00000000-0000-0000-0000-000000000000','deleteable','false')"
    field.Fname      = "users - deleteable"
    ok = CheckField(field)
    if !ok {return false}

    //add admin to role admin status
    masterUUID := utils.Generate()
    field.Fconn      = "masterConn"
    field.Ftable     = "masterconfig"
    field.Fquery     = "select config_param from masterconfig where config_param='id'"
    field.Finsert    = "insert into masterconfig (config_uniqueid,config_param,config_value) values ('master','id','"+masterUUID+"')"
    field.Fname      = "masterconfig - id"
    ok = CheckField(field)
    if !ok {return false}

    //add admin to role admin status
    field.Fconn      = "masterConn"
    field.Ftable     = "userRoles"
    field.Fquery     = "select ur_param from userRoles where ur_param='role' and ur_value='admin'"
    field.Finsert    = "insert into userRoles (ur_uniqueid,ur_param,ur_value) values ('00000000-0000-0000-0000-000000000001','role','admin')"
    field.Fname      = "userRoles - role"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "userRoles"
    field.Fquery     = "select ur_param from userRoles where ur_param='permissions'"
    field.Finsert    = "insert into userRoles (ur_uniqueid,ur_param,ur_value) values ('00000000-0000-0000-0000-000000000001','permissions','get,put,post,delete')"
    field.Fname      = "userRoles - role"
    ok = CheckField(field)
    if !ok {return false}

    //add admin to group admin status
    field.Fconn      = "masterConn"
    field.Ftable     = "userGroups"
    field.Fquery     = "select ug_param from userGroups where ug_param='group' and ug_value='admin'"
    field.Finsert    = "insert into userGroups (ug_uniqueid,ug_param,ug_value) values ('00000000-0000-0000-0000-000000000002','group','admin')"
    field.Fname      = "userGroups - group"
    ok = CheckField(field)
    if !ok {return false}

    //create usergrouproles admin values
    field.Fconn      = "masterConn"
    field.Ftable     = "usergrouproles"
    field.Fquery     = "select ugr_uniqueid from usergrouproles where ugr_uniqueid='00000000-0000-0000-0000-000000000003' and ugr_param='user' and ugr_value='00000000-0000-0000-0000-000000000000'"
    field.Finsert    = "insert into usergrouproles (ugr_uniqueid,ugr_param,ugr_value) values ('00000000-0000-0000-0000-000000000003','user','00000000-0000-0000-0000-000000000000')"
    field.Fname      = "usergrouproles - user"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "usergrouproles"
    field.Fquery     = "select ugr_uniqueid from usergrouproles where ugr_uniqueid='00000000-0000-0000-0000-000000000003' and ugr_param='role' and ugr_value='00000000-0000-0000-0000-000000000001'"
    field.Finsert    = "insert into usergrouproles (ugr_uniqueid,ugr_param,ugr_value) values ('00000000-0000-0000-0000-000000000003','role','00000000-0000-0000-0000-000000000001')"
    field.Fname      = "usergrouproles - role"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "usergrouproles"
    field.Fquery     = "select ugr_value from usergrouproles where ugr_uniqueid='00000000-0000-0000-0000-000000000004' and ugr_value='00000000-0000-0000-0000-000000000002'"
    field.Finsert    = "insert into usergrouproles (ugr_uniqueid,ugr_param,ugr_value) values ('00000000-0000-0000-0000-000000000004','group','00000000-0000-0000-0000-000000000002')"
    field.Fname      = "usergrouproles - group"
    ok = CheckField(field)
    if !ok {return false}
    field.Fconn      = "masterConn"
    field.Ftable     = "usergrouproles"
    field.Fquery     = "select ugr_value from usergrouproles where ugr_uniqueid='00000000-0000-0000-0000-000000000004' and ugr_value='00000000-0000-0000-0000-000000000000'"
    field.Finsert    = "insert into usergrouproles (ugr_uniqueid,ugr_param,ugr_value) values ('00000000-0000-0000-0000-000000000004','user','00000000-0000-0000-0000-000000000000')"
    field.Fname      = "user - usergrouproles"
    ok = CheckField(field)
    if !ok {return false}

    //add dispatcher status
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='status' and plugin_uniqueid='dispatcher'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('dispatcher','status','disabled')"
    field.Fname      = "dispatcher - status"
    ok = CheckField(field)
    if !ok {
        return false
    }

    //Create default Zeek values
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='nodeConfig'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','nodeConfig','')"
    field.Fname      = "plugin - nodeConfig"
    ok = CheckField(field)
    if !ok {
        return false
    }
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='networksConfig'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','networksConfig','')"
    field.Fname      = "plugin - networksConfig"
    ok = CheckField(field)
    if !ok {
        return false
    }
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='policiesMaster'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','policiesMaster','')"
    field.Fname      = "plugin - policiesMaster"
    ok = CheckField(field)
    if !ok {
        return false
    }
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='policiesNode'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','policiesNode','')"
    field.Fname      = "plugin - policiesNode"
    ok = CheckField(field)
    if !ok {
        return false
    }
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='variables1'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','variables1','')"
    field.Fname      = "plugin - variables1"
    ok = CheckField(field)
    if !ok {
        return false
    }
    field.Fconn      = "masterConn"
    field.Ftable     = "plugins"
    field.Fquery     = "select plugin_param from plugins where plugin_param='variables2'"
    field.Finsert    = "insert into plugins (plugin_uniqueid,plugin_param,plugin_value) values ('zeek','variables2','')"
    field.Fname      = "plugin - variables2"
    ok = CheckField(field)
    if !ok {
        return false
    }

    return true
}

func CheckDB(conn string)(ok bool) {
    loadDataSQL := map[string]map[string]string{}
    loadDataSQL[conn] = map[string]string{}
    loadDataSQL[conn]["path"] = ""
    loadDataSQL, err := utils.GetConf(loadDataSQL)
    if err != nil {
        logs.Error("Configuration -> Can't get "+conn+" path from main.conf")
        return false
    }
    dbpath := loadDataSQL[conn]["path"]

    exists := DbExists(dbpath)

    if exists {
        logs.Warn("Configuration -> db "+dbpath+" exists")
        return true
    } else {
        logs.Warn("Configuration -> db "+dbpath+" does not exist, ... Creating")
        err = DbCreate(dbpath)
        if err != nil {
            return false
        }
    }
    return true
}

func CheckField(field Field)(ok bool){
    loadDataSQL := map[string]map[string]string{}
    loadDataSQL[field.Fconn] = map[string]string{}
    loadDataSQL[field.Fconn]["path"] = ""
    loadDataSQL, err := utils.GetConf(loadDataSQL)
    if err != nil {logs.Error("Configuration -> Can't get DB "+field.Fconn+" path from main.conf"); return false}
    
    dbpath := loadDataSQL[field.Fconn]["path"]

    exists := FieldExists(dbpath, field.Fquery)
    if !exists {
        logs.Warn("Configuration -> Field "+field.Fname+" doesn't exist on Table/DB "+field.Ftable+"/"+field.Fconn+" ...Creating")
        created := FieldCreate(dbpath, field.Finsert, field.Fname)
        if !created {
            return false
        }
        return true
    }

    logs.Info("Configuration -> Field "+field.Fname+" exists on Table/DB "+field.Ftable+"/"+field.Fconn)
    return true
}

func FieldExists(dbpath, qry string)(ok bool){
    dblink, err := sql.Open("sqlite3", dbpath)
    if err != nil {
        logs.Error("Configuration -> Check Field -> db " + dbpath + " can't be opened -> err: "+err.Error())
        return false
    }
    defer dblink.Close()
    row := dblink.QueryRow(qry)

    var fieldname string
    switch err := row.Scan(&fieldname); err {
    case sql.ErrNoRows:
        return false
    case nil:
        return true
    default:
        return false
    }
    return true
}

func FieldCreate (dbpath string, insert string, name string)(ok bool){
    logs.Info("Configuration -> Creating field "+name+" in "+dbpath)

    dblink, err := sql.Open("sqlite3", dbpath)
    if err != nil {
        logs.Error("Configuration -> Check Field -> db " + dbpath + " can't be opened -> err: "+err.Error())
        return false
    }
    defer dblink.Close()
    _, err = dblink.Exec(insert)
    if err != nil {
        logs.Error("Configuration -> Creating field " + name + " failed -> err: "+err.Error())
        return false
    }
    return true
}

func CheckTable(table Table)(ok bool){
    loadDataSQL := map[string]map[string]string{}
    loadDataSQL[table.Tconn] = map[string]string{}
    loadDataSQL[table.Tconn]["path"] = ""
    loadDataSQL, err := utils.GetConf(loadDataSQL)
    if err != nil {
        logs.Error("Configuration -> Can't get "+table.Tconn+" path from main.conf")
        return false
    }
    dbpath := loadDataSQL[table.Tconn]["path"]

    exists := TableExists(dbpath, table.Tname)
    if !exists {
        logs.Warn("Configuration -> Table "+table.Tname+" doesn't exist on DB "+table.Tconn+" ...Creating")
        created := TableCreate(table.Tconn,table.Tname,table.Tcreate)
        if !created {
            return false
        }
        return true
    }

    logs.Info("Configuration -> Table "+table.Tname+" exists on DB "+table.Tconn)
    return true
}



func DbExists(db string)(exists bool){
    if _, err := os.Stat(db); os.IsNotExist(err) {
        logs.Error("Configuration -> Check DB -> db " + db + " not found -> err: " + err.Error())
        return false
    }else{
        dblink, err := sql.Open("sqlite3", db)
        if err != nil {
            logs.Error("Configuration -> Check DB -> db " + db + " can't be opened -> err: "+err.Error())
            return false
        }
        defer dblink.Close()
        return true
    }
    logs.Error("Configuration -> Check DB -> db " + db + " something went wrong, can't find a reason")
    return false
}

func TableExists(db string, table string)(exists bool){
    dblink, err := sql.Open("sqlite3", db)
    if err != nil {
        logs.Error("Configuration -> Check Table -> db " + db + " can't open -> err: "+err.Error())
        return false
    }
    defer dblink.Close()
    qry := "SELECT name FROM sqlite_master WHERE type='table' AND name=$1;"
    row := dblink.QueryRow(qry, table)

    var tablename string
    switch err := row.Scan(&tablename); err {
    case sql.ErrNoRows:
        return false
    case nil:
        return true
    default:
        return false
    }

    return true
}



func TableCreate(conn string, tablename string, create string)(ok bool){
    logs.Info("Configuration -> Creating table "+tablename+" in "+conn)
    loadDataSQL := map[string]map[string]string{}
    loadDataSQL[conn] = map[string]string{}
    loadDataSQL[conn]["path"] = ""
    loadDataSQL, err := utils.GetConf(loadDataSQL)
    if err != nil {
        logs.Error("Configuration -> Can't get "+conn+" path from main.conf -> "+err.Error())
        return false
    }
    dbpath := loadDataSQL[conn]["path"]
    db, err := sql.Open("sqlite3",dbpath)
    if err != nil {
        logs.Error("Configuration -> "+dbpath+" Open Failed -> err: "+err.Error())
        return false
    }
    _, err = db.Exec(create)
    if err != nil {
        logs.Error("Configuration -> Creating table " +tablename + " failed -> err: "+err.Error())
        return false
    }
    return true 
}

func DbCreate(db string)(err error) {
    logs.Warn ("Configuration -> Creating DB file -> "+db)
    _, err = os.OpenFile(db, os.O_CREATE, 0644)
    if err != nil {
        logs.Error("Configuration -> Creating DB File "+ db +" err: "+err.Error())
        return err
    }
    return nil
}