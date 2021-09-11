package logsLibs

import (
	"log"

	environmentLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/environment"
)

func LogIfDebug(debugVariable interface{}) {
	if !environmentLibs.IsRunningOnProductionEnvironment() {
		log.Println(debugVariable)
	}
}

func LogAndTerminate(debugVariable interface{}) {
	if !environmentLibs.IsRunningOnProductionEnvironment() {
		log.Fatal(debugVariable)
	}
}
