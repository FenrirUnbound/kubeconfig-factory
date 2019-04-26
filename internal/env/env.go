package env

import "os"

type Env interface {
	Get(string) string
	Set(string, string)
}

func Get(envName string) string {
	return os.Getenv(envName)
}

func Set(envName, value string) {
	os.Setenv(envName, value)
}
