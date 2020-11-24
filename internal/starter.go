package internal

import (
	"miraigo-robot/config"
	"miraigo-robot/internal/app"
	"miraigo-robot/utils"
)

func Start() {
	config.Init()
	utils.WriteLogToFS()
	app.Init()
}
