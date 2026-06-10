package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getCallerFile() string {
	_, file, line, _ := runtime.Caller(2)
	wd, err := os.Getwd()
	if err == nil {
		if rel, err := filepath.Rel(wd, file); err == nil {
			file = rel
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
