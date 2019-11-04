package ruleset

import(
    "fmt"
    "github.com/astaxie/beego/logs"
    "bufio"
    "regexp"
    "os"
	"os/exec"
	"encoding/json"
    "owlhmaster/database"
    "owlhmaster/utils"
    "owlhmaster/node"
    "errors"
	"database/sql"
    "strings"
    "time"
	"strconv"
	"io/ioutil"
)

type LinesID struct {
	Counter     int `json:"counter"`
	Values		[]Values `json:"values"`
}
type Values struct {
	Path     	string `json:"path"`
	FileName    string `json:"fileName"`
	Line		string `json:"line"`
}


//read rule raw data
func ReadSID(sid map[string]string)(sidLine map[string]string ,err error){
    sidMap := sid["sid"]
    uuidMap := sid["uuid"]
    path, err := ndb.GetRulesetPath(uuidMap)
    data, err := os.Open(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return
	}

    var validID = regexp.MustCompile(`sid:`+sidMap+`;`)
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        if validID.MatchString(scanner.Text()){
            sidLine := make(map[string]string)
            sidLine["raw"] = scanner.Text()
            return sidLine,err
        }
    }
    return nil,err
}

//Read ruleset rules data
func ReadRuleset(path string)(rules map[string]map[string]string, err error) {
    data, err := os.Open(path)
    if err != nil {fmt.Println("File reading error", err) }

    var validID = regexp.MustCompile(`sid:(\d+);`)
    var msgfield = regexp.MustCompile(`msg:\"([^"]+)\"`)
    var ipfield = regexp.MustCompile(`^([^\(]+)\(`)
	var enablefield = regexp.MustCompile(`^#`)

    scanner := bufio.NewScanner(data)
    rules = make(map[string]map[string]string)
    for scanner.Scan(){
        if validID.MatchString(scanner.Text()){
            sid := validID.FindStringSubmatch(scanner.Text())
            msg := msgfield.FindStringSubmatch(scanner.Text())
            ip := ipfield.FindStringSubmatch(scanner.Text())
            rule := make(map[string]string)
            if enablefield.MatchString(scanner.Text()){
                rule["enabled"]="Disabled"
            }else{
                rule["enabled"]="Enabled"
            }
            rule["sid"]=sid[1]
            rule["msg"]=msg[1]
            rule["ip"]=ip[1]
            rule["raw"]=scanner.Text()
            rules[sid[1]]=rule
        }
    }
    return rules,err
}

