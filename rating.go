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

func cleanMap(m *map[string]int,s []string) map[string]int {
  newMap := make(map[string]int)
  var fadd bool

  for key, value := range *m{
    fadd = false
    for val := range s{
      if key == s[val]{
        fadd = true
        break
      }
    // Customer, Financial, Reliability, Safety
      if key < s[val] {
        break
      }

    }
    if fadd{
      newMap[key]=value
    }
  }

  return newMap
}
  

func fail400(w http.ResponseWriter){


  w.WriteHeader(http.StatusBadRequest)
  w.Write([]byte("400 Nope."))

}

type ResultData struct{
  Source map[string] int
  Scores map[string] float64
}

func calcResult(mp *map[string]int) ResultData {
  var m map[string]int = *mp
  var res = ResultData{}
  res.Source = m

  average :=  float64(m["Customer"]+
              m["Financial"]+
              m["Reliability"]+ 
              m["Safety"])/float64(4.0)
  
  res.Scores["Average"]= average

  return res
}

func processData(w http.ResponseWriter, body []byte){
  if json.Valid(body){
    fmt.Fprintln(w, "Body is Valid")
    a := make(map[string]int)
    err := json.Unmarshal(body, &a)
    if err == nil{
      fmt.Fprintln(w,"Success!");
      fmt.Fprintln(w, "Length:", len(a))
      nm := cleanMap(&a, []string{"Customer", "Financial", "Reliability", "Safety"});
      for key, value := range nm {
          fmt.Fprint(w,key, ": ",strconv.Itoa(value), "\r\n")
          }
      res := calcResult(&nm)
      b,_ := json.Marshal(res)
      w.Write(b)
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
