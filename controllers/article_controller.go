package controllers

import (
	"backend_article/config"
	"backend_article/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article models.Post

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if article.Status != "draft" {
		if len(strings.TrimSpace(article.Title)) < 20 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Title minimal 20 karakter",
			})
			return
		}
	
		if len(strings.TrimSpace(article.Content)) < 200 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Content minimal 200 karakter",
			})
			return
		}
	
		if len(strings.TrimSpace(article.Category)) < 3 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Category minimal 3 karakter",
			})
			return
		}
	}

	validStatuses := map[string]bool{"publish": true, "draft": true, "thrash": true}
	if !validStatuses[article.Status] {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error": "Status harus 'publish', 'draft', atau 'thrash'",
		})
		return
	}

	config.DB.Create(&article)
	c.JSON(http.StatusOK, article)
}


func DraftArticle(c *gin.Context) {
	var article models.Post

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	article.Status = "draft"

	config.DB.Create(&article)
	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")
	status := c.Query("status") 

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

	query := config.DB.Order("created_date DESC").Limit(limit).Offset(offset)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Find(&articles)

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

	if article.Status != "draft" {
		if len(strings.TrimSpace(article.Title)) < 20 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Title minimal 20 karakter",
			})
			return
		}
	
		if len(strings.TrimSpace(article.Content)) < 200 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Content minimal 200 karakter",
			})
			return
		}
	
		if len(strings.TrimSpace(article.Category)) < 3 {
			c.JSON(http.StatusOK, gin.H{
				"status": 400,
				"error": "Category minimal 3 karakter",
			})
			return
		}
	}

	validStatuses := map[string]bool{"publish": true, "draft": true, "thrash": true}
	if !validStatuses[article.Status] {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error": "Status harus 'publish', 'draft', atau 'thrash'",
		})
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

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	
	var requestBody struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	validStatuses := map[string]bool{"publish": true, "draft": true, "thrash": true}
	if !validStatuses[requestBody.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status harus 'publish', 'draft', atau 'thrash'"})
		return
	}

	var article models.Post
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artikel tidak ditemukan"})
		return
	}

	article.Status = requestBody.Status
	config.DB.Save(&article)

	c.JSON(http.StatusOK, gin.H{"message": "Status berhasil diperbarui", "article": article})
}
