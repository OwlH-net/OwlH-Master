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