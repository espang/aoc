package aoc

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"strings"
)

type SplitOp int

const (
	Comma SplitOp = iota + 1
	CommaSpace
	Newline
)

// SplitBy splits an input by the given operation and will also trim
// any whitespace around the input.
func SplitBy(input string, op SplitOp) []string {
	var sep string
	switch op {
	case Comma:
		sep = ","
	case CommaSpace:
		sep = ", "
	case Newline:
		sep = "\n"
	}
	return strings.Split(strings.TrimSpace(input), sep)
}

func SplitByComma(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

func SplitByLine(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func LoadFrom(filename string) string {
	return mustLoadFile(filename)
}

func MustMakeInputAvailable(ctx context.Context, year, day int) (filename string) {
	rootPath := mustGetFromEnv(EnvRootPath)
	filename = filepath.Join(rootPath, "inputs", fmt.Sprintf("%d_%d.txt", year, day))
	logger := slog.Default().With(
		slog.Int("year", year),
		slog.Int("day", day),
		slog.String("filename", filename),
	)
	logger.DebugContext(ctx, "loading input")

	if fileExists(filename) {
		logger.DebugContext(ctx, "file exists; loading data from FS")
		return filename
	}

	logger.DebugContext(ctx, "input doesn't exist; requesting data from 'adventofcode.com'")
	content := downloadFile(ctx, year, day)
	mustStoreFile(filename, content)
	return filename
}

func downloadFile(ctx context.Context, year, day int) string {
	sessionValue := os.Getenv(EnvSessionCookie)
	if sessionValue == "" {
		slog.Warn("set the session value from the cookie", slog.String("env-var-name", EnvSessionCookie))
		panic("set session cookie to download inputs")
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionValue,
	}

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:       jar,
		Transport: &http.Transport{},
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Warn("failed to construct http Request", slog.Any("error", err))
		panic(err)
	}
	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		slog.Warn("failed to make http Request", slog.Any("error", err))
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Warn("unexpected status code", slog.String("status", resp.Status), slog.Int("status-code", resp.StatusCode))
		panic("unexpected status code")
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Warn("failed to read response", slog.Any("error", err))
		panic(err)
	}

	return string(content)
}

func mustLoadFile(filename string) string {
	buf, err := os.ReadFile(filename)
	if err != nil {
		slog.Warn("failed to load file", slog.String("filename", filename), slog.Any("error", err))
		panic(err)
	}
	return string(buf)
}

func mustStoreFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		slog.Warn("failed to write file", slog.String("filename", filename), slog.Any("error", err))
		panic(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return !info.IsDir()
}
