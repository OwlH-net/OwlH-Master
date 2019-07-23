package rulesetSource

import (
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
	"errors"
	"owlhmaster/utils"
	"strings"
	"os"
	"io/ioutil"
	"regexp"
	"path/filepath"
	"bufio"
)

func GetFileUUIDfromRulesetUUID(value string)(uuid string, err error){
	var uniqueid string
	uuidRules, err := ndb.Rdb.Query("select rule_uniqueid from rule_files where rule_value='"+value+"'")
	if err != nil {
        logs.Error("ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
        return "",err
	}
	defer uuidRules.Close()
	for uuidRules.Next() {
		if err = uuidRules.Scan(&uniqueid); err != nil {
			logs.Error("GetAllRulesetSource rows.Scan: %s", err.Error())
			return "",err
		}
	}
	return uniqueid,nil
}


func CreateRulesetSource(n map[string]string) (err error) {
	rulesetSourceKey := utils.Generate()
	if n["name"] == "" {
        return errors.New("Name is empty")
	}
	if n["desc"] == "" {
        return errors.New("Description is empty")
    }
	if n["url"] == "" {
        return errors.New("URL is empty")
    }
    if err := rulesetSourceExists(rulesetSourceKey); err != nil {
		logs.Error("rulesetSource exist: "+err.Error())
        return errors.New("rulesetSource exist")
	}
	
	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	path := sourceDownload["ruleset"]["sourceDownload"]


	if _, err := os.Stat(path+n["name"]); !os.IsNotExist(err) {
		return errors.New("The folder "+n["name"]+" already exists. Use other name for the new ruleset source.")
	}
	
	nameWithoutSpaces := strings.Replace(n["name"], " ", "_", -1)
	n["path"] = path + nameWithoutSpaces +"/"+ n["fileName"]
	    
    for key, value := range n {
        err = ndb.RulesetSourceKeyInsert(rulesetSourceKey, key, value)
    }
    if err != nil {
        return err
    }
    return nil
}

func CreateCustomRulesetSource(n map[string]string)(err error){
	if n["name"] == "" {
        return errors.New("Name is empty")
	}
	if n["desc"] == "" {
        return errors.New("Description is empty")
	}
	customRulesets := map[string]map[string]string{}
	customRulesets["ruleset"] = map[string]string{}
	customRulesets["ruleset"]["customRulesets"] = ""
	customRulesets,err = utils.GetConf(customRulesets)
	path := customRulesets["ruleset"]["customRulesets"]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	if err != nil {
		logs.Error("Error checking the path: "+err.Error())
        return err
	}
	
	newUUID := utils.Generate()

	nameWithoutSpaces := strings.Replace(n["name"], " ", "_", -1)
	nameFile := nameWithoutSpaces+".rules"
	n["path"] = path + nameFile

	if _, err := os.Stat(n["path"]); !os.IsNotExist(err) {
		return errors.New("The custom file "+n["name"]+" already exists. Use other name for the new custom ruleset source.")
	}

	err = ioutil.WriteFile(n["path"], []byte(""), 0755)
    if err != nil {
        return errors.New("Can't create the custom rule file.")
	}

	for key, value := range n {
        err = ndb.RulesetSourceKeyInsert(newUUID, key, value)
	}
	if err != nil {
        return errors.New("Error adding custom rule file data to database.")
	}
	
	md5,err := utils.CalculateMD5(n["path"])
	if err != nil {return errors.New("Error Checking MD5 for CreateCustomRulesetSource.")}

	//insert file into rule_files
	uuid := utils.Generate()
	err = ndb.InsertRulesetSourceRules(uuid, "name", n["name"])
	err = ndb.InsertRulesetSourceRules(uuid, "path", n["path"])
	err = ndb.InsertRulesetSourceRules(uuid, "sourceType", n["sourceType"])
	err = ndb.InsertRulesetSourceRules(uuid, "fileName", nameFile)
	err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
	err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", newUUID)
	err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
	err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
	err = ndb.InsertRulesetSourceRules(uuid, "md5", md5)
	if err != nil {
        return errors.New("Error adding custom rule file data to database.")
	}
	
    return nil
}

func rulesetSourceExists(sourceID string) (err error) {
    if ndb.Rdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    sql := "SELECT * FROM ruleset where ruleset_uniqueid = '"+sourceID+"';"
    rows, err := ndb.Rdb.Query(sql)
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

func GetAllRulesetSource()(sources map[string]map[string]string, err error){
	var allsources = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
    var uuidSource string
    if ndb.Rdb == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
    }
	sqlUUID := "select ruleset_uniqueid from ruleset where ruleset_param='type' and ruleset_value = 'source';"
	uuidRows, err := ndb.Rdb.Query(sqlUUID)
	defer uuidRows.Close()
	if err != nil {
        logs.Error("ndb.Rdb.Query Error checking uuid for take the uuid list for ruleset_source: %s", err.Error())
        return nil, err
    }
	for uuidRows.Next() {
		if err = uuidRows.Scan(&uuidSource); err != nil {
            logs.Error("GetAllRulesetSource UUIDSource uuidRows.Scan: %s", err.Error())
            return nil, err
		}
		sql := "select ruleset_uniqueid, ruleset_param, ruleset_value from ruleset where ruleset_uniqueid='"+uuidSource+"';"
		rows, err := ndb.Rdb.Query(sql)
		if err != nil {
			logs.Error("ndb.Rdb.Query Error : %s", err.Error())
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
	}

    return allsources, nil
}

//Delete specific file for ruleset
func DeleteRulesetFile(uuid string) (err error) {
	var pathToDelete string
	if ndb.Rdb == nil {
        logs.Error("DeleteRulesetSource -- Can't acces to database")
        return errors.New("DeleteRulesetSource -- Can't acces to database")
	}

	//delete file selected
	uuidPath, err := ndb.Rdb.Query("select rule_value from rule_files where rule_uniqueid = '"+uuid+"' and rule_param='path'")
	if err != nil {
		logs.Error("ndb.Rdb.Query Error checking path for delete a file on rule_files: %s", err.Error())
		return err
	}
	defer uuidPath.Close()
	for uuidPath.Next() {
		if err = uuidPath.Scan(&pathToDelete); err != nil {
			logs.Error("DeleteRulesetSource for delete path rows.Scan: %s", err.Error())
			return err
		}
	
		err = os.Remove(pathToDelete)
	}
	
	//delete a ruleset source in ruleset table
	sourceSQL, err := ndb.Rdb.Prepare("delete from rule_files where rule_uniqueid = ?")
    if err != nil {
        logs.Error("Error DeleteRulesetSource Prepare delete a file -> %s", err.Error())
        return err
	}
    _, err = sourceSQL.Exec(&uuid)
    if err != nil {
        logs.Error("Error DeleteRulesetSource deleting a file -> %s", err.Error())
        return err
	}
	return nil
}

func DeleteRulesetSource(anode map[string]string) (err error) {
	rulesetSourceUUID := anode["uuid"]
	sourceType := anode["sourceType"]
	// var pathToDelete string
	var uniqueid string
	var uuidArray []string
	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]

	if ndb.Rdb == nil {
        logs.Error("DeleteRulesetSource -- Can't acces to database")
        return errors.New("DeleteRulesetSource -- Can't acces to database")
	}

	if sourceType == "url" {
		name,err := ndb.GetRulesetSourceValue(rulesetSourceUUID, "name")
		nameWithoutSpaces := strings.Replace(name, " ", "_", -1)
		err = os.RemoveAll(pathDownloaded+nameWithoutSpaces)
		if err != nil {
			logs.Error("DeleteRulesetSource Error deleting path for URL source type: %s", err.Error())
			return err
		}
	}else{
		path,err := ndb.GetRulesetSourceValue(rulesetSourceUUID, "path")
		err = os.RemoveAll(path)
		if err != nil {
			logs.Error("DeleteRulesetSource Error deleting path for CUSTOM source type: %s", err.Error())
		}
	}
	
	//delete a ruleset source in ruleset table
	sourceSQL, err := ndb.Rdb.Prepare("delete from ruleset where ruleset_uniqueid = ?")
    if err != nil {
        logs.Error("DeleteRulesetSource Prepare delete a ruleset source -> %s", err.Error())
        return err
	}
    _, err = sourceSQL.Exec(&rulesetSourceUUID)
    if err != nil {
        logs.Error("DeleteRulesetSource deleting a ruleset source -> %s", err.Error())
        return err
	}
	
	//delete all ruleset source rules
	uuidRules, err := ndb.Rdb.Query("select rule_uniqueid from rule_files where rule_value='"+rulesetSourceUUID+"'")
	if err != nil {
        logs.Error("ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
        return err
	}
	defer uuidRules.Close()
	for uuidRules.Next() {
		if err = uuidRules.Scan(&uniqueid); err != nil {
			logs.Error("GetAllRulesetSource rows.Scan: %s", err.Error())
			return err
		}
		uuidArray = append(uuidArray, uniqueid)
	}
	for x := range uuidArray{
		DeleteUUIDPrepare, err := ndb.Rdb.Prepare("delete from rule_files where rule_uniqueid = ?")
		if err != nil {
			logs.Error("ndb.Rdb.Query Error deleting by rule_uniqueid for rule_files: %s", err.Error())
			return err
		}
		_, err = DeleteUUIDPrepare.Exec(&uuidArray[x])
		if err != nil {
			logs.Error("DeleteRulesetSource deleting a ruleset source -> %s", err.Error())
			return err
		}
	}

	return nil
}

func EditRulesetSource(data map[string]string) (err error) { 
	var sourceuuid = data["uuid"]
    if ndb.Rdb == nil {
		logs.Error("no access to database")
        return errors.New("no access to database")
	}
	for k,v := range data { 
		if k == "sourceuuid"{
			continue
		}else{
			//update ruleset source
			err = ndb.UpdateRuleset(sourceuuid, k, v)
			if err != nil {
				logs.Error("Error updating ruleset source data")
				return err
			}

		}
	}
		
	return nil
}

func OverwriteDownload(data map[string]string) (err error) {
	var fileExtension = regexp.MustCompile(`(\w+).rules$`)
	var newFilesDownloaded = make(map[string]string)
	var newFilesDB = make(map[string]map[string]string)

	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]

	splitPath := strings.Split(data["url"], "/")
	fileDownloaded := splitPath[len(splitPath)-1]

	_ = os.RemoveAll(pathDownloaded+data["name"])

	//download file
	if _, err := os.Stat(pathDownloaded+data["name"]); os.IsNotExist(err) {
		os.MkdirAll(pathDownloaded+data["name"], os.ModePerm)

		err = utils.DownloadFile(pathDownloaded + data["name"] + "/" + fileDownloaded, data["url"])
		if err != nil {
			logs.Error("Error downloading file from RulesetSource-> %s", err.Error())
			_ = os.RemoveAll(pathDownloaded+data["name"])

			// update ruleset "exists" field
			_ = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "false")
			return err
		}
	
		// err = utils.ExtractTarGz(data["path"], pathDownloaded, data["name"])
		err = utils.ExtractTarGz(pathDownloaded + data["name"] + "/" + fileDownloaded, pathDownloaded+data["name"])
		if err != nil {
			logs.Error("Error unzipping file downloaded: "+err.Error())
			err = os.RemoveAll(pathDownloaded+data["name"])
			
			// update ruleset "exists" field
			err = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "false")
			if err != nil {
				logs.Error("UpdateRuleset Error from RulesetSource-> %s", err.Error())
				return err
			}

			return err
		}
	}

	_ = ndb.UpdateRuleset(data["uuid"], "path", pathDownloaded + data["name"] + "/" + fileDownloaded)
	_ = ndb.UpdateRuleset(data["uuid"], "url", data["url"])

	//get map with new files downloaded
	err = filepath.Walk(pathDownloaded + data["name"],
		func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileExtension.MatchString(info.Name()){
			newFilesDownloaded[info.Name()] = file
		}
		return nil
	})

	//Create map
	dbFiles,err := ndb.GetRulesFromRuleset(data["uuid"])
	for y := range dbFiles { 
		var values = make(map[string]string)
		for z := range dbFiles[y] {
			values[z] = dbFiles[y][z]
		}
		values["count"] = "0" //DB contains a file, but the new download don't contain a current file 
		values["uuid"] = y
		newFilesDB[dbFiles[y]["file"]] = values
	}

	//assign stauts
	for r := range newFilesDownloaded {
		var values = make(map[string]string)
		values["filePath"] = newFilesDownloaded[r]
		if _, found := newFilesDB[r]; !found {
			values["count"] = "1" //new file downloaded -- !DB not contain the new file
			newFilesDB[r] = values
		}else{// _, found := newFilesDB[w]; found{
			values["count"] = "2"//File and DB OK
			newFilesDB[r] = values
		}
	}

	// check status
	for w := range newFilesDB {
		

		if newFilesDB[w]["count"] == "0" {
			err = ndb.UpdateRuleFiles(newFilesDB[w]["uuid"], "exists", "false")
			if err != nil{
				logs.Error("OverwriteDownload UPDATE error for update isDownloaded -- "+err.Error())
				return err
			}
		}
		
		if newFilesDB[w]["count"] == "1" {
			//add md5 for every file
			md5,err := utils.CalculateMD5(newFilesDB[w]["filePath"])
			if err != nil{
				logs.Error("OverwriteDownload Error calculating md5: "+err.Error())
				return err
			}

			uuid := utils.Generate()
			err = ndb.InsertRulesetSourceRules(uuid, "name", data["name"])
			err = ndb.InsertRulesetSourceRules(uuid, "path", newFilesDB[w]["filePath"])
			err = ndb.InsertRulesetSourceRules(uuid, "file", w)
			err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
			err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", data["uuid"])
			err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
			err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
			err = ndb.InsertRulesetSourceRules(uuid, "md5", md5)
		}
	}

	logs.Info("Overwrite complete")

	return nil
}

