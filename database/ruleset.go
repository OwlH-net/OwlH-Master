package ndb

import (
    "github.com/astaxie/beego/logs"
	"database/sql"
	"owlhmaster/utils"
	"os"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

var (
    Rdb *sql.DB
)

func RConn() {
	var err error
	loadDataSQL := map[string]map[string]string{}
	loadDataSQL["rulesetConn"] = map[string]string{}
	loadDataSQL["rulesetConn"]["path"] = ""
	loadDataSQL["rulesetConn"]["cmd"] = "" 
	loadDataSQL, err = utils.GetConf(loadDataSQL)    
    path := loadDataSQL["rulesetConn"]["path"]
    cmd := loadDataSQL["rulesetConn"]["cmd"]
	if err != nil {
		logs.Error("RConn Error getting data from main.conf at master: "+err.Error())
	}
	_, err = os.Stat(path) 
	if err != nil {
		panic("ruleset/Ruleset DB -- DB Open Failed: "+err.Error())
	}	
    Rdb, err = sql.Open(cmd,path)
    if err != nil {
		logs.Error("ruleset/Ruleset DB -- SQL openning Error: "+err.Error()) 
    }
    logs.Info("ruleset/Ruleset DB -- DB -> sql.Open, DB Ready") 
}

func InsertRulesetSourceRules(nkey string, key string, value string) (err error) {
    if Rdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := Rdb.Prepare("insert into rule_files (rule_uniqueid, rule_param, rule_value) values(?,?,?)")
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

func UpdateRuleset(uuid string, param string, value string)(err error){
	logs.Debug(param+" --> "+value)
	updateRulesetNode, err := Rdb.Prepare("update ruleset set ruleset_value = ? where ruleset_uniqueid = ? and ruleset_param = ?;")
        if (err != nil){
            logs.Error("UPDATE prepare error for update isDownloaded -- "+err.Error())
            return err
        }
        _, err = updateRulesetNode.Exec(&value, &uuid, &param)
        defer updateRulesetNode.Close()
        if (err != nil){
            logs.Error("UPDATE error for update isDownloaded -- "+err.Error())
            return err
		}
	return nil
}

func GetRulesetSourcePath(uuid string, param string)(path string, err error){
	var value string
	sql := "select ruleset_value from ruleset where ruleset_uniqueid='"+uuid+"' and ruleset_param = '"+param+"';"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query Error : %s", err.Error())
		return "", err
	}
	for rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetRulesetSourcePath rows.Scan: %s", err.Error())
			return "", err
		}
	} 
	return value,nil
}