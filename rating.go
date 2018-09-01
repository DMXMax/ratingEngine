// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
//  "encoding/json"
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

func handle(w http.ResponseWriter, r *http.Request) {
  if r.Header.Get("Content-Type") == "application/json" && r.Method == "POST"{
    body, err := ioutil.ReadAll(r.Body)
    if err==nil{
      fmt.Println(err)
      fail400(w)
    }

	  fmt.Fprintln(w, "Hello From the Ratings Engine!",body)
  }else{
    fail400(w)
  }
}
