package apiai

type ApiTestClient struct {
	config                 *ClientConfig
	QueryResponse          *QueryResponse
	TtsResponse            string
	GetContextResponse     *Context
	GetContextsResponse    []Context
	GetEntitiesResponse    []EntityDescription
	UpdateEntitiesResponse []Entity
	GetEntityResponse      *Entity
	CreateEntityResponse   *CreationResponse
	GetIntentResponse      *Intent
	CreateIntentResponse   *CreationResponse
	GetIntentsResponse     []IntentDescription
	Err                    error
}

func (tc *ApiTestClient) Query(Query) (*QueryResponse, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.QueryResponse, nil
}
func (tc *ApiTestClient) Tts(text string) (string, error) {
	if tc.Err != nil {
		return "", tc.Err
	}

	return tc.TtsResponse, nil
}
func (tc *ApiTestClient) GetContext(name string) (*Context, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetContextResponse, nil
}
func (tc *ApiTestClient) CreateContext(context *Context) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) DeleteContext(name string) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) GetContexts() ([]Context, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetContextsResponse, nil
}
func (tc *ApiTestClient) DeleteContexts() error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) GetEntities() ([]EntityDescription, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetEntitiesResponse, nil
}
func (tc *ApiTestClient) UpdateEntities(entities []Entity) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) GetEntity(idOrName string) (*Entity, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetEntityResponse, nil
}
func (tc *ApiTestClient) CreateEntity(entity Entity) (*CreationResponse, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.CreateEntityResponse, nil
}
func (tc *ApiTestClient) UpdateEntity(idOrName string, entity Entity) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) DeleteEntity(idOrName string) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) AddEntries(idOrName string, entries []Entry) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) UpdateEntries(idOrName string, entries []Entry) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) DeleteEntries(idOrName string, entries []string) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) GetIntent(id string) (*Intent, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetIntentResponse, nil
}
func (tc *ApiTestClient) CreateIntent(intent Intent) (*CreationResponse, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.CreateIntentResponse, nil
}
func (tc *ApiTestClient) UpdateIntent(id string, intent Intent) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) DeleteIntent(id string) error {
	if tc.Err != nil {
		return tc.Err
	}

	return nil
}
func (tc *ApiTestClient) GetIntents() ([]IntentDescription, error) {
	if tc.Err != nil {
		return nil, tc.Err
	}

	return tc.GetIntentsResponse, nil
}
