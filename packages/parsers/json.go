package json

import (
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type Message struct {
  Test string `json:"test"`
}

func Parse(request *http.Request) Message {
  var result Message
  body, _ := ioutil.ReadAll(request.Body)
  json.Unmarshal(body, &result)
  return result
}
