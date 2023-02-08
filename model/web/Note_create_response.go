package web

type NoteCreateResponse struct {
	Status string `json:"status"`
	Error  error  `json:"errors"`
}
