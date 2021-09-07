package administrations

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "example.com/schools/pkg/models"
  "example.com/schools/pkg/validations"
)

var user models.User

func RegisterStudents(writer http.ResponseWriter, request *http.Request) {
  reqBody, _ := ioutil.ReadAll(request.Body)
  user.Roles = "student"
  user.DescriptionRoles = "Ini Murid"
  json.Unmarshal(reqBody, &user)
  if validErrs := validations.RegisterValidate(user); len(validErrs) > 0 {
    err := validations.GetValidateMessages(writer, validErrs)
    json.NewEncoder(writer).Encode(err)
    return
  }
  result := models.CreateUser(user)
  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
  response, err := json.Marshal(result)
  if err != nil {
    panic(err)
  }
  writer.Write(response)
}


func GetAllStudents(writer http.ResponseWriter, request *http.Request) {
  writer.Write([]byte(`{"messages": "Yay"}`))
}
