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
	n["path"] = path + nameWithoutSpaces +".rules"

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
	
	//insert file into rule_files
	uuid := utils.Generate()
	err = ndb.InsertRulesetSourceRules(uuid, "name", n["name"])
	err = ndb.InsertRulesetSourceRules(uuid, "path", n["path"])
	err = ndb.InsertRulesetSourceRules(uuid, "sourceType", n["sourceType"])
	err = ndb.InsertRulesetSourceRules(uuid, "file", nameWithoutSpaces +".rules")
	err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
	err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", newUUID)
	err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
	err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
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

	// //delete path 
	// uuidPath, err := ndb.Rdb.Query("select ruleset_value from ruleset where ruleset_uniqueid = '"+rulesetSourceUUID+"' and ruleset_param='path'")
	// if err != nil {
	// 	logs.Error("ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
	// 	return err
	// }
	// defer uuidPath.Close()
	// for uuidPath.Next() {
	// 	if err = uuidPath.Scan(&pathToDelete); err != nil {
	// 		logs.Error("DeleteRulesetSource for delete path rows.Scan: %s", err.Error())
	// 		return err
	// 	}
	// 	logs.Info(pathToDelete)
		
	// 	// if sourceType == "custom" {
	// 	// 	nameWithoutSpaces := strings.Replace(n["name"], " ", "_", -1)

	// 	// }
	// 	// }else{

	// 	// }

	// 	err = os.RemoveAll(pathToDelete)
	// 	if err = uuidPath.Scan(&pathToDelete); err != nil {
	// 		logs.Error("DeleteRulesetSource Error deleting path: %s", err.Error())
	// 		return err
	// 	}
	// }
	
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

// //update ruleset data from ruleset source page
// func UpdateRulesetSource(param string, value string, sourceuuid string)(err error){
// 	editSource, err := ndb.Rdb.Prepare("update ruleset set ruleset_value = ? where ruleset_param = ? and ruleset_uniqueid = ?")
// 	if err != nil {
// 		logs.Error("Prepare EditRulesetSource-> %s", err.Error())
// 		return err
// 	}
// 	_, err = editSource.Exec(&value, &param, &sourceuuid)
// 	if err != nil {
// 		logs.Error("Execute EditRulesetSource-> %s", err.Error())
// 		return err
// 	}

// 	return nil
// }

