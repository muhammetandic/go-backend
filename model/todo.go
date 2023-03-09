package model

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Todo string `json:"todo"`
    IsCompleted bool `json:"isCompleted"`
    Description string `json:"description"`
}
