package ruleset

import(
    // "fmt"
    "github.com/astaxie/beego/logs"
    "bufio"
    "regexp"
    "os"
    "os/exec"
    "encoding/json"
    "owlhmaster/database"
    "owlhmaster/nodeclient"
    "owlhmaster/rulesetSource"
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
    Values      []Values `json:"values"`
}
type Values struct {
    Path        string `json:"path"`
    Enabled     string `json:"enabled"`
    FileName    string `json:"fileName"`
    Line        string `json:"line"`
    Sid         string `json:"sid"`
}


//read rule raw data
func ReadSID(sid map[string]string)(sidLine map[string]string ,err error){
    logs.Info("SID detail %v",sid)
    sidMap := sid["sid"]
    uuidMap := sid["uuid"]
    path, err := ndb.GetRulesetPath(uuidMap)
    data, err := os.Open(path)
    if err != nil {
        logs.Error("File reading error: %s", err.Error())
        return
    }

    var validID = regexp.MustCompile(`sid:\s?`+sidMap+`;`)
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
    if err != nil {
        logs.Error("File reading error %s", err.Error()) 
    }
    
    var validID = regexp.MustCompile(`sid:\s?(\d+);`)
    var msgfield = regexp.MustCompile(`msg:\s?\"([^"]+)\"`)
    var ipfield = regexp.MustCompile(`^([^\(]+)\(`)
    var enablefield = regexp.MustCompile(`^#`)

    scanner := bufio.NewScanner(data)
    rules = make(map[string]map[string]string)
    for scanner.Scan(){
        replaceFirst := strings.Replace(scanner.Text(), "“", "\"", -1)
        replaceSecond := strings.Replace(replaceFirst, "”", "\"", -1)
        if validID.MatchString(replaceSecond){
            sid := validID.FindStringSubmatch(replaceSecond); if len(sid) == 0 { logs.Error("ReadRuleset error: SID not found "+replaceSecond); continue }
            msg := msgfield.FindStringSubmatch(replaceSecond); if len(msg) == 0 { logs.Error("ReadRuleset error: MSG not found "+replaceSecond); continue }
            ip := ipfield.FindStringSubmatch(replaceSecond); if len(ip) == 0 { logs.Error("ReadRuleset error: RULE header not found "+replaceSecond); continue }

            rule := make(map[string]string)
            if enablefield.MatchString(replaceSecond){
                rule["enabled"]="Disabled"
            }else{
                rule["enabled"]="Enabled"
            }
            rule["sid"]=sid[1]
            rule["msg"]=msg[1]
            rule["ip"]=ip[1]
            rule["raw"]=replaceSecond
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
            logs.Warn("GetRuleSelected -> There is no ruleset with thie UUID %s", nid)
            return "", err
        }
        if err != nil {
            logs.Warn("GetRuleSelected -> row.Scan error %s", err.Error())
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
            logs.Warn("GetRuleName -> param or param doesn't exists")
            return "", err
        }
        if err != nil {
            logs.Warn("GetRuleName -> row.Scan error %s", err.Error())
            return "", err
        }
        return nameRule, nil
    }else {
        logs.Warn("GetRuleName -> no access to database")
        return "", err
    }
}

