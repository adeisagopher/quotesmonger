package controllers

import (
	"net/http"

	"github.com/adeisagopher/quotesmonger/model"
	"github.com/adeisagopher/quotesmonger/utility"
	"github.com/gin-gonic/gin"
)

// Cors middleware.
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

// @Summary return all quotes stored in the database.
// @ID get-all-quotes
// @Produce json
// @Success 200 {arrays} model.Quotes
// @Router /quotes [get]
func GetQuotes(ctx *gin.Context) {

	var quotes []model.Quotes

	utility.DB.Find(&quotes)

	ctx.JSON(http.StatusOK, gin.H{
		"quotes": quotes,
	})

}

// @Summary get a quote by ID.
// @ID get-quote-by-id
// @Produce json
// @Param id path string true "quote ID"
// @Success 200 {arrays} model.Quotes
// @Router /quote/{id} [get]
func GetQuoteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var quote []model.Quotes

	utility.DB.First(&quote, id)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})
}

// @Summary get a quote by category.
// @ID get-quote-by-category
// @Produce json
// @Param name query string true "quote category name"
// @Success 200 {arrays} model.Quotes
// @Router /quotes/category [get]
func GetQuoteByCategory(ctx *gin.Context) {

	category := ctx.Query("name")

	var quote model.Quotes

	utility.DB.Where("category = ?", category).Find(&quote)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})

}

// @Summary post a quote to the database.
// @ID post-quote
// @Produce json
// @Param data body model.Quotes true "quote data"
// @Success 200 {object} model.Quotes
// @Router /quote [post]
func PostQuotes(ctx *gin.Context) {
	var Body struct {
		Quote    string `json:"quote"`
		Author   string `json:"author"`
		Category string `json:"category"`
	}

	ctx.Bind(&Body)

	quote := model.Quotes{Quote: Body.Quote, Author: Body.Author, Category: Body.Category}

	utility.DB.Create(&quote)

	ctx.JSON(http.StatusOK, gin.H{
		"quote": quote,
	})

}
