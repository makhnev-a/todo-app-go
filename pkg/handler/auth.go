package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/makhnev-a/todo-app-go"
	"net/http"
)

func (h *Handler) signIn(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {

}
