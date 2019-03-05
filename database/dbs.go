package ndb

import (
    "github.com/astaxie/beego/logs"
    "database/sql"
    "strconv"
//    "fmt"
//   "time"
    _ "github.com/mattn/go-sqlite3"
)

var (
    Db *sql.DB
)

type Master struct {
    Id       string
    Name     string
    Ip       string
    Port     int
}

func init() {
    logs.Info ("DB Init ()")
}

func Conn() {
    var err error
    Db, err = sql.Open("sqlite3", "database/node.db")
    if err != nil {
        panic("DB Open Failed ")
    }
}


func Get_master() (err error) {
    logs.Info("DB -> Get Master")
    var (
        id string
        name string
        ip string
        port int
    )

    if Db != nil {
        row := Db.QueryRow("SELECT * FROM master WHERE master_id=1;")
        logs.Info ("DB -> Row %s", row)
        err = row.Scan(&id, &name, &ip, &port)
        if err == sql.ErrNoRows {
            logs.Warn("DB -> No Rows found")
        }
        if err != nil {
            logs.Warn("DB -> No results after Scan")
        }
        logs.Info("DB -> Master : id - %s, name - %s, ip - %s, port - %d", id,name,ip,port)
    } else {
        panic("DB -> there is no DB")
    }
    return nil
}



func Insert_master(mname string, mip string, mport int) (err error) {

    if Db != nil {
        err = Db.Ping()
        if err != nil {
            logs.Info ("DB -> No access to batabase")
            return err
        }
        stmt, err := Db.Prepare("INSERT INTO master(master_name, master_ip, master_port) values(?,?,?)")
        if err != nil {
            panic("DB Prepare failed")
        }
        logs.Info("DB -> db.Prepare, err = %s", err)

        res, err := stmt.Exec(mname, mip, strconv.Itoa(mport))
        logs.Info("DB -> info", res)
        if err != nil {
            panic("DB bad Query ")
            return err
        }
    } else {
        panic("DB -> there is no DB")
    }
    return nil
}

func Close() {
    Db.Close()
}