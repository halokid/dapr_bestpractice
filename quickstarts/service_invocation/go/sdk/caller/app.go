package main

import (
  "context"
  "fmt"
  "time"

  dapr "github.com/dapr/go-sdk/client"
)

// TODO: dapr.exe run --app-id caller --log-level debug  -- go run .

func main() {
  // create the client
  ctx := context.Background()
  client, err := dapr.NewClient()
  if err != nil {
    panic(any(err))
  }
  defer client.Close()

  content := &dapr.DataContent{
    ContentType: "text/plain",
    Data:        []byte("hellow"),
    // Data: []byte(`{"orderId":99}`),
  }
  for i := 0; i < 100; i++ {
    resp, err := client.InvokeMethodWithContent(ctx, "serving", "echo", "post", content)
    if err != nil {
      panic(any(err))
    }
    fmt.Printf("service method invoked, response: %s, i: %d\n", string(resp), i)
    time.Sleep(2 * time.Second)
  }

  fmt.Println("DONE (CTRL+C to Exit)")

  // TODO: if want the service show in the zipkin for the `service route call`, must let the service run in
  // TODO: daemon model, can not exist the service program, if not, the service `caller` will now show in the
  // TODO: zipkin service route

  time.Sleep(1000 * time.Second)
}
