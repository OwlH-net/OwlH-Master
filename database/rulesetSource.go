package ndb

import (
  "database/sql"
  "os"

  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
  _ "github.com/mattn/go-sqlite3"
)

var (
  RSdb *sql.DB
)

func RSConn() {
  var err error
  path, err := utils.GetKeyValueString("rulesetSourceConn", "path")
  if err != nil {
    logs.Error("SConn Error getting data from main.conf at master: " + err.Error())
  }
  cmd, err := utils.GetKeyValueString("rulesetSourceConn", "cmd")
  if err != nil {
    logs.Error("SConn Error getting data from main.conf at master: " + err.Error())
  }

  _, err = os.Stat(path)
  if err != nil {
    panic("database/RulesetSource DB -- DB Open Failed: " + err.Error())
  }
  RSdb, err = sql.Open(cmd, path)
  if err != nil {
    logs.Error("database/RulesetSource -- SQL openning Error: " + err.Error())
  }
  logs.Info("database/RulesetSource -- DB -> sql.Open, DB Ready")
}
