package validation

import (
	"github.com/astaxie/beego/logs"
	// jwt "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "errors"
    "owlhmaster/database"
    // "owlhmaster/utils"
)

// Check user privileges
func UserPrivilegeValidation(uuid string, privilege string) (val bool, err error) {
	// users, err := ndb.GetLoginData(); if err != nil {logs.Error("UserPrivilegeValidation error getting users: %s",err); return false, err}
	// if users[uuid]["privilege"] == "/"{
	// 	return true, nil
	// }else{
		allPrivileges, err := ndb.GetUserPrivileges(); if err != nil {logs.Error("UserPrivilegeValidation error getting privileges: %s",err); return false, err}
		hasPrivileges := false
		for x := range allPrivileges{
			if allPrivileges[x]["user"] == uuid {
				if allPrivileges[x]["privilege"] == "/"{
					return true, nil
				}
				if allPrivileges[x]["privilege"] == privilege{
					hasPrivileges = true
				}
			}
		}
		if hasPrivileges{
			return true, nil
		}else{
			return false, nil
		}
	// }

	return false, nil
}