package controllers

import (
	"net/http"

	"github.com/devesh/gin-gorm-crud/models"
	"github.com/gin-gonic/gin"
)

// ----- for validation of  user input while creating the book
type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author"  binding:"required"`
	Description string `json:"description" binding:"required"`
}

// ----- validation for user input while upadating the book
type UpdateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author"  binding:"required"`
	Description string `json:"description" binding:"required"`
}

// -------- handler for get all books avilable in database .
func FindBooks(c *gin.Context) {
	var books []models.Book // database acessing variable for books find
	models.DB.Find(&books)  //----- finding from database

	c.JSON(http.StatusOK, gin.H{"data ": books}) //  fethched data returning to client
}

// -------- creating the book and posting to database
func CreateBook(c *gin.Context) {
	// --validation input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//---create book
	book := models.Book{Title: input.Title, Author: input.Author, Description: input.Description}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// ----- Find a book from books database .
//
//	----- /book/:id
func FindBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// -------- updating the books by perticuler id
func UpdateBook(c *gin.Context) {
	//----- first we have to find a book by id
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found "})
		return
	}

	//------ after we have to validate the user input whilre updating the books

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data ": book})

}

///-------------deleat books by id

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})

	models.DB.Delete(&book)
}