func DownloadFile(data map[string]string) (err error) {

	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]

	value,err := ndb.GetRulesetSourceValue(data["uuid"], "path")
	if err != nil {
		logs.Error("Error Getting path for download file from RulesetSource-> %s", err.Error())
		return err
	}

	splitPath := strings.Split(value, "/")
	pathSelected := splitPath[len(splitPath)-2]

	if _, err := os.Stat(pathDownloaded+pathSelected); os.IsNotExist(err) {
		os.MkdirAll(pathDownloaded+pathSelected, os.ModePerm)

		err = utils.DownloadFile(data["path"], data["url"])
		if err != nil {
			logs.Error("Error downloading file from RulesetSource-> %s", err.Error())
			_ = os.RemoveAll(pathDownloaded+pathSelected)
			_ = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "false")
			return err
		}
	
		err = utils.ExtractTarGz(data["path"], pathDownloaded+pathSelected)
		if err != nil {
			logs.Error("Error unzipping file downloaded: "+err.Error())
			err = os.RemoveAll(pathDownloaded+pathSelected)
			return err
		}

		//insert into DB
		ruleFiles, err := Details(data)
		if err != nil{
			logs.Error("DownloadFile Error getting rule files: "+err.Error())
			return err
		}
		for k,_ := range ruleFiles["files"] {
			//add md5 for every file
			md5,err := utils.CalculateMD5(ruleFiles["files"][k])
			if err != nil{
				logs.Error("DownloadFile Error calculating md5: "+err.Error())
				return err
			}

			uuid := utils.Generate()
			err = ndb.InsertRulesetSourceRules(uuid, "name", data["name"])
			err = ndb.InsertRulesetSourceRules(uuid, "path", ruleFiles["files"][k])
			err = ndb.InsertRulesetSourceRules(uuid, "file", k)
			err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
			err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", data["uuid"])
			err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
			err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
			err = ndb.InsertRulesetSourceRules(uuid, "md5", md5)
		}
		if err != nil {
			logs.Error("DownloadFile Error from RulesetSource-> %s", err.Error())
			return err
		}
		
		//update isDownlaoded at ruleset
		err = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "true")
		if err != nil {
			logs.Error("UpdateRuleset Error from RulesetSource-> %s", err.Error())
			return err
		}

	}else{
		return errors.New("The folder "+data["name"]+" already exists. Use other name for the new ruleset source.")
	}
	logs.Info("Download and extract complete")

	return nil
}

