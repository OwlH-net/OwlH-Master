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

func GetSchedulerByValue(val string)(data string, err error){
    var value string
	sql := "select scheduler_uniqueid from scheduler where scheduler_value='"+val+"';"
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
        logs.Error("no access to database InsertScheduler")
        return errors.New("no access to database InsertScheduler")
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

func DeleteScheduler(uuid string)(err error){
	if Rdb == nil {
        logs.Error("no access to database DeleteSchedulerLog")
        return errors.New("no access to database DeleteSchedulerLog")
    }
	sourceSQL, err := Rdb.Prepare("delete from scheduler where scheduler_uniqueid = ?")
    if err != nil {
        logs.Error("Error DeleteScheduler Prepare delete a file -> %s", err.Error())
        return err
	}
    _, err = sourceSQL.Exec(&uuid)
    if err != nil {
        logs.Error("Error DeleteScheduler deleting a file -> %s", err.Error())
        return err
	}
	return nil
}


//select by uuid
func GetSchedulerLogByUniqueid(uuid string)(data map[string]map[string]string, err error){
	var allScheduleLogDetails = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
	sql := "select log_uniqueid, log_param, log_value from scheduler_log where log_uniqueid='"+uuid+"';"
	rows, err := Rdb.Query(sql)
	if err != nil {
		logs.Error("Rdb.Query GetSchedulerLogByUniqueid Error : %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&uniqid, &param, &value); err != nil {
			logs.Error("GetSchedulerLogByUniqueid rows.Scan: %s", err.Error())
			return nil, err
		}
		if allScheduleLogDetails[uniqid] == nil { allScheduleLogDetails[uniqid] = map[string]string{}}
		allScheduleLogDetails[uniqid][param]=value
	} 
	return allScheduleLogDetails, nil
}

//select *
func GetSchedulerLogByValue(val string)(data map[string]map[string]string, err error){
	var allScheduleLogDetails = map[string]map[string]string{}
    var uniqid string
    var param string
    var value string
	sql := "select log_uniqueid, log_param, log_value from scheduler_log where log_value='"+val+"';"
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
		if allScheduleLogDetails[uniqid] == nil { allScheduleLogDetails[uniqid] = map[string]string{}}
		allScheduleLogDetails[uniqid][param]=value
	} 
	return allScheduleLogDetails, nil
}

//insert
func InsertSchedulerLog(uuid string, key string, value string) (err error) {
    if Rdb == nil {
        logs.Error("no access to database InsertSchedulerLog")
        return errors.New("no access to database InsertSchedulerLog")
    }
    stmt, err := Rdb.Prepare("insert into scheduler_log (log_uniqueid, log_param, log_value) values (?,?,?);")
    if err != nil {
        logs.Error("InsertSchedulerLog Prepare -> %s", err.Error())
        return err
    }
    _, err = stmt.Exec(&uuid, &key, &value)
    if err != nil {
        logs.Error("InsertSchedulerLog Execute -> %s", err.Error())
        return err
    }
    return nil
}

//delete
func DeleteSchedulerLog(uuid string)(err error){
	if Rdb == nil {
        logs.Error("no access to database DeleteSchedulerLog")
        return errors.New("no access to database DeleteSchedulerLog")
    }
	sourceSQL, err := Rdb.Prepare("delete from scheduler_log where log_uniqueid = ?")
    if err != nil {
        logs.Error("Error DeleteSchedulerLog Prepare delete a file -> %s", err.Error())
        return err
	}
    _, err = sourceSQL.Exec(&uuid)
    if err != nil {
        logs.Error("Error DeleteSchedulerLog deleting a file -> %s", err.Error())
        return err
	}
	return nil
}