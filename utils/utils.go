package utils

import (
    "encoding/json"
    "github.com/astaxie/beego/logs"
    "io/ioutil"
    "io"
	"net/http"
	"crypto/tls"
)

//Read map data
//leer json del fichero para obtener el path del bpf
func GetConf(loadData map[string]map[string]string)(loadDataReturn map[string]map[string]string, err error) { 
    // confFilePath := "/etc/owlh/conf/main.conf"
    confFilePath := "./conf/main.conf"
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



func NewRequestHTTP(order string, url string, values io.Reader)(resp *http.Response, err error){
	req, err := http.NewRequest(order, url, values)
	if err != nil {
		logs.Error("Error Executing HTTP new request")
	}
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},}
    client := &http.Client{Transport: tr}
	resp, err = client.Do(req)
	if err != nil {
		logs.Error("Error Retrieving response from client HTTP new request")
	}
	return resp, err
}