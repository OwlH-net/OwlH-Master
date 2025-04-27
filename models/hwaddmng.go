package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  "github.com/OwlH-net/OwlH-Master/hwaddmng"
)

func AddMacIp(id string, data map[string]string, username string) (err error) {
  err = hwaddmng.AddMacIp(id, data)
  changecontrol.ChangeControlInsertData(err, "AddMacIp", username)
  return err
}

func LoadConfig(id string, anode map[string]string, username string) (data map[string]string, err error) {
  data, err = hwaddmng.LoadConfig(id, anode)
  changecontrol.ChangeControlInsertData(err, "LoadConfig", username)
  return data, err
}

func Config(id string, anodeIface map[string]interface{}, username string) (err error) {
  err = hwaddmng.Config(id, anodeIface)
  changecontrol.ChangeControlInsertData(err, "ConfigPost", username)
  return err
}

func Db(id string, anodeIface map[string]string, username string) (err error) {
  err = hwaddmng.Db(id, anodeIface)
  changecontrol.ChangeControlInsertData(err, "Db", username)
  return err
}

func ConfigGet(id string, username string) (data map[string]string, err error) {
  data, err = hwaddmng.ConfigGet(id)
  changecontrol.ChangeControlInsertData(err, "ConfigGet", username)
  return data, err
}
