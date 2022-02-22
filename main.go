package main

import (
	"log"
	"lupus/patapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error lodading .env file")
	}

	router := gin.Default()
	//config.NewConnect()
	routes.Routes(router)
	router.Run(":4545") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// psql, err := config.NewConnect()
	// fmt.Println("psql", psql)
	// fmt.Println(psql.GetAllPatients(context.TODO()))
}
