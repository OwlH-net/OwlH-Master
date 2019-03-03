package main

import (

    "github.com/astaxie/beego/logs"
    _ "owlhmaster/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "owlhmaster/database"
    //"owlhmaster/models"
    //"owlhmaster/utils"
    //"fmt"
    //"io/ioutil"
)


func main() {

    ndb.Conn()
    ndb.RConn()
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




