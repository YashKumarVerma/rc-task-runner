package controllers

import (
	"github.com/YashKumarVerma/rc-task-runner/internal/dispatcher"
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
	var response string
	ctx.BindJSON(&body)

	if !itemExists(dispatcher.ValidQuestions, body.ProgramID) {
		response = "no question with that id"
	} else {
		response = dispatcher.DispatchOutput(body.ProgramID, body.Input)
	}

	ctx.JSON(200, gin.H{
		"id":     body.ProgramID,
		"input":  body.Input,
		"output": response,
	})
}

// function to check if item exists in array
func itemExists(list []string, item string) bool {
	for _, element := range list {
		if item == element {
			return true
		}
	}
	return false
}
