package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/src/application/converter"
	"github.com/go-rest-api/src/application/dto"
	"github.com/go-rest-api/src/application/form"
	"github.com/go-rest-api/src/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	PORT = 8080
	DSN  = "root:root@tcp(127.0.0.1:3306)/simplize_dev?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	r := gin.Default()

	// to-do-item
	r.GET("/v1/items", FindAll(db))
	r.GET("/v1/items/:id", FindById(db))
	r.POST("/v1/items", CreateItem(db))
	r.PATCH("/v1/items/:id", UpdateById(db))
	r.DELETE("/v1/items/:id", DeleteById(db))

	r.Run(fmt.Sprintf(":%d", PORT))
}

func FindAll(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
		size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))

		var modelList []domain.TodoItem

		db = db.Where("status = ?", 1)

		var total int64

		if err := db.Table(domain.TodoItem{}.TableName()).Count(&total).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := db.Order("id DESC").
			Offset(page * size).
			Limit(size).
			Find(&modelList).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		dtoList := []dto.TodoItemDto{}

		for _, v := range modelList {
			dto := converter.ConvertToDtdto(v)
			dtoList = append(dtoList, dto)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success",
			"total":   total,
			"data":    dtoList,
		})
	}
}

func FindById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var model domain.TodoItem

		if err := db.Where("id = ?", id).First(&model).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		dto := converter.ConvertToDtdto(model)

		ctx.JSON(http.StatusOK, gin.H{
			"data": dto,
		})
	}
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var request form.TodoItemCreate

		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		model := converter.ConvertCreateToModel(request)

		if err := db.Create(&model).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		dto := converter.ConvertToDtdto(model)

		ctx.JSON(http.StatusOK, gin.H{
			"data": dto.Id,
		})
	}
}

func UpdateById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		var request form.TodoItemUpdate

		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		model := converter.ConvertUpdateToModel(request)

		if err := db.Where("id = ?", id).Updates(&model).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := db.Table(domain.TodoItem{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
