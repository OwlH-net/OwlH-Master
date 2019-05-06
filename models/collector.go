package models 

import (
    "owlhmaster/collector"
)

func PlayCollector(uuid string) (err error) {
    err = collector.PlayCollector(uuid)
    return err
}
func StopCollector(uuid string) (err error) {
	err = collector.StopCollector(uuid)
    return err
}
func ShowCollector(uuid string) (data string, err error) {
    data, err = collector.ShowCollector(uuid)
    return data, err
}