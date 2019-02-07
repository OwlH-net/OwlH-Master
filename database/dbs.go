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
    logs.Info("DB -> sql.Open, let's try to be Ready")
    var err error
    Db, err = sql.Open("sqlite3", "database/node.db")
    if err != nil {
        panic("DB Open Failed ")
    }
    logs.Info("DB -> sql.Open, DB Ready") 
}

func Insert_user() {
    logs.Info("DB -> Insert User")
    if Db != nil {
        stmt, err := Db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
        if err != nil {
            panic("DB Prepare failed")
        }
        logs.Info("DB -> db.Prepare, err = %s", err)

        res, err := stmt.Exec("astaxie2", "dep test2", "2012-12-09")
        logs.Info("DB -> info", res)
        if err != nil {
            panic("DB bad Query ")
        }
    } else {
        panic("DB -> there is no DB")
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
//        rows, err := Db.Query("SELECT * FROM master WHERE master_id=1;")
        row := Db.QueryRow("SELECT * FROM master WHERE master_id=1;")
//  for rows.Next() {
//    var id int
//    var firstName string
//    err = rows.Scan(&id, &firstName)
//    if err != nil {
//      // handle this error
//      panic(err)
//    }
//    fmt.Println(id, firstName)
//  }
        logs.Info ("DB -> Row %s", row)
        err = row.Scan(&id, &name, &ip, &port)
        if err == sql.ErrNoRows {
            logs.Warn("DB -> No encuentro na, el SQL no ha devuelto rows")
        }
        if err != nil {
            logs.Warn("DB -> no hemos leido bien los campos de scan")
        }
        logs.Info("DB -> Master : id - %s, name - %s, ip - %s, port - %d", id,name,ip,port)
    } else {
        panic("DB -> there is no DB")
    }
    return nil
}



func Insert_master(mname string, mip string, mport int) (err error) {
    logs.Info("DB -> Insert Master")
    logs.Info("DB -> name - %s, ip - %s, port - %d", mname, mip, mport)

    if Db != nil {
        err = Db.Ping()
        if err != nil {
            logs.Info ("DB -> no hay conexion a la DB")
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