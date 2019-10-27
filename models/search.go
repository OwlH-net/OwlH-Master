package models 

import (
    "owlhmaster/search"
)

func GetRulesetsBySearch(anode map[string]string)(data interface{}, err error) {
	data, err = search.GetRulesetsBySearch(anode)
	return data, err
}