func CompareFiles(uuid string) (mapData map[string]map[string]string, err error) {	
	rulesetFile,err := ndb.GetRuleFilesByUniqueid(uuid)
	var sourcePath string
	var rulesetPath string
	for r := range rulesetFile {
		rulesetPath = rulesetFile[r]["path"]
		sourceFile,err := ndb.GetRuleFilesByUniqueid(rulesetFile[r]["sourceFileUUID"])
		if err != nil {
			logs.Error("CompareFiles failed getting sourceFileUUID data: %s", err)
			return nil,err
		}
		for t := range sourceFile {
			sourcePath = sourceFile[t]["path"]
		}
	}

	fileOrig,err := utils.MapFromFile(sourcePath)
	fileDst,err := utils.MapFromFile(rulesetPath)

	var returnMap = make(map[string]map[string]string)
	lineExist := false

	//check if all the new lines are in old file
	for x,_ := range fileDst { 
		returnLines := make(map[string]string)
		for y,_ := range fileOrig { 			
			if x == y {
				lineExist = true
				if (fileDst[x]["Line"] != fileOrig[y]["Line"] || fileDst[x]["Enabled"] != fileOrig[y]["Enabled"]) {
					returnLines["new"] = fileDst[x]["Line"]
					returnLines["old"] = fileOrig[y]["Line"]
					returnLines["enabled-new"] = fileDst[x]["Enabled"]
					returnLines["enabled-old"] = fileOrig[y]["Enabled"]
					returnLines["sid"] = x
					returnMap[x] = returnLines
				}
				continue
			}
		}
		if !lineExist {
			returnLines["new"] = fileDst[x]["Line"]
			returnLines["old"] = "N/A"
			returnLines["enabled-new"] = fileDst[x]["Enabled"]
			returnLines["enabled-old"] = "N/A"
			returnLines["sid"] = x
			returnMap[x] = returnLines
		}

		lineExist = false
	}

	//check if all the old lines are in new file
	for y,_ := range fileOrig { 
		returnLines := make(map[string]string)
		for x,_ := range fileDst {
			if y == x {
				lineExist = true
				continue
			}
		}
		if !lineExist {
			returnLines["new"] = "N/A"
			returnLines["old"] = fileOrig[y]["Line"]
			returnLines["enabled-new"] = "N/A"
			returnLines["enabled-old"] = fileOrig[y]["Enabled"]
			returnLines["sid"] = y
			returnMap[y] = returnLines
		}

		lineExist = false
	}

	
	return returnMap, nil
}


