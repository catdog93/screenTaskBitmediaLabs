package main

import (
	"github.com/gin-gonic/gin"
	"github.com/screenTaskBitmediaLabs/controller"
	"github.com/screenTaskBitmediaLabs/repository"
	"log"
	"os"
)

const ()

var (
	dbHost = getEnvVariableOrDie("DB_HOST")
	dbName = getEnvVariableOrDie("DB_NAME")

	//usersService =
	//gamesService =
)

func main() {
	log.Println(repository.DBExist(dbHost, dbName))
	createDefaultRouter()
}

func createDefaultRouter() {
	router := gin.Default()

	router.POST("users", controller.GetUsers)
	router.GET("statistic", controller.GetStatistic)
	router.GET("rating", controller.GetRating)

	router.Run()
}

// function tries to get env variable for DB, it'll stop app if variable isn't found
func getEnvVariableOrDie(envVariableName string) string {
	env, found := os.LookupEnv(envVariableName)
	if !found {
		log.Fatalf("Failed to load env variable, %v", envVariableName)
	}
	return env
}
