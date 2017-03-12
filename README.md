# Go SDK for api.ai
[![Go Report Card](https://goreportcard.com/badge/github.com/marcossegovia/apiai-go)](https://goreportcard.com/report/github.com/marcossegovia/apiai-go)
[![GoDoc](https://godoc.org/github.com/marcossegovia/apiai-go?status.svg)](https://godoc.org/github.com/marcossegovia/apiai-go)

This library allows you to integrate Api.ai natural language processing service with your Go application.
See the [docs](https://docs.api.ai/docs) for more information.

## Installation

```bash
go get marcossegovia/apiai-go
```

## Usage

```go
import (
        "fmt"

        "github.com/marcossegovia/apiai-go"
)

func main() {
        client, err := apiai.NewClient(
                &ClientConfig{
                        token:      "YOUR-API-AI-TOKEN",
                        sessionId:  "YOUR_USER_SESSION_ID",
                        queryLang:  "en",    //Default en
                        speechLang: "en-US", //Default en-US
                },
        )
        if err != nil{
                fmt.Printf("%v", err)
        }

        qr, err := client.Query(Query{Query: []string{"My name is Marcos and I live in Barcelona"}})
        if err != nil {
                fmt.Printf("%v", err)
        }
        fmt.Printf("%v", qr.Result.Fulfillment.Speech)
}
```
## Bugs & Issues

See [CONTRIBUTING](CONTRIBUTING.md)
