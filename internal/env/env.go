package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment string

const (
	EnvironmentKey = "ENVIRONMENT"

	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

var environment Environment

func Load() (err error) {
	// Load .env if exists else fail silently
	if err = godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return
	}

	environment = Environment(os.Getenv(EnvironmentKey))
	switch environment {
	case EnvironmentDevelopment, EnvironmentProduction:
		return
	default:
		return fmt.Errorf(
			"unexpected %s value: expected: %q or %q, got %q",
			EnvironmentKey,
			EnvironmentDevelopment,
			EnvironmentProduction,
			environment,
		)
	}

}

func Getenv[T ~string](key T, defaultValue T) T {
	if value, ok := os.LookupEnv(string(key)); ok {
		return T(value)
	}
	return defaultValue
}

func IsDev() bool { return environment == EnvironmentDevelopment }
