package util

import (
	"os"
	"strings"
)

func GetLines(file string) []string {
	content, _ := os.ReadFile(file)
	return strings.Split(string(content), "\n")
}
