package apiai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Platform struct {
	Source string            `json:"source"`
	Data   map[string]string `json:"data"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Event struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type Query struct {
	Query           []string            `json:"query"`
	Event           Event               `json:"event"`
	Version         string              `json:"-"`
	SessionId       string              `json:"sessionId"`
	Language        string              `json:"lang"`
	Contexts        []Context           `json:"contexts"`
	ResetContexts   bool                `json:"resetContexts"`
	Entities        []EntityDescription `json:"entities"`
	Timezone        string              `json:"timezone"`
	Location        Location            `json:"location"`
	OriginalRequest Platform            `json:"originalRequest"`
}

type CreationResponse struct {
	Id     string `json:"id"`
	Status Status `json:"status"`
}

type CardButton struct {
	Text     string
	Postback string
}

type Metadata struct {
	IntentId                  string `json:"intentId"`
	WebhookUsed               string `json:"webhookUsed"`
	WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
	IntentName                string `json:"intentName"`
}

type Message struct {
	Type     interface{}  `json:"type"`
	Speech   string       `json:"speech"`
	ImageUrl string       `json:"imageUrl"`
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle"`
	Buttons  []CardButton `json:"buttons"`
	Replies  []string     `json:"replies"`
	Payload  interface{}  `json:"payload"`
}

type Fulfilment struct {
	Speech      string    `json:"speech"`
	DisplayText string    `json:"displayText"`
	Messages    []Message `json:"messages"`
}

type Status struct {
	Code         int    `json:"code"`
	ErrorType    string `json:"errorType"`
	ErrorId      string `json:"errorId"`
	ErrorDetails string `json:"errorDetails"`
}

type Result struct {
	Source           string                 `json:"source"`
	ResolvedQuery    string                 `json:"resolvedQuery"`
	Action           string                 `json:"action"`
	ActionIncomplete bool                   `json:"actionIncomplete"`
	Params           map[string]interface{} `json:"parameters"`
	Contexts         []Context              `json:"contexts"`
	Fulfillment      Fulfilment             `json:"fulfillment"`
	Score            float64                `json:"score"`
	Metadata         Metadata               `json:"metadata"`
}

type QueryResponse struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Language  string    `json:"lang"`
	Result    Result    `json:"result"`
	Status    Status    `json:"status"`
	SessionId string    `json:"sessionId"`
}

func (c *ApiClient) Query(q Query) (*QueryResponse, error) {
	q.Version = c.config.Version
	q.Language = c.config.QueryLang
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(q)

	if err != nil {
		return nil, fmt.Errorf("apiai: error on request, %v", err)
	}

	req, err := http.NewRequest("POST", c.buildUrl("query", nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-type", "application/json, charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	if len(c.config.ProxyURL) > 0 {
		url, err := url.Parse(c.config.ProxyURL)
		if err == nil {
			httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(url)}
		}
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response *QueryResponse
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&response)
		if err != nil {
			return nil, err
		}
		return response, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}
