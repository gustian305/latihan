package routes

import (
	"time"

	"latihan/controller/handler"
	"latihan/controller/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	// 1. Inisialisasi engine Gin
	r := gin.Default()

	// 2. Global Middleware (Opsional: misalnya CORS atau Logger)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 3. Pengelompokan Route (Versioning)
	api := r.Group("/api")
	{
		studentService := services.NewStudentService()
		studentHandler := handlers.NewStudentHandler(studentService)

		// Grouping untuk resource "student"
		studentRoutes := api.Group("/student")
		{
			studentRoutes.POST("/", studentHandler.CreateStudent)
			studentRoutes.GET("/", studentHandler.GetAllStudents)
		}
	}

	return r
}
