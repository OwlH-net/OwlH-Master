package elk

import (
    //elk7 "github.com/elastic/go-elasticsearch"
    //"github.com/elastic/go-elasticsearch/esapi"
    "github.com/astaxie/beego/logs"    
    "time"
    // "encoding/json"
    // "context"
    // "strings"
    "owlhmaster/utils"
    "strconv"
)

func Init(jsonRules []string)(){
    cont := 0
    for x := range jsonRules{
        WriteToELK(jsonRules[x],cont)
        cont++
    }
}

func WriteToELK(t string, cont int){

    // loadElk := map[string]map[string]string{}
    // loadElk["elk"] = map[string]string{}
    // loadElk["elk"]["ip"] = ""
    // loadElk["elk"]["port"] = ""
    // loadElk["elk"]["indexName"] = ""
    // loadElk["elk"]["status"] = ""
    // loadElk,err := utils.GetConf(loadElk)
    // if err != nil {logs.Error("Elastic/WriteToELK error readding data from main.conf: "+err.Error()); return}
    // ip := loadElk["elk"]["ip"]
    // port := loadElk["elk"]["port"]
    // indexName := loadElk["elk"]["indexName"]
    // status := loadElk["elk"]["status"]

    // if status == "disabled" {return}
    return 

    // cfg := elk7.Config{
    //     Addresses: []string{
    //         "http://"+ip+":"+port+"",
    //     },
    // }
    // es, err := elk7.NewClient(cfg)
    // if err != nil {logs.Error("Elastic search error: "+err.Error()); return}

    // // logs.Info(t)
    // req := esapi.IndexRequest{
    //     Index:      indexName,
    //     DocumentID: "item"+string(cont),
    //     Body:       strings.NewReader(t),
    //     Refresh:    "true",
    // }

    // // Perform the request with the client.
    // res, err := req.Do(context.Background(), es)
    // if err != nil {
    //     logs.Error("Error getting response: %s", err)
    //     logs.Error("RULE: %s", t)
    //     return
    // }    
    // defer res.Body.Close()

    // if res.IsError() {
    //     logs.Error("[%s] Error indexing document ID=%d", res.Status(), 5)
    // } else {
    //     // Deserialize the response into a map.
    //     var r map[string]interface{}
    //     if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    //           logs.Error("Error parsing the response body: %s", err)
    //     } else {
    //           // Print the response status and indexed document version.
    //         //   logs.Info("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
    //     }
    // }
}

func StartElasticSync()(){
    for {
        t,err := utils.GetKeyValueString("loop", "elastic")
        if err != nil {logs.Error("Search Error: Cannot load Elastic information.")}
        tDuration, err := strconv.Atoi(t)
        
        time.Sleep(time.Minute * time.Duration(tDuration))
    }
}