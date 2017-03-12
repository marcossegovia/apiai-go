# Go SDK for api.ai

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
        var client = apiai.NewClient(
                &ClientConfig{
                        token:      "YOUR-API-AI-TOKEN",
                        sessionId:  "YOUR_USER_SESSION_ID",
                        queryLang:  "en",    //Default en
                        speechLang: "en-US", //Default en-US
                },
        )

        qr, err := client.Query(Query{Query: []string{"My name is Marcos and I live in Barcelona"}})
        if err != nil {
                fmt.Printf("%v", err)
        }
        fmt.Printf("%v", qr.Result.Fulfillment.Speech)
}
```
## Bugs & Issues

See [CONTRIBUTING](CONTRIBUTING.md)
