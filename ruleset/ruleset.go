package ruleset

import(
    //"io/ioutil"
    "fmt"
    "github.com/astaxie/beego/logs"
    "bufio" //read line by line the doc
    "regexp"
    "os"
    "os/exec"
    //"strconv"
    "owlhmaster/utils"
    "owlhmaster/database"
    "errors"
    "database/sql"
    "strings"
    "time"
    "strconv"
)

func ReadSID(sid map[string]string)( sidLine map[string]string ,err error){

    sidMap := sid["sid"]
    uuidMap := sid["uuid"]

    logs.Info("SID: "+sidMap+" UID: "+uuidMap)

    path, err := GetRulesetPath(uuidMap)

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

func Read(path string)(rules map[string]map[string]string, err error) {

    data, err := os.Open(path)
    
    if err != nil {
        fmt.Println("File reading error", err)
        return 
    }

    var validID = regexp.MustCompile(`sid:(\d+);`)
    var ipfield = regexp.MustCompile(`^([^\(]+)\(`)
    var msgfield = regexp.MustCompile(`msg:([^;]+);`)
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

func AddRuleset(n map[string]string) (err error) {
    logs.Info("ADD RULESET -> IN")
    //crear UUID
    rulesetID := utils.Generate()
    logs.Info(n)

    //Verificar que nos llegan los params
    if _, ok := n["name"]; !ok {
        return errors.New("Name is empty")
    }
    if _, ok := n["path"]; !ok {
        return errors.New("Path is empty")
    }
    //Verificar que no existe el ruleset
    if err := rulesetExists(rulesetID); err != nil {
        return err
    }
    //Meter en DB
    for key, value := range n {
        err = rulesetInsert(rulesetID, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

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

func rulesetInsert(nkey string, key string, value string) (err error) {
    if ndb.Rdb == nil {
        logs.Error("rulesetInsert -- Can't access to database")
        return errors.New("rulesetInsert -- Can't access to database")
    }
    logs.Info("nkey: %s, key: %s, value: %s", nkey, key, value)
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

func GetAllRulesets() (rulesets *map[string]map[string]string, err error) {
    var allrulesets = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    if ndb.Rdb == nil {
        logs.Error("ruleset/GetAllRulesets -- Can't access to database")
        return nil, errors.New("ruleset/GetAllRulesets -- Can't access to database")
    }
    sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset;"
    rows, err := ndb.Rdb.Query(sql)
    if err != nil {
        logs.Error("ruleset/GetAllRulesets -- Query error: %s", err.Error())
        return nil, err
    }
    for rows.Next() {
        if err = rows.Scan(&uniqid, &param, &value); err != nil {
            logs.Error("ruleset/GetAllRulesets -- Query return error: %s", err.Error())
            return nil, err
        }
        //logs.Info("uniqid: %s, param: %s, value: %s", uniqid,param,value)
        if allrulesets[uniqid] == nil { allrulesets[uniqid] = map[string]string{}}
        allrulesets[uniqid][param]=value
    } 
    return &allrulesets, nil
}

func GetRulesetPath(uuid string) (n string, err error) {
    logs.Info("DB RULESET -> Get path"+uuid)
    var path string
    if ndb.Rdb != nil {
        row := ndb.Rdb.QueryRow("SELECT ruleset_value FROM ruleset WHERE ruleset_uniqueid=$1 and ruleset_param=\"path\";",uuid)
        err = row.Scan(&path)
        if err == sql.ErrNoRows {
            logs.Warn("DB RULESET -> There is no ruleset with id %s",uuid)
            return "", errors.New("DB RULESET -> There is no ruleset with id "+uuid)
        }
        if err != nil {
            logs.Warn("DB RULESET -> rows.Scan Error -> %s", err.Error())
            return "", errors.New("DB RULESET -> -> rows.Scan Error -> " + err.Error())
        }
        return path, nil
    } else {
        logs.Info("DB RULESET -> No access to database")
        return "", errors.New("DB RULESET -> No access to database")
    }
}

func GetRulesetRules(uuid string)(r map[string]map[string]string, err error){
    logs.Info("DB RULESET -> GetRulesetRules"+uuid)
    rules := make(map[string]map[string]string)
    path,err := GetRulesetPath(uuid) 
    rules,err = Read(path) 
    for rule, _ := range rules{
        retrieveNote := make(map[string]string)
        retrieveNote["uuid"] = uuid
        retrieveNote["sid"] = rule
        rules[rule]["note"], _ = GetRuleNote(retrieveNote)
    }
    return rules, err
}

func SetRuleSelected(n map[string]string) (err error) {
    node_uniqueid_ruleset := n["nid"]
    ruleset_uniqueid := n["rule_uid"]

    if ndb.Rdb == nil {
        logs.Error("SetRuleSelected -- Can't access to database")
        return errors.New("SetRuleSelected -- Can't access to database")
    }
    //rows, err := ndb.Rdb.Query("SELECT ruleset_uniqueid, node_uniqueid FROM ruleset_node WHERE ruleset_uniqueid=$1 and node_uniqueid=$2;",ruleset_uniqueid,node_uniqueid_ruleset)
    sqlQuery := "SELECT * FROM ruleset_node WHERE node_uniqueid = \""+node_uniqueid_ruleset+"\";"
    logs.Info("SQL Query : "+sqlQuery)
    rows, err := ndb.Rdb.Query(sqlQuery)
    if err != nil {
        logs.Info("Put SetRuleSelecteda query error %s",err.Error())
        return err
    }

    defer rows.Close()
    if rows.Next() {
        rows.Close()
        logs.Info("ruleset/SetRuleSelected UPDATE")
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
        logs.Info("RETURN NILL UPDATE")
        return nil

    } else {
        logs.Info("ruleset/SetRuleSelected INSERT")
        insertRulesetNode, err := ndb.Rdb.Prepare("insert into ruleset_node (ruleset_uniqueid, node_uniqueid) values (?,?);")
        _, err = insertRulesetNode.Exec(&ruleset_uniqueid, &node_uniqueid_ruleset)  
        defer insertRulesetNode.Close()

        if (err != nil){
            logs.Info("error insertRulesetNode en ruleset/rulesets--> "+err.Error())
            return err
        }

        return nil 
    }
    return err
}


func GetRuleSelected(nid string)(ruleset string, err error){
    var ruleSelected string
    if ndb.Rdb != nil {
        row := ndb.Rdb.QueryRow("SELECT ruleset_uniqueid FROM ruleset_node WHERE node_uniqueid = \""+nid+"\";")
        err = row.Scan(&ruleSelected)
        if err == sql.ErrNoRows{
            logs.Warn("GetRuleName -> There is no ruleset with thie UUID %s", nid)
            return "", err
        }
        if err != nil {
            logs.Warn("GetRuleName -> row.Scan error %s", err.Error())
            return "", err
        }
        return ruleSelected, nil
    }else {
        logs.Info("GetRuleSelected -> No access to database")
        return "", err
    }
}

func GetRuleName(nid string)(ruleset string, err error){
    logs.Debug("PREUBAS -> GetRuleName: "+nid)
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
        logs.Info("GetRuleName -> no access to database")
        return "", err
    }
}

func SetClonedRuleset(ruleCloned map[string]string)(err error){
    if ndb.Rdb == nil {
        logs.Error("rulesetExists -- Can't access to database")
        return errors.New("rulesetExists -- Can't access to database")
    }
    clonedRuleset := ruleCloned["cloned"]
    newRuleset := ruleCloned["new"]

    newRuleset = strings.Replace(newRuleset, " ", "_", -1)

    clonedPath := ruleCloned["path"]
    // Path must be setup in main.conf file. 
    path := "/etc/owlh/ruleset/"
    pathNewRule := path+newRuleset+".rules"
    //pathOldRule := path+clonedRuleset
    newUUID := utils.Generate()

    logs.Info("SetClonedRuleset: PATH: "+clonedPath+" Clone: "+clonedRuleset+" new "+newRuleset+" New UUID --> "+newUUID)

    rows, err := ndb.Rdb.Query("SELECT * FROM ruleset WHERE ruleset_uniqueid = \""+newUUID+"\";")
    defer rows.Close()
    if !rows.Next(){

        cpCmd := exec.Command("cp", clonedPath, pathNewRule)
        err = cpCmd.Run()
        if err != nil{
            logs.Info ("SetClonedRuleset --> Error exec cmd command")
            return err
        }
    
        insertCloneName, err := ndb.Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values (?,?,?);")
        _, err = insertCloneName.Exec(&newUUID, "name", &newRuleset)  
        defer insertCloneName.Close()
        if (err != nil){
            logs.Info("error insertCloneName name ruleset/rulesets--> "+err.Error())
            return err
        }
        logs.Info("ruleset/SetClonedRuleset INSERT name done")
        insertCloneValue, err := ndb.Rdb.Prepare("insert into ruleset (ruleset_uniqueid, ruleset_param, ruleset_value) values (?,?,?);")
        _, err = insertCloneValue.Exec(&newUUID, "path", &pathNewRule)
        defer insertCloneValue.Close()
        if (err != nil){
            logs.Info("error insertCloneValue name  ruleset/rulesets--> "+err.Error())
            return err
        }
        logs.Info("ruleset/SetClonedRuleset INSERT path done")

        return nil
    }
    if err != nil {
        logs.Warn("SetClonedRuleset -> rows.Scan %s", err.Error())
        return err
    }

    return nil   
}

func SetRulesetAction(ruleAction map[string]string)(err error){
    sid := ruleAction["sid"]
    uuid := ruleAction["uuid"]
    action := ruleAction["action"]

    path, err := GetRulesetPath(uuid)
    logs.Info("Path del uuid para el Action: "+path)

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
        logs.Info("Put SetRuleNote query error %s",err.Error())
        return err
    }
    defer rows.Close()
    if rows.Next() {
        rows.Close()
        logs.Info("ruleset/SetRuleNote UPDATE - "+sid+" "+uuid+" "+noteTime+" "+note)
        updateNote, err := ndb.Rdb.Prepare("update rule_note set ruleNote = ?, note_date = ? where ruleset_uniqueid = ? and rule_sid = ? ;")
        _, err = updateNote.Exec(&note, &noteTime, &uuid, &sid)               
        defer updateNote.Close()

        if (err != nil){
            logs.Info("SetRuleNote UPDATE Error -- "+err.Error())
            return err
        }
        logs.Info("updateNote RETURN NILL UPDATE")
        return nil

    } else {
        logs.Info("ruleset/SetRuleNote INSERT")
        insertNote, err := ndb.Rdb.Prepare("insert into rule_note (ruleset_uniqueid, rule_sid, note_date, ruleNote) values (?,?,?,?);")
        _, err = insertNote.Exec(&uuid, &sid, &noteTime, &note)  
        defer insertNote.Close()

        if (err != nil){
            logs.Info("error insertRulesetNode ruleset/rulesets--> "+err.Error())
            return err
        }
        logs.Info("insertNote RETURN NILL UPDATE")
        return nil 
    }
    return err
}

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
        //logs.Warn("DB GetNote -> no hemos leido bien los campos de scan")
        return "", errors.New("DB GetNote -> row.Scan error")
    }
    return noteText, nil
}