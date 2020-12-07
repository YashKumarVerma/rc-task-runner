package main

import (
	"strconv"

	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/rc-task-runner/internal/config"
	"github.com/YashKumarVerma/rc-task-runner/internal/controllers"
	"github.com/YashKumarVerma/rc-task-runner/internal/dispatcher"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Initialize()
	dispatcher.CheckInventory()
	for _, element := range dispatcher.ValidQuestions {
		ui.ContextPrint("gear", element)
	}

	ui.ContextPrint("thumbs_up", config.Load.Name)

	// initialize web server
	handler := gin.Default()
	handler.GET("/", controllers.HelloWorld)
	handler.POST("/run", controllers.Runner)
	handler.GET("/sync", controllers.Sync)

	handler.Run(":" + strconv.Itoa(config.Load.Port))
}
