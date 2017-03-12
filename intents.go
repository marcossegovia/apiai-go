package apiai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Text        string `json:"text"`
	Meta        string `json:"meta"`
	Alias       string `json:"alias"`
	UserDefined bool   `json:"userDefined"`
}

type UserSays struct {
	Id         string `json:"id"`
	Data       []Data `json:"data"`
	IsTemplate bool   `json:"isTemplate"`
	Count      int    `json:"count"`
}

type IntentParameter struct {
	Name         string   `json:"name"`
	Value        string   `json:"value"`
	DefaultValue string   `json:"defaultValue"`
	Required     bool     `json:"required"`
	DataType     string   `json:"dataType"`
	Prompts      []string `json:"prompts"`
	IsList       bool     `json:"isList"`
}

type IntentResponse struct {
	Action           string            `json:"action"`
	ResetContexts    bool              `json:"resetContexts"`
	AffectedContexts []Context         `json:"resetContexts"`
	Params           []IntentParameter `json:"parameters"`
	Messages         []Message         `json:"messages"`
}

type CortanaCommand struct {
	NavigateOrService string `json:"navigateOrService"`
	Target            string `json:"target"`
}

type Intent struct {
	Id                    string           `json:"id"`
	Name                  string           `json:"name"`
	Auto                  bool             `json:"auto"`
	Contexts              []string         `json:"contexts"`
	Templates             []string         `json:"templates"`
	UserSays              []UserSays       `json:"userSays"`
	Responses             []IntentResponse `json:"responses"`
	Priority              int              `json:"priority"`
	WebhookUsed           bool             `json:"webhookUsed"`
	WebhookForSlotFilling bool             `json:"webhookForSlotFilling"`
	FallbackIntent        bool             `json:"fallbackIntent"`
	CortanaCommand        CortanaCommand   `json:"cortanaCommand"`
	Events                []Event          `json:"events"`
}

type IntentDescription struct {
	Id             string            `json:"id"`
	Name           string            `json:"name"`
	ContextIn      []string          `json:"contextIn"`
	ContextOut     []Context         `json:"contextOut"`
	Actions        []string          `json:"actions"`
	Params         []IntentParameter `json:"parameters"`
	Priority       int               `json:"priority"`
	FallbackIntent bool              `json:"fallbackIntent"`
}

func (c *ApiClient) GetIntents() ([]IntentDescription, error) {
	req, err := http.NewRequest("GET", c.buildUrl("intents", nil), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var intents []IntentDescription
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&intents)
		if err != nil {
			return nil, err
		}
		return intents, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) GetIntent(id string) (*Intent, error) {
	req, err := http.NewRequest("GET", c.buildUrl("intents/"+id, nil), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var intent *Intent
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&intent)
		if err != nil {
			return nil, err
		}
		return intent, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) CreateIntent(intent Intent) (*CreationResponse, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(intent)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.buildUrl("intents", nil), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cr *CreationResponse
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&cr)
		if err != nil {
			return nil, err
		}
		return cr, nil
	default:
		return nil, fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) UpdateIntent(id string, intent Intent) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(intent)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.buildUrl("intents/"+id, nil), body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}

func (c *ApiClient) DeleteIntent(id string) error {
	req, err := http.NewRequest("DELETE", c.buildUrl("intents/"+id, nil), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.config.Token)

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("apiai: wops something happens because status code is %v", resp.StatusCode)
	}
}