func CreateNewFile(data map[string]string) (err error) {
	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["backupPath"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	backupPath := sourceDownload["ruleset"]["backupPath"]

	splitPath := strings.Split(data["path"], "/")
	pathSelected := splitPath[len(splitPath)-2]


	err = utils.BackupFile(backupPath + pathSelected, "drop.rules")
	if err != nil {
		logs.Error("CreateNewFile: Error BackupFile from map --> "+err.Error())
        return err
	}

	err = utils.ReplaceLines(data)
	if err != nil {
		logs.Error("CreateNewFile: Error replacing lines from map --> "+err.Error())
        return err
	}
    
    return nil
}

//get data from local files for insert into DB
func Details(data map[string]string) (files map[string]map[string]string, err error) {
	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]

	splitPath := strings.Split(data["path"], "/")
	pathSelected := splitPath[:len(splitPath)-1]
	folder := pathSelected[len(pathSelected)-1]

	path := pathDownloaded+folder
	var fileExtension = regexp.MustCompile(`(\w+).rules$`)
	dataFiles := map[string]map[string]string{}
	dataFiles["files"] = map[string]string{}

	err = filepath.Walk(path,
		func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileExtension.MatchString(info.Name()){
			dataFiles["files"][info.Name()] = file
		}
		return nil
	})
	
	if err != nil {
		logs.Error("Error recursively file")
		return nil ,err
	}
	
	return dataFiles ,nil
}

