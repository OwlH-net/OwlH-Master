package main

import (

    "github.com/astaxie/beego/logs"
    // "github.com/astaxie/beego/context"
    _ "owlhmaster/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "owlhmaster/database"
    "owlhmaster/dispatcher"
    "owlhmaster/node"
    "owlhmaster/master"
    "owlhmaster/search"
    "owlhmaster/scheduler"
    "owlhmaster/utils"
    "owlhmaster/ruleset"
    "owlhmaster/configuration"
    "os"
    "crypto/tls"
    "bufio"
    "strings"
    "runtime"
    "os/signal"
    "syscall"
)

func main() {
    
    //Application version
    logs.Info("Version OwlH Master: 0.12.0.20200427")
    utils.Load()

    //get logger data
    filename, err := utils.GetKeyValueString("logs", "filename")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    maxlines, err := utils.GetKeyValueString("logs", "maxlines")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    maxsize, err := utils.GetKeyValueString("logs", "maxsize")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    daily, err := utils.GetKeyValueString("logs", "daily")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    maxdays, err := utils.GetKeyValueString("logs", "maxdays")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    rotate, err := utils.GetKeyValueString("logs", "rotate")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    level, err := utils.GetKeyValueString("logs", "level")
    if err != nil {logs.Error("Main Error getting data from main.conf for load Logger data: "+err.Error())}
    
    logs.NewLogger(10000)
    logs.SetLogger(logs.AdapterFile,`{"filename":"`+filename+`", "maxlines":`+maxlines+` ,"maxsize":`+maxsize+`, "daily":`+daily+`, "maxdays":`+maxdays+`, "rotate":`+rotate+`, "level":`+level+`}`)


    //operative system values
    data:=OperativeSystemValues()
    for x := range data {
        if (x == "ID" || x == "ID_LIKE" || x == "VERSION_ID"){
            logs.Info(x +" -- "+data[x])
        }
    }

    //check database values
    cancontinue := configuration.MainCheck()
    if !cancontinue {
        logs.Error("can't continue, see previous logs")
        // return 
    }

    //Init database connection
    ndb.Conn()
    ndb.RConn()
    // ndb.GConn()    
    ndb.MConn()    
    

    //CheckServicesStatus
    go ManageSignals()  
    master.CheckServicesStatus()
    ruleset.Init()
    //Init dispatcher at master
    go dispatcher.Init()
    //Init scheduler at master
    go scheduler.Init()
    //Load all rulesets
    go search.Init()
    //Synchronize users to every node
    go node.SyncAllUserData()

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

func ManageSignals() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

    go func() {
        sig := <-sigs
        logs.Info("Signal received: "+ sig.String())

        //kill plugins
        master.StopPluginsGracefully()

        //stop node
        logs.Critical("Stopping Node...")
        os.Exit(0)
    }()
}