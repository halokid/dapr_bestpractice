package main

import (
  "context"
  "fmt"

  dapr "github.com/dapr/go-sdk/client"
)

func main() {
  // create the client
  ctx := context.Background()
  client, err := dapr.NewClient()
  if err != nil {
    panic(err)
  }
  defer client.Close()

  content := &dapr.DataContent{
    ContentType: "text/plain",
    Data:        []byte("hellow"),
  }
  resp, err := client.InvokeMethodWithContent(ctx, "neon_schedule", "ping", "post", content)
  if err != nil {
    panic(err)
  }
  fmt.Printf("service method invoked, response: %s\n", string(resp))

  in := &dapr.InvokeBindingRequest{
    Name:      "example-http-binding",
    Operation: "create",
  }
  if err := client.InvokeOutputBinding(ctx, in); err != nil {
    panic(err)
  }
  fmt.Println("output binding invoked")

  fmt.Println("DONE (CTRL+C to Exit)")
}
