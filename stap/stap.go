package stap

import(
    //"io/ioutil"
    // "fmt"
	"github.com/astaxie/beego/logs"
	"bytes"
	"encoding/json"
	"net/http"
	// "owlhmaster/node"
    // "bufio" //read line by line the doc
    // "regexp"
    // "os"
    // "os/exec"
    //"strconv"
	"owlhmaster/utils"
	"crypto/tls"
    // "owlhmaster/database"
    // "errors"
    // "database/sql"
    // "strings"
    // "time"
    // "strconv"
)


func AddServer(data map[string]string)(err error) {
	
	logs.Info("stap/stap.go stap AddServer()")

	uuid := data["uuid"]
	ipuuid,portuuid,err := utils.ObtainPortIp(uuid)

	logs.Info(ipuuid+"  *  *  *  "+portuuid)

	url := "https://"+ipuuid+":"+portuuid+"/node/stap/"

	valuesJSON,err := json.Marshal(data)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(valuesJSON))
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
    resp, err := client.Do(req)

    logs.Info("Request newBuffer(JSON) -------> ",req.Body)

    if err != nil {
        return err
	}
	
    defer resp.Body.Close()

    return nil
}