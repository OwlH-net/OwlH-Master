package main

import (

    "github.com/astaxie/beego/logs"
    _ "owlhmaster/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "owlhmaster/database"
    "owlhmaster/dispatcher"
    "owlhmaster/utils"
)


func main() {
	//Configuration for the logger
	var err error
	loadDataLogger := map[string]map[string]string{}
	loadDataLogger["logs"] = map[string]string{}
	loadDataLogger["logs"]["filename"] = ""
	loadDataLogger["logs"]["maxlines"] = ""
	loadDataLogger["logs"]["maxsize"] = ""
	loadDataLogger["logs"]["daily"] = ""
	loadDataLogger["logs"]["maxdays"] = ""
	loadDataLogger["logs"]["rotate"] = ""
	loadDataLogger["logs"]["level"] = ""
	loadDataLogger, err = utils.GetConf(loadDataLogger)    
    filename := loadDataLogger["logs"]["filename"]
	maxlines := loadDataLogger["logs"]["maxlines"]
	maxsize := loadDataLogger["logs"]["maxsize"]
	daily := loadDataLogger["logs"]["daily"]
	maxdays := loadDataLogger["logs"]["maxdays"]
	rotate := loadDataLogger["logs"]["rotate"]
	level := loadDataLogger["logs"]["level"]
	if err != nil {
		logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())
	}
	logs.NewLogger(10000)
	logs.SetLogger(logs.AdapterFile,`{"filename":"`+filename+`", "maxlines":`+maxlines+` ,"maxsize":`+maxsize+`, "daily":`+daily+`, "maxdays":`+maxdays+`, "rotate":`+rotate+`, "level":`+level+`}`)

	//Application version
	logs.Error("Version: 0.6.190528.0947")

	//Init database connection
    ndb.Conn()
    ndb.RConn()
	ndb.GConn()
	// ndb.RSConn()
	
	//Init dispatcher at master
    go dispatcher.Init()
	
	//Beego API documentation
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }

    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
        AllowCredentials: true,
    }))

    beego.Run()
}




