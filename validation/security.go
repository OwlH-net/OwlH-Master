package validation

import (
	"github.com/astaxie/beego/logs"
	// jwt "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "errors"
	"owlhmaster/database"
	// "strings"
    // "owlhmaster/utils"
)

func UserPermissionsValidation(uuidUser string, permissionRequest string) (val bool, err error) {
	allRelations, err := ndb.GetUserGroupRoles(); if err != nil {logs.Error("UserPermissionsValidation error getting permissions: %s",err); return false, err}
	rolePerm, err := ndb.GetRolePermissions(); if err != nil {logs.Error("UserPermissionsValidation error getting user rolePermissions: %s",err); return false, err}
	// allPerm, err := ndb.GetPermissions(); if err != nil {logs.Error("UserPermissionsValidation error getting user GetPermissions: %s",err); return false, err}

	// //check if permission exists
	// permExists := false
	// for x := range allPerm{
	// 	if x == permissionRequest {permExists = true}
	// }
	// if !permExists {logs.Error("Permissions validation error - This permission don't exists"); return false,err}

	for x := range allRelations{
		if allRelations[x]["user"] == uuidUser{
			//Check if user role has admin permissions
			for w := range rolePerm{
				if allRelations[x]["role"] == rolePerm[w]["role"] {
					if rolePerm[w]["permission"] == "admin" {
						return true, nil
					}					
				}
			}
			//Check if user role has requested permissions
			for w := range rolePerm{
				if allRelations[x]["role"] == rolePerm[w]["role"] {
					if rolePerm[w]["permission"] == permissionRequest {
						return true, nil
					}
				}
			}

			//Compare with role permissions for groups
			for y := range allRelations{
				if allRelations[x]["group"] == allRelations[y]["group"]{
					for w := range rolePerm{
						if allRelations[y]["role"] == rolePerm[w]["role"] {
							if rolePerm[w]["permission"] == permissionRequest {
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