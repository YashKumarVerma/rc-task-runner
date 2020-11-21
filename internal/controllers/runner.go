package controllers

import (
	"github.com/YashKumarVerma/rc-task-runner/internal/models"
	"github.com/gin-gonic/gin"
)

// HelloWorld to well, say hello to the world
func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"alive": true,
	})
}

// Runner to execute code
func Runner(ctx *gin.Context) {
	var body models.RequestBody
	ctx.BindJSON(&body)
	ctx.JSON(200, gin.H{
		"id":    body.ProgramID,
		"input": body.Input,
	})
}
