package envsupport

import (
	"os"
	"sync"
)

const ProfileEnvVar = "profile"

var once sync.Once

var profile string

func Profile() string {
	once.Do(loadFunc)
	return profile
}

func loadFunc() {
	profile = os.Getenv(ProfileEnvVar)
	if profile == "" {
		panic("environment variable " + ProfileEnvVar + " not set")
	}
}
