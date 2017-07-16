package apiai

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetIntents(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse []IntentDescription
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `[
   {
      "id": "8ce9a60f-4a1e-468c-9718-5030de11eb91",
      "name": "hobby",
      "contextIn": [
         "hobby"
      ],
      "parameters": [
         {
            "dataType": "@hobby",
            "name": "hobby",
            "value": "$hobby"
         }
      ],
      "contextOut": [
         {
            "name": "hobby",
            "lifespan": 0
         }
      ],
      "actions": [
         "hobby"
      ],
      "priority": 500000,
      "fallbackIntent": false
   },
   {
      "id": "c261cb5f-4b31-4390-af4a-1e58620d462c",
      "name": "greetings",
      "contextIn": [
         "greetings"
      ],
      "parameters": [
         {
            "required": true,
            "dataType": "@sys.given-name",
            "name": "name",
            "value": "$name",
            "prompts": [
               "Hi! What is your name?"
            ]
         },
         {
            "dataType": "@sys.number",
            "name": "age",
            "value": "$age",
            "prompts": [],
            "defaultValue": "unknown"
         }
      ],
      "contextOut": [
         {
            "name": "greetings",
            "lifespan": 10
         },
         {
            "name": "hobby",
            "lifespan": 5
         }
      ],
      "actions": [
         "greetings"
      ],
      "priority": 500000,
      "fallbackIntent": false
   }
]`),
			expectedResponse: []IntentDescription{
				{
					Id:        "8ce9a60f-4a1e-468c-9718-5030de11eb91",
					Name:      "hobby",
					ContextIn: []string{"hobby"},
					ContextOut: []Context{
						{
							Name:     "hobby",
							Lifespan: 0,
						},
					},
					Actions: []string{"hobby"},
					Params: []IntentParameter{
						{
							DataType: "@hobby",
							Name:     "hobby",
							Value:    "$hobby",
						},
					},
					Priority:       500000,
					FallbackIntent: false,
				},
				{
					Id:        "c261cb5f-4b31-4390-af4a-1e58620d462c",
					Name:      "greetings",
					ContextIn: []string{"greetings"},
					ContextOut: []Context{
						{
							Name:     "greetings",
							Lifespan: 10,
						},
						{
							Name:     "hobby",
							Lifespan: 5,
						},
					},
					Actions: []string{"greetings"},
					Params: []IntentParameter{
						{
							Required: true,
							DataType: "@sys.given-name",
							Name:     "name",
							Value:    "$name",
							Prompts:  []string{"Hi! What is your name?"},
						},
						{
							DataType:     "@sys.number",
							Name:         "age",
							Value:        "$age",
							Prompts:      []string{},
							DefaultValue: "unknown",
						},
					},
					Priority:       500000,
					FallbackIntent: false,
				},
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
		httpmock.RegisterResponder("GET", c.buildUrl("intents", nil), tc.responder)

		r, err := c.GetIntents()

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestGetIntent(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse *Intent
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
  "id": "51ee06e9-9ff5-428b-aafd-733bbd7e9978",
  "name": "greetings",
  "auto": true,
  "contexts": [
    "greetings"
  ],
  "templates": [
    "hi I am @sys.given-name:name and I am @sys.number:age ",
    "hi my name is @sys.given-name:name and yours",
    "hi"
  ],
  "userSays": [
    {
      "id": "a4fd3a1d-fa7b-4b45-8fb8-f8b1f556872d",
      "data": [
        {
          "text": "hi I am "
        },
        {
          "text": "Alex",
          "alias": "name",
          "meta": "@sys.given-name",
          "userDefined": false
        },
        {
          "text": " and I am "
        },
        {
          "text": "15",
          "alias": "age",
          "meta": "@sys.number",
          "userDefined": false
        }
      ],
      "isTemplate": false,
      "count": 1
    },
    {
      "id": "e16ec309-e2b2-4218-b515-d90832ca74e3",
      "data": [
        {
          "text": "hi my name is "
        },
        {
          "text": "Sam",
          "alias": "name",
          "meta": "@sys.given-name",
          "userDefined": false
        },
        {
          "text": " and yours"
        }
      ],
      "isTemplate": false,
      "count": 1
    },
    {
      "id": "4d2c6dcc-6b8d-47b5-bc68-39b4f5aa1184",
      "data": [
        {
          "text": "hi"
        }
      ],
      "isTemplate": false,
      "count": 0
    }
  ],
  "responses": [
    {
      "resetContexts": false,
      "action": "greetings",
      "affectedContexts": [],
      "parameters": [
        {
          "required": true,
          "dataType": "@sys.given-name",
          "name": "name",
          "value": "$name",
          "defaultValue": "unknown",
          "prompts": [
            "Hi! What is your name?"
          ],
          "isList": false
        },
        {
          "dataType": "@sys.number",
          "name": "age",
          "value": "$age",
          "isList": false
        }
      ],
      "messages": [
        {
          "type": "0",
          "speech": "Hi! Nice to meet you, $name! What is your hobby?"
        }
      ]
    }
  ],
  "priority": 500000,
  "cortanaCommand": {
    "navigateOrService": "NAVIGATE",
    "target": ""
  },
  "webhookUsed": false,
  "fallbackIntent": false
}`),
			expectedResponse: &Intent{
				Id:        "51ee06e9-9ff5-428b-aafd-733bbd7e9978",
				Name:      "greetings",
				Auto:      true,
				Contexts:  []string{"greetings"},
				Templates: []string{"hi I am @sys.given-name:name and I am @sys.number:age ", "hi my name is @sys.given-name:name and yours", "hi"},
				UserSays: []UserSays{
					{
						Id: "a4fd3a1d-fa7b-4b45-8fb8-f8b1f556872d",
						Data: []Data{
							{Text: "hi I am "},
							{Text: "Alex", Alias: "name", Meta: "@sys.given-name", UserDefined: false},
							{Text: " and I am "},
							{Text: "15", Alias: "age", Meta: "@sys.number", UserDefined: false},
						},
						IsTemplate: false,
						Count:      1,
					},
					{
						Id: "e16ec309-e2b2-4218-b515-d90832ca74e3",
						Data: []Data{
							{Text: "hi my name is "},
							{Text: "Sam", Alias: "name", Meta: "@sys.given-name", UserDefined: false},
							{Text: " and yours"},
						},
						IsTemplate: false,
						Count:      1,
					},
					{
						Id: "4d2c6dcc-6b8d-47b5-bc68-39b4f5aa1184",
						Data: []Data{
							{Text: "hi"},
						},
						IsTemplate: false,
						Count:      0,
					},
				},
				Responses: []IntentResponse{
					{
						ResetContexts:    false,
						Action:           "greetings",
						AffectedContexts: []Context{},
						Params: []IntentParameter{
							{
								Required:     true,
								DataType:     "@sys.given-name",
								Name:         "name",
								Value:        "$name",
								DefaultValue: "unknown",
								Prompts:      []string{"Hi! What is your name?"},
								IsList:       false,
							},
							{
								DataType: "@sys.number",
								Name:     "age",
								Value:    "$age",
								IsList:   false,
							},
						},
						Messages: []Message{
							{
								Type:   "0",
								Speech: "Hi! Nice to meet you, $name! What is your hobby?",
							},
						},
					},
				},
				Priority:       500000,
				CortanaCommand: CortanaCommand{NavigateOrService: "NAVIGATE", Target: ""},
				WebhookUsed:    false,
				FallbackIntent: false,
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
		httpmock.RegisterResponder("GET", c.buildUrl("intents/51ee06e9-9ff5-428b-aafd-733bbd7e9978", nil), tc.responder)

		r, err := c.GetIntent("51ee06e9-9ff5-428b-aafd-733bbd7e9978")

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestCreateIntent(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse *CreationResponse
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
"id": "613de225-65b2-4fa8-9965-c14ae7673826",
"status": {
  "code": 200,
  "errorType": "success"
}
}`),
			expectedResponse: &CreationResponse{
				Id:     "613de225-65b2-4fa8-9965-c14ae7673826",
				Status: Status{Code: 200, ErrorType: "success"},
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
		httpmock.RegisterResponder("POST", c.buildUrl("intents", nil), tc.responder)

		r, err := c.CreateIntent(Intent{
			Name:      "change appliance state",
			Auto:      true,
			Contexts:  nil,
			Templates: []string{"turn @state:state the @appliance:appliance ", "switch the @appliance:appliance @state:state"},
			UserSays: []UserSays{
				{
					Data: []Data{
						{
							Text: "turn ",
						},
						{
							Text:  "on",
							Alias: "state",
							Meta:  "@state",
						},
						{
							Text: "the ",
						},
						{
							Text:  "kitchen lights",
							Alias: "appliance",
							Meta:  "@appliance",
						},
					},
					IsTemplate: false,
					Count:      0,
				},
				{
					Data: []Data{
						{
							Text: "switch the ",
						},
						{
							Text:  "heating",
							Alias: "appliance",
							Meta:  "@appliance",
						},
						{
							Text: " ",
						},
						{
							Text:  "off",
							Alias: "state",
							Meta:  "@state",
						},
					},
					IsTemplate: false,
					Count:      0,
				},
			},
			Responses: []IntentResponse{
				{
					ResetContexts: false,
					Action:        "set-appliance",
					AffectedContexts: []Context{
						{
							Name:     "house",
							Lifespan: 10,
						},
					},
					Params: []IntentParameter{
						{
							DataType: "@appliance",
							Name:     "appliance",
							Value:    "$appliance",
						},
						{
							DataType: "@state",
							Name:     "state",
							Value:    "$state",
						},
					},
					Messages: []Message{
						{
							Type:   0,
							Speech: "Turning the $appliance $state!",
						},
					},
				},
			},
			Priority: 500000,
		})

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestUpdateIntent(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedError error
	}{
		{
			description:   "api ai success, no errors",
			responder:     httpmock.NewStringResponder(200, `{}`),
			expectedError: nil,
		}, {
			description:   "api ai failed with an error 400",
			responder:     httpmock.NewStringResponder(http.StatusBadRequest, `{}`),
			expectedError: fmt.Errorf("apiai: wops something happens because status code is 400"),
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("PUT", c.buildUrl("intents/613de225-65b2-4fa8-9965-c14ae7673826", nil), tc.responder)

		err := c.UpdateIntent("613de225-65b2-4fa8-9965-c14ae7673826", Intent{
			Name:      "set appliance on or off",
			Auto:      true,
			Contexts:  nil,
			Templates: []string{"turn @state:state the @appliance:appliance ", "switch the @appliance:appliance @state:state"},
			UserSays: []UserSays{
				{
					Data: []Data{
						{
							Text: "turn ",
						},
						{
							Text:  "on",
							Alias: "state",
							Meta:  "@state",
						},
						{
							Text: "the ",
						},
						{
							Text:  "kitchen lights",
							Alias: "appliance",
							Meta:  "@appliance",
						},
					},
					IsTemplate: false,
					Count:      0,
				},
				{
					Data: []Data{
						{
							Text: "switch the ",
						},
						{
							Text:  "heating",
							Alias: "appliance",
							Meta:  "@appliance",
						},
						{
							Text: " ",
						},
						{
							Text:  "off",
							Alias: "state",
							Meta:  "@state",
						},
					},
					IsTemplate: false,
					Count:      0,
				},
			},
			Responses: []IntentResponse{
				{
					ResetContexts: false,
					Action:        "set-appliance",
					AffectedContexts: []Context{
						{
							Name:     "house",
							Lifespan: 10,
						},
					},
					Params: []IntentParameter{
						{
							DataType: "@appliance",
							Name:     "appliance",
							Value:    "$appliance",
						},
						{
							DataType: "@state",
							Name:     "state",
							Value:    "$state",
						},
					},
					Messages: []Message{
						{
							Type:   0,
							Speech: "Turning the $appliance $state!",
						},
					},
				},
			},
			Priority: 500000,
		},
		)

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestDeleteIntent(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description   string
		responder     httpmock.Responder
		expectedError error
	}{
		{
			description:   "api ai success, no errors",
			responder:     httpmock.NewStringResponder(200, `{}`),
			expectedError: nil,
		}, {
			description:   "api ai failed with an error 400",
			responder:     httpmock.NewStringResponder(http.StatusBadRequest, `{}`),
			expectedError: fmt.Errorf("apiai: wops something happens because status code is 400"),
		},
	}

	for _, tc := range tests {
		httpmock.RegisterResponder("DELETE", c.buildUrl("intents/80f817e8-23fb-4e8e-ba62-eca1fcef7c3a", nil), tc.responder)

		err := c.DeleteIntent("80f817e8-23fb-4e8e-ba62-eca1fcef7c3a")

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
