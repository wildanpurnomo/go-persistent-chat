package environmentLibs

import "os"

func IsRunningOnProductionEnvironment() bool {
	return os.Getenv("APP_MODE") == "production"
}
