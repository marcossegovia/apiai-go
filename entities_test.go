package apiai

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetEntities(t *testing.T) {
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
		expectedResponse []EntityDescription
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `[
  {
    "id": "33868522-5747-4a31-88fb-3cd13bd18684",
    "name": "Appliances",
    "count": 11,
    "preview": "Coffee Maker <= (coffee maker, coffee machine, coffee), ..."
  },
  {
    "id": "6d6b7d50-7510-4fec-927b-ac3c3aaff009",
    "name": "Utilities",
    "count": 4,
    "preview": "Electricity <= (electricity, electrical), ..."
  }
]`),
			expectedResponse: []EntityDescription{
				{
					Id:      "33868522-5747-4a31-88fb-3cd13bd18684",
					Name:    "Appliances",
					Count:   11,
					Preview: "Coffee Maker <= (coffee maker, coffee machine, coffee), ...",
				},
				{
					Id:      "6d6b7d50-7510-4fec-927b-ac3c3aaff009",
					Name:    "Utilities",
					Count:   4,
					Preview: "Electricity <= (electricity, electrical), ...",
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
		httpmock.RegisterResponder("GET", c.buildUrl("entities", nil), tc.responder)

		r, err := c.GetEntities()

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestGetEntity(t *testing.T) {
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
		expectedResponse *Entity
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
  "id": "1de251bf-46a6-4056-af9c-96b6ca89dfd0",
  "name": "appliances",
  "entries": [
    {
      "value": "coffee maker",
      "synonyms": [
        "coffee maker",
        "coffee"
      ]
    },
    {
      "value": "thermostat",
      "synonyms": [
        "thermostat",
        "heat",
        "air conditioning"
      ]
    }
  ],
  "isEnum": false,
  "automatedExpansion": false
}`),
			expectedResponse: &Entity{
				Id:   "1de251bf-46a6-4056-af9c-96b6ca89dfd0",
				Name: "appliances",
				Entries: []Entry{
					{
						Value:    "coffee maker",
						Synonyms: []string{"coffee maker", "coffee"},
					},
					{
						Value:    "thermostat",
						Synonyms: []string{"thermostat", "heat", "air conditioning"},
					},
				},
				IsEnum:             false,
				AutomatedExpansion: false,
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
		httpmock.RegisterResponder("GET", c.buildUrl("entities/1de251bf-46a6-4056-af9c-96b6ca89dfd0", nil), tc.responder)

		r, err := c.GetEntity("1de251bf-46a6-4056-af9c-96b6ca89dfd0")

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestCreateEntity(t *testing.T) {
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
		expectedResponse *CreationResponse
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
  "id": "6d6b7d50-7510-4fec-927b-ac3c3aaff009",
  "status": {
    "code": 200,
    "errorType": "success"
  }
}`),
			expectedResponse: &CreationResponse{
				Id:     "6d6b7d50-7510-4fec-927b-ac3c3aaff009",
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
		httpmock.RegisterResponder("POST", c.buildUrl("entities", nil), tc.responder)

		r, err := c.CreateEntity(Entity{
			Name: "appliances",
			Entries: []Entry{
				{
					Value:    "coffee maker",
					Synonyms: []string{"coffee maker", "coffee"},
				},
				{
					Value:    "thermostat",
					Synonyms: []string{"thermostat", "heat", "air conditioning"},
				},
			},
		})

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestAddEntries(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("POST", c.buildUrl("entities/6d6b7d50-7510-4fec-927b-ac3c3aaff009/entries", nil), tc.responder)

		err := c.AddEntries(
			"6d6b7d50-7510-4fec-927b-ac3c3aaff009",
			[]Entry{
				{
					Value:    "coffee maker",
					Synonyms: []string{"coffee maker", "coffee"},
				},
				{
					Value:    "thermostat",
					Synonyms: []string{"thermostat", "heat", "air conditioning"},
				},
			},
		)

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestUpdateEntities(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("PUT", c.buildUrl("entities", nil), tc.responder)

		err := c.UpdateEntities([]Entity{
			{
				Name: "appliances",
				Entries: []Entry{
					{
						Value:    "coffee maker",
						Synonyms: []string{"coffee maker", "coffee"},
					},
					{
						Value:    "thermostat",
						Synonyms: []string{"thermostat", "heat", "air conditioning"},
					},
					{
						Value:    "computer",
						Synonyms: []string{"pc", "laptop", "sammy"},
					},
				},
			},
		},
		)

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestUpdateEntity(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("PUT", c.buildUrl("entities/1de251bf-46a6-4056-af9c-96b6ca89dfd0", nil), tc.responder)

		err := c.UpdateEntity("1de251bf-46a6-4056-af9c-96b6ca89dfd0", Entity{
			Name: "appliances",
			Entries: []Entry{
				{
					Value:    "coffee maker",
					Synonyms: []string{"coffee maker", "coffee"},
				},
				{
					Value:    "thermostat",
					Synonyms: []string{"thermostat", "heat", "air conditioning"},
				},
				{
					Value:    "computer",
					Synonyms: []string{"pc", "laptop", "sammy"},
				},
			},
		},
		)

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestUpdateEntries(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("PUT", c.buildUrl("entities/1de251bf-46a6-4056-af9c-96b6ca89dfd0/entries", nil), tc.responder)

		err := c.UpdateEntries("1de251bf-46a6-4056-af9c-96b6ca89dfd0", []Entry{
			{
				Value:    "coffee maker",
				Synonyms: []string{"coffee maker", "coffee"},
			},
			{
				Value:    "thermostat",
				Synonyms: []string{"thermostat", "heat", "air conditioning"},
			},
			{
				Value:    "computer",
				Synonyms: []string{"pc", "laptop", "sammy"},
			},
		},
		)

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestDeleteEntity(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("DELETE", c.buildUrl("entities/1de251bf-46a6-4056-af9c-96b6ca89dfd0", nil), tc.responder)

		err := c.DeleteEntity("1de251bf-46a6-4056-af9c-96b6ca89dfd0")

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestDeleteEntries(t *testing.T) {
	c, err := NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
		httpmock.RegisterResponder("DELETE", c.buildUrl("entities/1de251bf-46a6-4056-af9c-96b6ca89dfd0/entries", nil), tc.responder)

		err := c.DeleteEntries("1de251bf-46a6-4056-af9c-96b6ca89dfd0", []string{"coffee maker", "thermostat"})

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
