package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

//var a string
//var b chan int

// TODO: the `DAPR_HTTP_PORT` is defined in the command line start server
// TODO: dapr run --app-id checkout --app-protocol http --dapr-http-port 3500 -- go run .
// TODO: defined by `--dapr-http-port 3500`,  not read from OS env
// TODO: if `DAPR_HTTP_PORT` not defined in the command line, dapr will give a random port
// TODO: for it.

func main() {
  var DAPR_HOST, DAPR_HTTP_PORT string
  var okHost, okPort bool
  log.Printf("DAPR_HTTP_PORT 1 -->>> %+v", DAPR_HTTP_PORT)
  if DAPR_HOST, okHost = os.LookupEnv("DAPR_HOST"); !okHost {
    DAPR_HOST = "http://localhost"
  }
  if DAPR_HTTP_PORT, okPort = os.LookupEnv("DAPR_HTTP_PORT"); !okPort {
    log.Printf("defined DAPR_HTTP_PORT 1 -->>> %+v", DAPR_HTTP_PORT)
    DAPR_HTTP_PORT = "3500"
    // DAPR_HTTP_PORT = "3501"
  }
  log.Printf("DAPR_HTTP_PORT 2 -->>> %+v", DAPR_HTTP_PORT)
  for i := 1; i <= 10; i++ {
    order := "{\"orderId\":" + strconv.Itoa(i) + "}"
    client := &http.Client{}

    // DAPR_HTTP_PORT = "3501"    // TODO: just direct call service, so no service track
    url := DAPR_HOST + ":" + DAPR_HTTP_PORT + "/orders"

    log.Printf("url -->>> %+v", url)
    req, err := http.NewRequest("POST", url, strings.NewReader(order))
    // req, err := http.NewRequest("POST", DAPR_HOST+":"+DAPR_HTTP_PORT+"/orders", strings.NewReader(order))
    if err != nil {
      fmt.Print(err.Error())
      os.Exit(1)
    }

    // TODO: so `service A` call `service B` process is, you can call every service in the
    // TODO: dapr service list scope, if you use http, the service you can is support http by
    // TODO: defined `DAPR_HTTP_PORT`, then add the  `dapr-app-id` in the header special which
    // TODO: service you want to invoke, if your entry service is A, then will create a route
    // TODO: track in zipkin `A -> B`, if the entry is `B` itself, no route track will show,
    // TODO: just show call the `B` in zipkin

    // Adding app id as part of th header
    req.Header.Add("dapr-app-id", "order-processor")

    // Invoking a service
    response, err := client.Do(req)

    if err != nil {
      fmt.Print(err.Error())
      os.Exit(1)
    }

    result, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println("Order passed: ", string(result))
    time.Sleep(5 * time.Second)
  }

  // block the service for daemon
  block := make(chan int)
  <-block
}
