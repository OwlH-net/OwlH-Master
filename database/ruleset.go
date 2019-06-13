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

func RulesetSourceKeyInsert(nkey string, key string, value string) (err error) {
    if Rdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values (?,?,?);")
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
            logs.Error("UpdateRuleset UPDATE prepare error for update isDownloaded -- "+err.Error())
            return err
        }
        _, err = updateRulesetNode.Exec(&value, &uuid, &param)
        defer updateRulesetNode.Close()
        if (err != nil){
            logs.Error("UpdateRuleset UPDATE error for update isDownloaded -- "+err.Error())
            return err
		}
	return nil
}

func UpdateRuleFiles(uuid string, param string, value string)(err error){
	updateRulesetNode, err := Rdb.Prepare("update rule_files set rule_value = ? where rule_uniqueid = ? and rule_param = ?;")
	defer updateRulesetNode.Close()
	if (err != nil){
		logs.Error("UpdateRuleFiles UPDATE prepare error for update isDownloaded -- "+err.Error())
		defer updateRulesetNode.Close()
		return err
	}
	// defer updateRulesetNode.Close()
	_, err = updateRulesetNode.Exec(&value, &uuid, &param)
	if (err != nil){
		logs.Error("UpdateRuleFiles UPDATE error for update isDownloaded -- "+err.Error())
		defer updateRulesetNode.Close()
		return err
	}
	
	return nil
}

func GetRulesetSourceValue(uuid string, param string)(path string, err error){
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

func GetAllCustomRulesetDB()(path []string, err error){
	var customData []string
    var uniqid string

	sql := "select ruleset_uniqueid from ruleset where ruleset_value='custom' and ruleset_param='sourceType'";
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllCustomRuleset Rdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid); err != nil {
            logs.Error("GetAllCustomRuleset -- Query return error: %s", err.Error())
            return nil, err
		}
		customData = append(customData, uniqid)
	} 
	return customData,nil
}

func GetAllDataCustomRulesetDB(uuid string)(path map[string]map[string]string, err error){
	var customData = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string

	sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset where ruleset_uniqueid='"+uuid+"'";
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllCustomRuleset Rdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("GetAllCustomRuleset -- Query return error: %s", err.Error())
            return nil, err
		}
        if customData[uniqid] == nil { customData[uniqid] = map[string]string{}}
        customData[uniqid][param]=value
	} 
	return customData,nil
}

//Get a specific ruleset path
func GetRulesetPath(uuid string) (n string, err error) {
	var path string
    if Rdb != nil {
        row := Rdb.QueryRow("SELECT rule_value FROM rule_files WHERE rule_uniqueid=$1 and rule_param=\"path\";",uuid)
		err = row.Scan(&path)

        if err == sql.ErrNoRows {
            logs.Error("DB RULESET -> There is no ruleset with id %s",uuid)
            return "", errors.New("DB RULESET -> There is no ruleset with id "+uuid)
        }
        if err != nil {
            logs.Error("DB RULESET -> rows.Scan Error -> %s", err.Error())
            return "", errors.New("DB RULESET -> -> rows.Scan Error -> " + err.Error())
        }
        return path, nil
    } else {
        logs.Info("DB RULESET -> No access to database")
        return "", errors.New("DB RULESET -> No access to database")
    }
}

func GetRuleFilesValue(uuid string, param string)(path string, err error){
	var value string
	sql := "select rule_value from rule_files where rule_uniqueid='"+uuid+"' and rule_param = '"+param+"';"
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