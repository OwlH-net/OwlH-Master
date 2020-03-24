package validation

import (
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"errors"
    "owlhmaster/database"
    "owlhmaster/utils"
)

// Encode generates a jwt.
func Encode(uuid string, user string, secret string) (val string, err error) {

	type MyCustomClaims struct {
		Uuid string `json:"uuid"`
		User string `json:"user"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		uuid,
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "OwlH",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {logs.Error(err); return "", err}
	return tokenString, err
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))	
	if err != nil {logs.Error(err); return false, err}
    return true, nil
}

func CheckToken(token string, user string, uuid string, requestType string)(hasPrivileges bool, err error){
	users,err := ndb.GetLoginData()
	for x := range users{
		if (x == uuid) && (users[x]["user"] == user){
			tkn, err := Encode(uuid, users[x]["user"], users[x]["secret"])
			if err != nil {
				logs.Error("Error checking token: %s", err); return false,err
			}else{
				if token == tkn {
					status,err := UserPrivilegeValidation(uuid, requestType); if err != nil {logs.Error("requestType error: %s",err); return false,err}
					if status{
						masterID,err := ndb.LoadMasterID(); if err != nil {logs.Error("Error getting Master information: %s",err); return false,err}
						utils.TokenMasterUuid = masterID
						utils.TokenMasterUser = x
						return true,nil
					}else{
						return false,nil
					}
				}else{
					return false,errors.New("The token retrieved is false")
				}
			}
		}
	}
	return false,errors.New("There are not token. Error creating Token")
}

func VerifyToken(token string, user string, uuid string)(err error){
	users,err := ndb.GetLoginData()
	for x := range users{
		if (x == uuid) && (users[x]["user"] == user){
			tkn, err := Encode(uuid, users[x]["user"], users[x]["secret"])
			if err != nil {
				logs.Error("Error checking token: %s", err); return err
			}else{
				if token == tkn {					
					return nil					
				}else{
					return errors.New("The token retrieved is false")
				}
			}
		}
	}
	return errors.New("There are not token. Error creating Token")
}

func VerifyPermissions(uuidUser string, object string, permissions []string)(hasPermissions bool, err error){
	for x := range permissions{
		status,err := UserPrivilegeValidation2(uuidUser, permissions[x]); if err != nil {logs.Error("requestType error: %s",err); return false,err}
		if status{
			masterID,err := ndb.LoadMasterID(); if err != nil {logs.Error("Error getting Master information: %s",err); return false,err}
			utils.TokenMasterUuid = masterID
			utils.TokenMasterUser = uuidUser
			return true,nil
		}		
	}
	return false, nil
}