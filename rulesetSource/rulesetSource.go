package rulesetSource

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
	"errors"
	"owlhmaster/utils"
)


func CreateRulesetSource(n map[string]string) (err error) {
	rulesetSourceKey := utils.Generate()
    if _, ok := n["name"]; !ok {
		logs.Error("name source empty: "+err.Error())
        return errors.New("name empty")
    }
    if _, ok := n["desc"]; !ok {
		logs.Error("desc source empty: "+err.Error())
        return errors.New("desc empty")
    }
    if _, ok := n["path"]; !ok {
		logs.Error("path source empty: "+err.Error())
        return errors.New("path empty")
    }

    if err := rulesetSourceExists(rulesetSourceKey); err != nil {
		logs.Error("rulesetSource exist: "+err.Error())
        return errors.New("rulesetSource exist")
    }
    
    for key, value := range n {
        err = rulesetSourceKeyInsert(rulesetSourceKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func rulesetSourceExists(sourceID string) (err error) {
    if ndb.RSdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    sql := "SELECT * FROM ruleset_source where source_uniqueid = '"+sourceID+"';"
    rows, err := ndb.RSdb.Query(sql)
    if err != nil {
        logs.Error("Error on query rulesetSourceExist at rulesetSource.go "+err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return errors.New("Ruleset source Exists " + sourceID)
    } else {
        return nil
    }
}

func rulesetSourceKeyInsert(nkey string, key string, value string) (err error) {
    if ndb.RSdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := ndb.RSdb.Prepare("insert into ruleset_source (source_uniqueid, source_param, source_value) values(?,?,?)")
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

func GetAllRulesetSource()(sources map[string]map[string]string, err error){
	var allsources = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    if ndb.RSdb == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
    }
    sql := "select source_uniqueid, source_param, source_value from ruleset_source;"
    rows, err := ndb.RSdb.Query(sql)
    if err != nil {
        logs.Error("ndb.RSdb.Query Error : %s", err.Error())
        return nil, err
    }
    for rows.Next() {
        if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("GetAllRulesetSource rows.Scan: %s", err.Error())
            return nil, err
        }
        if allsources[uniqid] == nil { allsources[uniqid] = map[string]string{}}
        allsources[uniqid][param]=value
	} 
    return allsources, nil
}

func DeleteRulesetSource(rulesetSourceUUID string) (err error) {
	if ndb.RSdb == nil {
        logs.Error("DeleteRulesetSource -- Can't acces to database")
        return errors.New("DeleteRulesetSource -- Can't acces to database")
    }
	sourceSQL, err := ndb.RSdb.Prepare("delete from ruleset_source where source_uniqueid = ?")
    if err != nil {
        logs.Error("Prepare DeleteRulesetSource -> %s", err.Error())
        return err
    }
    _, err = sourceSQL.Exec(&rulesetSourceUUID)
    if err != nil {
        logs.Error("Execute DeleteRulesetSource -> %s", err.Error())
        return err
    }
	return nil
}

func EditRulesetSource(data map[string]string) (err error) { 
	var sourceuuid = data["sourceuuid"]
    if ndb.RSdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
	}
	logs.Debug(data)
	for k,v := range data { 
		if k == "sourceuuid"{
			continue
		}else{
			err = InsertRulesetSource(k, v, sourceuuid)
			if err != nil {
				logs.Error("Error inserting edit files from source ruleset")
				return err
			}
		}
    }
	return nil
}

func InsertRulesetSource(param string, value string, sourceuuid string)(err error){
	editSource, err := ndb.RSdb.Prepare("update ruleset_source set source_value = ? where source_param = ? and source_uniqueid = ?")
	if err != nil {
		logs.Error("Prepare EditRulesetSource-> %s", err.Error())
		return err
	}
	_, err = editSource.Exec(&value, &param, &sourceuuid)
	if err != nil {
		logs.Error("Execute EditRulesetSource-> %s", err.Error())
		return err
	}
	return nil
}

func DownloadFile(data map[string]string) (err error) {	
	err = utils.DownloadFile(data["path"], data["url"])
	if err != nil {
		logs.Error("Error downloading file from RulesetSource-> %s", err.Error())
		return err
	}

	err = utils.ExtractTarGz(data["path"])
	if err != nil {
		logs.Error("Error unzipping file downloaded: "+err.Error())
        return err
	}
	logs.Info("Extract complete!")
	return nil
}

func CompareFiles(data map[string]string) (mapData map[string]map[string]string, err error) {	
	file1, err := utils.MapFromFile(data["new"])
	file2, err := utils.MapFromFile(data["old"])
	if err != nil {
		logs.Error("Error getting file from map: "+err.Error())
        return nil, err
	}
	var returnMap = make(map[string]map[string]string)
	
	lineExist := false

	//check if all the new lines are in old file
	for x,_ := range file1 { 
		returnLines := make(map[string]string)
		for y,_ := range file2 { 
			if x == y {
				lineExist = true
				if (file1[x]["Line"] != file2[y]["Line"] || file1[x]["Enabled"] != file2[y]["Enabled"]) {
					returnLines["new"] = file1[x]["Line"]
					returnLines["old"] = file2[y]["Line"]
					returnLines["enabled-new"] = file1[x]["Enabled"]
					returnLines["enabled-old"] = file2[y]["Enabled"]
					returnLines["sid"] = x
					returnMap[x] = returnLines
				}
				continue
			}
		}
		if !lineExist {
			returnLines["new"] = file1[x]["Line"]
			returnLines["old"] = "N/A"
			returnLines["enabled-new"] = file1[x]["Enabled"]
			returnLines["enabled-old"] = "N/A"
			returnLines["sid"] = x
			returnMap[x] = returnLines
		}

		lineExist = false
	}

	//check if all the old lines are in new file
	for y,_ := range file2 { 
		returnLines := make(map[string]string)
		for x,_ := range file1 {
			if y == x {
				lineExist = true
				continue
			}
		}
		if !lineExist {
			returnLines["new"] = "N/A"
			returnLines["old"] = file2[y]["Line"]
			returnLines["enabled-new"] = "N/A"
			returnLines["enabled-old"] = file2[y]["Enabled"]
			returnLines["sid"] = y
			returnMap[y] = returnLines
		}

		lineExist = false
	}

	
	return returnMap, nil
}


func CreateNewFile(data map[string]string) (err error) {
	logs.Info(data);
	//create an old file backup 
	//create new file with
    
    return nil
}