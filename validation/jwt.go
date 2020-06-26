package validation

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "github.com/astaxie/beego/logs"
    jwt "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "io"
    "owlhmaster/database"
    "owlhmaster/utils"
    // "encoding/base64"
    // "encoding/hex"
)

// Encode generates a jwt.
func Encode(user string, secret string) (val string, err error) {

    type MyCustomClaims struct {
        User string `json:"user"`
        jwt.StandardClaims
    }

    // Create the Claims
    claims := MyCustomClaims{
        user,
        jwt.StandardClaims{
            ExpiresAt: 15000,
            Issuer:    "OwlH",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        logs.Error(err)
        return "", err
    }
    return tokenString, err
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password string, hash string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        logs.Error(err)
        return false, err
    }
    return true, nil
}

func VerifyToken(token string, user string) (err error) {
    users, err := ndb.GetLoginData()
    for x := range users {
        if users[x]["user"] == user {
            tkn, err := Encode(users[x]["user"], users[x]["secret"])
            if err != nil {
                logs.Error("Error checking token: %s", err)
                return err
            } else {
                if token == tkn {
                    return nil
                } else {
                    return errors.New("The token retrieved is false")
                }
            }
        }
    }
    return errors.New("There are not token. Error creating Token")
}

func VerifyPermissions(user string, object string, permissions []string) (hasPermissions bool, err error) {
    for x := range permissions {
        status, err := UserPermissionsValidation(user, permissions[x])
        if err != nil {
            logs.Error("requestType error: %s", err)
            return false, err
        }
        if status {
            utils.TokenMasterUser = user
            return true, nil
        }
    }
    return false, nil
}

func CanContinue(token, object string, permissions []string) (can bool, details map[string]string) {
    return true, nil

}

// ENCRYPT AND DECRYPT PASSWORD FOR RULESET
func createHash(key string) string {
    hasher := md5.New()
    hasher.Write([]byte(key))
    return hex.EncodeToString(hasher.Sum(nil))
}
func Encrypt(data []byte, passphrase string) []byte {
    block, _ := aes.NewCipher([]byte(createHash(passphrase)))
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return ciphertext
}
func Decrypt(data []byte, passphrase string) []byte {
    key := []byte(createHash(passphrase))
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err.Error())
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err.Error())
    }
    nonceSize := gcm.NonceSize()
    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        panic(err.Error())
    }
    return plaintext
}
