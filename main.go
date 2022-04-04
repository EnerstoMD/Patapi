package main

import (
	"log"
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

	dbConn := db.NewDbConnect()

	patientlister := service.NewPatService(dbConn)
	calEventLister := service.NewCalService(dbConn)
	userService := service.NewUserService(dbConn)
	ph := handler.NewPatientHandler(patientlister)
	ch := handler.NewCalendarHandler(calEventLister)
	uh := handler.NewUserHandler(userService)

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	handler.InitRoutes(router, ph, ch, uh)
	router.Run(os.Getenv("PATAPI_PORT")) // listen and serve on 0.0.0.0:4545
}
