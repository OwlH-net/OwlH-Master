package validation

import (
	"github.com/astaxie/beego/logs"
	// jwt "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "errors"
	"owlhmaster/database"
	"strings"
    // "owlhmaster/utils"
)

// Check user privileges
func UserPrivilegeValidation(uuid string, privilege string) (val bool, err error) {
	logs.Debug(privilege)
		allPrivileges, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting privileges: %s",err); return false, err}
		for x := range allPrivileges{
			if allPrivileges[x]["user"] == uuid{
				//Compare with group privileges
				groups, err := ndb.GetUserGroups(); if err != nil {logs.Error("UserPrivilegeValidation error getting user groups: %s",err); return false, err}
				allPrivilegesGroups := strings.Split(groups[allPrivileges[x]["group"]]["privileges"], ",")
				for p := range allPrivilegesGroups{
					logs.Warn("Group: "+allPrivilegesGroups[p])
					if allPrivilegesGroups[p] == privilege{
						return true, nil
					}
				}
				//Compare with role privileges
				roles, err := ndb.GetUserRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting user roles: %s",err); return false, err}
				allPrivilegesRoles := strings.Split(roles[allPrivileges[x]["role"]]["privileges"], ",")
				for r := range allPrivilegesRoles{
					logs.Notice("Role: "+allPrivilegesRoles[r])
					if allPrivilegesRoles[r] == privilege{
						return true, nil
					}
				}
			}
		}
		// allPrivileges, err := ndb.GetUserPrivileges(); if err != nil {logs.Error("UserPrivilegeValidation error getting privileges: %s",err); return false, err}
		// hasPrivileges := false
		// for x := range allPrivileges{
		// 	if allPrivileges[x]["user"] == uuid {
		// 		if allPrivileges[x]["privilege"] == "/"{
		// 			return true, nil
		// 		}
		// 		if allPrivileges[x]["privilege"] == privilege{
		// 			hasPrivileges = true
		// 		}
		// 	}
		// }
		// if hasPrivileges{
		// 	return true, nil
		// }else{
		// 	return false, nil
		// }

	return false, nil
}