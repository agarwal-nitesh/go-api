package models

type ResponseStruct struct {
	Data       interface{} `json:"data"`
	Debug      string      `json:"debug,omitempty"`
	Desc       string      `json:"desc,omitempty"`
	Exception  string      `json:"exception,omitempty"`
	Message    string      `json:"message,omitempty"`
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"statusCode"`
}
