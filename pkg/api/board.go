package api

import (
	"context"
	"kanban/pkg/database"
	"kanban/pkg/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var boardCollection *mongo.Collection = database.GetCollection(database.DB, "boards")
var validate = validator.New()

func GetAllBoards() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Query("userId")
		var boardList []models.BoardResponse
		defer cancel()

		results, err := boardCollection.Find(ctx, bson.M{"user_id": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var board models.Board
			boardTask := make(map[string]models.BoardTaskResponse)
			if err = results.Decode(&board); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
			}

			for _, boardtask := range board.BoardTask {
				boardTask[boardtask.Id] = models.BoardTaskResponse{
					TaskName: boardtask.TaskName,
					TaskList: boardtask.TaskList,
					TagColor: boardtask.TagColor,
				}
			}

			responseBoard := models.BoardResponse{
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

func GetBoardById(ctx context.Context, boardId string) (models.Board, error) {
	var board models.Board

	objId, _ := primitive.ObjectIDFromHex(boardId)
	filterBoardId := bson.M{"_id": objId}

	err := boardCollection.FindOne(ctx, filterBoardId).Decode(&board)
	if err != nil {
		return models.Board{}, err
	}
	return board, nil
}

func CreateNewBoard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Query("userId")
		var req models.Board
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if validationErr := validate.Struct(&req); validationErr != nil {
			c.JSON(http.StatusBadRequest, validationErr.Error())
			return
		}

		newBoard := models.Board{
			Id:        primitive.NewObjectID(),
			Name:      req.Name,
			UserId:    userId,
			BoardTask: []models.BoardTask{},
		}

		_, err := boardCollection.InsertOne(ctx, newBoard)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, newBoard)
	}
}

func CreateNewBoardTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var board models.Board
		var req models.CreateBoardTaskRequest
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

		responseBoardTask := models.BoardTask{
			Id:       strconv.Itoa(len(board.BoardTask) + 1),
			TaskName: req.TaskName,
			TaskList: []models.Task{},
			TagColor: req.TagColor,
		}

		board.BoardTask = append(board.BoardTask, responseBoardTask)

		updatedBoardTask := bson.M{
			"$set": bson.M{
				"board_task": board.BoardTask,
			},
		}

		_, err = boardCollection.UpdateOne(
			ctx,
			bson.M{"id": board.Id},
			updatedBoardTask,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, board)
	}
}
