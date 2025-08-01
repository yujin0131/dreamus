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

func findBuildByID(id string) (*models.Build, error) {
	var build models.Build
	if err := db.DB.First(&build, id).Error; err != nil {
		return nil, err
	}
	return &build, nil
}

func (ctrl *BuildController) Get(c *gin.Context) {
	id := c.Param("id")

	build, err := findBuildByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Build not found"})
		return
	}

	c.JSON(http.StatusOK, build)
}

func (ctrl *BuildController) Update(c *gin.Context) {
	id := c.Param("id")
	build, err := findBuildByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.BuildUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(build).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, build)

}

func (ctrl *BuildController) Delete(c *gin.Context) {
	id := c.Param("id")
	build, err := findBuildByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Delete(&build).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
