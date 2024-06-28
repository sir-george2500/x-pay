// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sir-george2500/x-pay/config"
	_ "github.com/sir-george2500/x-pay/config"
	_ "github.com/sir-george2500/x-pay/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title x-pay API
// @version 1.0
// @description APIs for managing payments by BITS.

// @contact.name BITS Support
// @contact.email support@bits.com

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @Summary Show the status of the x-pay server.
// @Description get the status of the x-pay server.
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func main() {
   config.InitDB()  // Initialize the database connection

    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, World!",
        })
    })

    // Database connection test route
    router.GET("/ping", func(c *gin.Context) {
        if err := config.PingDB(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "status": "Database connection failed",
                "error":  err.Error(),
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "status": "Database connection successful",
        })
    })

    // Serve Swagger UI
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    router.Run(":8080")
}

