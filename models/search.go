package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  "github.com/OwlH-net/OwlH-Master/search"
)

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/search/getRulesetsBySearch \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "search": "v",
//      "rulesetName": "v"
//   }
func GetRulesetsBySearch(anode map[string]string, username string) (data interface{}, err error) {
  data, err = search.GetRulesetsBySearch(anode)
  changecontrol.ChangeControlInsertData(err, "GetRulesetsBySearch", username)
  return data, err
}
