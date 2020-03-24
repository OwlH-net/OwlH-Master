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
func UserPrivilegeValidation(uuidUser string, requestType string) (val bool, err error) {
	allRelations, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting permissions: %s",err); return false, err}
	roles, err := ndb.GetUserRoles(); if err != nil {logs.Error("UserPrivilegeValidation error getting user roles: %s",err); return false, err}
	for x := range allRelations{
		if allRelations[x]["user"] == uuidUser{
			//Compare with role permissions
			allPermissionsRole := strings.Split(roles[allRelations[x]["role"]]["permissions"], ",")
			for r := range allPermissionsRole{
				if allPermissionsRole[r] == requestType{
					return true, nil
				}
			}	

			//Compare with role permissions for groups
			for y := range allRelations{
				if allRelations[x]["group"] == allRelations[y]["group"]{
					allPermissionsRole = strings.Split(roles[allRelations[y]["role"]]["permissions"], ",")
					for r := range allPermissionsRole{
						if allPermissionsRole[r] == requestType{
							return true, nil
						}
					}	
				}
			}
		}
	}

	return false, nil
}

func UserPrivilegeValidation2(uuidUser string, requestType string) (val bool, err error) {
	allRelations, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPrivilegeValidation2 error getting permissions: %s",err); return false, err}
	rolePerm, err := ndb.GetRolePermissions(); if err != nil {logs.Error("UserPrivilegeValidation2 error getting user rolePermissions: %s",err); return false, err}
	for x := range allRelations{
		if allRelations[x]["user"] == uuidUser{
			//Compare with role permissions
			for w := range rolePerm{
				if allRelations[x]["role"] == rolePerm[w]["role"] {
					if rolePerm[w]["permission"] == requestType {
						return true, nil
					}
				}
			}

			//Compare with role permissions for groups
			for y := range allRelations{
				if allRelations[x]["group"] == allRelations[y]["group"]{
					for w := range rolePerm{
						if allRelations[y]["role"] == rolePerm[w]["role"] {
							if rolePerm[w]["permission"] == requestType {
								return true, nil
							}
						}
					}
				}
			}
		}
	}

	return false, nil
}