package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var logData LoggingData

func handleCreate(c *gin.Context) (int, interface{}) {
	var jsonBody SampleEntity
	logData = LoggingData{}
	if err := c.ShouldBind(&jsonBody); err != nil {
		Log(logData, "info", "BindErrorCreateRequest", err.Error())
		return JsonapiErrorResp(http.StatusNotAcceptable, fmt.Sprintf("Request json not compatible %s", err.Error()))
	}
	logData.RawInterface = jsonBody
	Log(logData, "info", "ReceivedCreateRequest", "Request received to create entity")

	return 201, jsonBody
}

func handleRead(c *gin.Context) (int, interface{}) {
	logData = LoggingData{}
	//var offset int64
	//var limit int64 = 100
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	log.Println(offsetStr, limitStr)
	Log(logData, "info", "ReceivedReadRequest", "Request received to read")

	return 200, viper.Get("name")
}