//clone ruleset
func SetClonedRuleset(ruleCloned map[string]string)(err error){
    if ndb.Rdb == nil {logs.Error("rulesetExists -- Can't access to database"); return errors.New("rulesetExists -- Can't access to database")}

    copy, err := utils.GetKeyValueString("execute", "copy")
    if err != nil {logs.Error("SetClonedRuleset Error getting data from main.conf"); return err}
    path, err := utils.GetKeyValueString("ruleset", "path")
    if err != nil {logs.Error("SetClonedRuleset Error getting data from main.conf for load data: "+err.Error())}

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
        cpCmd := exec.Command(copy, clonedPath, pathNewRule)
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
    cmd, err := utils.GetKeyValueString("execute", "command")
    if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return err}
    param, err := utils.GetKeyValueString("execute", "param")
    if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return err}

    sid := ruleAction["sid"]
    uuid := ruleAction["uuid"]
    action := ruleAction["action"]
    path, err := ndb.GetRulesetPath(uuid)
    if (action == "Enable"){
        val := "sed -i '/sid:"+sid+"/s/^#//' "+path+""
        _, err := exec.Command(cmd, param, val).Output()
        if err == nil {
            return nil
        }
    }else{
        val := "sed -i '/sid:"+sid+"/s/^/#/' "+path+""
        _, err := exec.Command(cmd, param, val).Output()
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

    localRulesetFiles, err := utils.GetKeyValueString("ruleset", "localRulesets")
    if err != nil {logs.Error("DeleteRuleset Error getting data from main.conf for load data: "+err.Error()); return err}

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

    //select all groups
    groups,err := ndb.GetAllGroups()    
    if err != nil {logs.Error("DeleteRulesetNodeByUniqueid -> ERROR getting all groups: "+err.Error());return err}
    groupsRulesets,err := ndb.GetAllGroupRulesets()    
    if err != nil {logs.Error("DeleteRulesetNodeByUniqueid -> ERROR getting all grouprulesets: "+err.Error());return err}
    for id := range groups {
        for grid := range groupsRulesets {
            if groupsRulesets[grid]["groupid"] == id && groupsRulesets[grid]["rulesetid"] == uuid {
                //delete a node ruleset
                err = ndb.DeleteGroupRulesetByValue("groupid", id)
                if err != nil {logs.Error("DeleteRulesetNodeByUniqueid -> ERROR deleting grouprulesets: "+err.Error());return err}
                err = ndb.DeleteGroupRulesetByValue("rulesetid", uuid)
                if err != nil {logs.Error("DeleteRulesetNodeByUniqueid -> ERROR deleting grouprulesets: "+err.Error());return err}
            }
        }
        
    }

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
                logs.Error("GetAllRuleData rows.Scan: %s", err.Error())
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
            values.Enabled = sidLines[y]["enabled"]
            values.Path = data[x]["filePath"]
            values.FileName = data[x]["fileName"]
            values.Line = sidLines[y]["raw"]
            values.Sid = sidLines[y]["sid"]
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
        for w := range allSidsResult{
            CheckLinesID := LinesID{}
            var bothEnabled string
            var shouldDelete bool
            count := 0
            CheckLinesID = allSidsResult[w] 
            for r := range CheckLinesID.Values{
                if count == 0 {
                    count++
                    bothEnabled = CheckLinesID.Values[r].Enabled
                }else{
                    if ((bothEnabled != CheckLinesID.Values[r].Enabled) || ((bothEnabled == CheckLinesID.Values[r].Enabled) && (bothEnabled != "Enabled"))){
                        shouldDelete = true
                    }
                }
            }
            if shouldDelete {
                allSidsResult[w] = LinesID{}
                delete(allSidsResult, w)
                shouldDelete = false
            }
        }

        if len(allSidsResult) == 0{return nil,nil}    

        LinesOutput, err := json.Marshal(allSidsResult)
        if err != nil {logs.Error("ERROR Marshal allSidsResult --> "+err.Error()); return nil,err}
        return LinesOutput, nil
    }
}

