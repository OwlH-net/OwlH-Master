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
		// defer updateRulesetNode.Close()
		return err
	}
	// defer updateRulesetNode.Close()
	_, err = updateRulesetNode.Exec(&value, &uuid, &param)
	if (err != nil){
		logs.Error("UpdateRuleFiles UPDATE error for update isDownloaded -- "+err.Error())
		// defer updateRulesetNode.Close()
		return err
	}
	return nil
}

func GetRulesetSourceValue(uuid string, param string)(val string, err error){
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

func GetAllLocalRulesetDB()(path []string, err error){
	var customData []string
    var uniqid string

	sql := "select ruleset_uniqueid from ruleset where ruleset_value='local' and ruleset_param='type'";
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllLocalRulesetDB Rdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid); err != nil {
            logs.Error("GetAllLocalRulesetDB -- Query return error: %s", err.Error())
            return nil, err
		}
		customData = append(customData, uniqid)
	} 
	return customData,nil
}

func GetAllDataRulesetDB(uuid string)(data map[string]map[string]string, err error){
	var customData = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string

	sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset where ruleset_uniqueid='"+uuid+"'";
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllDataRulesetDB Rdb.Query Error : %s", err.Error())
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("GetAllDataRulesetDB -- Query return error: %s", err.Error())
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

func GetRulesFromRuleset(uuid string) (data map[string]map[string]string, err error){
	var allRuleDetails = map[string]map[string]string{}
	var uniqid string
    var param string
    var value string
	var uuidSource string
    if Rdb == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
	}
	sqlUUID := "select rule_uniqueid from rule_files where rule_param='sourceUUID' and rule_value = '"+uuid+"';"
	uuidRows, err := Rdb.Query(sqlUUID)
	if err != nil {
		logs.Error("Rdb.Query Error checking uuid for take the uuid list for GetDetails: %s", err.Error())
        return nil, err
    }
	defer uuidRows.Close()
	for uuidRows.Next() {
		if err = uuidRows.Scan(&uuidSource); err != nil {
            logs.Error("GetDetails UUIDSource uuidRows.Scan: %s", err.Error())
            return nil, err
		}		
		sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='"+uuidSource+"';"
		rows, err := Rdb.Query(sql)
		if err != nil {
			logs.Error("Rdb.Query Error : %s", err.Error())
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&uniqid, &param, &value); err != nil {
				logs.Error("GetDetails rows.Scan: %s", err.Error())
				return nil, err
			}
			if allRuleDetails[uniqid] == nil { allRuleDetails[uniqid] = map[string]string{}}
			allRuleDetails[uniqid][param]=value
		} 
	}
	return allRuleDetails, nil
}

func GetRuleFilesByUniqueid(uuid string)(data map[string]map[string]string, err error){
	var allRuleDetails = map[string]map[string]string{}
	var uniqid string
    var param string
    var value string
	sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='"+uuid+"';"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetDetails rows.Scan: %s", err.Error())
			return nil, err
		}
		if allRuleDetails[uniqid] == nil { allRuleDetails[uniqid] = map[string]string{}}
		allRuleDetails[uniqid][param]=value
	} 
	return allRuleDetails, nil
}

func DeleteRulesetByUniqueid(uuid string)(err error){
	deleteRulesetQuery, err := Rdb.Prepare("delete from ruleset where ruleset_uniqueid = ?;")
	_, err = deleteRulesetQuery.Exec(&uuid)
	defer deleteRulesetQuery.Close()
    if err != nil {
		logs.Error("DB DeleteRulese/deleteRulesetQueryt -> ERROR on table Ruleset...")
        return errors.New("DB DeleteRuleset/deleteRulesetQuery -> ERROR on table Ruleset...")
	}
	return nil
}

func DeleteRulesetNodeByUniqueid(uuid string)(err error){
	deleteRulesetNodeQuery, err := Rdb.Prepare("delete from ruleset_node where ruleset_uniqueid = ?;")
	_, err = deleteRulesetNodeQuery.Exec(&uuid)
	defer deleteRulesetNodeQuery.Close()
    if err != nil {
		logs.Error("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
        return errors.New("DB DeleteRuleset/deleteRulesetNodeQuery -> ERROR on table Ruleset_node...")
	}
	return nil
}

func DeleteRuleFilesByUuid(uuid string)(err error){
	DeleteUUIDPrepare, err := Rdb.Prepare("delete from rule_files where rule_uniqueid = ?")
	if err != nil {
		logs.Error("DeleteRulese Rdb.Query Error deleting by rule_uniqueid for rule_files: %s", err.Error())
		return err
	}
	_, err = DeleteUUIDPrepare.Exec(&uuid)
	if err != nil {
		logs.Error("DeleteRulese deleting a ruleset source -> %s", err.Error())
		return err
	}
	return nil
}

func GetAllRuleFiles()(data map[string]map[string]string, err error){
	var allRuleDetails = map[string]map[string]string{}
	var uniqid string
    var param string
    var value string
	sql := "select rule_uniqueid, rule_param, rule_value from rule_files;"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllRuleFiles Rdb.Query Error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetAllRuleFiles rows.Scan error: %s", err.Error())
			return nil, err
		}
		if allRuleDetails[uniqid] == nil { allRuleDetails[uniqid] = map[string]string{}}
		allRuleDetails[uniqid][param]=value
	} 
	return allRuleDetails, nil
}

func GetAllNodes()(data map[string]string, err error){
	var ruleset string
	var node string
	values := make(map[string]string)
	sql := "select ruleset_uniqueid, node_uniqueid from ruleset_node;"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("GetAllNodes Rdb.Query Error: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&ruleset, &node); err != nil {
			logs.Error("GetAllNodes rows.Scan error: %s", err.Error())
			return nil, err
		}
		values[ruleset]=node
	} 
	return values, nil
}