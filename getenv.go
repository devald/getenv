package getenv

import "os"

// Variable retrieves the value of the environment variable with error handling
func Variable(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}
