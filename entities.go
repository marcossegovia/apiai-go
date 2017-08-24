package apiai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EntityDescription struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Preview string `json:"preview"`
}

type Entry struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

type Entity struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Entries            []Entry `json:"entries"`
	IsEnum             bool    `json:"isEnum"`
	AutomatedExpansion bool    `json:"automatedExpansion"`
}

func (c *ApiClient) GetEntities() ([]EntityDescription, error) {
	resp, err := c.getResponseApiai(http.MethodGet, "entities", nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entities []EntityDescription
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&entities)
		if err != nil {
			return nil, err
		}
		return entities, nil
	default:
		return nil, fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) GetEntity(idOrName string) (*Entity, error) {
	resp, err := c.getResponseApiai(http.MethodGet, "entities/"+idOrName, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var entity *Entity
	switch resp.StatusCode {
	case http.StatusOK:
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&entity)
		if err != nil {
			return nil, err
		}
		return entity, nil
	default:
		return nil, fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) CreateEntity(entity Entity) (*CreationResponse, error) {

	resp, err := c.getResponseApiai(http.MethodPost, "entities", nil, entity)
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
		return nil, fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) AddEntries(idOrName string, entries []Entry) error {

	resp, err := c.getResponseApiai(http.MethodPost, "entities/"+idOrName+"/entries", nil, entries)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntities(entities []Entity) error {

	resp, err := c.getResponseApiai(http.MethodPut, "entities", nil, entities)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntity(idOrName string, entity Entity) error {

	resp, err := c.getResponseApiai(http.MethodPut, "entities/"+idOrName, nil, entity)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) UpdateEntries(idOrName string, entries []Entry) error {

	resp, err := c.getResponseApiai(http.MethodPut, "entities/"+idOrName+"/entries", nil, entries)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) DeleteEntity(idOrName string) error {

	resp, err := c.getResponseApiai(http.MethodDelete, "entities/"+idOrName, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}

func (c *ApiClient) DeleteEntries(idOrName string, entries []string) error {

	resp, err := c.getResponseApiai(http.MethodDelete, "entities/"+idOrName+"/entries", nil, entries)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf(ErrorApiAiRequestMsg, resp.StatusCode)
	}
}
