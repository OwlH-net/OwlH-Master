package search

import (
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "owlhmaster/database"
    "owlhmaster/elk"
    "owlhmaster/ruleset"
    "owlhmaster/utils"
    "regexp"
    "strconv"
    "strings"
    "time"
)

//Struct for search over a ruleset files
type Rule struct {
    Sid      string `json:"sid"`
    Msg      string `json:"msg"`
    Rulesets []Ruleset
}
type Ruleset struct {
    File       string   `json:"file"`
    Status     string   `json:"status"`
    Name       string   `json:"name"`
    Uuid       string   `json:"uuid"`
    SourceType string   `json:"sourceType"`
    Type       string   `json:"type"`
    Node       []string `json:"node"`
}

var RulesIndex []Rule = nil

func GetRulesetsBySearch(anode map[string]string) (data []Rule, err error) {
    var matchingRules []Rule = nil
    var isSID = regexp.MustCompile(`(\d+)`)
    sid := isSID.FindStringSubmatch(anode["search"])

    for w := range RulesIndex {
        if sid != nil {
            if strings.Contains(RulesIndex[w].Sid, anode["search"]) {
                currentRulesets := RulesIndex[w].Rulesets
                if anode["rulesetName"] == "" {
                    matchingRules = append(matchingRules, RulesIndex[w])
                } else {
                    for z := range currentRulesets {
                        if currentRulesets[z].Name == anode["rulesetName"] {
                            matchingRules = append(matchingRules, RulesIndex[w])
                        }
                    }
                }
            }
        } else {
            if strings.Contains(strings.ToLower(RulesIndex[w].Msg), strings.ToLower(anode["search"])) {
                currentRulesets := RulesIndex[w].Rulesets
                if anode["rulesetName"] == "" {
                    matchingRules = append(matchingRules, RulesIndex[w])
                } else {
                    for z := range currentRulesets {
                        if currentRulesets[z].Name == anode["rulesetName"] {
                            matchingRules = append(matchingRules, RulesIndex[w])
                        }
                    }
                }
            }
        }
    }
    return matchingRules, err
}

func Init() {
    t, err := utils.GetKeyValueString("loop", "search")
    if err != nil {
        logs.Error("Search Error: Cannot load search information.")
    }
    tDuration, err := strconv.Atoi(t)
    for {
        BuildRuleIndex()
        time.Sleep(time.Minute * time.Duration(tDuration))
    }
}

func BuildRuleIndex() {
    BuildRuleIndexLocal()
    //BuildRuleIndexElastic()
}

func BuildRuleIndexElastic() {
    var jsonRules []string
    allRulesets, err := ndb.GetAllRuleFiles()
    if err != nil {
        logs.Error("BuildRuleIndexElastic Error getting all rule files: " + err.Error())
        return
    }

    for x := range allRulesets {
        if allRulesets[x]["type"] == "source" {
            continue
        }
        currentRules, err := ruleset.ReadRuleset(allRulesets[x]["path"])
        if err != nil {
            logs.Error("BuildRuleIndexElastic Error getting rule file content: " + err.Error())
            continue
        }

        for r := range currentRules {
            currentRuleData := make(map[string]string)
            //regexp
            var validID = regexp.MustCompile(`([^\(]+)\((.*)\)`)
            sid := validID.FindStringSubmatch(currentRules[r]["raw"])
            if len(sid) != 3 {
                logs.Error("Rule header bad sintax: " + currentRules[r]["raw"])
                continue
            }

            match := strings.Split(sid[1], " ")
            //Check if rules are empty
            if len(match) != 7 {
                continue
            }

            //create rule map
            currentRuleData["raw"] = currentRules[r]["raw"]
            currentRuleData["status"] = currentRules[r]["enabled"]
            currentRuleData["type"] = strings.Trim(match[0], "#")
            currentRuleData["proto"] = match[1]
            currentRuleData["srcip"] = match[2]
            currentRuleData["srcport"] = match[3]
            currentRuleData["direction"] = match[4]
            currentRuleData["dstip"] = match[5]
            currentRuleData["dstport"] = match[6]
            currentRuleData["fileName"] = allRulesets[x]["path"]
            currentRuleData["RulesetName"] = allRulesets[x]["name"]

            matchContent := strings.Split(sid[2], ";")
            for h := range matchContent {
                if matchContent[h] == "" {
                    continue
                }
                if strings.Contains(matchContent[h], ":") {
                    keyValue := strings.Split(matchContent[h], ":")

                    if len(keyValue) != 2 {
                        logs.Error("Error key value not found: " + currentRules[r]["raw"])
                        continue
                    }
                    keyValue[0] = strings.Replace(keyValue[0], " ", "", -1)
                    if keyValue[0] == "" {
                        continue
                    }

                    if _, ok := currentRuleData[keyValue[0]]; ok {
                        currentRuleData[keyValue[0]] = currentRuleData[keyValue[0]] + " -- " + keyValue[1]
                    } else {
                        currentRuleData[keyValue[0]] = keyValue[1]
                    }
                } else {
                    matchContent[h] = strings.Replace(matchContent[h], " ", "", -1)
                    currentRuleData[matchContent[h]] = ""
                }
            }

            ruleElasticOutput, err := json.Marshal(currentRuleData)
            if err != nil {
                logs.Error("BuildRuleIndexElastic Error creating json file: " + err.Error())
            }

            jsonRules = append(jsonRules, string(ruleElasticOutput))
        }
    }

    elk.Init(jsonRules)
    logs.Info("Elastic data loaded")
}

func BuildRuleIndexLocal() {
    RulesIndex = nil
    exists := false
    cont := 0
    allRulesets, err := ndb.GetAllRuleFiles()
    if err != nil {
        logs.Error("Search/Init error: %s", err.Error())
        return
    }
    for x := range allRulesets {
        rset := Ruleset{}
        currentRules, err := ruleset.ReadRuleset(allRulesets[x]["path"])
        if err != nil {
            logs.Error("BuildRuleIndexLocal Error readding rulesets: " + err.Error())
            continue
        }
        rset.File = allRulesets[x]["path"]
        rset.Name = allRulesets[x]["name"]
        rset.SourceType = allRulesets[x]["sourceType"]
        rset.Type = allRulesets[x]["type"]
        for y := range currentRules {
            rset.Status = currentRules[y]["enabled"]
            rset.Uuid = x
            rule := Rule{}
            rule.Rulesets = append(rule.Rulesets, rset)
            rule.Sid = currentRules[y]["sid"]
            rule.Msg = currentRules[y]["msg"]
            exists = false

            for w := range RulesIndex {
                if RulesIndex[w].Sid == rule.Sid {
                    RulesIndex[w].Rulesets = append(RulesIndex[w].Rulesets, rset)
                    exists = true
                    break
                }
            }
            if !exists {
                RulesIndex = append(RulesIndex, rule)
            }
            cont++
        }
    }

    logs.Info("Ruleset list loaded")
}
