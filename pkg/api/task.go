package api

import (
	"context"
	"kanban/pkg/models"
	"net/http"
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

		newSubTask := []models.SubTask{}
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

		boardTask := board.BoardTask[req.BoardTaskId]
		boardTask.TaskList = append(boardTask.TaskList, task)
		board.BoardTask[req.BoardTaskId] = boardTask

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

func EditTaskListSameColumn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var req models.EditTaskListSameColumnRequest
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

		boardTask := board.BoardTask[req.BoardTaskId]
		boardTask.TaskList = req.TaskList
		board.BoardTask[req.BoardTaskId] = boardTask

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

		c.JSON(http.StatusOK, board)
	}
}

func EditTaskListDifferentColumn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var req models.EditTaskListDifferentColumnRequest
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

		sourceBoardTask := board.BoardTask[req.SourceId]
		destBoardTask := board.BoardTask[req.DestinationId]

		sourceBoardTask.TaskList = req.SourceItems
		destBoardTask.TaskList = req.DestinationItems

		board.BoardTask[req.SourceId] = sourceBoardTask
		board.BoardTask[req.DestinationId] = destBoardTask

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

		c.JSON(http.StatusOK, board)
	}
}

func EditSubtaskStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var req models.EditSubtaskStatusRequest
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

		boardTask := board.BoardTask[req.BoardTaskId]
		for _, task := range boardTask.TaskList {
			if task.Id == req.TaskId {
				task.Subtasks[req.SubTaskIndex].IsDone = req.Status
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

		c.JSON(http.StatusOK, board)
	}
}
