package entity

type Todo struct {
	ID     uint   `json:"id"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}
