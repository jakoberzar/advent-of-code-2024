package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFileOrDie(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}
