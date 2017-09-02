package apiai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Context struct {
	Name     string            `json:"name"`
	Lifespan int               `json:"lifespan"`
	Params   map[string]string `json:"parameters"`
}

func (c *ApiClient) GetContexts(sessionId string) ([]Context, error) {

	resp, err := c.getApiaiResponse(http.MethodGet, "contexts", map[string]string{"sessionId": sessionId}, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var contexts []Context
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&contexts)
		if err != nil {
			return nil, err
		}
		return contexts, nil
	default:
		return nil, fmt.Errorf(DefaultErrorMsg, resp.StatusCode)
	}
}

func (c *ApiClient) GetContext(name, sessionId string) (*Context, error) {
	resp, err := c.getApiaiResponse(http.MethodGet, "contexts/"+url.QueryEscape(name), map[string]string{"sessionId": sessionId}, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var context *Context
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&context)
		if err != nil {
			return nil, err
		}
		return context, nil
	default:
		return nil, fmt.Errorf(DefaultErrorMsg, resp.StatusCode)
	}
}

func (c *ApiClient) CreateContext(context Context, sessionId string) error {

	resp, err := c.getApiaiResponse(http.MethodPost, "contexts", map[string]string{"sessionId": sessionId}, context)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(DefaultErrorMsg, resp.StatusCode)
	}
}

func (c *ApiClient) DeleteContexts(sessionId string) error {
	resp, err := c.getApiaiResponse(http.MethodDelete, "contexts", map[string]string{"sessionId": sessionId}, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(DefaultErrorMsg, resp.StatusCode)
	}
}

func (c *ApiClient) DeleteContext(name, sessionId string) error {

	resp, err := c.getApiaiResponse(http.MethodDelete, "contexts/"+url.QueryEscape(name), map[string]string{"sessionId": sessionId}, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(DefaultErrorMsg, resp.StatusCode)
	}
}
