package main

import (
	"log"
	"lupus/patapi/pkg/db"
	handler "lupus/patapi/pkg/http"
	service "lupus/patapi/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
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

	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	//cors shouldnot be allowing every orign
	router.SetTrustedProxies(nil)
	handler.InitRoutes(router, ph, ch, uh)
	router.Run(":4545") // listen and serve on 0.0.0.0:4545
}
