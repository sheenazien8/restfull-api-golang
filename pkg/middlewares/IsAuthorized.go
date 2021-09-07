package middlewares

import (
  "fmt"
  "net/http"
  "strings"
  "example.com/schools/pkg/config"
  jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(config.GetConfig().SigningKey)

func IsAuthorized(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tokenString := extractToken(r)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if tokenString != "" {
      token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	  return nil, fmt.Errorf("There was an error")
	}
	return mySigningKey, nil
      })

      if err != nil {
	w.WriteHeader(401)
	fmt.Fprintf(w, `{"messages": "Token is expired"}`)
      }
      if token.Valid {
	next.ServeHTTP(w, r)
      }
    } else {
      w.WriteHeader(401)
      fmt.Fprintf(w, `{"messages": "Token not provided"}`)
    }
  })
}

func extractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  //normally Authorization the_token_xxx
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}
