package apiai

import (
	"fmt"
	"net/url"
)

const baseUrl = "https://api.api.ai/v1/"
const defaultVersion = "20150910"
const defaultQueryLang = "en"
const defaultSpeechLang = "en-US"

var queryLang = []string{"pt-BR", "zh-HK", "zh-CN", "zh-TW", "en", "nl", "fr", "de", "it", "ja", "ko", "pt", "ru", "es", "uk"}
var speechLang = []string{"en-US", "en-AU", "en-CA", "en-GB", "en-IN", "ru-RU", "de-DE", "es-ES", "pt-PT", "pt-BR", "zh-CN", "zh-TW", "zh-HK", "ja-JP", "fr-FR"}

type ClientConfig struct {
	Token      string //a9a9a9a9a9a9aa9a9a9a9a9a9a9a9a9a
	Version    string //YYYYMMDD
	QueryLang  string
	SpeechLang string
}

type ApiClient struct {
	config *ClientConfig
}

type Client interface {
	Query(Query) (*QueryResponse, error)
	Tts(text string) (string, error)
	GetContext(name string) (*Context, error)
	CreateContext(context Context) error
	DeleteContext(name string) error
	GetContexts() ([]Context, error)
	DeleteContexts() error
	GetEntities() ([]EntityDescription, error)
	UpdateEntities(entities []Entity) error
	GetEntity(idOrName string) (*Entity, error)
	CreateEntity(entity Entity) (*CreationResponse, error)
	UpdateEntity(idOrName string, entity Entity) error
	DeleteEntity(idOrName string) error
	AddEntries(idOrName string, entries []Entry) error
	UpdateEntries(idOrName string, entries []Entry) error
	DeleteEntries(idOrName string, entries []string) error
	GetIntent(id string) (*Intent, error)
	CreateIntent(intent Intent) (*CreationResponse, error)
	UpdateIntent(id string, intent Intent) error
	DeleteIntent(id string) error
	GetIntents() ([]IntentDescription, error)
}

func NewClient(conf *ClientConfig) (*ApiClient, error) {
	if conf.Token == "" {
		return nil, fmt.Errorf("%v", "You have to provide a Token")
	}
	if conf.Version == "" {
		conf.Version = defaultVersion
	}
	if conf.QueryLang == "" {
		conf.QueryLang = defaultQueryLang
	}
	if conf.SpeechLang == "" {
		conf.SpeechLang = defaultSpeechLang
	}
	if !languageAvailable(conf.QueryLang, queryLang) {
		return nil, fmt.Errorf("%v", "You have to provide a valid query language, see https://docs.api.ai/docs/languages")
	}
	if !languageAvailable(conf.SpeechLang, speechLang) {
		return nil, fmt.Errorf("%v", "You have to provide a valid speech language, see https://docs.api.ai/docs/tts#headers")
	}

	return &ApiClient{conf}, nil
}

func languageAvailable(inputLang string, languages []string) bool {
	for _, lang := range languages {
		if lang == inputLang {
			return true
		}
	}
	return false
}

func (c *ApiClient) buildUrl(endpoint string, params map[string]string) string {
	u := baseUrl + endpoint + "?v=" + c.config.Version
	if params != nil {
		for i, v := range params {
			u += "&" + i + "=" + url.QueryEscape(v)
		}
	}
	return u
}
