package models 

import (
    "owlhmaster/changeControl"
)

// curl -X GET \
//   https://52.47.197.22:50001/v1/changecontrol/ \
// }
func GetChangeControl(username string)(data map[string]map[string]string, err error) {
    data, err = changecontrol.GetChangeControl()
    changecontrol.ChangeControlInsertData(err, "GetChangeControl", username)
    return data, err
}