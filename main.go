package main

import (

    "github.com/astaxie/beego/logs"
    // "github.com/astaxie/beego/context"
    _ "owlhmaster/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "owlhmaster/database"
    "owlhmaster/dispatcher"
    "owlhmaster/master"
    "owlhmaster/search"
    "owlhmaster/scheduler"
    "owlhmaster/utils"
    "owlhmaster/configuration"
    // "owlhmaster/validation"
    // "owlhmaster/controllers"
    "os"
    "crypto/tls"
    "bufio"
    "strings"
    "runtime"
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
    logs.Info("Version OwlH Master: 0.12.0.20200203")

    cancontinue := configuration.MainCheck()
    if !cancontinue {
        logs.Error("can't continue, see previous logs")
        // return 
    }

    //operative system values
    data:=OperativeSystemValues()
    for x := range data {
        if (x == "ID" || x == "ID_LIKE" || x == "VERSION_ID"){
            logs.Info(x +" -- "+data[x])
        }
    }
    //Init database connection
    ndb.Conn()
    ndb.RConn()
    // ndb.GConn()    
    ndb.MConn()    
    
    //CheckServicesStatus
    master.CheckServicesStatus()

    //Init dispatcher at master
    go dispatcher.Init()
    //Init scheduler at master
    go scheduler.Init()
    //Load all rulesets
    go search.Init()

    //Beego API documentation
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
        beego.BConfig.WebConfig.Session.SessionOn = true
    }


    beego.BeeApp.Server.TLSConfig = &tls.Config{    CipherSuites: []uint16{
                                                        tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
                                                        tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
                                                    },
                                                    MinVersion:               tls.VersionTLS12,
                                                    PreferServerCipherSuites: true,
                                                }

    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "token", "user", "uuid"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
        AllowCredentials: true,
    }))

    // var TokenValidation = func(ctx *context.Context) {
    //     err := validation.CheckToken(ctx.Input.Header("token"), ctx.Input.Header("user"), ctx.Input.Header("uuid"))
    //     if err != nil {            
    //         logs.Error(err)
    //         beego.Router("/login", &controllers.MasterController{})
    //     // }else{
    //         // ctx.Redirect(302, "/nodes")            
    //         // logs.Notice(ctx.Input.Header("token"))
    //     }
        
    //     // if strings.HasPrefix(ctx.Input.URL(), "/login") {
    //     //     return
    //     // }
        
    //     // _, ok := ctx.Input.Session("uid").(int)
    //     // if !ok {
    //     //     ctx.Redirect(302, "/login")
    //     // }
    // }
    // beego.InsertFilter("/*", beego.BeforeRouter, TokenValidation)
    // beego.InsertFilter("^(/login)$", beego.BeforeRouter, TokenValidation)
    beego.Run()
}

func OperativeSystemValues()(values map[string]string){
    if (runtime.GOOS == "linux"){
        logs.Info("============"+runtime.GOOS+"============")
        var OSmap = make(map[string]string)
        file, err := os.Open("/etc/os-release")
        if err != nil {logs.Error("No os-release file")}
        defer file.Close()
        
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            if (scanner.Text() != ""){
                sidsSplit := strings.Split(scanner.Text(), "=")
                str := strings.Replace(sidsSplit[1], "\"", "", -1)
                OSmap[sidsSplit[0]] = str
            }            
        }
        return OSmap
    }else{
        return nil
    }
}