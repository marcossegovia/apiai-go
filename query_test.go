package apiai

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse *QueryResponse
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
  "id": "b340a1f7-abee-4e13-9bdd-5e8938a48b7d",
  "timestamp": "1992-02-04T00:00:00Z",
  "lang": "en",
  "result": {
    "source": "agent",
    "resolvedQuery": "my name is Marcos and I live in Barcelona",
    "action": "greetings",
    "actionIncomplete": false,
    "parameters": {},
    "contexts": [],
    "metadata": {
      "intentId": "a123a123",
      "webhookUsed": "false",
      "webhookForSlotFillingUsed": "false",
      "intentName": "greetings"
    },
    "fulfillment": {
      "speech": "Hi Marcos! Nice to meet you!",
      "messages": [
        {
          "type": 0,
          "speech": "Hi Marcos! Nice to meet you!"
        }
      ]
    },
    "score": 1
  },
  "status": {
    "code": 200,
    "errorType": "success"
  },
  "sessionId": "123454321"
}`),
			expectedResponse: &QueryResponse{
				Id:        "b340a1f7-abee-4e13-9bdd-5e8938a48b7d",
				Timestamp: time.Date(1992, time.February, 04, 00, 00, 00, 0, time.UTC),
				Language:  "en",
				Result: Result{
					Source:        "agent",
					ResolvedQuery: "my name is Marcos and I live in Barcelona",
					Action:        "greetings",
					Contexts:      []Context{},
					Fulfillment: Fulfilment{
						Speech: "Hi Marcos! Nice to meet you!",
						Messages: []Message{
							{Type: 0, Speech: "Hi Marcos! Nice to meet you!"},
						},
					},
					Score: 1,
					Metadata: Metadata{
						IntentId:                  "a123a123",
						WebhookUsed:               "false",
						WebhookForSlotFillingUsed: "false",
						IntentName:                "greetings",
					},
				},
				Status:    Status{Code: http.StatusOK, ErrorType: "success"},
				SessionId: "123454321",
			},
			expectedError: nil,
		}, {
			description:      "api ai failed with an error 400",
			responder:        httpmock.NewStringResponder(http.StatusBadRequest, `{}`),
			expectedResponse: nil,
			expectedError:    fmt.Errorf("apiai: wops something happens because status code is 400"),
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("POST", c.buildUrl("query", nil), tc.responder)

		r, err := c.Query(Query{Query: []string{"my name is Marcos and I live in Barcelona"}})

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
