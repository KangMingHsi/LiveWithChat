package cmd

import "os"

// Gets environment string
func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}