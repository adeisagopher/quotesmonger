package main

import (
	"github.com/adeisagopher/quotesmonger/controllers"
	"github.com/adeisagopher/quotesmonger/utility"
	"github.com/gin-gonic/gin"

	_ "github.com/adeisagopher/quotesmonger/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utility.ConnectToDB()
}

// @title Quotesmonger API
// @version 1.0
// @description A simple quotes API with multiple categories written in GO.

// @contact.name adeisagopher
// @contact.url https://github.com/adeisagopher
// @contact.email adeisagopher@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	r := gin.Default()

	// Swagger docs route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CORs middleware
	r.Use(controllers.CORSMiddleware())

	r.GET("/quotes", controllers.GetQuotes)
	r.GET("/quote/:id", controllers.GetQuoteByID)
	r.GET("/quotes/category", controllers.GetQuoteByCategory)
	r.POST("/quote", controllers.PostQuotes)
	r.Run()
}

// nodemon --exec go run main.go --signal SIGTERM
