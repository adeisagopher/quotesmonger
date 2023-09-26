package controllers

import (
	"net/http"

	"github.com/adeisagopher/quotesmonger/model"
	"github.com/adeisagopher/quotesmonger/utility"
	"github.com/gin-gonic/gin"
)

// Return all quotes stored in the database.
func GetQuotes(ctx *gin.Context) {

	var quotes []model.Quotes

	utility.DB.Find(&quotes)

	ctx.JSON(http.StatusOK, gin.H{
		"quotes": quotes,
	})

}

// Find and return a single quote by ID.
func GetQuoteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var quote []model.Quotes

	utility.DB.First(&quote, id)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})
}

// Find and return quotes by Category
func GetQuoteByCategory(ctx *gin.Context) {

	category := ctx.Query("name")

	var quote model.Quotes

	utility.DB.Where("category = ?", category).Find(&quote)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})

}

// Post a quote to the database.
func PostQuotes(ctx *gin.Context) {
	var Body struct {
		Quote    string
		Author   string
		Category string
	}

	ctx.Bind(&Body)

	quote := model.Quotes{Quote: Body.Quote, Author: Body.Author, Category: Body.Category}

	utility.DB.Create(&quote)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})

}
