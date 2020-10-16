package utils

import (
	"github.com/astaxie/beego/logs"
)

func Logger() {
    var err error
	//get logger data
	filepath, err := GetKeyValueString("logs", "filepath")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	filename, err := GetKeyValueString("logs", "filename")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	maxlines, err := GetKeyValueString("logs", "maxlines")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	maxsize, err := GetKeyValueString("logs", "maxsize")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	daily, err := GetKeyValueString("logs", "daily")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	maxdays, err := GetKeyValueString("logs", "maxdays")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	rotate, err := GetKeyValueString("logs", "rotate")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	level, err := GetKeyValueString("logs", "level")
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: " + err.Error())
	}
	maxfiles, err := GetKeyValueInt("logs", "maxfiles")
	if err != nil {
		logs.Error("Error getting data from main.conf for load Logger data: " + err.Error())
	}

	//transform maxsize to bytes
	newMaxSize,_ := GetBytesFromSizeType(maxsize)
	//clear older log files
	err = ClearOlderLogFiles(filepath, filename+"." , maxfiles)
	if err != nil {logs.Error(err.Error())}
	//create logger
	logs.NewLogger(10000)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+filepath+filename+`", "maxlines":`+maxlines+` ,"maxsize":`+newMaxSize+`, "daily":`+daily+`, "maxdays":`+maxdays+`, "rotate":`+rotate+`, "level":`+level+`}`)
}