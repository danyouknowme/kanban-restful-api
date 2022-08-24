package api

import (
	"context"
	"kanban/pkg/database"
	"kanban/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var boardCollection *mongo.Collection = database.GetCollection(database.DB, "boards")

func GetAllBoards() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var boardList []ResponseBoard
		defer cancel()

		results, err := boardCollection.Find(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var board models.Board
			boardTask := make(map[string]BoardTask)
			if err = results.Decode(&board); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}

			for _, boardtask := range board.BoardTask {
				boardTask[boardtask.Id] = BoardTask{
					TaskName: boardtask.TaskName,
					TaskList: boardtask.TaskList,
					TagColor: boardtask.TagColor,
				}
			}

			responseBoard := ResponseBoard{
				Id:        board.Id,
				Name:      board.Name,
				UserId:    board.UserId,
				BoardTask: boardTask,
			}

			boardList = append(boardList, responseBoard)
		}
		c.JSON(http.StatusOK, boardList)
	}
}

type ResponseBoard struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id"`
	Name      string               `bson:"name" json:"boardName"`
	UserId    string               `bson:"user_id" json:"userId"`
	BoardTask map[string]BoardTask `bson:"board_task" json:"boardTask"`
}

type BoardTask struct {
	TaskName string        `bson:"task_name" json:"taskName"`
	TaskList []models.Task `bson:"task_list" json:"taskList"`
	TagColor string        `bson:"tag_color" json:"tagColor"`
}