//
func GetDetails(uuid string) (data map[string]map[string]string, err error){
	var checked string
	
	dbFiles,err := ndb.GetRulesFromRuleset(uuid)
	
	for x := range dbFiles{
		checked = utils.VerifyPathExists(dbFiles[x]["path"])
		err = ndb.UpdateRuleFiles(x, "exists", checked)
		if err != nil {
			logs.Error("ndb.UpdateRuleFiles Error : %s", err.Error())
			return nil, err
		}
		dbFiles[x]["exists"]=checked

		//get sourceUUID from rule_files files by uuid
		sourceFile,err := ndb.GetRuleFilesByUniqueid(dbFiles[x]["sourceFileUUID"])
		if err != nil {
			logs.Error("ndb.GetRuleFilesByUniqueid Error : %s", err.Error())
			return nil, err
		}
		
		//check differences for md5
		for n := range sourceFile{
			dbFiles[x]["existsSourceFile"]=utils.VerifyPathExists(sourceFile[n]["path"])
			md5S,err := utils.CalculateMD5(sourceFile[n]["path"])
			if (sourceFile[n]["md5"] != md5S){
				dbFiles[x]["isUpdated"]="true"
				err = ndb.UpdateRuleFiles(n, "isUpdated", "true")
				err = ndb.UpdateRuleFiles(x, "isUpdated", "true")
				if err != nil {
					logs.Error("Error updating after compare md5 GetDetails: %s", err.Error())
					return nil, err
				}
			}else{
				dbFiles[x]["isUpdated"]="false"
				err = ndb.UpdateRuleFiles(n, "isUpdated", "false")
				err = ndb.UpdateRuleFiles(x, "isUpdated", "false")
				if err != nil {
					logs.Error("Error updating after compare md5 GetDetails: %s", err.Error())
					return nil, err
				}
			}
		}
	}
	
	return dbFiles, nil
}

