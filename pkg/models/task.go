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

type CreateTaskRequest struct {
	BoardTaskId string   `json:"boardTaskId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Subtasks    []string `json:"subtasks"`
}

type EditTaskListSameColumnRequest struct {
	BoardTaskId string `json:"boardTaskId"`
	TaskList    []Task `json:"taskList"`
}

type EditTaskListDifferentColumnRequest struct {
	SourceId         string `json:"sourceId"`
	DestinationId    string `json:"destinationId"`
	SourceItems      []Task `json:"sourceItems"`
	DestinationItems []Task `json:"destinationItems"`
}
