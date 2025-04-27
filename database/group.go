package ndb

import (
  "database/sql"
  "os"

  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
  _ "github.com/mattn/go-sqlite3"
)

var (
  Gdb *sql.DB
)

func GConn() {
  var err error
  path, err := utils.GetKeyValueString("groupConn", "path")
  if err != nil {
    logs.Error("GConn Error getting data from main.conf at master: " + err.Error())
  }
  cmd, err := utils.GetKeyValueString("groupConn", "cmd")
  if err != nil {
    logs.Error("GConn Error getting data from main.conf at master: " + err.Error())
  }

  _, err = os.Stat(path)
  if err != nil {
    panic("database/Group DB -- DB Open Failed: " + err.Error())
  }
  Gdb, err = sql.Open(cmd, path)
  if err != nil {
    logs.Error("database/Group -- SQL openning Error: " + err.Error())
  }
  logs.Info("database/Group -- DB -> sql.Open, DB Ready")
}
