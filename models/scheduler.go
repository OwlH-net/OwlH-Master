package models 

import (
    "owlhmaster/scheduler"
    "owlhmaster/changeControl"
)

// curl -X PUT \
//   https://52.47.197.22:50002/v1/scheduler/add \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v",
//     "update": "v",
//     "period": "v",
//     "day": "v",
//     "month": "v",
//     "year": "v",
//     "hour": "v",
//     "minute": "v",
//     "status": "v",
//     "type": "v"
// }
func SchedulerTask(data map[string]string) (err error) {
    err = scheduler.SchedulerTask(data)
    changecontrol.ChangeControlInsertData(err, "SchedulerTask")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/scheduler/add \
//   -H 'Content-Type: application/json' \
//   -d '{
//     "uuid": "v"
// }
func StopTask(data map[string]string) (err error) {
    err = scheduler.StopTask(data)
    changecontrol.ChangeControlInsertData(err, "StopTask")
    return err
}

// curl -X POST \
//   https://52.47.197.22:50002/v1/scheduler/log/:uuid \
func GetLog(uuid string) (logReg map[string]map[string]string ,err error) {
    logReg,err = scheduler.GetLog(uuid)
    changecontrol.ChangeControlInsertData(err, "GetLog")
    return logReg,err
}