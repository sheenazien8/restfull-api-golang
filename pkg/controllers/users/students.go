package users

import (
  "encoding/json"
  "net/http"
  "example.com/schools/pkg/models"
)

func UsersAll(writer http.ResponseWriter, request *http.Request) {
  Users := models.GetAllUser()
  writer.WriteHeader(http.StatusOK)
  rest, err := json.Marshal(Users)
  if err != nil {
    panic(err)
  }
  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
  writer.WriteHeader(http.StatusOK)
  writer.Write(rest)
}
