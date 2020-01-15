package utils

import (
	"crypto/hmac"
    "crypto/sha256"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"
    "errors"
    "strings"
	// "encoding/json"
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

// // Encode generates a jwt.
// func Encode(payload Payload, secret string) string {
	// // type Header struct {
	// // 	Alg string `json:"alg"`
	// // 	Typ string `json:"typ"`
	// // } 
	// header := Header{
	// 	Alg: "HS256",
	// 	Typ: "JWT",
	// }
	// // str, _ := json.Marshal(header)
	// // header = Base64Encode(string(str))
	// // encodedPayload, _ := json.Marshal(payload)
	// // signatureValue := header + "." + 
	// // Base64Encode(string(encodedPayload))
	// // return signatureValue + "." + Hash(signatureValue, secret)

	// // Create a new token object, specifying signing method and the claims
	// // you would like it to contain.
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })

	// // Sign and get the complete encoded token as a string using the secret
	// tokenString, err := token.SignedString(hmacSampleSecret)
// }

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