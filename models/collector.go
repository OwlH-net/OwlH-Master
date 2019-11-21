package models 

import (
    "owlhmaster/collector"
)

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/play/:uuid \
// }
func PlayCollector(uuid string) (err error) {
    err = collector.PlayCollector(uuid)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/stop/:uuid \
// }
func StopCollector(uuid string) (err error) {
    err = collector.StopCollector(uuid)
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/show/:uuid \
// }
func ShowCollector(uuid string) (data string, err error) {
    data, err = collector.ShowCollector(uuid)
    return data, err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/playMasterCollector \
// }
func PlayMasterCollector() (err error) {
    err = collector.PlayMasterCollector()
    return err
}

// curl -X PUT \
//   https://52.47.197.22:50002/v1/node/stopMasterCollector \
// }
func StopMasterCollector() (err error) {
    err = collector.StopMasterCollector()
    return err
}

// curl -X GET \
//   https://52.47.197.22:50002/v1/node/showMasterCollector \
// }
func ShowMasterCollector() (data string, err error) {
    data, err = collector.ShowMasterCollector()
    return data, err
}