package ndb

import (
    "github.com/astaxie/beego/logs"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllScheduler()(data map[string]map[string]string, err error){
	var allScheduleDetails = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
	sql := "select scheduler_uniqueid,scheduler_param, scheduler_value from scheduler;"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query GetAllScheduler Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetAllScheduler rows.Scan: %s", err.Error())
			return nil, err
		}
		if allScheduleDetails[uniqid] == nil { allScheduleDetails[uniqid] = map[string]string{}}
		allScheduleDetails[uniqid][param]=value
	} 
	return allScheduleDetails, nil
}

func GetSchedulerByUniqueid(uuid string)(data map[string]map[string]string, err error){
	var allScheduleDetails = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
	sql := "select scheduler_uniqueid,scheduler_param, scheduler_value from scheduler where scheduler_uniqueid='"+uuid+"';"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query GetScheduleByUniqueid Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetScheduleByUniqueid rows.Scan: %s", err.Error())
			return nil, err
		}
		if allScheduleDetails[uniqid] == nil { allScheduleDetails[uniqid] = map[string]string{}}
		allScheduleDetails[uniqid][param]=value
	} 
	return allScheduleDetails, nil
}

func GetSchedulerByValue(uuid string)(data string, err error){
    var value string
	sql := "select scheduler_uniqueid from scheduler where scheduler_value='"+uuid+"';"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query GetSchedulerByValue Error : %s", err.Error())
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&value); err != nil {
			logs.Error("GetSchedulerByValue rows.Scan: %s", err.Error())
			return "", err
		}
	} 
	return value, nil
}

func InsertScheduler(uuid string, key string, value string) (err error) {
    if Rdb == nil {
        logs.Error("no access to database")
        return errors.New("no access to database")
    }
    stmt, err := Rdb.Prepare("insert into scheduler (scheduler_uniqueid, scheduler_param, scheduler_value) values (?,?,?);")
    if err != nil {
        logs.Error("InsertRulesetSchedule Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&uuid, &key, &value)
    if err != nil {
        logs.Error("InsertRulesetSchedule Execute -> %s", err.Error())
        return err
    }
    return nil
}

func UpdateScheduler(uuid string, param string, value string)(err error){
	update,err := Rdb.Prepare("update scheduler set scheduler_value = ? where scheduler_uniqueid = ? and scheduler_param = ?;")
	if (err != nil){
		logs.Error("UpdateSchedule prepare UPDATE error -- "+err.Error())
		return err
	}
	_, err = update.Exec(&value, &uuid, &param)
	defer update.Close()
	if (err != nil){
		logs.Error("UpdateSchedule UPDATE error-- "+err.Error())
		return err
	}
	return nil
}