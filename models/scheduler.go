package models 

import (
    "owlhmaster/scheduler"
)

func SchedulerTask(data map[string]string) (err error) {
    err = scheduler.SchedulerTask(data)
    return err
}

func StopTask(data map[string]string) (err error) {
    err = scheduler.StopTask(data)
    return err
}

func GetLog(uuid string) (logReg map[string]map[string]string ,err error) {
    logReg,err = scheduler.GetLog(uuid)
    return logReg,err
}