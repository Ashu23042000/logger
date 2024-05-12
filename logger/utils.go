package logger

import (
	"fmt"
	"runtime"
)

func getCallerFile() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}
