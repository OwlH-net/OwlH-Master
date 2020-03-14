package utils

import (
    "github.com/astaxie/beego/logs"
	"encoding/json"
	"io/ioutil"
	"reflect"
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