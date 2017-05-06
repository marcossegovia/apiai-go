# Go SDK for api.ai
[![Go Report Card](https://goreportcard.com/badge/github.com/marcossegovia/apiai-go)](https://goreportcard.com/report/github.com/marcossegovia/apiai-go)
[![GoDoc](https://godoc.org/github.com/marcossegovia/apiai-go?status.svg)](https://godoc.org/github.com/marcossegovia/apiai-go)
[![Build Status](https://travis-ci.org/marcossegovia/apiai-go.svg?branch=master)](https://travis-ci.org/marcossegovia/apiai-go)
[![codecov](https://codecov.io/gh/marcossegovia/apiai-go/branch/master/graph/badge.svg)](https://codecov.io/gh/marcossegovia/apiai-go)

This library allows you to integrate Api.ai natural language processing service with your Go application.
For more information see the [docs](https://docs.api.ai/docs).

We encourage you to follow this official guide to [get started in 5 steps](https://docs.api.ai/docs/get-started) before proceeding. 

## Installation

```bash
go get github.com/marcossegovia/apiai-go
```

## Usage

```go
import (
        "fmt"

        "github.com/marcossegovia/apiai-go"
)

func main() {
    client, err := apiai.NewClient(
        &apiai.ClientConfig{
            Token:      "YOUR-API-AI-TOKEN",
            QueryLang:  "en",    //Default en
            SpeechLang: "en-US", //Default en-US
        },
    )
    if err != nil {
        fmt.Printf("%v", err)
    }
    //Set the query string and your current user identifier.
    qr, err := client.Query(apiai.Query{Query: []string{"My name is Marcos and I live in Barcelona"}, SessionId: "123454321"})
    if err != nil {
        fmt.Printf("%v", err)
    }
    fmt.Printf("%v", qr.Result.Fulfillment.Speech)
}
```
## Bugs & Issues

See [CONTRIBUTING](CONTRIBUTING.md)
