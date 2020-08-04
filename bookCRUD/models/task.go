// models/book.go

package models

type Task struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
}
