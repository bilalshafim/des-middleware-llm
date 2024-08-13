package es

// FulfillmentMessage represents a message in the fulfillmentMessages array
type FulfillmentMessage struct {
    Text Text `json:"text"`
}

// Text represents the structure within the Text field of FulfillmentMessage
type Text struct {
    Text []string `json:"text"`
}
