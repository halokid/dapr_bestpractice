package main

import (
  "context"
  "fmt"

  dapr "github.com/dapr/go-sdk/client"
)

// TODO: dapr run --app-id caller --log-level debug go run .

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
  }
  //resp, err := client.InvokeMethodWithContent(ctx, "neon_schedule", "ping", "post", content)
  resp, err := client.InvokeMethodWithContent(ctx, "order-processor", "orders", "post", content)
  if err != nil {
    panic(any(err))
  }
  fmt.Printf("service method invoked, response: %s\n", string(resp))

  in := &dapr.InvokeBindingRequest{
    Name:      "example-http-binding",
    Operation: "create",
  }
  if err := client.InvokeOutputBinding(ctx, in); err != nil {
    panic(any(err))
  }
  fmt.Println("output binding invoked")

  fmt.Println("DONE (CTRL+C to Exit)")
}





