package utilities

import "net/http"
import "reflect"

func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
  for _, mw := range middleware {
    h = mw(h)
  }
  return h
}

func InArray(val interface{}, array interface{}) (exists bool) {
    exists = false

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                exists = true
                return
            }
        }
    }

    return
}
