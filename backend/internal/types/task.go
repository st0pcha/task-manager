package types

type TaskResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"is_done"`
}
