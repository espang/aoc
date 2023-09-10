package aoc

import (
	"log/slog"
	"os"
)

const (
	EnvSessionCookie = "AOC_SESSION"
	EnvRootPath      = "AOC_ROOT_FOLDER"
)

func mustGetFromEnv(envVarName string) string {
	val := os.Getenv(envVarName)
	if val == "" {
		slog.Warn("missing required env var", slog.String("var-name", envVarName))
		panic("missing required env var")
	}
	return val
}
