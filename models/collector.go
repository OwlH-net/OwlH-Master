package models 

import (
    "owlhmaster/collector"
    "owlhmaster/changeControl"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/play/:uuid \
// }
func PlayCollector(uuid string) (err error) {
    err = collector.PlayCollector(uuid)
    changecontrol.ChangeControlInsertData(err, "PlayCollector")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/stop/:uuid \
// }
func StopCollector(uuid string) (err error) {
    err = collector.StopCollector(uuid)
    changecontrol.ChangeControlInsertData(err, "StopCollector")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/show/:uuid \
// }
func ShowCollector(uuid string) (data string, err error) {
    data, err = collector.ShowCollector(uuid)
    changecontrol.ChangeControlInsertData(err, "ShowCollector")
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/playMasterCollector \
// }
func PlayMasterCollector() (err error) {
    err = collector.PlayMasterCollector()
    changecontrol.ChangeControlInsertData(err, "PlayMasterCollector")
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/stopMasterCollector \
// }
func StopMasterCollector() (err error) {
    err = collector.StopMasterCollector()
    changecontrol.ChangeControlInsertData(err, "StopMasterCollector")
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/showMasterCollector \
// }
func ShowMasterCollector() (data string, err error) {
    data, err = collector.ShowMasterCollector()
    changecontrol.ChangeControlInsertData(err, "ShowMasterCollector")
    return data, err
}