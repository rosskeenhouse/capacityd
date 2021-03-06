package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
//    "encoding/json"
    "github.com/golang/protobuf/jsonpb"
    capacity_protocol "github.com/rosskeenhouse/capacity/server/proto"
)

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  marshaler := jsonpb.Marshaler{}
  status := &capacity_protocol.Capacity{ State: capacity_protocol.Capacity_HEALTHY }
  j, e := marshaler.MarshalToString(status)
  if e != nil {
    //t.Errorf("%s: marshaling error: %v", tt.desc, err)
    return
  }
  fmt.Fprint(w, j)
}

func main() {
    router := httprouter.New()
    router.GET("/", Health)
    router.GET("/health/:name", Health)

    log.Fatal(http.ListenAndServe(":8080", router))
}
