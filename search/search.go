package search

import(
	// "regexp"
    // "os"
	// "os/exec"
	"encoding/json"
	"owlhmaster/database"
	"time"
    "owlhmaster/ruleset"
    "owlhmaster/elk"
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
	Uuid    	string		`json:"uuid"`
	Node    	[]string	`json:"node"`
}
var RulesIndex []Rule = nil

func GetRulesetsBySearch(anode map[string]string)(data []Rule, err error) {
	var matchingRules []Rule = nil
	var isSID = regexp.MustCompile(`(\d+)`)
	sid := isSID.FindStringSubmatch(anode["search"])
	
	for w := range RulesIndex {
		if sid != nil{
			if strings.Contains(RulesIndex[w].Sid, anode["search"]){
				currentRulesets := RulesIndex[w].Rulesets
				if anode["rulesetName"] == ""{
					// for z := range currentRulesets {
						matchingRules = append(matchingRules, RulesIndex[w])
					// }
				}else{					
					for z := range currentRulesets {
						if currentRulesets[z].Name == anode["rulesetName"]{
							// if RulesIndex[w].Rulesets.Uuid != currentRulesets[z].Uuid{
								matchingRules = append(matchingRules, RulesIndex[w])
							// }
						}
					}
				}
			}
		}else {
			if strings.Contains(strings.ToLower(RulesIndex[w].Msg), strings.ToLower(anode["search"])){
				currentRulesets := RulesIndex[w].Rulesets
				if anode["rulesetName"] == ""{
					// for z := range currentRulesets {
					// 	Rsets := Ruleset{}
					// 	Rsets.Name = currentRulesets[z].Name
					// 	Rsets.Status = currentRulesets[z].Status
					// 	Rsets.File = currentRulesets[z].File
					// 	Rsets.Uuid = currentRulesets[z].Uuid
						
					// 	var NewRule Rule
					// 	NewRule.Sid = RulesIndex[w].Sid
					// 	NewRule.Msg = RulesIndex[w].Msg
					// 	NewRule.Rulesets = append(NewRule.Rulesets, Rsets) 

						// matchingRules = append(matchingRules, NewRule)
						matchingRules = append(matchingRules, RulesIndex[w])
					// }
				}else{
					for z := range currentRulesets {
						if currentRulesets[z].Name == anode["rulesetName"]{
							// Rsets := Ruleset{}
							// Rsets.Name = currentRulesets[z].Name
							// Rsets.Status = currentRulesets[z].Status
							// Rsets.File = currentRulesets[z].File
							// Rsets.Uuid = currentRulesets[z].Uuid

							// var NewRule Rule
							// NewRule.Sid = RulesIndex[w].Sid
							// NewRule.Msg = RulesIndex[w].Msg
							// NewRule.Rulesets = append(NewRule.Rulesets, Rsets) 
	
							// matchingRules = append(matchingRules, NewRule)
							matchingRules = append(matchingRules, RulesIndex[w])
						}
					}
				}
			}	
		}
	}
	return matchingRules, err
}

func Init()(){
	for {
		BuildRuleIndex()
		time.Sleep(5 * time.Minute)
	}	
	// for {
	// 	RulesIndex = nil
	// 	exists := false
	// 	allRulesets,err := ndb.GetAllRuleFiles()
	// 	if err != nil {logs.Error("Search/Init error: %s", err.Error())}
	// 	for x,_ := range allRulesets {	
	// 		rset := Ruleset{}
	// 		currentRules, _ := ruleset.ReadRuleset(allRulesets[x]["path"])
	// 		rset.File = allRulesets[x]["path"]
	// 		rset.Name = allRulesets[x]["name"]
	// 		for y := range currentRules {				
	// 			rset.Status = currentRules[y]["enabled"]
	// 			rset.Uuid = x
	// 			rule := Rule{}
	// 			rule.Rulesets = append(rule.Rulesets, rset)
	// 			rule.Sid = currentRules[y]["sid"]
	// 			rule.Msg = currentRules[y]["msg"]
	// 			exists = false

	// 			for w := range RulesIndex {
	// 				if RulesIndex[w].Sid == rule.Sid{	
	// 					RulesIndex[w].Rulesets = append(RulesIndex[w].Rulesets, rset)
	// 					exists=true
	// 					break
	// 				}
	// 			}
	// 			if !exists {
	// 				RulesIndex = append(RulesIndex, rule)
	// 			}
	// 		}
	// 	}

	// 	//load nodes with rulesets
	// 	nodes,err := ndb.GetAllNodes() 
	// 	for f := range nodes{
	// 		nodeName,err := ndb.ObtainNodeName(nodes[f])
	// 	}

	// 	logs.Info("Ruleset list has been updated.")
	// 	time.Sleep(5 * time.Minute)
	// }
}

func BuildRuleIndex()(){
	BuildRuleIndexLocal()
	BuildRuleIndexElastic()
}

