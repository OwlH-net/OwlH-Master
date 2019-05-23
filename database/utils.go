
package ndb

import (
    "github.com/astaxie/beego/logs"
	"regexp"
	"bufio"
	"os"
)

func CreateNewRuleFile(uuid string)(RulesData []byte, err error){
	var uniqueid string
	var rulePath string
	var uuidArray []string
	var validID = regexp.MustCompile(`sid:(\d+);`)

	//read rule uuid
	uuidRules, err := Rdb.Query("select rule_uniqueid from rule_files where rule_value='"+uuid+"'")
	if err != nil {
		logs.Error("CreateNewRuleFile ndb.Rdb.Query Error checking rule_uniqueid for rule_files: %s", err.Error())
		return nil, err
	}
	defer uuidRules.Close()
	for uuidRules.Next() {
		if err = uuidRules.Scan(&uniqueid); err != nil {
			logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error())
			return nil, err
		}
		uuidArray = append(uuidArray, uniqueid)
	}

	//read files paths and
	for x := range uuidArray{
		rules, err := Rdb.Query("select rule_value from rule_files where rule_param = 'path' and rule_uniqueid= '"+uuidArray[x]+"'")
		if err != nil {
			logs.Error("CreateNewRuleFile ndb.Rdb.Query Error loading files paths: %s", err.Error())
			return nil, err
		}
		defer rules.Close()
		for rules.Next() {
			if err = rules.Scan(&rulePath); err != nil {
				logs.Error("CreateNewRuleFile rows.Scan: %s", err.Error())
				return nil,err
			}
			file, err := os.Open(rulePath)
			if err != nil {
				logs.Error("File reading error: %s", err.Error())
				return nil, err
			}
			scanner := bufio.NewScanner(file)
			for scanner.Scan(){
				if validID.MatchString(scanner.Text()){
					// RulesData = append(RulesData, scanner.Bytes())
					RulesData[len(RulesData)+1] = scanner.Bytes()
				}
			}
		}
	}
	logs.Warn(RulesData)
	return RulesData,nil
}