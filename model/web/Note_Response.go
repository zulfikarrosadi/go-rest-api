package web

type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NoteResponse struct {
	Status string `json:"status"`
	Error  string `json:"errors"`
	Data   []Note `json:"data"`
}
