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
		allPermissions, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting permissions: %s",err); return false, err}
		for x := range allPermissions{
			if allPermissions[x]["user"] == uuid{
				//Compare with role permissions
				roles, err := ndb.GetUserRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting user roles: %s",err); return false, err}
				allPermissionsRole := strings.Split(roles[allPermissions[x]["role"]]["permissions"], ",")
				for r := range allPermissionsRole{
					if allPermissionsRole[r] == permission{
						return true, nil
					}
				}
				//Compare with group permissions
				groups, err := ndb.GetUserGroups(); if err != nil {logs.Error("UserPrivilegeValidation error getting user groups: %s",err); return false, err}
				allPermissionsGroups := strings.Split(groups[allPermissions[x]["group"]]["permissions"], ",")
				for g := range allPermissionsGroups{
					if allPermissionsGroups[g] == permission{
						return true, nil
					}
				}
			}
		}

	return false, nil
}