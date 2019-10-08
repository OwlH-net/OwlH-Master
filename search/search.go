package search

import(
	// "regexp"
    // "os"
	// "os/exec"
	// "encoding/json"
	"owlhmaster/database"
	"time"
    "owlhmaster/ruleset"
    // "errors"
	// "database/sql"
    "strings"
	// "strconv"
	"regexp"
	// "io/ioutil"
    "github.com/astaxie/beego/logs"
)


//Struct for search over a ruleset files
type Rule struct {
	Sid 	  	string		`json:"sid"`
	Msg       	string		`json:"msg"`
	Rulesets  	[]Ruleset	
}
type Ruleset struct {
	File      	string		`json:"file"`
	Status    	string		`json:"status"`
	Name    	string		`json:"name"`
}
var RulesIndex []Rule = nil

func GetRulesetsBySearch(anode map[string]string)(data []Rule, err error) {
	logs.Debug(anode)
	var matchingRules []Rule = nil
	var isSID = regexp.MustCompile(`(\d+)`)
	sid := isSID.FindStringSubmatch(anode["search"])
	for w := range RulesIndex {
		if sid != nil{
			if strings.Contains(RulesIndex[w].Sid, anode["search"]){
				// matchingRules = append(matchingRules, RulesIndex[w])
				currentRulesets := RulesIndex[w].Rulesets
				for z := range currentRulesets {
					if currentRulesets[z].Name == anode["rulesetName"]{
						logs.Info(currentRulesets[z].Name)
						matchingRules = append(matchingRules, RulesIndex[w])
					}
				}
				// for s := range RulesIndex.Rulesets {
				// 	logs.Notice(s)
				// 	// logs.Notice(RulesIndex[w])
				// 	Ruleset = append(Ruleset, RulesIndex[w])
				// }
			}
		}else {
			if strings.Contains(strings.ToLower(RulesIndex[w].Msg), strings.ToLower(anode["search"])){
				currentRulesets := RulesIndex[w].Rulesets
				for z := range currentRulesets {
					if currentRulesets[z].Name == anode["rulesetName"]{
						Rsets := Ruleset{}
						Rsets.Name = currentRulesets[z].Name
						Rsets.Status = currentRulesets[z].Status
						Rsets.File = currentRulesets[z].File
						
						var NewRule Rule
						NewRule.Sid = RulesIndex[w].Sid
						NewRule.Msg = RulesIndex[w].Msg
						NewRule.Rulesets = append(NewRule.Rulesets, Rsets) 

						matchingRules = append(matchingRules, NewRule)
					}
				}
			}	
		}
	}

	logs.Notice(matchingRules)

	return matchingRules, err
}

func Init()(){
	for {
		exists := false
		allRulesets,err := ndb.GetAllRuleFiles()
		if err != nil {logs.Error("Search/Init error: %s", err.Error())}
		for x,_ := range allRulesets {	
			rset := Ruleset{}
			currentRules, _ := ruleset.ReadRuleset(allRulesets[x]["path"])
			rset.File = allRulesets[x]["path"]
			rset.Name = allRulesets[x]["name"]
			for y := range currentRules {
				rule := Rule{}
				rset.Status = currentRules[y]["enabled"]
				rule.Rulesets = append(rule.Rulesets, rset)
				rule.Sid = currentRules[y]["sid"]
				rule.Msg = currentRules[y]["msg"]
				exists = false
			
				for w := range RulesIndex {
					if RulesIndex[w].Sid == rule.Sid{	
						RulesIndex[w].Rulesets = append(RulesIndex[w].Rulesets, rset)
						exists=true
						break
					}
				}
				if !exists {
					RulesIndex = append(RulesIndex, rule)
				}
			}
		}
		time.Sleep(60 * time.Minute)
		logs.Info("Ruleset list has been updated.")
	}
}