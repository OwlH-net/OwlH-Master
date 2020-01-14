package utils

import (
	"crypto/hmac"
    "crypto/sha256"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
    // "errors"
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
// 	type Header struct {
// 		Alg string `json:"alg"`
// 		Typ string `json:"typ"`
// 	} 
// 	header := Header{
// 		Alg: "HS256",
// 		Typ: "JWT",
// 	}
// 	str, _ := json.Marshal(header)
// 	header = Base64Encode(string(str))
// 	encodedPayload, _ := json.Marshal(payload)
// 	signatureValue := header + "." + 
// 	Base64Encode(string(encodedPayload))
// 	return signatureValue + "." + Hash(signatureValue, secret)
// }

func Decode(jwtToken string, secret string) (payload interface{}, err error) {
	// check if the jwt token contains header, payload and token
	token := strings.Split(jwtToken, ".")
	if len(token) != 3 {err := errors.New("Invalid token: token should contain header, payload and secret"); return nil, err}
	logs.Info(token[0])
	logs.Info(token[1])
	logs.Info(token[2])

   	// decode payload
	decodedPayload, err := Base64Decode(token[1])
	if err != nil {return nil, err}

	// parses payload from string to a struct
   	err = json.Unmarshal([]byte(decodedPayload), &payload)
	if err != nil {return nil, err}

	// checks if the token has expired.
	// if payload.Exp != 0 && time.Now().Unix() > payload.Exp {return nil, errors.New("Expired token: token has expired")}
	// signatureValue := token[0] + "." + token[1]

	// verifies if the header and signature is exactly whats in
	// the signature
	// if CompareHmac(token[0]+"."+token[1], token[2], secret) == false {return nil, errors.New("Invalid JWT signature")}

	// err = jwt.Verify(token[0]+"."+token[1], token[2], secret)
	// logs.Error(err)

	// // Parse the token
	// token, err := jwt.ParseWithClaims(jwtToken, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
	// 	// since we only use the one private key to sign the tokens,
	// 	// we also only use its public counter part to verify
	// 	return verifyKey, nil
	// })
	// if err != nil {logs.Error(err)}
	// logs.Debug(token)

	return payload, nil
}

func CreateToken(tok string)(token string ,err error){
	token, err := jwt.createToken("foo")
	return token, err
}