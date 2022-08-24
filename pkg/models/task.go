package models

type Task struct {
	Id          string    `bson:"id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Subtasks    []SubTask `bson:"subtasks" json:"subtasks"`
}

type SubTask struct {
	Name   string `bson:"name" json:"name"`
	IsDone bool   `bson:"is_done" json:"isDone"`
}
