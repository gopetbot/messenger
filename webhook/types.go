package webhook

type intent struct {
	DisplayName string `json:"displayName"`
}

type queryResult struct {
	Intent intent `json:"intent"`
}

type text struct {
	Text []string `json:"text"`
}

type message struct {
	Text text `json:"text"`
}
type webhookRequest struct {
	Session     string      `json:"session"`
	ResponseID  string      `json:"responseId"`
	QueryResult queryResult `json:"queryResult"`
}

type webhookResponse struct {
	FulfillmentMessages []message `json:"fulfillmentMessages"`
}

