package controllers

import (
	"dreamus/internal/db"
	"dreamus/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BuildController struct {
}

func NewBuildController() *BuildController {
	return &BuildController{}
}

func (ctrl *BuildController) Create(c *gin.Context) {
	var build models.Build
	if err := c.ShouldBindJSON(&build); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&build).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, build)
}

func (ctrl *BuildController) Get(c *gin.Context) {

}

func (ctrl *BuildController) Update(c *gin.Context) {

}

func (ctrl *BuildController) Delete(c *gin.Context) {

}
