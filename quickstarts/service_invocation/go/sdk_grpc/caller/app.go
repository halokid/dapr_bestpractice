package main

import (
  "context"
  "fmt"
  dapr "github.com/dapr/go-sdk/client"
  "time"
)

// TODO:  dapr run --app-id caller  --log-level debug -- go run .
func main() {
  // just for this demo
  ctx := context.Background()
  //json := `{ "message": "hello" }`
  //data := []byte(json)
  //store := "statestore"
  //pubsub := "messages"

  // create the client
  client, err := dapr.NewClient()
  if err != nil {
    panic(any(err))
  }
  defer client.Close()

  // invoke a method called EchoMethod on another dapr enabled service
  content := &dapr.DataContent{
    ContentType: "text/plain",
    Data:        []byte("hellow"),
  }
  for i := 0; i < 100; i++ {
    resp, err := client.InvokeMethodWithContent(ctx, "serving", "echo", "post", content)
    if err != nil {
      panic(any(err))
    }
    fmt.Printf("service method invoked, response: %s\n", string(resp))
    time.Sleep(1 * time.Second)
  }

  fmt.Println("DONE (CTRL+C to Exit)")
  time.Sleep(100 * time.Second)
}



