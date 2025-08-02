package controllers

import (
	"dreamus/internal/db"
	"dreamus/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BuildController struct {
}

func NewBuildController() *BuildController {
	return &BuildController{}
}

// @Summary Create new build
// @Description Create a new build record
// @Tags builds
// @Accept json
// @Produce json
// @Param build body models.Build true "Build Info"
// @Success 201 {object} models.Build
// @Failure 400 {object} models.BadRequestError
// @Failure 500 {object} models.InternalError
// @Router /builds [post]
func (ctrl *BuildController) Create(c *gin.Context) {
	var build models.Build
	if err := c.ShouldBindJSON(&build); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestError{Message: err.Error()})
		return
	}

	if err := db.DB.Create(&build).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, build)
}

func findBuildByID(id string) (*models.Build, error) {
	var build models.Build
	if err := db.DB.First(&build, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NotFoundError{Message: "build not found"}
		}
		return nil, err
	}
	return &build, nil
}

// @Summary Get a build by ID
// @Description Get a build record by its ID
// @Tags builds
// @Produce json
// @Param id path int true "Build ID"
// @Success 200 {object} models.Build
// @Failure 404 {object} models.NotFoundError
// @Failure 500 {object} models.InternalError
// @Router /builds/{id} [get]
func (ctrl *BuildController) Get(c *gin.Context) {
	id := c.Param("id")

	build, err := findBuildByID(id)
	if err != nil {
		if notFoundErr, ok := err.(models.NotFoundError); ok {
			c.JSON(http.StatusNotFound, notFoundErr)
		} else {
			c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, build)
}

// @Summary Update a build
// @Description Update fields of an existing build
// @Tags builds
// @Accept json
// @Produce json
// @Param id path int true "Build ID"
// @Param build body models.BuildUpdateInput true "Updated Build Info"
// @Success 200 {object} models.Build
// @Failure 400 {object} models.BadRequestError
// @Failure 404 {object} models.NotFoundError
// @Failure 500 {object} models.InternalError
// @Router /builds/{id} [put]
func (ctrl *BuildController) Update(c *gin.Context) {
	id := c.Param("id")
	build, err := findBuildByID(id)
	if err != nil {
		if notFoundErr, ok := err.(models.NotFoundError); ok {
			c.JSON(http.StatusNotFound, notFoundErr)
		} else {
			c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		}
		return
	}

	var input models.BuildUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestError{Message: err.Error()})
		return
	}

	if err := db.DB.Model(build).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, build)

}

// @Summary Delete a build
// @Description Delete a build by its ID
// @Tags builds
// @Produce json
// @Param id path int true "Build ID"
// @Success 204
// @Failure 404 {object} models.NotFoundError
// @Failure 500 {object} models.InternalError
// @Router /builds/{id} [delete]
func (ctrl *BuildController) Delete(c *gin.Context) {
	id := c.Param("id")
	build, err := findBuildByID(id)
	if err != nil {
		if notFoundErr, ok := err.(models.NotFoundError); ok {
			c.JSON(http.StatusNotFound, notFoundErr)
		} else {
			c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		}
		return
	}

	if err := db.DB.Delete(&build).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalError{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
