// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "encoding/json"
  "strconv"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func fail400(w http.ResponseWriter){


  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte("400 Nope."))

}

func processData(w http.ResponseWriter, body []byte){
  if json.Valid(body){
    w.Write([]byte("Body is Valid"))
    a := make(map[string]int)
    err := json.Unmarshal(body, &a)
    if err == nil{
      w.Write([]byte("Success!"))
      w.Write([]byte(a["Safety"]))
      w.Write([]byte("Length:"))
      w.Write([]byte(strconv.Itoa((len(a)))))
      for key, value := range a {
          w.Write([]byte(key))
          w.Write([]byte(value))
          }
    }else{
      w.Write([]byte(err.Error()))
    }
  }else{
    w.Write([]byte("Invalid Body"))
  }

}


func handle(w http.ResponseWriter, r *http.Request) {
  if r.Header.Get("Content-Type") == "application/json" && r.Method == "POST"{
    body, err := ioutil.ReadAll(r.Body)
    if err==nil{
    

      fmt.Fprintln(w, "Basic Tests Passed. Data: ",string(body))
      processData(w, body)
    }else{
      fail400(w)
    }
  }
}
