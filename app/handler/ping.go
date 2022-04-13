package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sweeneyf/seal-it-api/pkg/logger"
)

func Ping(c *gin.Context) {

	reply := "hello from proof-of-deed @" + time.Now().Format("2006-01-02 15:04:05.000000")
	logger.Log.Info(reply)
	c.IndentedJSON(http.StatusCreated, reply)

}
