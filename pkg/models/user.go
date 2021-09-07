package models

import (
  "example.com/schools/pkg/config"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  Id int `json:"id"`
  Roles string `json:"roles"`
  DescriptionRoles string `json:"description_roles"`
  Username string `json:"username"`
  Password string `json:"password"`
  NickName string `json:"nick_name"`
  FullName string `json:"full_name"`
  Nis string `json:"nis"`
}

var Users []User

func GetAllUser() []User {
  Users = []User{
    {
      Id: 1,
      Username: "sheena",
      Roles: "students",
      DescriptionRoles: "Ini DescriptionRoles",
      Password: "Password",
      FullName: "FirstName",
      NickName: "LastName",
      Nis: "NIs",
    },
  }

  return Users
}

func CreateUser(userData User) User {
  db := config.ConnectToDB()
  password := []byte(userData.Password)
  // Hashing the password with the default cost of 10
  hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
  userData.Password = string(hashedPassword)
  if err != nil {
    panic(err.Error())
  }
  query := "INSERT INTO users(roles, description_roles, username, password, nick_name, full_name, nis) VALUES (?,?,?,?,?,?,?)"
  insert, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }
  recorded, err := insert.Exec(
    userData.Roles,
    userData.DescriptionRoles,
    userData.Username,
    userData.Password,
    userData.NickName,
    userData.FullName,
    userData.Nis,
  )
  if err != nil {
    panic(err.Error())
  }
  lastId, err := recorded.LastInsertId()
  if err != nil {
    panic(err.Error())
  }
  userData.Id = int(lastId)
  defer insert.Close()

  return userData
}

func GetUserByColumn(column string, value string) ([]User, int) {
  db := config.ConnectToDB()
  query := "SELECT * FROM users where " + column + "=?"
  result, err := db.Query(query, value)
  if err != nil {
    panic(err.Error())
  }
  user := User{}
  var response = []User{}
  var count = 0
  for result.Next() {
    count += 1
    result.Scan(
      &user.Id,
      &user.Roles,
      &user.DescriptionRoles,
      &user.Username,
      &user.Password,
      &user.NickName,
      &user.FullName,
      &user.Nis,
    )
    response = append(response, user)
  }
  defer db.Close()

  return response, count
}
