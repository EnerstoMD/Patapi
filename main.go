package main

import (
	"log"
	"lupus/patapi/pkg/db"
	handler "lupus/patapi/pkg/http"
	service "lupus/patapi/pkg/services"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
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

	//use https://github.com/itsjamie/gin-cors ??
	router := gin.New()

	router.Use(gin.Recovery(), gin.Logger(), cors.Middleware(cors.Config{
		// Origins:        "http://localhost:4200,http://localhost,http://51.15.205.164,http://51.15.205.164/",
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS, PATCH",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second,
		Credentials:    true,
		// ValidateHeaders: true,
	}))
	//cors shouldnot be allowing every orign
	router.SetTrustedProxies(nil)
	handler.InitRoutes(router, ph, ch, uh)
	router.Run(":4545") // listen and serve on 0.0.0.0:4545
}