func BuildRuleIndexElastic()(){
	var jsonRules []string
	allRulesets,err := ndb.GetAllRuleFiles()
	if err!=nil {logs.Error("BuildRuleIndexElastic Error getting all rule files: "+err.Error())}

	for x,_ := range allRulesets {
		if allRulesets[x]["type"] == "source" {continue}
		currentRules, err := ruleset.ReadRuleset(allRulesets[x]["path"])
		if err!=nil {logs.Error("BuildRuleIndexElastic Error getting rule file content: "+err.Error())}

		for r := range currentRules{
			currentRuleData := make(map[string]string)
			//regexp
			var validID = regexp.MustCompile(`([^\(]+)\((.*)\)`)
			sid := validID.FindStringSubmatch(currentRules[r]["raw"])

			match := strings.Split(sid[1], " ")
			currentRuleData["raw"] = 		currentRules[r]["raw"]
			currentRuleData["status"] = 	currentRules[r]["enabled"]
			currentRuleData["type"] = 		strings.Trim(match[0], "#")
			currentRuleData["proto"] = 		match[1]
			currentRuleData["srcip"] = 		match[2]
			currentRuleData["srcport"] = 	match[3]
			currentRuleData["direction"] = 	match[4]
			currentRuleData["dstip"] = 		match[5]
			currentRuleData["dstport"] = 	match[6]
			currentRuleData["fileName"] = 	allRulesets[x]["path"]
			currentRuleData["RulesetName"] =allRulesets[x]["name"]

			matchContent := strings.Split(sid[2], ";")
			for h := range matchContent{
				if matchContent[h] == "" {continue}
				if strings.Contains(matchContent[h], ":"){
					keyValue := strings.Split(matchContent[h], ":")
					keyValue[0] = strings.Replace(keyValue[0]," ","",-1)
					if keyValue[0] == "" {continue}

					if _, ok := currentRuleData[keyValue[0]]; ok {
						currentRuleData[keyValue[0]] = currentRuleData[keyValue[0]]+" -- "+keyValue[1]
					}else{
						currentRuleData[keyValue[0]] = keyValue[1]
					}
				}else{
					matchContent[h] = strings.Replace(matchContent[h]," ","",-1)
					currentRuleData[matchContent[h]] = ""
				}					
			}
			ruleElasticOutput, err := json.Marshal(currentRuleData)
			if err!=nil {logs.Error("BuildRuleIndexElastic Error creating json file: "+err.Error())}
	
			jsonRules = append(jsonRules, string(ruleElasticOutput))
		} 	
	}

	elk.Init(jsonRules)
	logs.Notice("Elastic data loaded")
}


func BuildRuleIndexLocal()(){
	RulesIndex = nil
	exists := false
	cont := 0
	allRulesets,err := ndb.GetAllRuleFiles()
	if err != nil {logs.Error("Search/Init error: %s", err.Error())}
	for x,_ := range allRulesets {	
		rset := Ruleset{}
		currentRules, _ := ruleset.ReadRuleset(allRulesets[x]["path"])
		rset.File = allRulesets[x]["path"]
		rset.Name = allRulesets[x]["name"]
		for y := range currentRules {			
			rset.Status = currentRules[y]["enabled"]
			rset.Uuid = x
			rule := Rule{}
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
			cont++
		}
	}
	// nodes,err := ndb.GetAllNodes()
	// for x,_ := range allRulesets {	
	// 	for f := range nodes{
	// 		currentRules, _ := ruleset.ReadRuleset(allRulesets[x]["path"])
	// 		if allRulesets[x]["sourceUUID"] == f {
	// 			rset := Ruleset{}
	// 			nodeName,err := ndb.ObtainNodeName(nodes[f])
	// 			rset.Node = append(rset.Node, nodeName)
	// 			rule.Rulesets = append(rule.Rulesets, rset)
	// 		}
	// 		if RulesIndex[w].Sid == currentRules[w]["sid"]{	
	// 			RulesIndex[w].Rulesets = append(RulesIndex[w].Rulesets, rset)
	// 			exists=true
	// 			break
	// 		}
	// 		for y := range currentRules {	
	// 			for w,l := range RulesIndex {
	// 				if contains(arr,nodeName){

	// 				}

	// 			}
	// 			if !exists {
	// 				RulesIndex = append(RulesIndex, rule)
	// 			}

	// 		}
	// 	}

	// }




	logs.Info("Ruleset list loaded")
}


// nodes,err := ndb.GetAllNodes()
// if err != nil {logs.Error("Error getting all ruleset values: "+err.Error())}
// for f := range nodes{
// 	if allRulesets[x]["sourceUUID"] == f{
// 		nodeName,err := ndb.ObtainNodeName(nodes[f])
// 		if err != nil {logs.Error("Error getting node name from their ruleset: "+err.Error())}
// 		logs.Warn(nodes[f]+"  ->  "+nodeName)
		
		
// 		// // isNodeYet := false
// 		// for r := range rset.Node{
// 		// 	if rset.Node[r] == nodeName{
// 		// 		// isNodeYet = true
// 		// 		break;
// 		// 	}else{
// 		// 		rset.Node = append(rset.Node, nodeName)
// 		// 	}
// 		// }
// 		// // if !isNodeYet{
// 		// // }
// 	}
// }