func OverwriteDownload(data map[string]string) (err error) {
	var fileExtension = regexp.MustCompile(`(\w+).rules$`)
	var newFilesDownloaded = make(map[string]string)
	var newFilesDB = make(map[string]map[string]string)

	sourceDownload := map[string]map[string]string{}
	sourceDownload["ruleset"] = map[string]string{}
	sourceDownload["ruleset"]["sourceDownload"] = ""
	sourceDownload,err = utils.GetConf(sourceDownload)
	pathDownloaded := sourceDownload["ruleset"]["sourceDownload"]
	
	pathAndFile,err := ndb.GetRulesetSourceValue(data["uuid"], "path")
	if err != nil {
		logs.Error("Error Getting path for download file from RulesetSource-> %s", err.Error())
		return err
	}

	splitPath := strings.Split(pathAndFile, "/")
	pathSelected := splitPath[len(splitPath)-2]

	_ = os.RemoveAll(pathDownloaded+pathSelected)

	//download file
	if _, err := os.Stat(pathDownloaded+pathSelected); os.IsNotExist(err) {
		os.MkdirAll(pathDownloaded+pathSelected, os.ModePerm)

		err = utils.DownloadFile(data["path"], data["url"])
		if err != nil {
			logs.Error("Error downloading file from RulesetSource-> %s", err.Error())
			_ = os.RemoveAll(pathDownloaded+pathSelected)

			// update ruleset "exists" field
			_ = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "false")

			return err
		}
	
		// err = utils.ExtractTarGz(data["path"], pathDownloaded, pathSelected)
		err = utils.ExtractTarGz(data["path"], pathDownloaded+pathSelected)
		if err != nil {
			logs.Error("Error unzipping file downloaded: "+err.Error())
			err = os.RemoveAll(pathDownloaded+pathSelected)
			
			// update ruleset "exists" field
			err = ndb.UpdateRuleset(data["uuid"], "isDownloaded", "false")
			if err != nil {
				logs.Error("UpdateRuleset Error from RulesetSource-> %s", err.Error())
				return err
			}

			return err
		}
	logs.Info("Download and extract complete")

	}

	//get map with new files downloaded
	err = filepath.Walk(pathDownloaded+pathSelected,
		func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileExtension.MatchString(info.Name()){
			newFilesDownloaded[info.Name()] = info.Name()
		}
		return nil
	})

	//Create map
	dbFiles,err := GetDetails(data["uuid"])
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
		

		// if newFilesDB[w]["count"] == "0" {
		// 	err = ndb.UpdateRuleFiles(newFilesDB[w]["uuid"], "exists", "false")
		// 	if (err != nil){
		// 		logs.Error("OverwriteDownload UPDATE error for update isDownloaded -- "+err.Error())
		// 		return err
		// 	}

		// 	// err = ndb.UpdateRuleFiles(newFilesDB[w]["uuid"], "isUpdated", "true")
		// 	// if (err != nil){
		// 	// 	logs.Error("OverwriteDownload UPDATE error for update isDownloaded -- "+err.Error())
		// 	// 	return err
		// 	// }

		// 	// var uuidFromNewFiles string
		// 	// sql := "select rule_uniqueid from rule_files where rule_param='sourceFileUUID' and rule_value='"+newFilesDB[w]["uuid"]+"';"
		// 	// rows, err := ndb.Rdb.Query(sql)
		// 	// for rows.Next() {
		// 	// 	if err = rows.Scan(&uuidFromNewFiles); err != nil {
		// 	// 		logs.Error("OverwriteDownload GetRulesetSourcePath rows.Scan: %s", err.Error())
		// 	// 		return err
		// 	// 	}
		// 	// 	err = ndb.UpdateRuleFiles(uuidFromNewFiles, "exists", "false")
		// 	// 	if (err != nil){
		// 	// 		logs.Error("OverwriteDownload UPDATE error for update isDownloaded -- "+err.Error())
		// 	// 		return err
		// 	// 	}
		// 	// } 
		// } else 
		if newFilesDB[w]["count"] == "1" {
			uuid := utils.Generate()
			err = ndb.InsertRulesetSourceRules(uuid, "name", pathSelected)
			err = ndb.InsertRulesetSourceRules(uuid, "path", pathAndFile)
			err = ndb.InsertRulesetSourceRules(uuid, "file", w)
			err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
			err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", data["uuid"])
			err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
			// err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
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
	
		// err = utils.ExtractTarGz(data["path"], pathDownloaded, data["name"])
		err = utils.ExtractTarGz(data["path"], pathDownloaded+pathSelected)
		if err != nil {
			logs.Error("Error unzipping file downloaded: "+err.Error())
			err = os.RemoveAll(pathDownloaded+pathSelected)
			return err
		}

		//insert into DB
		ruleFiles, err := Details(data)
		for k,_ := range ruleFiles["files"] {
			uuid := utils.Generate()
			err = ndb.InsertRulesetSourceRules(uuid, "name", data["name"])
			err = ndb.InsertRulesetSourceRules(uuid, "path", ruleFiles["files"][k])
			err = ndb.InsertRulesetSourceRules(uuid, "file", k)
			err = ndb.InsertRulesetSourceRules(uuid, "type", "source")
			err = ndb.InsertRulesetSourceRules(uuid, "sourceUUID", data["uuid"])
			err = ndb.InsertRulesetSourceRules(uuid, "exists", "true")
			// err = ndb.InsertRulesetSourceRules(uuid, "isUpdated", "false")
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
					returnLines["newFile"] = data["new"]
					returnLines["old"] = file2[y]["Line"]
					returnLines["oldFile"] = data["old"]
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
	var allRuleDetails = map[string]map[string]string{}
	var uniqid string
    var param string
    var value string
	var uuidSource string
	var checked string
    if ndb.Rdb == nil {
        logs.Error("no access to database")
        return nil, errors.New("no access to database")
    }
	sqlUUID := "select rule_uniqueid from rule_files where rule_param='sourceUUID' and rule_value = '"+uuid+"';"
	uuidRows, err := ndb.Rdb.Query(sqlUUID)
	if err != nil {
		logs.Error("ndb.Rdb.Query Error checking uuid for take the uuid list for GetDetails: %s", err.Error())
        return nil, err
    }
	defer uuidRows.Close()
	for uuidRows.Next() {
		if err = uuidRows.Scan(&uuidSource); err != nil {
            logs.Error("GetDetails UUIDSource uuidRows.Scan: %s", err.Error())
            return nil, err
		}		
		sql := "select rule_uniqueid, rule_param, rule_value from rule_files where rule_uniqueid='"+uuidSource+"';"
		rows, err := ndb.Rdb.Query(sql)
		if err != nil {
			logs.Error("ndb.Rdb.Query Error : %s", err.Error())
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

			// //check diff
			// allRuleDetails[uniqid]["isUpdated"]=verifyDiff(local,surce)

			// //check if file exist
			// if param == "path" {
			// 	checked = VerifyPathExists(allRuleDetails[uuidSource]["path"])								
			// }else if param == "exists" {
			// 	allRuleDetails[uuidSource]["exists"]=checked
			// }else{
				// allRuleDetails[uniqid][param]=value
			// }
		} 
	}
	
	for x := range allRuleDetails{
		checked = VerifyPathExists(allRuleDetails[x]["path"])
		logs.Warn(allRuleDetails[x]["file"]+" -*- "+checked)
		allRuleDetails[x]["exists"]=checked
		err = ndb.UpdateRuleFiles(x, "exists", checked)
		if err != nil {
			logs.Error("ndb.UpdateRuleFiles Error : %s", err.Error())
			return nil, err
		}
	}
	return allRuleDetails, nil
}

func VerifyPathExists(path string)(stauts string){
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "false"
	}else{
		return "true"
	}
}