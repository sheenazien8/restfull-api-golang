package routes

import (
  "fmt"
  "net/http"
  "example.com/schools/pkg/controllers/administrations"
  "example.com/schools/pkg/controllers/auth"
  "example.com/schools/pkg/middlewares"
  "example.com/schools/pkg/utilities"
  "github.com/gorilla/mux"
)

func RequestHandlers() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/check", func(writer http.ResponseWriter, request *http.Request) {
    var jsonData = []byte(`{ "messages": "api tersedia" }`)
    writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
    writer.WriteHeader(http.StatusOK)
    writer.Write(jsonData)
  }).Methods("GET")

  router.Handle("/administrations/students", utilities.Middleware(
    http.HandlerFunc(administrations.RegisterStudents),
    middlewares.IsAuthorized,
  )).Methods("POST")
  router.Handle("/administrations/students", utilities.Middleware(
    http.HandlerFunc(administrations.GetAllStudents),
    middlewares.IsAuthorized,
  ))
  router.Handle("/administrations/students/{id}", utilities.Middleware(
    http.HandlerFunc(administrations.GetStudents),
    middlewares.IsAuthorized,
  ))

  // Auth
  router.HandleFunc("/login", auth.Login).Methods("POST")
  router.Handle("/logout", utilities.Middleware(
    http.HandlerFunc(auth.Logout), 
    middlewares.IsAuthorized,
  )).Methods("DELETE")

  fmt.Println("Server running on portn 8080")
  http.ListenAndServe(":8080", router)
}
