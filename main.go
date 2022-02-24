package main

import (
	"log"
	"lupus/patapi/pkg/db"
	handler "lupus/patapi/pkg/http"
	patientfile "lupus/patapi/pkg/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error lodading .env file")
	}

	patientlister := patientfile.NewService(patientfile.NewService(db.NewDbConnect()))
	ph := handler.NewHandler(patientlister)

	//patientHandler := handler.NewHandler(patientfile.NewService(db.NewDbConnect()))

	router := gin.Default()
	handler.InitRoutes(router, ph)
	router.Run(":4545") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// psql, err := config.NewConnect()
	// fmt.Println("psql", psql)
	// fmt.Println(psql.GetAllPatients(context.TODO()))
}
