package models

// RequestBody defines structure of request to execute code online
type RequestBody struct {
	ProgramID string `json:"id"`
	Input     string `json:"input"`
}