//Add new ruleset
func AddRuleset(n map[string]string) (err error) {
    rulesetID := utils.Generate()
    //Verify parameters
	if n["name"] == "" {
        return errors.New("Name is empty")
	}
	if n["desc"] == "" {
        return errors.New("Description is empty")
    }
    //Verify that the ruleset exists
    if err := rulesetExists(rulesetID); err != nil {
        return err
    }
    //Insert new ruleset into DB
    for key, value := range n {
        err = rulesetInsert(rulesetID, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

//Check if a specific ruleset exists
func rulesetExists(rulesetID string) (err error) {
    if ndb.Rdb == nil {
        logs.Error("rulesetExists -- Can't access to database")
        return errors.New("rulesetExists -- Can't access to database")
    }
    sql := "SELECT * FROM ruleset where ruleset_uniqueid = '"+rulesetID+"';"
    rows, err := ndb.Rdb.Query(sql)
    if err != nil {
        logs.Error(err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return errors.New("rulesetExists -- RulesetId exists")
    } else {
        return nil
    }
}

//Insert a new ruleset
func rulesetInsert(nkey string, key string, value string) (err error) {
    if ndb.Rdb == nil {
        logs.Error("rulesetInsert -- Can't access to database")
        return errors.New("rulesetInsert -- Can't access to database")
    }
    stmt, err := ndb.Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values(?,?,?)")
    if err != nil {
        logs.Error("rulesetInsert -- Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&nkey, &key, &value)
    if err != nil {
        logs.Error("rulesetInsert -- Execute -> %s", err.Error())
        return err
    }
    return nil
}

//Get all rulesets from DB
func GetAllRulesets() (rulesets map[string]map[string]string, err error) {
    var allrulesets = map[string]map[string]string{}
    var uniqid string
    var param string
	var value string
	var uuidLocal string
    if ndb.Rdb == nil {
        logs.Error("ruleset/GetAllRulesets -- Can't access to database")
        return nil, errors.New("ruleset/GetAllRulesets -- Can't access to database")
    }
    sql := "select ruleset_uniqueid from ruleset where ruleset_param='type' and ruleset_value = 'local';"
    rows, err := ndb.Rdb.Query(sql)
    if err != nil {
        logs.Error("ruleset/GetAllRulesets -- Query error: %s", err.Error())
        return nil, err
	}
	defer rows.Close()
    for rows.Next() {
        if err = rows.Scan(&uuidLocal); err != nil {
            logs.Error("ruleset/GetAllRulesets -- Query return error: %s", err.Error())
            return nil, err
		}
		sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset where ruleset_uniqueid='"+uuidLocal+"';"
			rowsLocal, err := ndb.Rdb.Query(sql)
			if err != nil {
				logs.Error("GetAllRulesets ndb.Rdb.Query Error : %s", err.Error())
				return nil, err
			}
			defer rowsLocal.Close()
			for rowsLocal.Next() {
				if err = rowsLocal.Scan(&uniqid, &param, &value); err != nil {
					logs.Error("GetAllRulesets rowsLocal.Scan: %s", err.Error())
					return nil, err
				}
				if allrulesets[uniqid] == nil { allrulesets[uniqid] = map[string]string{}}
				allrulesets[uniqid][param]=value

				//add scheduler status
				schedulerUUID,err := ndb.GetSchedulerByValue(uniqid)
				if err != nil {
					logs.Error("GetAllRulesets GetSchedulerByValue: %s", err.Error())
					return nil, err
				}
				schedulerData,err := ndb.GetSchedulerByUniqueid(schedulerUUID)
				if err != nil {
					logs.Error("GetAllRulesets GetSchedulerByUniqueid: %s", err.Error())
					return nil, err
				}
				allrulesets[uniqid]["status"] = schedulerData[schedulerUUID]["status"]

			} 
	}	

	return allrulesets, nil
}

//Get rules from specific ruleset
func GetRulesetRules(uuid string)(r map[string]map[string]string, err error){
    rules := make(map[string]map[string]string)
	path,err := ndb.GetRulesetPath(uuid)
    rules,err = ReadRuleset(path)
    for rule, _ := range rules{
        retrieveNote := make(map[string]string)
        retrieveNote["uuid"] = uuid
        retrieveNote["sid"] = rule
        rules[rule]["note"], _ = GetRuleNote(retrieveNote)
		sourceType,err := ndb.GetRuleFilesValue(uuid, "sourceType")
		if err != nil {
			logs.Error("GetRulesetRules--> GetRuleFilesValue query error %s",err.Error())
			return nil,err
		}
		rules[rule]["sourceType"] = sourceType
    }
    return rules, err
}

//Set a selected ruleset to node
func SetRuleSelected(n map[string]string) (err error) {
    node_uniqueid_ruleset := n["nid"]
    ruleset_uniqueid := n["rule_uid"]

    if ndb.Rdb == nil {
        logs.Error("SetRuleSelected -- Can't access to database")
        return errors.New("SetRuleSelected -- Can't access to database")
    }
    sqlQuery := "SELECT * FROM ruleset_node WHERE node_uniqueid = \""+node_uniqueid_ruleset+"\";"
    rows, err := ndb.Rdb.Query(sqlQuery)
    if err != nil {
        logs.Error("Put SetRuleSelecteda query error %s",err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        rows.Close()
        updateRulesetNode, err := ndb.Rdb.Prepare("update ruleset_node set ruleset_uniqueid = ? where node_uniqueid = ?;")
        if (err != nil){
            logs.Error("SetRuleSelected UPDATE prepare error -- "+err.Error())
            return err
        }
        _, err = updateRulesetNode.Exec(&ruleset_uniqueid, &node_uniqueid_ruleset)
        defer updateRulesetNode.Close()

        if (err != nil){
            logs.Error("SetRuleSelected UPDATE Error -- "+err.Error())
            return err
        }
        return nil
    } else {
        insertRulesetNode, err := ndb.Rdb.Prepare("insert into ruleset_node (ruleset_uniqueid, node_uniqueid) values (?,?);")
        _, err = insertRulesetNode.Exec(&ruleset_uniqueid, &node_uniqueid_ruleset)
        defer insertRulesetNode.Close()

        if (err != nil){
            logs.Error("error insertRulesetNode en ruleset/rulesets--> "+err.Error())
            return err
        }
        return nil
    }
    return err
}

//Get a specific ruleset
func GetRuleSelected(nid string)(ruleset string, err error){
    var ruleSelected string
    if ndb.Rdb != nil {
        row := ndb.Rdb.QueryRow("SELECT ruleset_uniqueid FROM ruleset_node WHERE node_uniqueid = \""+nid+"\";")
        err = row.Scan(&ruleSelected)
        if err == sql.ErrNoRows{
            logs.Error("GetRuleName -> There is no ruleset with thie UUID %s", nid)
            return "", err
        }
        if err != nil {
            logs.Error("GetRuleName -> row.Scan error %s", err.Error())
            return "", err
        }
        return ruleSelected, nil
    }else {
        logs.Error("GetRuleSelected -> No access to database")
        return "", err
    }
}

//Get a specific rule name
func GetRuleName(nid string)(ruleset string, err error){
    var nameRule string
    if ndb.Rdb != nil {
        row := ndb.Rdb.QueryRow("SELECT ruleset_value FROM ruleset WHERE ruleset_uniqueid = \""+nid+"\" and ruleset_param = \"name\";")
        err = row.Scan(&nameRule)
        if err == sql.ErrNoRows{
            logs.Error("GetRuleName -> param or param doesn't exists")
            return "", err
        }
        if err != nil {
            logs.Error("GetRuleName -> row.Scan error %s", err.Error())
            return "", err
        }
        return nameRule, nil
    }else {
        logs.Error("GetRuleName -> no access to database")
        return "", err
    }
}

//clone ruleset
func SetClonedRuleset(ruleCloned map[string]string)(err error){
    if ndb.Rdb == nil {
        logs.Error("rulesetExists -- Can't access to database")
        return errors.New("rulesetExists -- Can't access to database")
	}
	//load path from main.conf
	loadPath := map[string]map[string]string{}
	loadPath["ruleset"] = map[string]string{}
	loadPath["ruleset"]["path"] = ""
	loadPath,err = utils.GetConf(loadPath)
	path := loadPath["ruleset"]["path"]

	// clonedRuleset := ruleCloned["cloned"]
    newName := ruleCloned["newName"]
    newFile := ruleCloned["newFile"]
    newDesc := ruleCloned["newDesc"]
    clonedPath := ruleCloned["path"]
    newRulesetFile := strings.Replace(newFile, " ", "_", -1)

    pathNewRule := path+newRulesetFile+".rules"
    newUUID := utils.Generate()

    rows, err := ndb.Rdb.Query("SELECT * FROM ruleset WHERE ruleset_uniqueid = \""+newUUID+"\";")
    defer rows.Close()
    if !rows.Next(){
        cpCmd := exec.Command("cp", clonedPath, pathNewRule)
        err = cpCmd.Run()
        if err != nil{
            logs.Error("SetClonedRuleset --> Error exec cmd command: "+err.Error())
            return err
		}
		
		err = insertRulesetValues(newUUID, "name", newName)
		err = insertRulesetValues(newUUID, "file", newFile)
		err = insertRulesetValues(newUUID, "desc", newDesc)
		err = insertRulesetValues(newUUID, "path", pathNewRule)
		err = insertRulesetValues(newUUID, "type", "local")
        if (err != nil){
            logs.Error("error insertRulesetValues values on ruleset/rulesets--> "+err.Error())
            return err
        }

        return nil
    }
    if err != nil {
        logs.Error("SetClonedRuleset -> rows.Scan %s", err.Error())
        return err
    }
    return nil
}

//insert values to ruleset table
func insertRulesetValues(uuid string, param string, value string)(err error){
	insertRulesetValues, err := ndb.Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values (?,?,?);")
	_, err = insertRulesetValues.Exec(&uuid, &param, &value)
	defer insertRulesetValues.Close()
	if (err != nil){
		return err
	}
	return nil
}

//Change rule status to enabled or disabled
func SetRulesetAction(ruleAction map[string]string)(err error){
    sid := ruleAction["sid"]
    uuid := ruleAction["uuid"]
    action := ruleAction["action"]
    path, err := ndb.GetRulesetPath(uuid)
    if (action == "Enable"){
        cmd := "sed -i '/sid:"+sid+"/s/^#//' "+path+""
        _, err := exec.Command("bash", "-c", cmd).Output()
        if err == nil {
            return nil
        }
    }else{
        cmd := "sed -i '/sid:"+sid+"/s/^/#/' "+path+""
        _, err := exec.Command("bash", "-c", cmd).Output()
        if err == nil {
            return nil
        }
    }
    return err
}

//Add notes to specific rule
func SetRuleNote(ruleNote map[string]string)(err error){
    if ndb.Rdb == nil {
        logs.Error("SetRuleNote -- Can't access to database")
        return errors.New("SetRuleNote -- Can't access to database")
    }
    sid := ruleNote["sid"]
    uuid := ruleNote["uuid"]
    note := ruleNote["note"]
    t := time.Now()
    noteTime := strconv.FormatInt(t.Unix(), 10)

    sqlQuery := "SELECT * FROM rule_note WHERE ruleset_uniqueid = \""+uuid+"\" and rule_sid = \""+sid+"\";"
    rows, err := ndb.Rdb.Query(sqlQuery)
    if err != nil {
        logs.Error("Put SetRuleNote query error %s",err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        rows.Close()
        updateNote, err := ndb.Rdb.Prepare("update rule_note set ruleNote = ?, note_date = ? where ruleset_uniqueid = ? and rule_sid = ? ;")
        _, err = updateNote.Exec(&note, &noteTime, &uuid, &sid)
        defer updateNote.Close()

        if (err != nil){
            logs.Error("SetRuleNote UPDATE Error -- "+err.Error())
            return err
        }
        return nil

    } else {
        insertNote, err := ndb.Rdb.Prepare("insert into rule_note (ruleset_uniqueid, rule_sid, note_date, ruleNote) values (?,?,?,?);")
        _, err = insertNote.Exec(&uuid, &sid, &noteTime, &note)
        defer insertNote.Close()

        if (err != nil){
            logs.Error("error insertRulesetNode ruleset/rulesets--> "+err.Error())
            return err
        }
        return nil
    }
    return err
}

//Get note from specific rule
func GetRuleNote(ruleGetNote map[string]string)(note string, err error){
    sidMap := ruleGetNote["sid"]
    uuidMap := ruleGetNote["uuid"]
    var noteText string
    if ndb.Rdb == nil {
        logs.Error("GetRuleNote -- Can't access to database")
        return "", errors.New("GetRuleNote -- Can't access to database")
    }
    row := ndb.Rdb.QueryRow("SELECT ruleNote FROM rule_note WHERE ruleset_uniqueid=\""+uuidMap+"\" and rule_sid=\""+sidMap+"\";")
	err = row.Scan(&noteText)

    if err != nil {
        // logs.Error("DB GetNote -> Can't read query result: "+err.Error())
        return "", err
    }
    return noteText, nil
}

//Delete specific ruleset and all their asociated files
func DeleteRuleset(rulesetMap map[string]string)(err error){
	uuid := rulesetMap["uuid"]
	name := rulesetMap["name"]
	rulesetFolderName := strings.Replace(name, " ", "_", -1)
	// var uniqueid string
	// var uuidArray []string	

	localRulesets := map[string]map[string]string{}
	localRulesets["ruleset"] = map[string]string{}
	localRulesets["ruleset"]["localRulesets"] = ""
	localRulesets,err = utils.GetConf(localRulesets)
	localRulesetFiles := localRulesets["ruleset"]["localRulesets"]

	//delete LOG for scheduler
	err = ndb.DeleteSchedulerLog(uuid)
	if err != nil {
		logs.Error("Error deleting LOG DeleteSchedulerLog: "+err.Error())
        return err
	}

	//delete scheduler
	schedulerUUID,err := ndb.GetSchedulerByValue(uuid)
	if err != nil { logs.Error("Error getting scheduler uuid GetSchedulerByValue: "+err.Error()); return err}

	err = ndb.DeleteScheduler(schedulerUUID)
	if err != nil {logs.Error("Error deleting scheduler uuid DeleteSchedulerLog: "+err.Error()); return err}

	//delete ruleset
	err = ndb.DeleteRulesetByUniqueid(uuid)
	if err != nil {logs.Error("DeleteRulesetByUniqueid -> ERROR deleting ruleset: "+err.Error()); return err}

	//delete a node ruleset
	err = ndb.DeleteRulesetNodeByUniqueid(uuid)
	if err != nil {logs.Error("DeleteRulesetNodeByUniqueid -> ERROR deleting ruleset: "+err.Error());return err}

	//delete ruleset from path
	err = os.RemoveAll(localRulesetFiles+rulesetFolderName)
	if err != nil {
		logs.Error("DB DeleteRuleset/rm -> ERROR deleting ruleset from their path...")
		return errors.New("DB DeleteRuleset/rm -> ERROR deleting ruleset from their path...")
	}

	//delete all ruleset source rules for specific uuid
	rules,err := ndb.GetRulesFromRuleset(uuid)
	if err != nil {logs.Error("GetRulesFromRuleset -> ERROR getting all rule_files for delete local ruleset: "+err.Error());return err}
	
	for sourceUUID := range rules{
		err = ndb.DeleteRuleFilesByUuid(sourceUUID)
		if err != nil {logs.Error("DeleteRuleFilesByUuid -> ERROR deleting all local ruleset rule files associated: "+err.Error());return err}
	}

	//update to nil group ruleset
	rulesetsForGroups, err := ndb.GetAllGroupsBValue(uuid)
	if err != nil {logs.Error("GetAllGroupsBValue -> ERROR getting all groups by ruleset uuid: "+err.Error()); return err}
	
	for y := range rulesetsForGroups {
		err = ndb.UpdateGroupValue(y, "ruleset", "")
		if err != nil {logs.Error("Error updating to null rulesets into group table: "+err.Error()); return err}
		err = ndb.UpdateGroupValue(y, "rulesetID", "")
		if err != nil {logs.Error("Error updating to null rulesetsID into group table: "+err.Error()); return err}
	}


	// uuidRules, err := ndb.Rdb.Query("select rule_uniqueid from rule_files where rule_value='"+uuid+"'")
	// if err != nil {
	// 	logs.Error("DeleteRulese ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
	// 	return err
	// }
	// defer uuidRules.Close()
	// for uuidRules.Next() {
	// 	if err = uuidRules.Scan(&uniqueid); err != nil {
	// 		logs.Error("DeleteRulese rows.Scan: %s", err.Error())
	// 		return err
	// 	}
	// 	uuidArray = append(uuidArray, uniqueid)
	// }
	// for x := range uuidArray{
	// 	err = ndb.DeleteRuleFilesByUuid(uuidArray[x])
	// 	if err != nil {
	// 		logs.Error("DeleteRuleset ndb.Rdb.Query Error deleting by rule_uniqueid for rule_files: %s", err.Error())
	// 		return err
	// 	}
	// }

    return nil
}

//Get all source rulesets from DB
func GetAllRuleData()(data map[string]map[string]string,err error) {
    var allRuleDetails = map[string]map[string]string{}
    var uniqid string
	if ndb.Rdb == nil { logs.Error("ruleset/GetAllRuleData -- Can't access to database"); return nil, errors.New("ruleset/GetAllRuleData -- Can't access to database")}
		
    sql := "select rule_uniqueid from rule_files where rule_param='type' and rule_value='source';"
    rows, err := ndb.Rdb.Query(sql)
    if err != nil {
        logs.Error("ruleset/GetAllRuleData -- Query error: %s", err.Error())
        return nil, err
    }
    for rows.Next() {
		if err = rows.Scan(&uniqid); err != nil { logs.Error("ruleset/GetAllRuleData -- Query return error: %s", err.Error()); return nil, err}
		
		var uniqidSub string
		var paramSub string
		var valueSub string
		sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='"+uniqid+"';"
		rows, err := ndb.Rdb.Query(sql)
		if err != nil {
			logs.Error("Rdb.Query Error : %s", err.Error())
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&uniqidSub, &paramSub, &valueSub); err != nil {
				logs.Error("GetDetails rows.Scan: %s", err.Error())
				return nil, err
			}
			if allRuleDetails[uniqidSub] == nil { allRuleDetails[uniqidSub] = map[string]string{}}
			allRuleDetails[uniqidSub][paramSub]=valueSub
		} 
	}
	return allRuleDetails, nil
}

func FindDuplicatedSIDs(data map[string]map[string]string)(duplicated []byte, err error){
	allSids := make(map[string]LinesID)
	allSidsResult := make(map[string]LinesID)
	
	for x := range data {
		sidLines,err := ReadRuleset(data[x]["filePath"])
		if err != nil {
			logs.Error("ERROR --> "+err.Error())
			return nil,err
		}

		for y := range sidLines {
			values := Values{}
			linesID := LinesID{}
			values.Path = data[x]["filePath"]
			values.FileName = data[x]["fileName"]
			values.Line = sidLines[y]["raw"]
			if _, exists := allSids[y]; exists { //exist				
				linesID.Counter = allSids[y].Counter 
				linesID.Counter += 1
				linesID.Values = allSids[y].Values
				linesID.Values = append(linesID.Values, values)
				allSids[y]=linesID
			}else{ //not exist
				linesID.Counter = 1
				linesID.Values = append(linesID.Values, values)
				allSids[y]=linesID
			}
		}
	}
	//create response array
	for n := range allSids{
		if allSids[n].Counter > 1 {
			allSidsResult[n] = allSids[n]
		}
	}
	//check if response array is empty
	if len(allSidsResult) == 0{
		return nil,nil
	}else{
		LinesOutput, err := json.Marshal(allSidsResult)
		if err != nil {
			logs.Error("ERROR Marshal allSidsResult --> "+err.Error())
			return nil,err
		}
		return LinesOutput, nil
	}
}

//Add new ruleset to locale ruleset
func AddNewRuleset(data map[string]map[string]string)(duplicated []byte, err error) {
	//check for duplicated rule SIDs
	if duplicated,err = FindDuplicatedSIDs(data); duplicated != nil {
		return duplicated, nil
	}
	if err != nil {
		logs.Error("ruleset/AddNewRuleset -- duplicated error: %s", err.Error())
		return nil,err
	}

    if ndb.Rdb == nil {
        logs.Error("ruleset/AddNewRuleset -- Can't access to database")
        return nil,errors.New("ruleset/AddNewRuleset -- Can't access to database")
	}
	
	localRulesets := map[string]map[string]string{}
	localRulesets["ruleset"] = map[string]string{}
	localRulesets["ruleset"]["localRulesets"] = ""
	localRulesets,err = utils.GetConf(localRulesets)
	localFiles := localRulesets["ruleset"]["localRulesets"]

	rulesetUUID := utils.Generate()
	rulesetCreated := false
	
	for x := range data {		
		rulesetFolderName := strings.Replace(data[x]["rulesetName"], " ", "_", -1)
		path := localFiles + rulesetFolderName + "/" + data[x]["fileName"]

		if !rulesetCreated {
			err = insertRulesetValues(rulesetUUID, "type", "local")
			err = insertRulesetValues(rulesetUUID, "name", data[x]["rulesetName"])
			err = insertRulesetValues(rulesetUUID, "desc", data[x]["rulesetDesc"])
			if err != nil {
				logs.Error("ruleset/AddNewRuleset -- Insert error: %s", err.Error())
				return nil,err
			}
			rulesetCreated = true
		}
				
		//copy source file into new folder
		if _, err := os.Stat(localFiles + rulesetFolderName); os.IsNotExist(err) {
			os.MkdirAll(localFiles + rulesetFolderName, os.ModePerm)
		}
		
		//copyfile
		cpCmd := exec.Command("cp", data[x]["filePath"], path)
    	err = cpCmd.Run()
		if err != nil {
			logs.Error("ruleset/AddNewRuleset -- Error copying new file: %s", err.Error())
			return nil,err
		}

		//add md5 for every file
		md5,err := utils.CalculateMD5(path)
		if err != nil {logs.Error("ruleset/AddNewRuleset -- Error calculating md5: %s", err.Error());return nil,err}

		ruleFilesUUID := utils.Generate()
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "name", data[x]["rulesetName"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "path", path)
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "file", data[x]["fileName"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "type", "local")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceUUID", rulesetUUID)
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceFileUUID", x)
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "exists", "true")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "isUpdated", "false")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "md5", md5)
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceType", data[x]["sourceType"])
		if err != nil {logs.Error("ruleset/AddNewRuleset -- Error Inserting Ruleset: %s", err.Error());return nil,err}
	}

	return nil,nil
}

func GetAllCustomRulesets()(data map[string]map[string]string, err error) {
	customData := make(map[string]map[string]string)
	uuid,err := ndb.GetAllCustomRulesetDB()
	for x := range uuid {
		retrievedData,err := ndb.GetAllDataRulesetDB(uuid[x])
		customData[uuid[x]] = retrievedData[uuid[x]]
		if err != nil {
			logs.Error("GetAllCustomRulesets -- Error retrieving Custom ruleset data: %s", err.Error())
			return nil,err
		}
	}
	return customData, nil
}

func AddRulesToCustomRuleset(anode map[string]string)(duplicatedRules map[string]string, err error) {
	rulesDuplicated := make(map[string]string)
	sidsSplit := strings.Split(anode["sids"], ",")
	path,err := ndb.GetRulesetSourceValue(anode["dest"], "path") 
	for uuid := range sidsSplit{
		var validID = regexp.MustCompile(`sid:`+sidsSplit[uuid]+`;`)
		readSidsData := make(map[string]string)
		readSidsData["sid"] = sidsSplit[uuid]
		readSidsData["uuid"] = anode["orig"]
		readSidsData["action"] = "Disable"
		
		sidLine,err := ReadSID(readSidsData)
		if err != nil {
			logs.Error("AddRulesToCustomRuleset -- Error readding SID: %s", err.Error())
			return nil,err
		}

		// path,err := ndb.GetRulesetSourceValue(anode["dest"], "path") 
		if err != nil {
			logs.Error("AddRulesToCustomRuleset -- Error getting GetRulesetSourceValue: %s", err.Error())
			return nil,err
		}
		file, err := os.Open(path)
		defer file.Close()

		scanner := bufio.NewScanner(file) 
    	for scanner.Scan() { 
			if validID.MatchString(scanner.Text()){ 
				rulesDuplicated[sidsSplit[uuid]] = scanner.Text()
			}
		}

		if rulesDuplicated[sidsSplit[uuid]] == "" {
			var EnabledRule = regexp.MustCompile(`^#`)
			rulePath,err := ndb.GetRuleFilesValue(anode["orig"], "path")

			//change destiny status to Enable
			writeFile,err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer writeFile.Close()

			str := EnabledRule.ReplaceAllString(sidLine["raw"],"")

			_, err = writeFile.WriteString(str)
			_, err = writeFile.WriteString("\n")

			if err != nil {
				logs.Error("AddRulesToCustomRuleset -- Error getting origin path: %s", err.Error())
				return nil,err
			}

			originFile, err := os.Open(rulePath)
			defer originFile.Close()	

			//change origin status to Disable
			scanner := bufio.NewScanner(originFile) 
			for scanner.Scan() { 
				if validID.MatchString(scanner.Text()) {
					if !EnabledRule.MatchString(scanner.Text()) {
						err = SetRulesetAction(readSidsData)
						if err != nil {
							logs.Error("AddRulesToCustomRuleset -- SetRulesetAction Error writting data: %s", err.Error())
							return nil,err
						}	
					}
				}
			}
		}
	}

	valuesCustom,_ := ndb.GetAllDataRulesetDB(anode["dest"])
	for a,_ := range valuesCustom {
		md5,err := utils.CalculateMD5(valuesCustom[a]["path"])
		if err != nil {logs.Error("ruleset/AddRulesToCustomRuleset -- Error calculating md5: %s", err.Error());return nil,err}
	
		//add destination custom ruleset to locale ruleset who clone rules.
		ruleFilesUUID := utils.Generate()
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "name", valuesCustom[a]["name"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "path", valuesCustom[a]["path"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "file", valuesCustom[a]["fileName"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "type", "local")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceUUID", anode["ruleset"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceFileUUID", anode["dest"])
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "exists", "true")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "isUpdated", "false")
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "md5", md5)
		err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceType", valuesCustom[a]["sourceType"])
		if err != nil {logs.Error("AddRulesToCustomRuleset -- Error inserting custom ruleset into local ruleset: %s", err.Error()); return nil,err}
	}


	return rulesDuplicated, nil	
}

func ReadRulesetData(uuid string)(content map[string]string, err error) {
	path,err := ndb.GetRulesetPath(uuid)
	//get file path
	fileReaded, err := ioutil.ReadFile(path)
    if err != nil {
		logs.Error("ReadRulesetData Error reading ruleset content: "+err.Error())
        return nil,err
    }
	sendBackFile := make(map[string]string)
	sendBackFile["fileContent"] = string(fileReaded)
	return sendBackFile,nil
}

func SaveRulesetData(anode map[string]string)(err error) {
	uuid := anode["uuid"]
	content := anode["content"]
	
	path,err := ndb.GetRulesetPath(uuid)

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
    if err != nil {
		logs.Error("SaveRulesetData failed opening file: %s", err)
		return err
    }
	defer file.Close()
	file.Truncate(0)
	file.Seek(0,0)
    _, err = file.WriteAt([]byte(content), 0) // Write at 0 beginning
    if err != nil {
		logs.Error("SaveRulesetData failed writing to file: %s", err)
		return err
    }
	
	return nil
}

func SynchronizeAllRulesets()(err error){
	uuids,err := ndb.GetAllLocalRulesetDB()
	for x := range uuids{
		anode := make(map[string]string)
		anode["uuid"] = uuids[x]
		err = node.SyncRulesetToAllNodes(anode)
		if err != nil {logs.Error("Error SynchronizeAllRulesets: "+err.Error()); return err}
	}
    return nil
}

func UpdateRule(anode map[string]string)(err error) {

	if anode["sid"] == "" {logs.Error("UpdateRule error checking SID numbers: SID number is nil"); return errors.New("Error checking SID numbers: SID number is nil")}
	var numbers = regexp.MustCompile(`^[0-9]*$`)
	sidValue := numbers.FindStringSubmatch(anode["sid"])
	if sidValue == nil {logs.Error("UpdateRule error checking SID numbers: SID doesn't have only numbers"); return errors.New("Error checking SID numbers: SID doesn't have only numbers")}

	path,err := ndb.GetRulesetPath(anode["uuid"])
	if err != nil {logs.Error("UpdateRule/GetRulesetPath Error: "+err.Error()); return err}

	anode["line"] = strings.Replace(anode["line"], "/", "\\/", -1)
	cmd := "sed -i '/sid:"+anode["sid"]+"/s/.*/"+anode["line"]+"/' "+path+""
	_, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {logs.Error("UpdateRule Error: "+err.Error()); return err}

   return nil
}