func OverwriteRuleFile(uuid string)(err error){
	var sourceFile = map[string]map[string]string{}
	dbFiles,err := ndb.GetRuleFilesByUniqueid(uuid)
	for r := range dbFiles {
		sourceFile,err = ndb.GetRuleFilesByUniqueid(dbFiles[r]["sourceFileUUID"])
		if err != nil {
			logs.Error("SaveRulesetData failed writing to file: %s", err)
			return err
		}

		for t := range sourceFile {
			fileOrig, err := os.Open(sourceFile[t]["path"])
			fileDst,err := os.OpenFile(dbFiles[r]["path"], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				logs.Error("OverwriteRuleFile failed opening file: %s", err)
				return err
			}
			defer fileOrig.Close()
			defer fileDst.Close()

			reader := bufio.NewReader(fileOrig)
			content, _ := ioutil.ReadAll(reader)
	
			fileDst.Truncate(0)
			fileDst.Seek(0,0)
			_, err = fileDst.WriteAt([]byte(content), 0) // Write at 0 beginning
			if err != nil {
				logs.Error("OverwriteRuleFile failed writing to file: %s", err)
				return err
			}

			//change md5 and isUpdated status
			md5S,err := utils.CalculateMD5(sourceFile[t]["path"])
			md5L,err := utils.CalculateMD5(dbFiles[r]["path"])
			err = ndb.UpdateRuleFiles(t, "isUpdated", "false")	
			err = ndb.UpdateRuleFiles(t, "md5", md5S)					
			err = ndb.UpdateRuleFiles(r, "isUpdated", "false")	
			err = ndb.UpdateRuleFiles(r, "md5", md5L)					
			if err != nil {
				logs.Error("Error updating sourceFile after compare md5 OverwriteRuleFile: %s", err.Error())
				return err
			}
		}
	}
	return nil
}

func AddNewLinesToRuleset(uuid string)(err error){
	rulesetFile,err := ndb.GetRuleFilesByUniqueid(uuid)
	var sourcePath string
	var rulesetPath string
	for r := range rulesetFile {
		rulesetPath = rulesetFile[r]["path"]
		sourceFile,err := ndb.GetRuleFilesByUniqueid(rulesetFile[r]["sourceFileUUID"])
		if err != nil {
			logs.Error("AddNewLinesToRuleset failed getting sourceFileUUID data: %s", err)
			return err
		}
		for t := range sourceFile {
			sourcePath = sourceFile[t]["path"]
		}
	}

	fileOrig,err := utils.MapFromFile(sourcePath)
	fileDst,err := utils.MapFromFile(rulesetPath)
	if err != nil {
		logs.Error("AddNewLinesToRuleset failed getting map from file: %s", err)
		return err
	}

	for x := range fileOrig {
		isEquals := false
		for y := range fileDst {
			if x == y{
				isEquals = true
				continue
			}

		}
		if !isEquals{
			writeFile,err := os.OpenFile(rulesetPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			defer writeFile.Close()
		
			_, err = writeFile.WriteString(fileOrig[x]["Line"])
			_, err = writeFile.WriteString("\n")
		
			if err != nil {
				logs.Error("AddNewLinesToRuleset -- Error adding new lines: %s", err.Error())
				return err
			}
		}
	}

	//change md5 and isUpdated status
	for w := range rulesetFile{
		sourceFile,err := ndb.GetRuleFilesByUniqueid(rulesetFile[w]["sourceFileUUID"])
		if err != nil {
			logs.Error("Error getting sourceFile for update md5 OverwriteRuleFile: %s", err.Error())
			return err
		}
		for s := range sourceFile{
			//change md5 and isUpdated status
			md5S,err := utils.CalculateMD5(sourceFile[s]["path"])
			md5L,err := utils.CalculateMD5(rulesetFile[w]["path"])
			err = ndb.UpdateRuleFiles(w, "isUpdated", "false")	
			err = ndb.UpdateRuleFiles(w, "md5", md5L)					
			err = ndb.UpdateRuleFiles(s, "isUpdated", "false")	
			err = ndb.UpdateRuleFiles(s, "md5", md5S)					
			if err != nil {
				logs.Error("Error updating sourceFile after compare md5 OverwriteRuleFile: %s", err.Error())
				return err
			}
		}
	}

	return nil
}