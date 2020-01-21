package validation

import (
	// "crypto/hmac"
    // "crypto/sha256"
	// "encoding/base64"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	// "encoding/json"
	"errors"
	// "strings"
    "owlhmaster/database"
)
// //Enconde to Base64
// func Base64Encode(src string) string {
//     return strings.
//         TrimRight(base64.URLEncoding.
//             EncodeToString([]byte(src)), "=")
// }

// // Base64Encode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
// func Base64Decode(src string) (string, error) {
//     decoded, err := base64.URLEncoding.DecodeString(src)
//     if err != nil {logs.Error("Decoding Error: %s", err); return "", err}
//     return string(decoded), nil
// }

// // Hash generates a Hmac256 hash of a string using a secret
// func Hash(src string, secret string) string {
// 	key := []byte(secret)
// 	h := hmac.New(sha256.New, key)
// 	h.Write([]byte(src))
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }

// // isValidHash validates a hash againt a value
// func isValidHash(value string, hash string, secret string) bool {
// 	return hash == Hash(value, secret)
// }

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

// func Decode(jwtToken string, secret string) (err error) {
// 	token := strings.Split(jwtToken, ".")
// 	if len(token) != 3 {err := errors.New("Invalid token: token should contain header, payload and secret"); return err}
	
// 	jwtKey := []byte(secret) 
// 	tkn, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	logs.Notice(jwtToken)
// 	logs.Debug(tkn)
// 	logs.Error(err)
// 	if err != nil {err := errors.New("Invalid token"); return err}


// 	return nil
// }

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    logs.Info("NEW HASH PASSWD--> "+string(bytes))
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))	
	if err != nil {logs.Error(err); return false, err}
    return true, nil
}

func CheckToken(token string, user string, uuid string)(err error){
	users,err := ndb.GetLoginData()
	for x := range users{
		if (x == uuid) && (users[x]["user"] == user){
			// valid,err := Encode(uuid, user, users[x]["secret"])
			tkn, err := Encode(uuid, user, users[x]["secret"])
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