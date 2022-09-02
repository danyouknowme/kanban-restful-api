package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"kanban/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		oauthState := utils.GenerateStateOauthCookie(c.Writer)
		u := utils.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
		http.Redirect(c.Writer, c.Request, u, http.StatusTemporaryRedirect)
	}
}

func GoogleCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		oauthState, _ := c.Request.Cookie("oauthstate")
		state := c.Request.FormValue("state")
		code := c.Request.FormValue("code")
		c.Writer.Header().Add("content-type", "application/json")

		if state != oauthState.Value {
			http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
			return
		}

		token, err := utils.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		response, err := http.Get(utils.OauthGoogleUrlAPI + token.AccessToken)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		// Parse user data JSON Object
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Fprintln(c.Writer, string(contents))
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
