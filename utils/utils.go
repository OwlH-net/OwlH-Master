package utils

import (
    //"owlhnode/models"
    "encoding/json"
    // "strconv"
    //"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
    // "io"
    // "strings"
    // "os"
    // "time"
    // "os/exec"
    // "fmt"
    // "crypto/rand"
)

//Read map data
//leer json del fichero para obtener el path del bpf
func GetConf(loadData map[string]map[string]string)(loadDataReturn map[string]map[string]string, err error) { 
    confFilePath := "/etc/owlh/conf/main.conf"
    jsonPathBpf, err := ioutil.ReadFile(confFilePath)
    if err != nil {
        logs.Error("utils/GetConf -> can't open Conf file -> " + confFilePath)
        return nil, err
    }

    var anode map[string]map[string]string
    json.Unmarshal(jsonPathBpf, &anode)

    for k,y := range loadData { 
        for y,_ := range y {
            if v, ok := anode[k][y]; ok {
                loadData[k][y] = v
            }else{
                loadData[k][y] = "None"
            }
        }
    }
    
    return loadData, nil
}