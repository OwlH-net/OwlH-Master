package elastic

import (
    "github.com/astaxie/beego/logs"
    "io/ioutil"
	"encoding/json"
	"errors"
	"owlhmaster/utils"
	"bytes"
)

func Init()(){
	go startElasticSync()
}

func startElasticSync()(){
	for {
		
		time.Sleep(30 * time.Minute)
	}
}