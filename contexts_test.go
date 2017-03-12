package apiai

import (
        "fmt"
        "testing"
        "net/http"
        "net/url"

        "github.com/jarcoal/httpmock"
        "github.com/stretchr/testify/assert"
)

func TestGetContexts(t *testing.T) {
        var c = NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
  "parameters": [
    {
      "name": "option-1",
      "value": "yes"
    },
    {
      "name": "option-2",
      "value": "no"
    }
  ]
},
{
  "name": "Coffee time",
  "parameters": [
    {
      "name": "type-1",
      "value": "long"
    },
    {
      "name": "type-2",
      "value": "short"
    },
    {
      "name": "temperature-1",
      "value": "hot"
    },
    {
      "name": "temperature-2",
      "value": "cold"
    }
  ]
}
]`),
                        expectedResponse: []Context{
                                {
                                        Name: "Play game",
                                        Params: []ContextParameter{
                                                {
                                                        Name:  "option-1",
                                                        Value: "yes",
                                                },
                                                {
                                                        Name:  "option-2",
                                                        Value: "no",
                                                },
                                        },
                                },
                                {
                                        Name: "Coffee time",
                                        Params: []ContextParameter{
                                                {
                                                        Name:  "type-1",
                                                        Value: "long",
                                                },
                                                {
                                                        Name:  "type-2",
                                                        Value: "short",
                                                },
                                                {
                                                        Name:  "temperature-1",
                                                        Value: "hot",
                                                },
                                                {
                                                        Name:  "temperature-2",
                                                        Value: "cold",
                                                },
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
                        "sessionId": c.config.sessionId,
                }), tc.responder)

                r, err := c.GetContexts()

                assert.Equal(r, tc.expectedResponse, tc.description)
                assert.Equal(err, tc.expectedError, tc.description)

                httpmock.Reset()
        }
}

func TestGetContext(t *testing.T) {
        var c = NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
  "parameters": [
    {
      "name": "type-1",
      "value": "long"
    },
    {
      "name": "type-2",
      "value": "short"
    },
    {
      "name": "temperature-1",
      "value": "hot"
    },
    {
      "name": "temperature-2",
      "value": "cold"
    }
  ]
}`),
                        expectedResponse: &Context{
                                Name: "Coffee time",
                                Params: []ContextParameter{
                                        {
                                                Name:  "type-1",
                                                Value: "long",
                                        },
                                        {
                                                Name:  "type-2",
                                                Value: "short",
                                        },
                                        {
                                                Name:  "temperature-1",
                                                Value: "hot",
                                        },
                                        {
                                                Name:  "temperature-2",
                                                Value: "cold",
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
                httpmock.RegisterResponder("GET", c.buildUrl("contexts/"+url.QueryEscape("Coffee time"), map[string]string{
                        "sessionId": c.config.sessionId,
                }), tc.responder)

                r, err := c.GetContext("Coffee time")

                assert.Equal(r, tc.expectedResponse, tc.description)
                assert.Equal(err, tc.expectedError, tc.description)

                httpmock.Reset()
        }
}

func TestCreateContext(t *testing.T) {
        var c = NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
                        "sessionId": c.config.sessionId,
                }), tc.responder)

                err := c.CreateContext(Context{
                        Name: "Coffee time",
                        Params: []ContextParameter{
                                {
                                        Name:  "type-1",
                                        Value: "long",
                                },
                                {
                                        Name:  "type-2",
                                        Value: "short",
                                },
                                {
                                        Name:  "temperature-1",
                                        Value: "hot",
                                },
                                {
                                        Name:  "temperature-2",
                                        Value: "cold",
                                },
                        },
                })

                assert.Equal(err, tc.expectedError, tc.description)

                httpmock.Reset()
        }
}

func TestDeleteContexts(t *testing.T) {
        var c = NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
                        "sessionId": c.config.sessionId,
                }), tc.responder)

                err := c.DeleteContexts()

                assert.Equal(err, tc.expectedError, tc.description)

                httpmock.Reset()
        }
}

func TestDeleteContext(t *testing.T) {
        var c = NewClient(&ClientConfig{token: "fakeToken", sessionId: "123454321"})
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
                        "sessionId": c.config.sessionId,
                }), tc.responder)

                err := c.DeleteContext("Coffee time")

                assert.Equal(err, tc.expectedError, tc.description)

                httpmock.Reset()
        }
}
