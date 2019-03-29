package main

import (
	"fmt"
	"litchi-rest-go/api"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	router := gin.New()

	srvPort, hasPort := os.LookupEnv("SERVICE_PORT")
	if hasPort == false {
		//log.Fatal("Service port missing")
		srvPort = "8001"
	}
	dbInfo, hasDbInfo := os.LookupEnv("DATABASE_INFO")
	if hasDbInfo == false {
		//log.Fatal("Database name missing")
		dbInfo = "localhost"
	}
	dbName, hasDbName := os.LookupEnv("DATABASE_NAME")
	if hasDbName == false {
		//log.Fatal("Database name missing")
		dbName = "litchi"
	}

	viper.SetDefault("dbName", dbName)
	viper.SetDefault("dbInfo", dbInfo)
	viper.SetConfigName("general")
	viper.AddConfigPath("./api/config/")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Global middleware
	// Logger middleware by default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	// CORS settings
	router.Use(func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Add("Access-Control-Max-Age", "10000")
		context.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS")
		context.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization,Content-Type,Accept")
		context.Next()
	})

	v1 := router.Group("/v1")

	api.MountRoute(v1)

	router.Run(":" + srvPort)
}
