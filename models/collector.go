package models

import (
  changecontrol "github.com/OwlH-net/OwlH-Master/changeControl"
  "github.com/OwlH-net/OwlH-Master/collector"
)

//  curl -X GET \
//    https://52.47.197.22:50002/v1/node/play/:uuid \
//  }
func PlayCollector(uuid string, username string) (err error) {
  err = collector.PlayCollector(uuid)
  changecontrol.ChangeControlInsertData(err, "PlayCollector", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/node/stop/:uuid \
//  }
func StopCollector(uuid string, username string) (err error) {
  err = collector.StopCollector(uuid)
  changecontrol.ChangeControlInsertData(err, "StopCollector", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/node/show/:uuid \
//  }
func ShowCollector(uuid string, username string) (data string, err error) {
  data, err = collector.ShowCollector(uuid)
  changecontrol.ChangeControlInsertData(err, "ShowCollector", username)
  return data, err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/node/playMasterCollector \
//  }
func PlayMasterCollector(username string) (err error) {
  err = collector.PlayMasterCollector()
  changecontrol.ChangeControlInsertData(err, "PlayMasterCollector", username)
  return err
}

//  curl -X PUT \
//    https://52.47.197.22:50002/v1/node/stopMasterCollector \
//  }
func StopMasterCollector(username string) (err error) {
  err = collector.StopMasterCollector()
  changecontrol.ChangeControlInsertData(err, "StopMasterCollector", username)
  return err
}

//  curl -X GET \
//    https://52.47.197.22:50002/v1/node/showMasterCollector \
//  }
func ShowMasterCollector(username string) (data string, err error) {
  data, err = collector.ShowMasterCollector()
  changecontrol.ChangeControlInsertData(err, "ShowMasterCollector", username)
  return data, err
}
