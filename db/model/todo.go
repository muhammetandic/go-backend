package model

type Todo struct {
	Model
	Task        string `json:"task" binding:"required"`
	IsCompleted *bool  `json:"isCompleted" binding:"required"`
	Description string `json:"description"`
}
