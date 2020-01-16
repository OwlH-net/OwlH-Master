package utils

import (
	"crypto/hmac"
    "crypto/sha256"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
    "errors"
	"strings"
	"golang.org/x/crypto/bcrypt"
	// "encoding/json"
    "owlhmaster/database"
)
//Enconde to Base64
func Base64Encode(src string) string {
    return strings.
        TrimRight(base64.URLEncoding.
            EncodeToString([]byte(src)), "=")
}

// Base64Encode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
func Base64Decode(src string) (string, error) {
    decoded, err := base64.URLEncoding.DecodeString(src)
    if err != nil {logs.Error("Decoding Error: %s", err); return "", err}
    return string(decoded), nil
}

// Hash generates a Hmac256 hash of a string using a secret
func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// isValidHash validates a hash againt a value
func isValidHash(value string, hash string, secret string) bool {
	return hash == Hash(value, secret)
}

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

func Decode(jwtToken string, secret string) (err error) {
	token := strings.Split(jwtToken, ".")
	if len(token) != 3 {err := errors.New("Invalid token: token should contain header, payload and secret"); return err}
	logs.Info("Pass: "+secret)
	logs.Info(token[0])
	logs.Info(token[1])
	logs.Info(token[2])


	jwtKey := []byte("42isTheAnswer") //get user secret
	tkn, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
	})

	logs.Notice(tkn)
	logs.Error(err)

	return nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    logs.Info("HASH PSSWD--> "+string(bytes))
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
        if users[x]["user"] == data["user"]{
			check, err := utils.CheckPasswordHash(data["password"], users[x]["pass"])
			if err != nil{return err}
            if check{
				return nil
			}
		}
	}
	return errors.New("There are not token. Error creating Token")
}