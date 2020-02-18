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

// Check user permissions
func UserPrivilegeValidation(uuid string, permission string) (val bool, err error) {
		allPrivileges, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting permissions: %s",err); return false, err}
		for x := range allPrivileges{
			if allPrivileges[x]["user"] == uuid{
				//Compare with group permissions
				groups, err := ndb.GetUserGroups(); if err != nil {logs.Error("UserPrivilegeValidation error getting user groups: %s",err); return false, err}
				allPrivilegesGroups := strings.Split(groups[allPrivileges[x]["group"]]["permissions"], ",")
				for p := range allPrivilegesGroups{
					if allPrivilegesGroups[p] == permission{
						return true, nil
					}
				}
				//Compare with role permissions
				roles, err := ndb.GetUserRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting user roles: %s",err); return false, err}
				allPrivilegesRoles := strings.Split(roles[allPrivileges[x]["role"]]["permissions"], ",")
				for r := range allPrivilegesRoles{
					if allPrivilegesRoles[r] == permission{
						return true, nil
					}
				}
			}
		}
		// allPrivileges, err := ndb.GetUserPermissions(); if err != nil {logs.Error("UserPrivilegeValidation error getting permissions: %s",err); return false, err}
		// hasPrivileges := false
		// for x := range allPrivileges{
		// 	if allPrivileges[x]["user"] == uuid {
		// 		if allPrivileges[x]["permission"] == "/"{
		// 			return true, nil
		// 		}
		// 		if allPrivileges[x]["permission"] == permission{
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