//Add new ruleset to locale ruleset
func AddNewRuleset(data map[string]map[string]string)(duplicated []byte, err error) {
    //check for duplicated rule SIDs
    if duplicated,err = FindDuplicatedSIDs(data); duplicated != nil {
        return duplicated, nil
    }
    if err != nil {logs.Error("ruleset/AddNewRuleset -- duplicated error: %s", err.Error()); return nil,err}

    if ndb.Rdb == nil {logs.Error("ruleset/AddNewRuleset -- Can't access to database"); return nil,errors.New("ruleset/AddNewRuleset -- Can't access to database")}

    localFiles, err := utils.GetKeyValueString("ruleset", "localRulesets")
    if err != nil {logs.Error("DeleteRuleset Error getting data from main.conf for load data: "+err.Error()); return duplicated, err}

    rulesetUUID := utils.Generate()
    rulesetCreated := false
    for x := range data {        
        rulesetFolderName := strings.Replace(data[x]["rulesetName"], " ", "_", -1)
        path := localFiles + rulesetFolderName + "/" + data[x]["fileName"]

        if !rulesetCreated {
            err = insertRulesetValues(rulesetUUID, "type", "local")
            err = insertRulesetValues(rulesetUUID, "name", data[x]["rulesetName"])
            err = insertRulesetValues(rulesetUUID, "desc", data[x]["rulesetDesc"])
            if err != nil {logs.Error("ruleset/AddNewRuleset -- Insert error: %s", err.Error()); return nil,err}
            rulesetCreated = true
        }
                
        //copy source file into new folder
        if _, err := os.Stat(localFiles + rulesetFolderName); os.IsNotExist(err) {
            os.MkdirAll(localFiles + rulesetFolderName, os.ModePerm)
        }
        
        //copyfile
        copy, err := utils.GetKeyValueString("execute", "copy")
        if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return nil, err}

        cpCmd := exec.Command(copy, data[x]["filePath"], path)
        err = cpCmd.Run()
        if err != nil {logs.Error("ruleset/AddNewRuleset -- Error copying new file: %s", err.Error()); return nil,err}

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
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "linesAdded", "true")
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
        var validID = regexp.MustCompile(`sid:\s?`+sidsSplit[uuid]+`;`)
        readSidsData := make(map[string]string)
        readSidsData["sid"] = sidsSplit[uuid]
        readSidsData["uuid"] = anode["orig"]
        readSidsData["action"] = "Disable"
        
        sidLine,err := ReadSID(readSidsData)
        if err != nil {
            logs.Error("AddRulesToCustomRuleset -- Error readding SID: %s", err.Error())
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
            writeFile,err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
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

    //replace quotes
    replaceFirst := strings.Replace(content, "“", "\"", -1)
    replaceSecond := strings.Replace(replaceFirst, "”", "\"", -1)
    
    //Verify Rule content
    rulesArray := strings.Split(replaceSecond,"\n")
    rulesCount := 0 
    failsCount := 0
    rulesFailed := make(map[string]map[string][]string)
    for x := range rulesArray {
        if rulesArray[x] != "" {
            rulesCount++
            isSuccess,stats := parserule(rulesArray[x])
            if !isSuccess {
                failsCount++
                if rulesFailed[rulesArray[x]] == nil { rulesFailed[rulesArray[x]] = map[string][]string{} }
                rulesFailed[rulesArray[x]]["error"] = stats
            }
        }
    }

    logs.Notice("TOTAL RULES CHECKED --> "+strconv.Itoa(rulesCount))
    logs.Notice("TOTAL RULES FAILED --> "+strconv.Itoa(failsCount))
    for x := range rulesFailed{
        logs.Info("----------------------------")
        logs.Info("RULE: "+x)
        for w := range rulesFailed[x]["error"]{
            logs.Info(rulesFailed[x]["error"][w])
        }

    }

    path,err := ndb.GetRulesetPath(uuid)

    file, err := os.OpenFile(path, os.O_RDWR, 0644)
    if err != nil {logs.Error("SaveRulesetData failed opening file: %s", err); return err}

    defer file.Close()
    file.Truncate(0)
    file.Seek(0,0)
    _, err = file.WriteAt([]byte(replaceSecond), 0) // Write at 0 beginning
    if err != nil {logs.Error("SaveRulesetData failed writing to file: %s", err); return err}
    
    //update md5
    allFiles,err := ndb.GetAllRuleFiles()
    sourceMD5,err := utils.CalculateMD5(allFiles[uuid]["path"]); if err != nil {logs.Error("SaveRulesetData Error calculating source md5: %s", err.Error()); return err}    
    err = ndb.UpdateRuleFiles(uuid, "md5", sourceMD5); if err != nil {logs.Error("SaveRulesetData Error updating source md5: %s", err.Error()); return err}                  
    
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

    cmd, err := utils.GetKeyValueString("execute", "command")
    if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return err}
    param, err := utils.GetKeyValueString("execute", "param")
    if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return err}

    anode["line"] = strings.Replace(anode["line"], "/", "\\/", -1)
    val := "sed -i '/sid:"+anode["sid"]+"/s/.*/"+anode["line"]+"/' "+path+""
    _, err = exec.Command(cmd, param, val).Output()
    if err != nil {logs.Error("UpdateRule Error: "+err.Error()); return err}

   return nil
}

