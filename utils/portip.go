package utils

import (
//     "github.com/astaxie/beego/logs"
//     "strings"
// //    "database/sql"
//     // "fmt"
// //   "time"
// //    _ "github.com/mattn/go-sqlite3"
//     "owlhmaster/database"
//     "errors"
//     "owlhmaster/nodeclient"
//     "owlhmaster/ruleset"
//     "owlhmaster/utils"
//     "regexp"
//     "io/ioutil"
//     // "bufio"
//     // "os"
//     //"io"
//     "net/http"
//     // "net/url"
//     // "strconv"
//     "crypto/tls"
//     "bytes"
//     "encoding/json"
)

func obtainPortIp()()  {
	// //Take IP from specific uuid
	// sqlIP := "select node_value from nodes where node_param = 'ip' and node_uniqueid = '"+loadFile["uuid"]+"';"
	// logs.Info("Datos SQL IP --> "+sqlIP)
	// ip, err := ndb.Db.Query(sqlIP)
	// if err != nil {
	// 	logs.Error("Error al ejecutar la query UUID: %s", err.Error())
	// 	return voidArray, err
	// }
	// defer ip.Close()
	// if ip.Next() {
	// 	ip.Scan(&ipData)
	// }
	// logs.Info("Datos IP --> "+ipData)

	// //Take PORT from specific uuid
	// sqlPORT := "select node_value from nodes where node_param = 'port' and node_uniqueid = '"+loadFile["uuid"]+"';"
	// logs.Info("Datos SQL PORT --> "+sqlPORT)
	// port, err := ndb.Db.Query(sqlPORT)
	// if err != nil {
	// 	logs.Error("Error al ejecutar la query UUID: %s", err.Error())
	// 	return voidArray, err
	// }
	// defer port.Close()
	// if port.Next() {
	// 	if err = port.Scan(&portData); err != nil {
	// 		return voidArray, err
	// 	}
	// }
	// logs.Info("Datos PORT --> "+portData)
	// values["file"] = loadFile["file"]
}