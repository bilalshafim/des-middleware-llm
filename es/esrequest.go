package es

import (
	"encoding/json"
)

// DialogflowRequest represents the structure of the request from Dialogflow
type Request struct {
	ResponseID                  string          `json:"responseId"`
	Session                     string          `json:"session"`
	QueryResult                 QueryResult     `json:"queryResult"`
	OriginalDetectIntentRequest json.RawMessage `json:"originalDetectIntentRequest"` // Assuming it could be any structure
}

// QueryResult represents the structure within the QueryResult field
type QueryResult struct {
	QueryText                 string                 `json:"queryText"`
	Action                    string                 `json:"action"`
	Parameters                map[string]interface{} `json:"parameters"`
	AllRequiredParamsPresent  bool                   `json:"allRequiredParamsPresent"`
	FulfillmentText           string                 `json:"fulfillmentText"`
	FulfillmentMessages       []FulfillmentMessage   `json:"fulfillmentMessages"`
	OutputContexts            []OutputContext        `json:"outputContexts"`
	Intent                    Intent                 `json:"intent"`
	IntentDetectionConfidence float64                `json:"intentDetectionConfidence"`
	DiagnosticInfo            map[string]interface{} `json:"diagnosticInfo"` // Assuming it could be any structure
	LanguageCode              string                 `json:"languageCode"`
}

// OutputContext represents a context in the outputContexts array
type OutputContext struct {
	Name          string                 `json:"name"`
	LifespanCount int                    `json:"lifespanCount"`
	Parameters    map[string]interface{} `json:"parameters"`
}

// Intent represents the structure within the Intent field
type Intent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