//Add new ruleset to locale ruleset
func ModifyRuleset(data map[string]map[string]string)(duplicated []byte, err error) {
    //check for duplicated rule SIDs
    if duplicated,err = FindDuplicatedSIDs(data); duplicated != nil {
        return duplicated, nil
    }
    if err != nil {logs.Error("ruleset/ModifyRuleset -- duplicated error: %s", err.Error()); return nil,err}
    if ndb.Rdb == nil {logs.Error("ruleset/ModifyRuleset -- Can't access to database"); return nil,errors.New("ruleset/ModifyRuleset -- Can't access to database")}

    //delete all rule files
    for x := range data{
        err = ndb.DeleteRuleFileParamValue("sourceUUID", data[x]["uuid"]);  if err != nil {logs.Error("ruleset/ModifyRuleset -- delete files for update error: %s", err.Error()); return nil,err}
    }

    //insert all files 
    localFiles, err := utils.GetKeyValueString("ruleset", "localRulesets")
    if err != nil {logs.Error("DeleteRuleset Error getting data from main.conf for load data: "+err.Error()); return duplicated, err}

    rulesetModified := false
    for x := range data {       
        rulesetFolderName := strings.Replace(data[x]["rulesetName"], " ", "_", -1)
        path := localFiles + rulesetFolderName + "/" + data[x]["fileName"]

        if !rulesetModified {
            //change ruleset name and desc
            err = ndb.UpdateRuleset(data[x]["uuid"], "name", data[x]["rulesetName"]); if err != nil {logs.Error("ruleset/ModifyRuleset -- modify name error: %s", err.Error()); return nil,err}
            err = ndb.UpdateRuleset(data[x]["uuid"], "desc", data[x]["rulesetDesc"]); if err != nil {logs.Error("ruleset/ModifyRuleset -- modify desc error: %s", err.Error()); return nil,err}
            rulesetModified = true
        }
                
        //check source file folder
        if _, err := os.Stat(localFiles + rulesetFolderName); os.IsNotExist(err) {
            os.MkdirAll(localFiles + rulesetFolderName, os.ModePerm)
        }
        
        //copyfile
        copy, err := utils.GetKeyValueString("execute", "copy")
        if err != nil {logs.Error("SetRulesetAction Error getting data from main.conf"); return nil, err}

        cpCmd := exec.Command(copy, data[x]["filePath"], path)
        err = cpCmd.Run()
        if err != nil {logs.Error("ruleset/ModifyRuleset -- Error copying new file: %s", err.Error()); return nil,err}

        //add md5 for every file
        md5,err := utils.CalculateMD5(path)
        if err != nil {logs.Error("ruleset/ModifyRuleset -- Error calculating md5: %s", err.Error());return nil,err}

        ruleFilesUUID := utils.Generate()
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "name", data[x]["rulesetName"])
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "path", path)
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "file", data[x]["fileName"])
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "type", "local")
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceUUID", data[x]["uuid"])
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceFileUUID", x)
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "exists", "true")
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "isUpdated", "false")
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "md5", md5)
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "sourceType", data[x]["sourceType"])
        err = ndb.InsertRulesetSourceRules(ruleFilesUUID, "linesAdded", "true")
        if err != nil {logs.Error("ruleset/ModifyRuleset -- Error Inserting Ruleset: %s", err.Error());return nil,err}
    }

    return nil,nil
}


func SyncToAll(content map[string]string)(err error) {
    full := true
    if scope,ok := content["scope"]; ok {
        if scope == "local" { full = false }
    }

    if full {
        //Download
        err = SyncToAllDownload(content)
        if err != nil {logs.Error("SyncToAll Error getting nodes by ruleset: %s", err.Error()); return err}
    }

    //sync
    err = SyncToAllNodes(content)
    if err != nil {logs.Error("SyncToAll Error getting nodes by ruleset: %s", err.Error()); return err}
        
    return nil
}

//download
func SyncToAllDownload(content map[string]string)(err error) {
    //download/overwrite rulesets source content
    data,err := ndb.GetRulesFromRuleset(content["uuid"])
    if err != nil {logs.Error("SyncToAll Error getting GetRulesFromRuleset: %s", err.Error()); return err}

    var sources []string
    for id := range data {
        exists := false
        for x := range sources {
            if sources[x] == data[id]["sourceFileUUID"]{
                exists = true
            }
        }
        if !exists{
            sources = append(sources, data[id]["sourceFileUUID"])
        }
    }
    
    DloadOwrite := make(map[string]map[string]string)
    //get sourceFileUUID
    for f := range sources {
        DloadOwriteValues := make(map[string]string)

        sourceRulesets,err := ndb.GetRuleFilesByUniqueid(sources[f])
        if err != nil {logs.Error("SyncToAll Error getting sources IDs: %s", err.Error()); return err}

        rset,err := ndb.GetAllDataRulesetDB(sourceRulesets[sources[f]]["sourceUUID"])
        if err != nil {logs.Error("SyncToAll Error getting sources URLs: %s", err.Error()); return err}
        DloadOwriteValues["uuid"] = sourceRulesets[sources[f]]["sourceUUID"]
        DloadOwriteValues["path"] = rset[sourceRulesets[sources[f]]["sourceUUID"]]["path"]
        DloadOwriteValues["name"] = rset[sourceRulesets[sources[f]]["sourceUUID"]]["name"]
        DloadOwriteValues["url"] = rset[sourceRulesets[sources[f]]["sourceUUID"]]["url"]
        DloadOwriteValues["isDownloaded"] = rset[sourceRulesets[sources[f]]["sourceUUID"]]["isDownloaded"]
        DloadOwriteValues["sourceType"] = rset[sourceRulesets[sources[f]]["sourceUUID"]]["sourceType"]
        DloadOwrite[sourceRulesets[sources[f]]["sourceUUID"]] = DloadOwriteValues
    }

    for h := range DloadOwrite {
        if DloadOwrite[h]["sourceType"] == "custom"{continue}
        if DloadOwrite[h]["isDownloaded"] == "false"{
            logs.Info("Downloading...")
            err = rulesetSource.DownloadFile(DloadOwrite[h])    
            if err != nil {logs.Error("SyncToAll Error download ruleset: %s", err.Error()); return err}
        }else{
            logs.Info("Overwriting...")
            err = rulesetSource.OverwriteDownload(DloadOwrite[h])
            if err != nil {logs.Error("SyncToAll Error Overwrite ruleset: %s", err.Error()); return err}
        }
    }

    allRules,err := ndb.GetRulesFromRuleset(content["uuid"])
    for d := range allRules{
        if content["update"] == "add-lines" && content["update"] != "" {
            logs.Info("Adding new rules...")
            err = rulesetSource.AddNewLinesToRuleset(d)
            if err != nil {logs.Error("SyncToAllDownload Error AddNewLinesToRuleset ruleset: %s", err); return err}
        }else{
            logs.Info("Overwriting rules...")
            err = rulesetSource.OverwriteRuleFile(d)
            if err != nil {logs.Error("SyncToAllDownload Error OverwriteRuleFile ruleset: %s", err); return err}
        }
    }

    return nil
}

