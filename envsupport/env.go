package envsupport

import (
	"os"
	"sync"
)

const ProfileEnvVar = "profile"
const ConfigPathVar = "config_path"

var once sync.Once

var configPath string
var profile string

func Profile() string {
	once.Do(loadFunc)
	return profile
}

func ConfigPath() string {
	once.Do(loadFunc)
	return configPath
}

func loadFunc() {
	profile = os.Getenv(ProfileEnvVar)
	if profile == "" {
		panic("environment variable " + ProfileEnvVar + " not set")
	}
	configPath = os.Getenv(ConfigPathVar)
	if configPath == "" {
		panic("environment variable " + ConfigPathVar + " not set")
	}
}
