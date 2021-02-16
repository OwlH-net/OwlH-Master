package validation

import (
	"github.com/astaxie/beego/logs"

	"strings"
)

// UserPermissionsValidation - Validate user request
func UserPermissionsValidation(user string, permissionRequest string) (val bool, err error) {
	uuidUser, err := ndb.GetUserID(user)
	allRelations, err := ndb.GetUserGroupRoles()
	if err != nil {
		logs.Error("UserPermissionsValidation error getting permissions: %s", err)
		return false, err
	}
	rolePerm, err := ndb.GetRolePermissions()
	if err != nil {
		logs.Error("UserPermissionsValidation error getting user rolePermissions: %s", err)
		return false, err
	}
	allPerm, err := ndb.GetPermissions()
	if err != nil {
		logs.Error("UserPermissionsValidation error getting user GetPermissions: %s", err)
		return false, err
	}

	//check if permission exists
	permExists := false
	for x := range allPerm {
		if x == permissionRequest {
			permExists = true
		}
	}
	if !permExists {
		logs.Error("Permissions validation error - This permission don't exists")
		return false, err
	}

	for x := range allRelations {
		if allRelations[x]["user"] == uuidUser {

			//Check role permissions
			for w := range rolePerm {
				if allRelations[x]["role"] == rolePerm[w]["role"] && allRelations[x]["role"] != "" {
					//split permissions
					permissionsList := strings.Split(rolePerm[w]["permissions"], ",")
					//check if user is admin
					for p := range permissionsList {
						if permissionsList[p] == "admin" {
							return true, nil
						}
					}
					//check for permissions requested
					for s := range permissionsList {
						if permissionsList[s] == permissionRequest {
							return true, nil
						}
					}
				}
			}

			//check group roles for user groups
			for s := range allRelations {
				if allRelations[s]["group"] == allRelations[x]["group"] && allRelations[x]["group"] != "" {
					for roleID := range rolePerm {
						if rolePerm[roleID]["role"] == allRelations[s]["role"] {
							//split permissions
							permissionsList := strings.Split(rolePerm[roleID]["permissions"], ",")
							//check if user is admin
							for p := range permissionsList {
								if permissionsList[p] == "admin" {
									return true, nil
								}
							}
							//check for permissions requested
							for s := range permissionsList {
								if permissionsList[s] == permissionRequest {
									return true, nil
								}
							}
						}
					}
				}
			}
		}
	}

	logs.Error("Permission DENIED!!!!!!")

	return false, nil
}