func SyncToAllNodes(content map[string]string)(err error) {
    //synchronize to all nodes
    //get all nodes ID with this ruleset
    nodeList,err := ndb.GetNodeWithRulesetUUID(content["uuid"])
    if err != nil {logs.Error("SyncToAllSync Error getting nodes by ruleset: %s", err.Error()); return err}

    if len(nodeList) <= 0 { 
        logs.Warn("SyncToAllNodes: ruleset %s  -> No nodes asigned", content["uuid"])
        // return errors.New("SyncToAllNodes ERROR getting ruleset to synchronize, UUID "+content["uuid"]+" doesn't exist")
    }

    //get all groups
    allGroups,err := ndb.GetAllGroups()
    if err != nil {logs.Error("SyncToAllSync Error getting all groups: %s", err.Error()); return err}
    //get all rulesets into groups
    allGroupRset,err := ndb.GetAllGroupRulesets()
    if err != nil {logs.Error("SyncToAllSync Error getting all groups: %s", err.Error()); return err}
    //Get all group nodes
    allGroupNodes,err := ndb.GetAllGroupNodes()
    if err != nil {logs.Error("SyncToAllSync Error getting all groupNodes: %s", err.Error()); return err}
    rulesetName,err := ndb.GetRulesetSourceValue(content["uuid"], "name")

    for x := range allGroups {
        exists := false
        for y := range allGroupRset{
            //check if a group use this ruleset
            if allGroupRset[y]["groupid"] == x {
                if allGroupRset[y]["rulesetid"] == content["uuid"] {
                    exists = true
                }
            }
        }
    
        if exists {
            //get group nodes
            for z := range allGroupNodes{
                if allGroupNodes[z]["groupid"] == x {
                    nodeExists := false
                    for node := range nodeList {
                        if nodeList[node] == allGroupNodes[z]["nodesid"] {
                            nodeExists = true
                        }
                    }
                    if !nodeExists {
                        nodeList = append(nodeList, allGroupNodes[z]["nodesid"])
                    }
                }
            }
        }
    }

    logs.Info("Synchronizing ruleset...")
    //sync to all nodes
    for nodeID := range nodeList {
        values := make(map[string][]byte)
        //get node token
        err = ndb.GetTokenByUuid(nodeList[nodeID]); if err!=nil{logs.Error("SyncToAllSync Error loading node token: %s",err); return err}
        //get node ip and port
        ipnid,portnid,err := ndb.ObtainPortIp(nodeList[nodeID])
        if err != nil { logs.Error("SyncToAllSync ERROR Obtaining Port and Ip: "+err.Error()); return err}

        //get ruleset content
        rulesetData,err := node.CreateNewRuleFile(content["uuid"])
        if err != nil {logs.Error("SyncToAllSync error creating ruleset file: "+err.Error()); return err}
        
        values["data"] = rulesetData
        values["name"] = []byte(rulesetName)

        //send to 
        err = nodeclient.SyncGroupRulesetToNode(ipnid, portnid, values)
        if err != nil {logs.Error("SyncToAllSync error SyncGroupRulesetToNode: "+err.Error()); return err}
    }
        
    logs.Notice("Sync complete!")
    return err
}