// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
  if r.Header.Get("Content-Type") == "application/json" && r.Method == "POST"{
	  fmt.Fprintln(w, "Hello From the Ratings Engine!", r.Method,r.Header)
  }else{
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("400 Nope."))
    w.Write([]byte(r.Header.Get("Content-Type")))
  }
}
