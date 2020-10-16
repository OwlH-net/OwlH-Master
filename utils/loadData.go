package utils

import (
    "github.com/astaxie/beego/logs"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"fmt"
    "errors"
)

var mainconfData map[string]interface{}
func Load()(){
	confFilePath := "conf/main.conf"
	jsonPath, err := ioutil.ReadFile(confFilePath)	
	err = json.Unmarshal(jsonPath, &mainconfData)
	if err != nil {logs.Error(err.Error())}
}

func GetKeyValueString(key,sub string)(result string, err error){	
	keyValue,err := GetKeyValue(key,sub)
	if err != nil {logs.Error(err.Error()); return "", err}
	switch w:= reflect.ValueOf(keyValue); w.Kind() {
	case reflect.String:
		return w.String(), nil
	default:
		return "",errors.New("GetKeyValueString This value is not a String")
	}
}
func GetKeyValueSlice(key,sub string)(result []string, err error){
	keyValue,err := GetKeyValue(key,sub)
	if err != nil {logs.Error(err.Error()); return nil, err}
	switch w:= reflect.ValueOf(keyValue); w.Kind() {
	case reflect.Slice:
		return keyValue.([]string), nil
	default:
		return nil,errors.New("GetKeyValueSlice This value is not a Slice")
	}
}
func GetKeyValueBool(key, sub string) (result bool, err error) {
	keyValue, err := GetKeyValue(key, sub)
	if err != nil {
		logs.Error(err.Error())
		return false, err
	}
	switch w := reflect.ValueOf(keyValue); w.Kind() {
	case reflect.Bool:
		return keyValue.(bool), nil
	default:
		return false, errors.New("GetKeyValueBool This value is not a Slice")
	}
}
func GetKeyValueInt(key, sub string) (result int, err error) {
	keyValue, err := GetKeyValue(key, sub)
	if err != nil {
		logs.Error(err.Error())
		return -1, err
	}
	switch w := reflect.ValueOf(keyValue); w.Kind() {
	case reflect.Int:
		return keyValue.(int), nil
	case reflect.Float64:
		return int(keyValue.(float64)), nil
	default:
		nerr := fmt.Sprintf("GetKeyValueInt This value is not a Integer --> %v",w.Kind())
		return -1, errors.New(nerr)
	}
}

func GetKeyValue(key,sub string)(result interface{}, err error){
	if _,ok := mainconfData[key]; ok{
		newData := mainconfData[key].(map[string]interface{})
		if _,ok := newData[sub]; ok{
			return newData[sub], nil
		}else{
			return nil, errors.New("GetKeyValue subkey "+sub+" is not available")
		}
	}
	return nil, errors.New("GetKeyValue key "+key+" is not available")
}