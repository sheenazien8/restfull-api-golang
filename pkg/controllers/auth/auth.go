package auth

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "time"
  "example.com/schools/pkg/config"
  "example.com/schools/pkg/models"
  jwt "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
)

var mySigningKey = []byte(config.GetConfig().SigningKey)
var user models.User

func Login(writer http.ResponseWriter, request *http.Request)  {
  reqBody, _ := ioutil.ReadAll(request.Body)
  json.Unmarshal(reqBody, &user)
  result, count := models.GetUserByColumn("username", user.Username)
  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if count > 0 {
    res := result[0]
    password := []byte(user.Password)
    hashedPassword := []byte(res.Password)
    isNotMatched := bcrypt.CompareHashAndPassword(hashedPassword, password)
    if isNotMatched != nil {
      writer.WriteHeader(http.StatusBadRequest)
      writer.Write([]byte(`{ "messages": "Password anda salah!" }`))
      return
    }
    // generate jwt token
    token, err := GenerateJWT(res)
    if err != nil {
      panic(err.Error())
    }
    responseUser, err := json.Marshal(res)
    response := []byte(`{
      "messages": "Anda Sukses Login",
      "token": "`+ token +`",
      "data": `+ string(responseUser) +`,
    }`)
    if err != nil {
      panic(err)
    }
    writer.Write(response)
  } else {
    writer.WriteHeader(http.StatusBadRequest)
    writer.Write([]byte(`{ "messages": "Maaf User tidak ditemukan" }`))
  }
}

func Logout(writer http.ResponseWriter, request *http.Request) {
}

func GenerateJWT(userData models.User) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = userData.Username
    claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}

