package ndb

import (
	"database/sql"
	"os"
	"owlhmaster/utils"

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
	RSdb, err = sql.Open(cmd, path+"?cache=shared&mode=rwc&_busy_timeout=5000")
	if err != nil {
		logs.Error("database/RulesetSource -- SQL openning Error: " + err.Error())
	}
	logs.Info("database/RulesetSource -- DB -> sql.Open, DB Ready")
}
