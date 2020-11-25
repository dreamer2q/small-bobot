package internal

import (
	"miraigo-robot/internal/app"
	"miraigo-robot/utils"
)

func Start() {
	utils.WriteLogToFS()
	app.Init()
}
