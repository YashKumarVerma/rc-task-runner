package main

import (
	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"
	"github.com/YashKumarVerma/rc-task-runner/internal/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Initialize()
	ui.ContextPrint("thumbs_up", config.Load.Name)

	// initialize web server
	handler := gin.Default()
	handler.GET("/", controllers.HelloWorld)
	handler.POST("/run/:programID", controllers.Runner)

	handler.Run()
}
