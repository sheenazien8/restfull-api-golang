package validations

import (
  "net/http"
  "net/url"
  "example.com/schools/pkg/models"
)

func RegisterValidate(UserData models.User) url.Values {
  errs := url.Values{}
  // you can simplify this rule with loop or anything you want
  if UserData.NickName == "" {
    errs.Add("nickname", "Nick Name is required")
  }

  return errs
}

func GetValidateMessages(writer http.ResponseWriter, err url.Values) map[string]interface{} {
  writer.Header().Set("Content-type", "applciation/json")
  writer.WriteHeader(http.StatusBadRequest)

  errors := map[string]interface{}{"validationError": err}

  return errors
}
