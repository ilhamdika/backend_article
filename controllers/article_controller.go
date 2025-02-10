package controllers

import (
	"backend_article/config"
	"backend_article/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Post
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&article)
	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit := 10
	offset := 0
	var err1, err2 error

	if limitStr != "" {
		limit, err1 = strconv.Atoi(limitStr)
	}
	if offsetStr != "" {
		offset, err2 = strconv.Atoi(offsetStr)
	}

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Limit dan offset harus berupa angka"})
		return
	}

	var articles []models.Post
	config.DB.Order("created_date DESC").Limit(limit).Offset(offset).Find(&articles)

	c.JSON(http.StatusOK, articles)
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Post
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Post
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&article)
	c.JSON(http.StatusOK, article)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Post
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	config.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
}
