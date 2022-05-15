package main

import (
	"log"
	"lupus/patapi/pkg/auth"
	"lupus/patapi/pkg/db"
	handler "lupus/patapi/pkg/http"
	service "lupus/patapi/pkg/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error lodading .env file")
	}

	DbSources := db.NewDbConnect()

	patientlister := service.NewPatService(DbSources)
	calEventLister := service.NewCalService(DbSources)
	authService := auth.NewAuthService(DbSources)
	userService := service.NewUserService(DbSources, authService)
	ph := handler.NewPatientHandler(patientlister)
	ch := handler.NewCalendarHandler(calEventLister)
	uh := handler.NewUserHandler(userService)

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	handler.InitRoutes(router, ph, ch, uh, authService)
	router.Run(":4545") // listen and serve on 0.0.0.0:4545
}
