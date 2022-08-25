package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Board struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"boardName"`
	UserId    string             `bson:"user_id" json:"userId"`
	BoardTask []BoardTask        `bson:"board_task" json:"boardTask"`
}

type BoardTask struct {
	Id       string `bson:"board_task_id" json:"boardTaskId"`
	TaskName string `bson:"task_name" json:"taskName"`
	TaskList []Task `bson:"task_list" json:"taskList"`
	TagColor string `bson:"tag_color" json:"tagColor"`
}

type BoardResponse struct {
	Id        primitive.ObjectID           `bson:"_id" json:"id"`
	Name      string                       `bson:"name" json:"boardName"`
	UserId    string                       `bson:"user_id" json:"userId"`
	BoardTask map[string]BoardTaskResponse `bson:"board_task" json:"boardTask"`
}

type BoardTaskResponse struct {
	TaskName string `bson:"task_name" json:"taskName"`
	TaskList []Task `bson:"task_list" json:"taskList"`
	TagColor string `bson:"tag_color" json:"tagColor"`
}

type CreateBoardTaskRequest struct {
	TaskName string `bson:"task_name" json:"taskName"`
	TagColor string `bson:"tag_color" json:"tagColor"`
}
