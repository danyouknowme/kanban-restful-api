package api

import (
	"context"
	"kanban/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var req models.CreateTaskRequest
		userId := c.Query("userId")
		boardId := c.Query("boardId")
		defer cancel()

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		board, err := GetBoardById(ctx, boardId)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		if board.UserId != userId {
			c.JSON(http.StatusForbidden, gin.H{"message": "you don't allow to do this"})
			return
		}

		var newSubTask []models.SubTask
		for _, subTask := range req.Subtasks {
			newSubTask = append(newSubTask, models.SubTask{
				Name:   subTask,
				IsDone: false,
			})
		}

		task := models.Task{
			Id:          uuid.New().String(),
			Name:        req.Name,
			Description: req.Description,
			Subtasks:    newSubTask,
		}

		for _, boardTask := range board.BoardTask {
			if boardTask.Id == req.BoardTaskId {
				boardTask.TaskList = append(boardTask.TaskList, task)
				boardTaskId, _ := strconv.Atoi(boardTask.Id)
				board.BoardTask[boardTaskId-1] = boardTask
				break
			}
		}

		updatedBoardTask := bson.M{
			"$set": bson.M{
				"board_task": board.BoardTask,
			},
		}

		_, err = boardCollection.UpdateOne(
			ctx,
			bson.M{"_id": board.Id},
			updatedBoardTask,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, board)
	}
}
