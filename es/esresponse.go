package es

// DialogflowResponse represents the structure of the response to be sent
type Response struct {
	FulfillmentText     string                 `json:"fulfillmentText"`
	FulfillmentMessages []FulfillmentMessage   `json:"fulfillmentMessages"`
	Source              string                 `json:"source"`
	Payload             map[string]interface{} `json:"payload"`
	OutputContexts      []Context              `json:"outputContexts"`
	FollowUpEventInput  EventInput             `json:"followupEventInput"`
	SessionEntityTypes  []SessionEntityTypes   `json:"sessionEntityTypes"`
}

type Context struct {
	Name          string                 `json:"name"`
	LifespanCount int                    `json:"lifespanCount"`
	Parameters    map[string]interface{} `json:"parameters"`
}

type EventInput struct {
	Name         string                 `json:"name"`
	Parameters   map[string]interface{} `json:"parameters"`
	LanguageCode string                 `json:"languageCode"`
}

type EntityOverrideMode string

const (
	ENTITY_OVERRIDE_MODE_UNSPECIFIED EntityOverrideMode = "ENTITY_OVERRIDE_MODE_UNSPECIFIED"
	ENTITY_OVERRIDE_MODE_OVERRIDE    EntityOverrideMode = "ENTITY_OVERRIDE_MODE_OVERRIDE"
	ENTITY_OVERRIDE_MODE_SUPPLEMENT  EntityOverrideMode = "ENTITY_OVERRIDE_MODE_SUPPLEMENT"
)

type SessionEntityTypes struct {
	Name               string             `json:"name"`
	EntityOverrideMode EntityOverrideMode `json:"entityOverrideMode"`
	Entities           []Entity           `json:"entities"`
}

type Entity struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}
