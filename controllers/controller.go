package controllers

import (
	"context"
	"net/http"
	"parksInfoGo/apis"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDisneyQueues(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resp, err := apis.GetDisneyWaitTimesAllRides(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, resp)
}

func GetCaliforniaAdventureQueues(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resp, err := apis.GetCaliforniaAdventureWaitTimesAllRides(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(200, resp)
}
