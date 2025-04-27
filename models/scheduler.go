package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  "github.com/OwlH-net/OwlH-Master/scheduler"
)

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/scheduler/add \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v",
//      "update": "v",
//      "period": "v",
//      "day": "v",
//      "month": "v",
//      "year": "v",
//      "hour": "v",
//      "minute": "v",
//      "status": "v",
//      "type": "v"
//  }
func SchedulerTask(data map[string]string, username string) (err error) {
  err = scheduler.SchedulerTask(data)
  changecontrol.ChangeControlInsertData(err, "SchedulerTask", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/scheduler/add \
//    -H 'Content-Type: application/json' \
//    -d '{
//      "uuid": "v"
//  }
func StopTask(data map[string]string, username string) (err error) {
  err = scheduler.StopTask(data)
  changecontrol.ChangeControlInsertData(err, "StopTask", username)
  return err
}

//  curl -X POST \
//    https://52.47.197.22:50002/v1/scheduler/log/:uuid \
func GetLog(uuid string, username string) (logReg map[string]map[string]string, err error) {
  logReg, err = scheduler.GetLog(uuid)
  changecontrol.ChangeControlInsertData(err, "GetLog", username)
  return logReg, err
}
