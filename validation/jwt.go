package validation

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/md5"
  "crypto/rand"
  "encoding/hex"
  "encoding/json"
  "errors"
  "io"
  "strconv"
  "time"

  ndb "github.com/OwlH-net/OwlH-Master/database"
  "github.com/OwlH-net/OwlH-Master/utils"
  "github.com/astaxie/beego/logs"
  jwt "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
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

type Token struct {
  Secret    string `json:"secret"`
  Timestamp string `json:"timestamp"`
}

func SaveUserLoginData(user string, secret string) (err error) {
  var tokens = []Token{}
  var newToken = Token{}

  //get default time for check if timestamp has expired
  mainTimeout, nil := utils.GetKeyValueInt("master", "tokenTimestamp")
  if err != nil {
    mainTimeout = 300
  } //default time for check

  //get timestamp
  currentTime := time.Now().Unix()

  //get db struct
  users, err := ndb.GetLoginData()
  if err != nil {
    logs.Error("ERROR saving user login data: %s", err)
    return err
  }

  for x := range users {
    if users[x]["user"] == user {
      // keys := make(tokens,0)
      json.Unmarshal([]byte(users[x]["userTokens"]), &tokens)

      newToken.Secret = secret
      newToken.Timestamp = strconv.Itoa(int(currentTime) + int(mainTimeout))

      tokens = append(tokens, newToken)

      userTokens, _ := json.Marshal(tokens)

      err = ndb.UpdateUser(x, "userTokens", string(userTokens))
      if err != nil {
        logs.Error("ERROR updating user login data: %s", err)
        return err
      }
    }
  }

  //save new data into struct

  return nil
}

// //verify token for every user
// func VerifyUserToken(user string) (err error) {
//     var tokens = []Token{}
//     users, err := ndb.GetLoginData()

//     logs.Info("%+v",users)

//     for x := range users {
//         if users[x]["user"] == user {
//             // keys := make(tokens,0)
//             json.Unmarshal([]byte(users[x]["userTokens"]), &tokens)

//             for tokenID := range tokens{
//                 logs.Debug("Token -> %s, Timestamp -> %s", tokens[tokenID].Secret, tokens[tokenID].Timestamp)
//             }

//         }
//     }

//     return nil
// }

// verify master token
func VerifyToken(tokenRetrieved string, user string) (err error) {
  var tokens = []Token{}
  var tokensFiltered = []Token{}
  var newToken = Token{}
  currentTime := time.Now().Unix()

  //get default time for check if timestamp has expired
  mainTimeout, nil := utils.GetKeyValueInt("master", "tokenTimestamp")
  if err != nil {
    mainTimeout = 300
  } //default time for check

  users, err := ndb.GetLoginData()
  if err != nil {
    logs.Error("VerifyToken ERROR getting login data: %s", err)
    return err
  }

  for x := range users {

    if users[x]["user"] == user {
      json.Unmarshal([]byte(users[x]["userTokens"]), &tokens)

      //Delete invalid secret keys for token
      for tokenID := range tokens {
        ts, _ := strconv.ParseInt(tokens[tokenID].Timestamp, 10, 64)
        tkn, err := Encode(user, tokens[tokenID].Secret)
        if err != nil {
          logs.Error("VerifyToken ERROR checking token: %s", err)
          return err
        }

        if users[x]["expire"] == "true" {
          if int(ts) >= (int(currentTime)) {
            if tkn == tokenRetrieved {
              //save secret and timstamp
              newToken.Secret = tokens[tokenID].Secret
              newToken.Timestamp = strconv.Itoa(int(currentTime) + int(mainTimeout))
              tokensFiltered = append(tokensFiltered, newToken)
            } else {
              //save secret and timstamp
              newToken.Secret = tokens[tokenID].Secret
              newToken.Timestamp = tokens[tokenID].Timestamp
              tokensFiltered = append(tokensFiltered, newToken)
            }
          }
        } else {
          newToken.Secret = tokens[tokenID].Secret
          tokensFiltered = append(tokensFiltered, newToken)
          newToken.Timestamp = "0"
        }
      }

      //save tokens into db
      userTokens, _ := json.Marshal(tokensFiltered)
      err = ndb.UpdateUser(x, "userTokens", string(userTokens))
      if err != nil {
        logs.Error("ERROR updating user login data: %s", err)
        return err
      }

      //over all valid secret keys, check tokens
      for tokenID := range tokensFiltered {
        tkn, err := Encode(user, tokensFiltered[tokenID].Secret)
        if err != nil {
          logs.Error("Error checking token: %s", err)
          return err
        } else {
          if tokenRetrieved == tkn {
            return nil
          }
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
