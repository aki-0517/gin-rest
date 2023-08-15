package handler

import (
	"gin-rest/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Player interface {
	Create(c *gin.Context)
	Find(c *gin.Context)
	FindAll(c *gin.Context)
}

type playerHandler struct {
	usecase usecase.Player
}

func NewPlayer(u usecase.Player) Player {
	return &playerHandler{u}
}

type CreateRequestParam struct {
	Name string `json:"task" binding:"required,max=60"`
}

func (t *playerHandler) Create(c *gin.Context) {
	var req CreateRequestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.usecase.Create(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, nil)
}

type FindRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (t *playerHandler) Find(c *gin.Context) {
	var req FindRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := t.usecase.Find(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (t *playerHandler) FindAll(c *gin.Context) {
	res, err := t.usecase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}