package internal

import (
	"robot-hduin/config"
	"robot-hduin/internal/app"
	"robot-hduin/utils"
)

func Start() {
	config.Init()
	utils.WriteLogToFS()
	app.Init()
}
