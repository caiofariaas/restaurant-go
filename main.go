package main

import (
	"os"

	"golang-restaurant-management/database"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())

    // Rotas públicas (sem autenticação)
    routes.UserRoutes(router) // Inclui SignUp e Login

    // Grupo de rotas protegidas
    protectedRoutes := router.Group("/")
    protectedRoutes.Use(middleware.Authentication())

    routes.FoodRoutes(protectedRoutes)
    routes.MenuRoutes(protectedRoutes)
    routes.TableRoutes(protectedRoutes)
    routes.OrderRoutes(protectedRoutes)
    routes.OrderItemRoutes(protectedRoutes)
    routes.InvoiceRoutes(protectedRoutes)

    router.Run(":" + port)
}
