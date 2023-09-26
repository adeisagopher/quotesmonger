package main

import (
	"github.com/adeisagopher/quotesmonger/controllers"
	"github.com/adeisagopher/quotesmonger/utility"
	"github.com/gin-gonic/gin"
)

func init() {
	utility.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/quotes", controllers.GetQuotes)
	r.GET("/quote/:id", controllers.GetQuoteByID)
	r.GET("/quotes/category", controllers.GetQuoteByCategory)
	r.POST("/quote", controllers.PostQuotes)
	r.Run()
}

// nodemon --exec go run main.go --signal SIGTERM
