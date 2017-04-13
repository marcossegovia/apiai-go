package apiai

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetContexts(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken", SessionId: "123454321"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse []Context
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `[
{
 "name": "Play game",
 "parameters": {
  "any": "value",
  "number": "1"
 }
},
{
 "name": "Coffee time",
 "parameters": {
  "temperature": "cold"
 }
}
]`),
			expectedResponse: []Context{
				{
					Name: "Play game",
					Params: map[string]string{
						"any":    "value",
						"number": "1",
					},
				},
				{
					Name: "Coffee time",
					Params: map[string]string{
						"temperature": "cold",
					},
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
		httpmock.RegisterResponder("GET", c.buildUrl("contexts", map[string]string{
			"SessionId": c.config.SessionId,
		}), tc.responder)

		r, err := c.GetContexts()

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestGetContext(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken", SessionId: "123454321"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse *Context
		expectedError    error
	}{
		{
			description: "api ai success, no errors",
			responder: httpmock.NewStringResponder(200, `{
  "name": "Coffee time",
  "parameters": {
   "type-1": "long",
   "type-2": "short",
   "temperature-1": "hot",
   "temperature-2": "cold"
  }
}`),
			expectedResponse: &Context{
				Name: "Coffee time",
				Params: map[string]string{
					"type-1":        "long",
					"type-2":        "short",
					"temperature-1": "hot",
					"temperature-2": "cold",
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
		httpmock.RegisterResponder("GET", c.buildUrl("contexts/"+url.QueryEscape("Coffee time"), map[string]string{
			"SessionId": c.config.SessionId,
		}), tc.responder)

		r, err := c.GetContext("Coffee time")

		assert.Equal(r, tc.expectedResponse, tc.description)
		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestCreateContext(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken", SessionId: "123454321"})
	if err != nil {
		t.FailNow()
	}
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		description      string
		responder        httpmock.Responder
		expectedResponse *Context
		expectedError    error
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
		httpmock.RegisterResponder("POST", c.buildUrl("contexts", map[string]string{
			"SessionId": c.config.SessionId,
		}), tc.responder)

		err := c.CreateContext(Context{
			Name: "Coffee time",
			Params: map[string]string{
				"type-1":        "long",
				"type-2":        "short",
				"temperature-1": "hot",
				"temperature-2": "cold",
			},
		})

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestDeleteContexts(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken", SessionId: "123454321"})
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
		httpmock.RegisterResponder("DELETE", c.buildUrl("contexts", map[string]string{
			"SessionId": c.config.SessionId,
		}), tc.responder)

		err := c.DeleteContexts()

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}

func TestDeleteContext(t *testing.T) {
	c, err := NewClient(&ClientConfig{Token: "fakeToken", SessionId: "123454321"})
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
		httpmock.RegisterResponder("DELETE", c.buildUrl("contexts/"+url.QueryEscape("Coffee time"), map[string]string{
			"SessionId": c.config.SessionId,
		}), tc.responder)

		err := c.DeleteContext("Coffee time")

		assert.Equal(err, tc.expectedError, tc.description)

		httpmock.Reset()
	}
}
